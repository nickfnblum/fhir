#
# Copyright 2022 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""Utilities for working with Value Sets."""

import functools
import itertools
from typing import Iterable, Set

import sqlalchemy

from proto.google.fhir.proto.r4.core.resources import structure_definition_pb2
from proto.google.fhir.proto.r4.core.resources import value_set_pb2
from google.fhir.utils import fhir_package


class ValueSetResolver:
  """Utility for retrieving and resolving value set resources to the codes they contain.

  Attributes:
    package_manager: The FhirPackageManager object to use when retrieving
      resource definitions. The FhirPackage objects contained in package_manager
      will be consulted when value set resource definitions are needed. The
      package manager should contain common resources, for instance, ones from
      the US Core implementation guide, in order to ensure definitions for all
      relevant value sets may be found. If a requisite value set definition is
      not present in the package manager, the resolver will throw an error
      instead of attempting to retrieve it over the network.
  """

  def __init__(self, package_manager: fhir_package.FhirPackageManager) -> None:
    self.package_manager = package_manager

  def value_sets_from_fhir_package(
      self,
      package: fhir_package.FhirPackage) -> Iterable[value_set_pb2.ValueSet]:
    """Retrieves all value sets referenced by the given FHIR package.

    Finds all value set resources in the package as well as any value sets
    referenced by structure definitions in the package.

    Args:
      package: The FHIR package from which to retrieve value sets.

    Yields:
      All value sets referenced by the FHIR package.
    """
    value_sets_from_structure_definitions = itertools.chain.from_iterable(
        self.value_sets_from_structure_definition(structure_definition)
        for structure_definition in package.structure_definitions)
    value_sets = itertools.chain(
        package.value_sets,
        value_sets_from_structure_definitions,
    )
    yield from _unique_by_url(value_sets)

  def value_sets_from_structure_definition(
      self, structure_definition: structure_definition_pb2.StructureDefinition
  ) -> Iterable[value_set_pb2.ValueSet]:
    """Retrieves all value sets referenced by the given structure definition.

    Finds the union of value sets bound to any element from either the
    differential or snapshot definition.

    Args:
      structure_definition: The structure definition from which to retrieve
        value sets.

    Yields:
      All value sets referenced by the structure definition.
    """
    elements = itertools.chain(structure_definition.differential.element,
                               structure_definition.snapshot.element)
    value_set_urls = (
        element.binding.value_set.value
        for element in elements
        if element.binding.value_set.value)
    value_sets = (self.value_set_from_url(url) for url in value_set_urls)
    yield from _unique_by_url(value_sets)

  def value_set_from_url(self, url: str) -> value_set_pb2.ValueSet:
    """Retrieves the value set for the given URL.

    The value set is assumed to be a member of one of the packages contained in
    self.package_manager. This function will not attempt to look up resources
    over the network in other locations.

    Args:
      url: The url of the value set to retrieve.

    Returns:
      The value set for the given URL.
    """
    value_set = self.package_manager.get_resource(url)
    if value_set is None:
      raise ValueError('Unable to find value set in given resolver packages.')
    elif not isinstance(value_set, value_set_pb2.ValueSet):
      raise ValueError('URL does not refer to a value set, found: %s' %
                       value_set.DESCRIPTOR.name)
    else:
      return value_set


def valueset_codes_insert_statement_for(
    value_set: value_set_pb2.ValueSet,
    table: sqlalchemy.sql.expression.TableClause) -> sqlalchemy.sql.dml.Insert:
  """Builds an INSERT statement for placing the value set's codes into the given table.

  The INSERT may be used to build a valueset_codes table as described by:
  https://github.com/FHIR/sql-on-fhir/blob/master/sql-on-fhir.md#valueset-support

  Returns an sqlalchemy insert expression which inserts all of the value set's
  expanded codes into the given table which do not already exist in the table.
  The query will avoid inserting duplicate rows if some of the codes are already
  present in the given table. It will not attempt to perform an 'upsert' or
  modify any existing rows.

  Args:
    value_set: The expanded value set with codes to insert into the given table.
      The value set should have already been expanded by the
      ValueSetResolver.expand_value_set method.
    table: The SqlAlchemy table to receive the INSERT. May be an sqlalchemy
      Table or TableClause object. The table is assumed to have the columns
      'valueseturi', 'valuesetversion', 'system', 'code.'

  Returns:
    The sqlalchemy insert expression which you may execute to perform the actual
    database write.
  """
  if not value_set.expansion.contains:
    raise ValueError(
        'Value set must be expanded. Call ValueSetResolver.expand_value_set on '
        'the value set to expand it first.')

  # Build a SELECT statement for each code.
  code_literals = (
      _code_as_select_literal(value_set, code)
      for code in value_set.expansion.contains)
  # UNION each SELECT to build a single select subquery for all codes.
  codes = functools.reduce(lambda query, literal: query.union_all(literal),
                           code_literals).alias('codes')
  # Filter the codes to those not already present in `table` with a LEFT JOIN.
  new_codes = sqlalchemy.select((codes,)).select_from(
      codes.outerjoin(
          table,
          sqlalchemy.and_(
              codes.c.valueseturi == table.c.valueseturi,
              codes.c.valuesetversion == table.c.valuesetversion,
              codes.c.system == table.c.system,
              codes.c.code == table.c.code,
          ))).where(
              sqlalchemy.and_(
                  table.c.valueseturi.is_(None),
                  table.c.valuesetversion.is_(None),
                  table.c.system.is_(None),
                  table.c.code.is_(None),
              ))
  return table.insert().from_select(new_codes.columns, new_codes)


def _code_as_select_literal(
    value_set: value_set_pb2.ValueSet,
    code: value_set_pb2.ValueSet.Expansion.Contains) -> sqlalchemy.select:
  """Builds a SELECT statement for the literals in the given code."""
  return sqlalchemy.select((
      sqlalchemy.sql.expression.literal(
          value_set.url.value).label('valueseturi'),
      sqlalchemy.sql.expression.literal(
          value_set.version.value).label('valuesetversion'),
      sqlalchemy.sql.expression.literal(code.system.value).label('system'),
      sqlalchemy.sql.expression.literal(code.code.value).label('code'),
  ))


def _unique_by_url(
    value_sets: Iterable[value_set_pb2.ValueSet]
) -> Iterable[value_set_pb2.ValueSet]:
  """Filters value_sets to remove those with duplicate URLs.

  Args:
    value_sets: The value sets to filter.

  Yields:
    The value sets filtered to only those without duplicate URLs.
  """
  seen: Set[str] = set()

  for value_set in value_sets:
    url = value_set.url.value
    if url not in seen:
      seen.add(url)
      yield value_set
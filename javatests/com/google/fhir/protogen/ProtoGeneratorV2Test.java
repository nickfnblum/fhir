//    Copyright 2023 Google Inc.
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        https://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package com.google.fhir.protogen;

import static com.google.common.collect.ImmutableList.toImmutableList;
import static com.google.common.truth.extensions.proto.ProtoTruth.assertThat;

import com.google.common.collect.ImmutableList;
import com.google.errorprone.annotations.CanIgnoreReturnValue;
import com.google.fhir.common.InvalidFhirException;
import com.google.fhir.proto.Annotations.FhirVersion;
import com.google.fhir.proto.ProtoGeneratorAnnotations;
import com.google.fhir.proto.ProtogenConfig;
import com.google.fhir.r4.core.Bundle;
import com.google.fhir.r4.core.ContainedResource;
import com.google.fhir.r4.core.StructureDefinition;
import com.google.fhir.r4.core.StructureDefinitionKindCode;
import com.google.fhir.r4.core.TypeDerivationRuleCode;
import com.google.protobuf.DescriptorProtos.DescriptorProto;
import com.google.protobuf.DescriptorProtos.FileDescriptorProto;
import com.google.protobuf.Descriptors.Descriptor;
import com.google.testing.junit.testparameterinjector.TestParameter;
import com.google.testing.junit.testparameterinjector.TestParameter.TestParameterValuesProvider;
import com.google.testing.junit.testparameterinjector.TestParameterInjector;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import org.junit.Test;
import org.junit.runner.RunWith;

@RunWith(TestParameterInjector.class)
public final class ProtoGeneratorV2Test {
  private static FhirPackage r4Package = null;

  /**
   * Gets the R4 package. Lazy loaded to avoid questions of class/test/parameter initialization
   * order.
   */
  private static FhirPackage getR4Package() throws IOException, InvalidFhirException {
    if (r4Package == null) {
      r4Package = FhirPackage.load("spec/fhir_r4_package.zip");
    }
    return r4Package;
  }

  public static ProtoGeneratorV2 makeR4ProtoGenerator() throws Exception {
    ProtogenConfig config =
        ProtogenConfig.newBuilder()
            .setProtoPackage("google.fhir.r4.core")
            .setJavaProtoPackage("com.google.fhir.r4.core")
            .setLicenseDate("1995")
            .setSourceDirectory("proto/google/fhir/proto/r4/core")
            .setFhirVersion(FhirVersion.R4)
            .build();

    FhirPackage r4Package = getR4Package();
    ValueSetGeneratorV2 valueSetGenerator = new ValueSetGeneratorV2(config, r4Package);
    FileDescriptorProto codesFileDescriptor = valueSetGenerator.forCodesUsedIn(r4Package);
    FileDescriptorProto valueSetsFileDescriptor = valueSetGenerator.forValueSetsUsedIn(r4Package);

    return new ProtoGeneratorV2(
        config,
        r4Package,
        valueSetGenerator.getBoundCodeGenerator(codesFileDescriptor, valueSetsFileDescriptor));
  }

  /**
   * Sorts all messages in a file, to make it easier to compare. This is not technically necessary
   * if we compare with ignoringRepeatedFieldOrder, but the diffs are so large if order is wrong
   * that it is very hard to read the diff.
   */
  FileDescriptorProto sorted(FileDescriptorProto descriptor) {
    var messages = new ArrayList<>(descriptor.getMessageTypeList());
    messages.sort((a, b) -> a.getName().compareTo(b.getName()));
    return descriptor.toBuilder().clearMessageType().addAllMessageType(messages).build();
  }

  /**
   * Cleans up some elements of the file that cause false diffs, such as protogenerator annotations,
   * which are only used for printing comments, dependencies that are pruned by post-processing, and
   * datatypes that are hardcoded rather than generated.
   */
  FileDescriptorProto cleaned(FileDescriptorProto file) {
    var builder = file.toBuilder().clearName().clearDependency();
    builder.getOptionsBuilder().clearGoPackage();

    builder.clearMessageType();
    for (var message : file.getMessageTypeList()) {
      // Some datatypes are still hardcoded and added in the generation script, rather than by
      // the protogenerator.
      if (!message.getName().equals("CodingWithFixedCode")
          && !message.getName().equals("CodingWithFixedSystem")
          && !message.getName().equals("Extension")) {
        builder.addMessageType(clean(message.toBuilder()));
      }
    }
    return builder.build();
  }

  /**
   * Cleans up protogenerator annotations, which are only used for printing comments. Also removes
   * reserved ranges, since the ProtoGenerator represents these using the reservedReason annotation
   * (and so doesn't have reserved ranges until printing).
   */
  @CanIgnoreReturnValue
  DescriptorProto.Builder clean(DescriptorProto.Builder builder) {
    builder.getOptionsBuilder().clearExtension(ProtoGeneratorAnnotations.messageDescription);
    builder.clearReservedRange();

    var fields = new ArrayList<>(builder.getFieldBuilderList());
    builder.clearField();

    for (var field : fields) {
      field.getOptionsBuilder().clearExtension(ProtoGeneratorAnnotations.fieldDescription);
      if (!field.getOptions().hasExtension(ProtoGeneratorAnnotations.reservedReason)) {
        builder.addField(field);
      }
    }
    for (var nested : builder.getNestedTypeBuilderList()) {
      clean(nested);
    }
    return builder;
  }

  private static boolean isCoreResource(StructureDefinition def) {
    return def.getKind().getValue() == StructureDefinitionKindCode.Value.RESOURCE
        && def.getDerivation().getValue() == TypeDerivationRuleCode.Value.SPECIALIZATION
        && !def.getAbstract().getValue();
  }

  /**
   * Tests that generating the R4 Datatypes file using generateLegacyDatatypesFileDescriptor
   * generates descriptors that match the currently-checked in R4 datatypes. This serves as both a
   * whole lot of unit tests, and as a regression test to guard against checking in a change that
   * would alter the R4 Core protos.
   */
  @Test
  public void r4RegressionTest_generateLegacyDatatypesFileDescriptor() throws Exception {

    ImmutableList<String> resourceNames =
        ContainedResource.getDescriptor().getFields().stream()
            .map(field -> field.getMessageType().getName())
            .collect(toImmutableList());
    FileDescriptorProto descriptor =
        makeR4ProtoGenerator().generateLegacyDatatypesFileDescriptor(resourceNames);
    assertThat(sorted(cleaned(descriptor)))
        .ignoringRepeatedFieldOrder()
        .isEqualTo(
            sorted(cleaned(com.google.fhir.r4.core.String.getDescriptor().getFile().toProto())));
  }

  /** Parameter provider for getting all resources that have a file generated. */
  private static final class ResourceProvider implements TestParameterValuesProvider {
    private static final int NUM_RESOURCES = 145;

    @Override
    public List<StructureDefinition> provideValues() {
      FhirPackage r4Package;
      try {
        r4Package = getR4Package();
      } catch (IOException | InvalidFhirException e) {
        throw new AssertionError(e);
      }

      List<StructureDefinition> structDefs = new ArrayList<>();
      for (StructureDefinition structDef : r4Package.structureDefinitions()) {
        if (isCoreResource(structDef)
            // Skip bundle since it is part of bundle_and_contained_resource
            && !structDef.getName().getValue().equals("Bundle")) {
          structDefs.add(structDef);
        }
      }
      // Sanity check that we're actually testing all the resources.
      if (structDefs.size() != NUM_RESOURCES) {
        throw new AssertionError(
            "Expected " + NUM_RESOURCES + " resources, got " + structDefs.size());
      }
      return structDefs;
    }
  }

  /**
   * Tests that generating the R4 Resource files using generateResourceFileDescriptor generates
   * descriptors that match the currently-checked in R4 Resource files. This serves as both a whole
   * lot of unit tests, and as a regression test to guard against checking in a change that would
   * alter the R4 Core protos.
   *
   * <p>This test is parameterized on resource StructureDefinition, as provided by
   * `ResourceProvider`, which provides one resource StructureDefinition per test.
   */
  @Test
  public void r4RegressionTest_generateResourceFileDescriptor(
      @TestParameter(valuesProvider = ResourceProvider.class) StructureDefinition resource)
      throws Exception {
    ProtoGeneratorV2 protoGenerator = makeR4ProtoGenerator();
    protoGenerator.addSearchParameters();

    FileDescriptorProto file = protoGenerator.generateResourceFileDescriptor(resource);

    // Get the checked-in FileDescriptor for the resource.
    String name = file.getOptions().getJavaPackage() + "." + file.getMessageType(0).getName();
    FileDescriptorProto descriptor =
        ((Descriptor) Class.forName(name).getMethod("getDescriptor").invoke(null))
            .getFile()
            .toProto();

    assertThat(cleaned(file)).ignoringRepeatedFieldOrder().isEqualTo(cleaned(descriptor));
  }

  /**
   * Tests that generating the R4 bundle_and_contained_resource file using
   * generateBundleAndContainedResource generates file descriptors that match the currently-checked
   * in R4 bundle_and_contained_resources file. This serves as both a unit tests, and
   * as a regression test to guard against checking in a change that would alter the R4 Core protos.
   */
  @Test
  public void r4RegressionTest_generateBundleAndContainedResource() throws Exception {
    ProtoGeneratorV2 protoGenerator = makeR4ProtoGenerator();
    protoGenerator.addSearchParameters();

    List<String> resourceNames = new ArrayList<>();
    StructureDefinition bundle = null;
    for (StructureDefinition structDef : r4Package.structureDefinitions()) {
      if (structDef.getName().getValue().equals("Bundle")) {
        bundle = structDef;
      }
      if (isCoreResource(structDef)) {
        resourceNames.add(structDef.getName().getValue());
      }
    }

    if (bundle == null) {
      throw new AssertionError("No Bundle found");
    }

    assertThat(cleaned(protoGenerator.generateBundleAndContainedResource(bundle, resourceNames, 0)))
        .ignoringRepeatedFieldOrder()
        .isEqualTo(cleaned(Bundle.getDescriptor().getFile().toProto()));
  }
}
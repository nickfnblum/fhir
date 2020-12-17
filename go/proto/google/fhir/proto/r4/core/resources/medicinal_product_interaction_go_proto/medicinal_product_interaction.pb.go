//    Copyright 2019 Google Inc.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: proto/google/fhir/proto/r4/core/resources/medicinal_product_interaction.proto

package medicinal_product_interaction_go_proto

import (
	any "github.com/golang/protobuf/ptypes/any"
	_ "github.com/google/fhir/go/proto/google/fhir/proto/annotations_go_proto"
	datatypes_go_proto "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Auto-generated from StructureDefinition for MedicinalProductInteraction, last
// updated 2019-11-01T09:29:23.356+11:00. MedicinalProductInteraction. See
// http://hl7.org/fhir/StructureDefinition/MedicinalProductInteraction
type MedicinalProductInteraction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Logical id of this artifact
	Id *datatypes_go_proto.Id `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Metadata about the resource
	Meta *datatypes_go_proto.Meta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *datatypes_go_proto.Uri `protobuf:"bytes,3,opt,name=implicit_rules,json=implicitRules,proto3" json:"implicit_rules,omitempty"`
	// Language of the resource content
	Language *datatypes_go_proto.Code `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *datatypes_go_proto.Narrative `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	// Contained, inline Resources
	Contained []*any.Any `protobuf:"bytes,6,rep,name=contained,proto3" json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []*datatypes_go_proto.Extension `protobuf:"bytes,8,rep,name=extension,proto3" json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []*datatypes_go_proto.Extension `protobuf:"bytes,9,rep,name=modifier_extension,json=modifierExtension,proto3" json:"modifier_extension,omitempty"`
	// The medication for which this is a described interaction
	Subject []*datatypes_go_proto.Reference `protobuf:"bytes,10,rep,name=subject,proto3" json:"subject,omitempty"`
	// The interaction described
	Description *datatypes_go_proto.String                 `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	Interactant []*MedicinalProductInteraction_Interactant `protobuf:"bytes,12,rep,name=interactant,proto3" json:"interactant,omitempty"`
	// The type of the interaction e.g. drug-drug interaction, drug-food
	// interaction, drug-lab test interaction
	Type *datatypes_go_proto.CodeableConcept `protobuf:"bytes,13,opt,name=type,proto3" json:"type,omitempty"`
	// The effect of the interaction, for example "reduced gastric absorption of
	// primary medication"
	Effect *datatypes_go_proto.CodeableConcept `protobuf:"bytes,14,opt,name=effect,proto3" json:"effect,omitempty"`
	// The incidence of the interaction, e.g. theoretical, observed
	Incidence *datatypes_go_proto.CodeableConcept `protobuf:"bytes,15,opt,name=incidence,proto3" json:"incidence,omitempty"`
	// Actions for managing the interaction
	Management *datatypes_go_proto.CodeableConcept `protobuf:"bytes,16,opt,name=management,proto3" json:"management,omitempty"`
}

func (x *MedicinalProductInteraction) Reset() {
	*x = MedicinalProductInteraction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MedicinalProductInteraction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MedicinalProductInteraction) ProtoMessage() {}

func (x *MedicinalProductInteraction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MedicinalProductInteraction.ProtoReflect.Descriptor instead.
func (*MedicinalProductInteraction) Descriptor() ([]byte, []int) {
	return file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_rawDescGZIP(), []int{0}
}

func (x *MedicinalProductInteraction) GetId() *datatypes_go_proto.Id {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *MedicinalProductInteraction) GetMeta() *datatypes_go_proto.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *MedicinalProductInteraction) GetImplicitRules() *datatypes_go_proto.Uri {
	if x != nil {
		return x.ImplicitRules
	}
	return nil
}

func (x *MedicinalProductInteraction) GetLanguage() *datatypes_go_proto.Code {
	if x != nil {
		return x.Language
	}
	return nil
}

func (x *MedicinalProductInteraction) GetText() *datatypes_go_proto.Narrative {
	if x != nil {
		return x.Text
	}
	return nil
}

func (x *MedicinalProductInteraction) GetContained() []*any.Any {
	if x != nil {
		return x.Contained
	}
	return nil
}

func (x *MedicinalProductInteraction) GetExtension() []*datatypes_go_proto.Extension {
	if x != nil {
		return x.Extension
	}
	return nil
}

func (x *MedicinalProductInteraction) GetModifierExtension() []*datatypes_go_proto.Extension {
	if x != nil {
		return x.ModifierExtension
	}
	return nil
}

func (x *MedicinalProductInteraction) GetSubject() []*datatypes_go_proto.Reference {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *MedicinalProductInteraction) GetDescription() *datatypes_go_proto.String {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *MedicinalProductInteraction) GetInteractant() []*MedicinalProductInteraction_Interactant {
	if x != nil {
		return x.Interactant
	}
	return nil
}

func (x *MedicinalProductInteraction) GetType() *datatypes_go_proto.CodeableConcept {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *MedicinalProductInteraction) GetEffect() *datatypes_go_proto.CodeableConcept {
	if x != nil {
		return x.Effect
	}
	return nil
}

func (x *MedicinalProductInteraction) GetIncidence() *datatypes_go_proto.CodeableConcept {
	if x != nil {
		return x.Incidence
	}
	return nil
}

func (x *MedicinalProductInteraction) GetManagement() *datatypes_go_proto.CodeableConcept {
	if x != nil {
		return x.Management
	}
	return nil
}

// The specific medication, food or laboratory test that interacts
type MedicinalProductInteraction_Interactant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique id for inter-element referencing
	Id *datatypes_go_proto.String `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []*datatypes_go_proto.Extension `protobuf:"bytes,2,rep,name=extension,proto3" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []*datatypes_go_proto.Extension                `protobuf:"bytes,3,rep,name=modifier_extension,json=modifierExtension,proto3" json:"modifier_extension,omitempty"`
	Item              *MedicinalProductInteraction_Interactant_ItemX `protobuf:"bytes,4,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *MedicinalProductInteraction_Interactant) Reset() {
	*x = MedicinalProductInteraction_Interactant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MedicinalProductInteraction_Interactant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MedicinalProductInteraction_Interactant) ProtoMessage() {}

func (x *MedicinalProductInteraction_Interactant) ProtoReflect() protoreflect.Message {
	mi := &file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MedicinalProductInteraction_Interactant.ProtoReflect.Descriptor instead.
func (*MedicinalProductInteraction_Interactant) Descriptor() ([]byte, []int) {
	return file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MedicinalProductInteraction_Interactant) GetId() *datatypes_go_proto.String {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *MedicinalProductInteraction_Interactant) GetExtension() []*datatypes_go_proto.Extension {
	if x != nil {
		return x.Extension
	}
	return nil
}

func (x *MedicinalProductInteraction_Interactant) GetModifierExtension() []*datatypes_go_proto.Extension {
	if x != nil {
		return x.ModifierExtension
	}
	return nil
}

func (x *MedicinalProductInteraction_Interactant) GetItem() *MedicinalProductInteraction_Interactant_ItemX {
	if x != nil {
		return x.Item
	}
	return nil
}

// The specific medication, food or laboratory test that interacts
type MedicinalProductInteraction_Interactant_ItemX struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Choice:
	//	*MedicinalProductInteraction_Interactant_ItemX_Reference
	//	*MedicinalProductInteraction_Interactant_ItemX_CodeableConcept
	Choice isMedicinalProductInteraction_Interactant_ItemX_Choice `protobuf_oneof:"choice"`
}

func (x *MedicinalProductInteraction_Interactant_ItemX) Reset() {
	*x = MedicinalProductInteraction_Interactant_ItemX{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MedicinalProductInteraction_Interactant_ItemX) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MedicinalProductInteraction_Interactant_ItemX) ProtoMessage() {}

func (x *MedicinalProductInteraction_Interactant_ItemX) ProtoReflect() protoreflect.Message {
	mi := &file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MedicinalProductInteraction_Interactant_ItemX.ProtoReflect.Descriptor instead.
func (*MedicinalProductInteraction_Interactant_ItemX) Descriptor() ([]byte, []int) {
	return file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (m *MedicinalProductInteraction_Interactant_ItemX) GetChoice() isMedicinalProductInteraction_Interactant_ItemX_Choice {
	if m != nil {
		return m.Choice
	}
	return nil
}

func (x *MedicinalProductInteraction_Interactant_ItemX) GetReference() *datatypes_go_proto.Reference {
	if x, ok := x.GetChoice().(*MedicinalProductInteraction_Interactant_ItemX_Reference); ok {
		return x.Reference
	}
	return nil
}

func (x *MedicinalProductInteraction_Interactant_ItemX) GetCodeableConcept() *datatypes_go_proto.CodeableConcept {
	if x, ok := x.GetChoice().(*MedicinalProductInteraction_Interactant_ItemX_CodeableConcept); ok {
		return x.CodeableConcept
	}
	return nil
}

type isMedicinalProductInteraction_Interactant_ItemX_Choice interface {
	isMedicinalProductInteraction_Interactant_ItemX_Choice()
}

type MedicinalProductInteraction_Interactant_ItemX_Reference struct {
	Reference *datatypes_go_proto.Reference `protobuf:"bytes,1,opt,name=reference,proto3,oneof"`
}

type MedicinalProductInteraction_Interactant_ItemX_CodeableConcept struct {
	CodeableConcept *datatypes_go_proto.CodeableConcept `protobuf:"bytes,2,opt,name=codeable_concept,json=codeableConcept,proto3,oneof"`
}

func (*MedicinalProductInteraction_Interactant_ItemX_Reference) isMedicinalProductInteraction_Interactant_ItemX_Choice() {
}

func (*MedicinalProductInteraction_Interactant_ItemX_CodeableConcept) isMedicinalProductInteraction_Interactant_ItemX_Choice() {
}

var File_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto protoreflect.FileDescriptor

var file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66,
	0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x34, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x64, 0x69,
	0x63, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68,
	0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x34, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x0c, 0x0a, 0x1b,
	0x4d, 0x65, 0x64, 0x69, 0x63, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72,
	0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x0e, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x5f,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x55, 0x72, 0x69, 0x52, 0x0d, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x52,
	0x75, 0x6c, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x4e, 0x61, 0x72, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x32, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x64, 0x12, 0x3c, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x4d, 0x0a, 0x12, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x6f, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e,
	0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x42, 0x35, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x10, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x69, 0x6e,
	0x61, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x0a, 0x4d,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x09, 0x53,
	0x75, 0x62, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x3d, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x5e, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x61, 0x6e, 0x74, 0x18,
	0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66,
	0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x63, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x61, 0x6e, 0x74, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x61, 0x6e, 0x74,
	0x12, 0x38, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x63, 0x65, 0x70, 0x74, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x65, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74,
	0x52, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x42, 0x0a, 0x09, 0x69, 0x6e, 0x63, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70,
	0x74, 0x52, 0x09, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0a,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72,
	0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x52, 0x0a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x1a, 0xa9, 0x04, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x61,
	0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x3c, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72,
	0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a,
	0x12, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x5e, 0x0a, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x58, 0x42, 0x06,
	0xf0, 0xd0, 0x87, 0xeb, 0x04, 0x01, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x1a, 0xff, 0x01, 0x0a,
	0x05, 0x49, 0x74, 0x65, 0x6d, 0x58, 0x12, 0x90, 0x01, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x50, 0xf2, 0xff, 0xfc, 0xc2,
	0x06, 0x10, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x0a, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x09, 0x53, 0x75, 0x62, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x15, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x09,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x63, 0x6f, 0x64,
	0x65, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69,
	0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x61, 0x62,
	0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6f, 0x64,
	0x65, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x3a, 0x06, 0xa0, 0x83,
	0x83, 0xe8, 0x06, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x3a, 0x4f,
	0xc0, 0x9f, 0xe3, 0xb6, 0x05, 0x03, 0xb2, 0xfe, 0xe4, 0x97, 0x06, 0x43, 0x68, 0x74, 0x74, 0x70,
	0x3a, 0x2f, 0x2f, 0x68, 0x6c, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a,
	0x04, 0x08, 0x07, 0x10, 0x08, 0x42, 0x8d, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x50, 0x01, 0x5a, 0x6a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68, 0x69, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x34, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x98,
	0xc6, 0xb0, 0xb5, 0x07, 0x04, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_rawDescOnce sync.Once
	file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_rawDescData = file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_rawDesc
)

func file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_rawDescGZIP() []byte {
	file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_rawDescOnce.Do(func() {
		file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_rawDescData)
	})
	return file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_rawDescData
}

var file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_goTypes = []interface{}{
	(*MedicinalProductInteraction)(nil),                   // 0: google.fhir.r4.core.MedicinalProductInteraction
	(*MedicinalProductInteraction_Interactant)(nil),       // 1: google.fhir.r4.core.MedicinalProductInteraction.Interactant
	(*MedicinalProductInteraction_Interactant_ItemX)(nil), // 2: google.fhir.r4.core.MedicinalProductInteraction.Interactant.ItemX
	(*datatypes_go_proto.Id)(nil),                         // 3: google.fhir.r4.core.Id
	(*datatypes_go_proto.Meta)(nil),                       // 4: google.fhir.r4.core.Meta
	(*datatypes_go_proto.Uri)(nil),                        // 5: google.fhir.r4.core.Uri
	(*datatypes_go_proto.Code)(nil),                       // 6: google.fhir.r4.core.Code
	(*datatypes_go_proto.Narrative)(nil),                  // 7: google.fhir.r4.core.Narrative
	(*any.Any)(nil),                                       // 8: google.protobuf.Any
	(*datatypes_go_proto.Extension)(nil),                  // 9: google.fhir.r4.core.Extension
	(*datatypes_go_proto.Reference)(nil),                  // 10: google.fhir.r4.core.Reference
	(*datatypes_go_proto.String)(nil),                     // 11: google.fhir.r4.core.String
	(*datatypes_go_proto.CodeableConcept)(nil),            // 12: google.fhir.r4.core.CodeableConcept
}
var file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_depIdxs = []int32{
	3,  // 0: google.fhir.r4.core.MedicinalProductInteraction.id:type_name -> google.fhir.r4.core.Id
	4,  // 1: google.fhir.r4.core.MedicinalProductInteraction.meta:type_name -> google.fhir.r4.core.Meta
	5,  // 2: google.fhir.r4.core.MedicinalProductInteraction.implicit_rules:type_name -> google.fhir.r4.core.Uri
	6,  // 3: google.fhir.r4.core.MedicinalProductInteraction.language:type_name -> google.fhir.r4.core.Code
	7,  // 4: google.fhir.r4.core.MedicinalProductInteraction.text:type_name -> google.fhir.r4.core.Narrative
	8,  // 5: google.fhir.r4.core.MedicinalProductInteraction.contained:type_name -> google.protobuf.Any
	9,  // 6: google.fhir.r4.core.MedicinalProductInteraction.extension:type_name -> google.fhir.r4.core.Extension
	9,  // 7: google.fhir.r4.core.MedicinalProductInteraction.modifier_extension:type_name -> google.fhir.r4.core.Extension
	10, // 8: google.fhir.r4.core.MedicinalProductInteraction.subject:type_name -> google.fhir.r4.core.Reference
	11, // 9: google.fhir.r4.core.MedicinalProductInteraction.description:type_name -> google.fhir.r4.core.String
	1,  // 10: google.fhir.r4.core.MedicinalProductInteraction.interactant:type_name -> google.fhir.r4.core.MedicinalProductInteraction.Interactant
	12, // 11: google.fhir.r4.core.MedicinalProductInteraction.type:type_name -> google.fhir.r4.core.CodeableConcept
	12, // 12: google.fhir.r4.core.MedicinalProductInteraction.effect:type_name -> google.fhir.r4.core.CodeableConcept
	12, // 13: google.fhir.r4.core.MedicinalProductInteraction.incidence:type_name -> google.fhir.r4.core.CodeableConcept
	12, // 14: google.fhir.r4.core.MedicinalProductInteraction.management:type_name -> google.fhir.r4.core.CodeableConcept
	11, // 15: google.fhir.r4.core.MedicinalProductInteraction.Interactant.id:type_name -> google.fhir.r4.core.String
	9,  // 16: google.fhir.r4.core.MedicinalProductInteraction.Interactant.extension:type_name -> google.fhir.r4.core.Extension
	9,  // 17: google.fhir.r4.core.MedicinalProductInteraction.Interactant.modifier_extension:type_name -> google.fhir.r4.core.Extension
	2,  // 18: google.fhir.r4.core.MedicinalProductInteraction.Interactant.item:type_name -> google.fhir.r4.core.MedicinalProductInteraction.Interactant.ItemX
	10, // 19: google.fhir.r4.core.MedicinalProductInteraction.Interactant.ItemX.reference:type_name -> google.fhir.r4.core.Reference
	12, // 20: google.fhir.r4.core.MedicinalProductInteraction.Interactant.ItemX.codeable_concept:type_name -> google.fhir.r4.core.CodeableConcept
	21, // [21:21] is the sub-list for method output_type
	21, // [21:21] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() {
	file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_init()
}
func file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_init() {
	if File_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MedicinalProductInteraction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MedicinalProductInteraction_Interactant); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MedicinalProductInteraction_Interactant_ItemX); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*MedicinalProductInteraction_Interactant_ItemX_Reference)(nil),
		(*MedicinalProductInteraction_Interactant_ItemX_CodeableConcept)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_goTypes,
		DependencyIndexes: file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_depIdxs,
		MessageInfos:      file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_msgTypes,
	}.Build()
	File_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto = out.File
	file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_rawDesc = nil
	file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_goTypes = nil
	file_proto_google_fhir_proto_r4_core_resources_medicinal_product_interaction_proto_depIdxs = nil
}
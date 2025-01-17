// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: v1/org_policy_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PolicyType int32

const (
	PolicyType_POLICY_TYPE_UNSPECIFIED PolicyType = 0
	PolicyType_PIPELINE_APPROVAL       PolicyType = 1
	PolicyType_BACKUP_PLAN             PolicyType = 2
	PolicyType_SQL_REVIEW              PolicyType = 3
	PolicyType_SENSITIVE_DATA          PolicyType = 4
	PolicyType_ACCESS_CONTROL          PolicyType = 5
)

// Enum value maps for PolicyType.
var (
	PolicyType_name = map[int32]string{
		0: "POLICY_TYPE_UNSPECIFIED",
		1: "PIPELINE_APPROVAL",
		2: "BACKUP_PLAN",
		3: "SQL_REVIEW",
		4: "SENSITIVE_DATA",
		5: "ACCESS_CONTROL",
	}
	PolicyType_value = map[string]int32{
		"POLICY_TYPE_UNSPECIFIED": 0,
		"PIPELINE_APPROVAL":       1,
		"BACKUP_PLAN":             2,
		"SQL_REVIEW":              3,
		"SENSITIVE_DATA":          4,
		"ACCESS_CONTROL":          5,
	}
)

func (x PolicyType) Enum() *PolicyType {
	p := new(PolicyType)
	*p = x
	return p
}

func (x PolicyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PolicyType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_org_policy_service_proto_enumTypes[0].Descriptor()
}

func (PolicyType) Type() protoreflect.EnumType {
	return &file_v1_org_policy_service_proto_enumTypes[0]
}

func (x PolicyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PolicyType.Descriptor instead.
func (PolicyType) EnumDescriptor() ([]byte, []int) {
	return file_v1_org_policy_service_proto_rawDescGZIP(), []int{0}
}

type GetPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the policy to retrieve.
	// Format: {resource type}/{resource id}/policies/{policy type}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPolicyRequest) Reset() {
	*x = GetPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_org_policy_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolicyRequest) ProtoMessage() {}

func (x *GetPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_org_policy_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolicyRequest.ProtoReflect.Descriptor instead.
func (*GetPolicyRequest) Descriptor() ([]byte, []int) {
	return file_v1_org_policy_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetPolicyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListPoliciesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent, which owns this collection of policies.
	// Format: {resource type}/{resource id}/policies/{policy type}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of policies to return. The service may return fewer than
	// this value.
	// If unspecified, at most 50 policies will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `GetPolicies` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `GetPolicies` must match
	// the call that provided the page token.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListPoliciesRequest) Reset() {
	*x = ListPoliciesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_org_policy_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPoliciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPoliciesRequest) ProtoMessage() {}

func (x *ListPoliciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_org_policy_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPoliciesRequest.ProtoReflect.Descriptor instead.
func (*ListPoliciesRequest) Descriptor() ([]byte, []int) {
	return file_v1_org_policy_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListPoliciesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListPoliciesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPoliciesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListPoliciesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The policies from the specified request.
	Policies []*Policy `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListPoliciesResponse) Reset() {
	*x = ListPoliciesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_org_policy_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPoliciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPoliciesResponse) ProtoMessage() {}

func (x *ListPoliciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_org_policy_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPoliciesResponse.ProtoReflect.Descriptor instead.
func (*ListPoliciesResponse) Descriptor() ([]byte, []int) {
	return file_v1_org_policy_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListPoliciesResponse) GetPolicies() []*Policy {
	if x != nil {
		return x.Policies
	}
	return nil
}

func (x *ListPoliciesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the policy.
	// Format: {resource type}/{resource id}/policies/{policy type}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The system-assigned, unique identifier for a resource.
	Uid               string     `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	State             State      `protobuf:"varint,3,opt,name=state,proto3,enum=bytebase.v1.State" json:"state,omitempty"`
	InheritFromParent bool       `protobuf:"varint,4,opt,name=inherit_from_parent,json=inheritFromParent,proto3" json:"inherit_from_parent,omitempty"`
	Type              PolicyType `protobuf:"varint,5,opt,name=type,proto3,enum=bytebase.v1.PolicyType" json:"type,omitempty"`
	Payload           string     `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_org_policy_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_v1_org_policy_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_v1_org_policy_service_proto_rawDescGZIP(), []int{3}
}

func (x *Policy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Policy) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Policy) GetState() State {
	if x != nil {
		return x.State
	}
	return State_STATE_UNSPECIFIED
}

func (x *Policy) GetInheritFromParent() bool {
	if x != nil {
		return x.InheritFromParent
	}
	return false
}

func (x *Policy) GetType() PolicyType {
	if x != nil {
		return x.Type
	}
	return PolicyType_POLICY_TYPE_UNSPECIFIED
}

func (x *Policy) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

var File_v1_org_policy_service_proto protoreflect.FileDescriptor

var file_v1_org_policy_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62,
	0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x6f, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x6f, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0xd5, 0x01, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04,
	0xe2, 0x41, 0x01, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x5f, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2a, 0x89, 0x01, 0x0a, 0x0a, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x4f, 0x4c,
	0x49, 0x43, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x49, 0x50, 0x45, 0x4c, 0x49,
	0x4e, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x10, 0x02, 0x12, 0x0e,
	0x0a, 0x0a, 0x53, 0x51, 0x4c, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x03, 0x12, 0x12,
	0x0a, 0x0e, 0x53, 0x45, 0x4e, 0x53, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41,
	0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x43, 0x4f, 0x4e,
	0x54, 0x52, 0x4f, 0x4c, 0x10, 0x05, 0x32, 0xed, 0x04, 0x0a, 0x10, 0x4f, 0x72, 0x67, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa7, 0x02, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1d, 0x2e, 0x62, 0x79, 0x74, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0xe5, 0x01,
	0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0xd7, 0x01, 0x5a, 0x22,
	0x12, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f,
	0x2a, 0x7d, 0x5a, 0x26, 0x12, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x5a, 0x32, 0x12, 0x30, 0x2f, 0x76,
	0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x2f, 0x2a, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x5a, 0x3e,
	0x12, 0x3c, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x2a, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x15,
	0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xae, 0x02, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xd8, 0x01, 0xda, 0x41,
	0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0xce, 0x01, 0x5a, 0x22, 0x12, 0x20, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x2a, 0x7d, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x5a, 0x26, 0x12, 0x24,
	0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x5a, 0x32, 0x12, 0x30, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x3d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x2a, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x5a, 0x3e, 0x12, 0x3c, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v1_org_policy_service_proto_rawDescOnce sync.Once
	file_v1_org_policy_service_proto_rawDescData = file_v1_org_policy_service_proto_rawDesc
)

func file_v1_org_policy_service_proto_rawDescGZIP() []byte {
	file_v1_org_policy_service_proto_rawDescOnce.Do(func() {
		file_v1_org_policy_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_org_policy_service_proto_rawDescData)
	})
	return file_v1_org_policy_service_proto_rawDescData
}

var file_v1_org_policy_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_org_policy_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_org_policy_service_proto_goTypes = []interface{}{
	(PolicyType)(0),              // 0: bytebase.v1.PolicyType
	(*GetPolicyRequest)(nil),     // 1: bytebase.v1.GetPolicyRequest
	(*ListPoliciesRequest)(nil),  // 2: bytebase.v1.ListPoliciesRequest
	(*ListPoliciesResponse)(nil), // 3: bytebase.v1.ListPoliciesResponse
	(*Policy)(nil),               // 4: bytebase.v1.Policy
	(State)(0),                   // 5: bytebase.v1.State
}
var file_v1_org_policy_service_proto_depIdxs = []int32{
	4, // 0: bytebase.v1.ListPoliciesResponse.policies:type_name -> bytebase.v1.Policy
	5, // 1: bytebase.v1.Policy.state:type_name -> bytebase.v1.State
	0, // 2: bytebase.v1.Policy.type:type_name -> bytebase.v1.PolicyType
	1, // 3: bytebase.v1.OrgPolicyService.GetPolicy:input_type -> bytebase.v1.GetPolicyRequest
	2, // 4: bytebase.v1.OrgPolicyService.ListPolicies:input_type -> bytebase.v1.ListPoliciesRequest
	4, // 5: bytebase.v1.OrgPolicyService.GetPolicy:output_type -> bytebase.v1.Policy
	3, // 6: bytebase.v1.OrgPolicyService.ListPolicies:output_type -> bytebase.v1.ListPoliciesResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_org_policy_service_proto_init() }
func file_v1_org_policy_service_proto_init() {
	if File_v1_org_policy_service_proto != nil {
		return
	}
	file_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_org_policy_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPolicyRequest); i {
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
		file_v1_org_policy_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPoliciesRequest); i {
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
		file_v1_org_policy_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPoliciesResponse); i {
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
		file_v1_org_policy_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_org_policy_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_org_policy_service_proto_goTypes,
		DependencyIndexes: file_v1_org_policy_service_proto_depIdxs,
		EnumInfos:         file_v1_org_policy_service_proto_enumTypes,
		MessageInfos:      file_v1_org_policy_service_proto_msgTypes,
	}.Build()
	File_v1_org_policy_service_proto = out.File
	file_v1_org_policy_service_proto_rawDesc = nil
	file_v1_org_policy_service_proto_goTypes = nil
	file_v1_org_policy_service_proto_depIdxs = nil
}

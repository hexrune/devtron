// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.9.1
// source: applist.proto

package client

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ClusterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiServerUrl string `protobuf:"bytes,1,opt,name=apiServerUrl,proto3" json:"apiServerUrl,omitempty"`
	Token        string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ClusterId    int32  `protobuf:"varint,3,opt,name=clusterId,proto3" json:"clusterId,omitempty"`
	ClusterName  string `protobuf:"bytes,4,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
}

func (x *ClusterConfig) Reset() {
	*x = ClusterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_applist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterConfig) ProtoMessage() {}

func (x *ClusterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_applist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterConfig.ProtoReflect.Descriptor instead.
func (*ClusterConfig) Descriptor() ([]byte, []int) {
	return file_applist_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterConfig) GetApiServerUrl() string {
	if x != nil {
		return x.ApiServerUrl
	}
	return ""
}

func (x *ClusterConfig) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ClusterConfig) GetClusterId() int32 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *ClusterConfig) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

type AppListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters []*ClusterConfig `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *AppListRequest) Reset() {
	*x = AppListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_applist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppListRequest) ProtoMessage() {}

func (x *AppListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_applist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppListRequest.ProtoReflect.Descriptor instead.
func (*AppListRequest) Descriptor() ([]byte, []int) {
	return file_applist_proto_rawDescGZIP(), []int{1}
}

func (x *AppListRequest) GetClusters() []*ClusterConfig {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type DeployedAppList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeployedAppDetail []*DeployedAppDetail `protobuf:"bytes,1,rep,name=DeployedAppDetail,proto3" json:"DeployedAppDetail,omitempty"`
	ClusterId         int32                `protobuf:"varint,2,opt,name=clusterId,proto3" json:"clusterId,omitempty"`
	ErrorMsg          string               `protobuf:"bytes,3,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
	Errored           bool                 `protobuf:"varint,4,opt,name=errored,proto3" json:"errored,omitempty"`
}

func (x *DeployedAppList) Reset() {
	*x = DeployedAppList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_applist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployedAppList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployedAppList) ProtoMessage() {}

func (x *DeployedAppList) ProtoReflect() protoreflect.Message {
	mi := &file_applist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployedAppList.ProtoReflect.Descriptor instead.
func (*DeployedAppList) Descriptor() ([]byte, []int) {
	return file_applist_proto_rawDescGZIP(), []int{2}
}

func (x *DeployedAppList) GetDeployedAppDetail() []*DeployedAppDetail {
	if x != nil {
		return x.DeployedAppDetail
	}
	return nil
}

func (x *DeployedAppList) GetClusterId() int32 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *DeployedAppList) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

func (x *DeployedAppList) GetErrored() bool {
	if x != nil {
		return x.Errored
	}
	return false
}

type DeployedAppDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId        string               `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	AppName      string               `protobuf:"bytes,2,opt,name=appName,proto3" json:"appName,omitempty"`
	ChartName    string               `protobuf:"bytes,3,opt,name=chartName,proto3" json:"chartName,omitempty"`
	ChartAvatar  string               `protobuf:"bytes,4,opt,name=chartAvatar,proto3" json:"chartAvatar,omitempty"`
	Environment  *EnvironmentDetails  `protobuf:"bytes,5,opt,name=environment,proto3" json:"environment,omitempty"`
	LastDeployed *timestamp.Timestamp `protobuf:"bytes,6,opt,name=LastDeployed,proto3" json:"LastDeployed,omitempty"`
}

func (x *DeployedAppDetail) Reset() {
	*x = DeployedAppDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_applist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployedAppDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployedAppDetail) ProtoMessage() {}

func (x *DeployedAppDetail) ProtoReflect() protoreflect.Message {
	mi := &file_applist_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployedAppDetail.ProtoReflect.Descriptor instead.
func (*DeployedAppDetail) Descriptor() ([]byte, []int) {
	return file_applist_proto_rawDescGZIP(), []int{3}
}

func (x *DeployedAppDetail) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *DeployedAppDetail) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *DeployedAppDetail) GetChartName() string {
	if x != nil {
		return x.ChartName
	}
	return ""
}

func (x *DeployedAppDetail) GetChartAvatar() string {
	if x != nil {
		return x.ChartAvatar
	}
	return ""
}

func (x *DeployedAppDetail) GetEnvironment() *EnvironmentDetails {
	if x != nil {
		return x.Environment
	}
	return nil
}

func (x *DeployedAppDetail) GetLastDeployed() *timestamp.Timestamp {
	if x != nil {
		return x.LastDeployed
	}
	return nil
}

type EnvironmentDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string `protobuf:"bytes,1,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	ClusterId   int32  `protobuf:"varint,2,opt,name=clusterId,proto3" json:"clusterId,omitempty"`
	Namespace   string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *EnvironmentDetails) Reset() {
	*x = EnvironmentDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_applist_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentDetails) ProtoMessage() {}

func (x *EnvironmentDetails) ProtoReflect() protoreflect.Message {
	mi := &file_applist_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentDetails.ProtoReflect.Descriptor instead.
func (*EnvironmentDetails) Descriptor() ([]byte, []int) {
	return file_applist_proto_rawDescGZIP(), []int{4}
}

func (x *EnvironmentDetails) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *EnvironmentDetails) GetClusterId() int32 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *EnvironmentDetails) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_applist_proto protoreflect.FileDescriptor

var file_applist_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x89, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x70, 0x69, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x0e,
	0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x0f, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x40,
	0x0a, 0x11, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x41, 0x70, 0x70, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x64, 0x41, 0x70, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x11, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x41, 0x70, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x65, 0x64, 0x22, 0xfa, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x64, 0x41, 0x70, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68,
	0x61, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x72,
	0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x68, 0x61, 0x72, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x35, 0x0a, 0x0b, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x64, 0x22, 0x72, 0x0a, 0x12, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x32, 0x4f, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x0f, 0x2e, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x41, 0x70, 0x70, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x72, 0x6f, 0x6e, 0x2d, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x68, 0x65, 0x6c, 0x6d, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x62, 0x65, 0x61, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_applist_proto_rawDescOnce sync.Once
	file_applist_proto_rawDescData = file_applist_proto_rawDesc
)

func file_applist_proto_rawDescGZIP() []byte {
	file_applist_proto_rawDescOnce.Do(func() {
		file_applist_proto_rawDescData = protoimpl.X.CompressGZIP(file_applist_proto_rawDescData)
	})
	return file_applist_proto_rawDescData
}

var file_applist_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_applist_proto_goTypes = []interface{}{
	(*ClusterConfig)(nil),       // 0: ClusterConfig
	(*AppListRequest)(nil),      // 1: AppListRequest
	(*DeployedAppList)(nil),     // 2: DeployedAppList
	(*DeployedAppDetail)(nil),   // 3: DeployedAppDetail
	(*EnvironmentDetails)(nil),  // 4: EnvironmentDetails
	(*timestamp.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_applist_proto_depIdxs = []int32{
	0, // 0: AppListRequest.clusters:type_name -> ClusterConfig
	3, // 1: DeployedAppList.DeployedAppDetail:type_name -> DeployedAppDetail
	4, // 2: DeployedAppDetail.environment:type_name -> EnvironmentDetails
	5, // 3: DeployedAppDetail.LastDeployed:type_name -> google.protobuf.Timestamp
	1, // 4: ApplicationService.ListApplications:input_type -> AppListRequest
	2, // 5: ApplicationService.ListApplications:output_type -> DeployedAppList
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_applist_proto_init() }
func file_applist_proto_init() {
	if File_applist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_applist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterConfig); i {
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
		file_applist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppListRequest); i {
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
		file_applist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployedAppList); i {
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
		file_applist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployedAppDetail); i {
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
		file_applist_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentDetails); i {
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
			RawDescriptor: file_applist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_applist_proto_goTypes,
		DependencyIndexes: file_applist_proto_depIdxs,
		MessageInfos:      file_applist_proto_msgTypes,
	}.Build()
	File_applist_proto = out.File
	file_applist_proto_rawDesc = nil
	file_applist_proto_goTypes = nil
	file_applist_proto_depIdxs = nil
}

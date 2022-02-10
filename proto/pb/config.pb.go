// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/config.proto

package pb

import (
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

type Target int32

const (
	Target_k8sAllNodes     Target = 0
	Target_k8sControlPlane Target = 1
	Target_k8sComputePlane Target = 2
)

// Enum value maps for Target.
var (
	Target_name = map[int32]string{
		0: "k8sAllNodes",
		1: "k8sControlPlane",
		2: "k8sComputePlane",
	}
	Target_value = map[string]int32{
		"k8sAllNodes":     0,
		"k8sControlPlane": 1,
		"k8sComputePlane": 2,
	}
)

func (x Target) Enum() *Target {
	p := new(Target)
	*p = x
	return p
}

func (x Target) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Target) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_config_proto_enumTypes[0].Descriptor()
}

func (Target) Type() protoreflect.EnumType {
	return &file_proto_config_proto_enumTypes[0]
}

func (x Target) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Target.Descriptor instead.
func (Target) EnumDescriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{0}
}

type NodeType int32

const (
	NodeType_worker      NodeType = 0
	NodeType_master      NodeType = 1
	NodeType_apiEndpoint NodeType = 2
)

// Enum value maps for NodeType.
var (
	NodeType_name = map[int32]string{
		0: "worker",
		1: "master",
		2: "apiEndpoint",
	}
	NodeType_value = map[string]int32{
		"worker":      0,
		"master":      1,
		"apiEndpoint": 2,
	}
)

func (x NodeType) Enum() *NodeType {
	p := new(NodeType)
	*p = x
	return p
}

func (x NodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_config_proto_enumTypes[1].Descriptor()
}

func (NodeType) Type() protoreflect.EnumType {
	return &file_proto_config_proto_enumTypes[1]
}

func (x NodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeType.Descriptor instead.
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{1}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Manifest     string   `protobuf:"bytes,3,opt,name=manifest,proto3" json:"manifest,omitempty"`
	DesiredState *Project `protobuf:"bytes,4,opt,name=desiredState,proto3" json:"desiredState,omitempty"`
	CurrentState *Project `protobuf:"bytes,5,opt,name=currentState,proto3" json:"currentState,omitempty"`
	MsChecksum   []byte   `protobuf:"bytes,6,opt,name=msChecksum,proto3" json:"msChecksum,omitempty"`
	DsChecksum   []byte   `protobuf:"bytes,7,opt,name=dsChecksum,proto3" json:"dsChecksum,omitempty"`
	CsChecksum   []byte   `protobuf:"bytes,8,opt,name=csChecksum,proto3" json:"csChecksum,omitempty"`
	BuilderTTL   int32    `protobuf:"varint,9,opt,name=builderTTL,proto3" json:"builderTTL,omitempty"`
	SchedulerTTL int32    `protobuf:"varint,10,opt,name=schedulerTTL,proto3" json:"schedulerTTL,omitempty"`
	ErrorMessage string   `protobuf:"bytes,11,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Config) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Config) GetManifest() string {
	if x != nil {
		return x.Manifest
	}
	return ""
}

func (x *Config) GetDesiredState() *Project {
	if x != nil {
		return x.DesiredState
	}
	return nil
}

func (x *Config) GetCurrentState() *Project {
	if x != nil {
		return x.CurrentState
	}
	return nil
}

func (x *Config) GetMsChecksum() []byte {
	if x != nil {
		return x.MsChecksum
	}
	return nil
}

func (x *Config) GetDsChecksum() []byte {
	if x != nil {
		return x.DsChecksum
	}
	return nil
}

func (x *Config) GetCsChecksum() []byte {
	if x != nil {
		return x.CsChecksum
	}
	return nil
}

func (x *Config) GetBuilderTTL() int32 {
	if x != nil {
		return x.BuilderTTL
	}
	return 0
}

func (x *Config) GetSchedulerTTL() int32 {
	if x != nil {
		return x.SchedulerTTL
	}
	return 0
}

func (x *Config) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Clusters             []*K8Scluster `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	LoadBalancerClusters []*LBcluster  `protobuf:"bytes,3,rep,name=loadBalancerClusters,proto3" json:"loadBalancerClusters,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{1}
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetClusters() []*K8Scluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *Project) GetLoadBalancerClusters() []*LBcluster {
	if x != nil {
		return x.LoadBalancerClusters
	}
	return nil
}

type K8Scluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterInfo *ClusterInfo `protobuf:"bytes,1,opt,name=clusterInfo,proto3" json:"clusterInfo,omitempty"`
	Network     string       `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	Kubeconfig  string       `protobuf:"bytes,3,opt,name=kubeconfig,proto3" json:"kubeconfig,omitempty"`
	Kubernetes  string       `protobuf:"bytes,4,opt,name=kubernetes,proto3" json:"kubernetes,omitempty"`
}

func (x *K8Scluster) Reset() {
	*x = K8Scluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8Scluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8Scluster) ProtoMessage() {}

func (x *K8Scluster) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8Scluster.ProtoReflect.Descriptor instead.
func (*K8Scluster) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{2}
}

func (x *K8Scluster) GetClusterInfo() *ClusterInfo {
	if x != nil {
		return x.ClusterInfo
	}
	return nil
}

func (x *K8Scluster) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *K8Scluster) GetKubeconfig() string {
	if x != nil {
		return x.Kubeconfig
	}
	return ""
}

func (x *K8Scluster) GetKubernetes() string {
	if x != nil {
		return x.Kubernetes
	}
	return ""
}

type LBcluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterInfo *ClusterInfo `protobuf:"bytes,1,opt,name=clusterInfo,proto3" json:"clusterInfo,omitempty"`
	Roles       []*Role      `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	Dns         *DNS         `protobuf:"bytes,3,opt,name=dns,proto3" json:"dns,omitempty"`
	TargetedK8S string       `protobuf:"bytes,4,opt,name=targetedK8s,proto3" json:"targetedK8s,omitempty"`
}

func (x *LBcluster) Reset() {
	*x = LBcluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LBcluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LBcluster) ProtoMessage() {}

func (x *LBcluster) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LBcluster.ProtoReflect.Descriptor instead.
func (*LBcluster) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{3}
}

func (x *LBcluster) GetClusterInfo() *ClusterInfo {
	if x != nil {
		return x.ClusterInfo
	}
	return nil
}

func (x *LBcluster) GetRoles() []*Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *LBcluster) GetDns() *DNS {
	if x != nil {
		return x.Dns
	}
	return nil
}

func (x *LBcluster) GetTargetedK8S() string {
	if x != nil {
		return x.TargetedK8S
	}
	return ""
}

type ClusterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hash       string      `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	PublicKey  string      `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PrivateKey string      `protobuf:"bytes,4,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	NodePools  []*NodePool `protobuf:"bytes,5,rep,name=nodePools,proto3" json:"nodePools,omitempty"`
}

func (x *ClusterInfo) Reset() {
	*x = ClusterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterInfo) ProtoMessage() {}

func (x *ClusterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterInfo.ProtoReflect.Descriptor instead.
func (*ClusterInfo) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{4}
}

func (x *ClusterInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClusterInfo) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *ClusterInfo) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *ClusterInfo) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *ClusterInfo) GetNodePools() []*NodePool {
	if x != nil {
		return x.NodePools
	}
	return nil
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Protocol   string `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Port       int32  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	TargetPort int32  `protobuf:"varint,4,opt,name=targetPort,proto3" json:"targetPort,omitempty"`
	Target     Target `protobuf:"varint,5,opt,name=target,proto3,enum=platform.Target" json:"target,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{5}
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Role) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Role) GetTargetPort() int32 {
	if x != nil {
		return x.TargetPort
	}
	return 0
}

func (x *Role) GetTarget() Target {
	if x != nil {
		return x.Target
	}
	return Target_k8sAllNodes
}

type DNS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname  string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Providers []string `protobuf:"bytes,2,rep,name=providers,proto3" json:"providers,omitempty"`
}

func (x *DNS) Reset() {
	*x = DNS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNS) ProtoMessage() {}

func (x *DNS) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNS.ProtoReflect.Descriptor instead.
func (*DNS) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{6}
}

func (x *DNS) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *DNS) GetProviders() []string {
	if x != nil {
		return x.Providers
	}
	return nil
}

type NodePool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Region     string    `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	ServerType string    `protobuf:"bytes,3,opt,name=serverType,proto3" json:"serverType,omitempty"`
	Image      string    `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	DiskSize   uint32    `protobuf:"varint,5,opt,name=diskSize,proto3" json:"diskSize,omitempty"`
	Zone       string    `protobuf:"bytes,6,opt,name=zone,proto3" json:"zone,omitempty"`
	Count      uint32    `protobuf:"varint,9,opt,name=count,proto3" json:"count,omitempty"`
	Nodes      []*Node   `protobuf:"bytes,10,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Provider   *Provider `protobuf:"bytes,11,opt,name=provider,proto3" json:"provider,omitempty"`
	IsControl  bool      `protobuf:"varint,12,opt,name=isControl,proto3" json:"isControl,omitempty"`
}

func (x *NodePool) Reset() {
	*x = NodePool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodePool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodePool) ProtoMessage() {}

func (x *NodePool) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodePool.ProtoReflect.Descriptor instead.
func (*NodePool) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{7}
}

func (x *NodePool) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodePool) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *NodePool) GetServerType() string {
	if x != nil {
		return x.ServerType
	}
	return ""
}

func (x *NodePool) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *NodePool) GetDiskSize() uint32 {
	if x != nil {
		return x.DiskSize
	}
	return 0
}

func (x *NodePool) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *NodePool) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *NodePool) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *NodePool) GetProvider() *Provider {
	if x != nil {
		return x.Provider
	}
	return nil
}

func (x *NodePool) GetIsControl() bool {
	if x != nil {
		return x.IsControl
	}
	return false
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Private  string   `protobuf:"bytes,2,opt,name=private,proto3" json:"private,omitempty"`
	Public   string   `protobuf:"bytes,3,opt,name=public,proto3" json:"public,omitempty"`
	NodeType NodeType `protobuf:"varint,4,opt,name=nodeType,proto3,enum=platform.NodeType" json:"nodeType,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{8}
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Node) GetPrivate() string {
	if x != nil {
		return x.Private
	}
	return ""
}

func (x *Node) GetPublic() string {
	if x != nil {
		return x.Public
	}
	return ""
}

func (x *Node) GetNodeType() NodeType {
	if x != nil {
		return x.NodeType
	}
	return NodeType_worker
}

type Provider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Credentials string `protobuf:"bytes,2,opt,name=credentials,proto3" json:"credentials,omitempty"`
}

func (x *Provider) Reset() {
	*x = Provider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provider) ProtoMessage() {}

func (x *Provider) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provider.ProtoReflect.Descriptor instead.
func (*Provider) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{9}
}

func (x *Provider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Provider) GetCredentials() string {
	if x != nil {
		return x.Credentials
	}
	return ""
}

var File_proto_config_proto protoreflect.FileDescriptor

var file_proto_config_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0xfe,
	0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0c, 0x64, 0x65, 0x73,
	0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x35, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x73, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6d, 0x73, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x73, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x64, 0x73, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x73, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x73, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x54, 0x54, 0x4c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x54, 0x54, 0x4c, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x54, 0x54, 0x4c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x54, 0x54, 0x4c, 0x12, 0x22, 0x0a, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x98, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x30, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x4b, 0x38, 0x73,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x47, 0x0a, 0x14, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x4c, 0x42, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x14, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x0a, 0x4b,
	0x38, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0b, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1e, 0x0a, 0x0a,
	0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e, 0x0a, 0x0a,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x22, 0xad, 0x01, 0x0a,
	0x09, 0x4c, 0x42, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0b, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x03, 0x64, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x44, 0x4e, 0x53, 0x52, 0x03, 0x64, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x65, 0x64, 0x4b, 0x38, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x65, 0x64, 0x4b, 0x38, 0x73, 0x22, 0xa7, 0x01, 0x0a,
	0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x6f, 0x6c,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x09, 0x6e, 0x6f, 0x64,
	0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x6f, 0x72,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x3f, 0x0a,
	0x03, 0x44, 0x4e, 0x53, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x22, 0xa6,
	0x02, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x22, 0x7c, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x2e, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6e, 0x6f, 0x64,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x40, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2a, 0x43, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x6b, 0x38, 0x73, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x6b, 0x38, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x50, 0x6c, 0x61, 0x6e, 0x65, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x6b, 0x38, 0x73, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x10, 0x02, 0x2a, 0x33, 0x0a, 0x08,
	0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x10,
	0x02, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_config_proto_rawDescOnce sync.Once
	file_proto_config_proto_rawDescData = file_proto_config_proto_rawDesc
)

func file_proto_config_proto_rawDescGZIP() []byte {
	file_proto_config_proto_rawDescOnce.Do(func() {
		file_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_config_proto_rawDescData)
	})
	return file_proto_config_proto_rawDescData
}

var file_proto_config_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_config_proto_goTypes = []interface{}{
	(Target)(0),         // 0: platform.Target
	(NodeType)(0),       // 1: platform.NodeType
	(*Config)(nil),      // 2: platform.Config
	(*Project)(nil),     // 3: platform.Project
	(*K8Scluster)(nil),  // 4: platform.K8scluster
	(*LBcluster)(nil),   // 5: platform.LBcluster
	(*ClusterInfo)(nil), // 6: platform.ClusterInfo
	(*Role)(nil),        // 7: platform.Role
	(*DNS)(nil),         // 8: platform.DNS
	(*NodePool)(nil),    // 9: platform.NodePool
	(*Node)(nil),        // 10: platform.Node
	(*Provider)(nil),    // 11: platform.Provider
}
var file_proto_config_proto_depIdxs = []int32{
	3,  // 0: platform.Config.desiredState:type_name -> platform.Project
	3,  // 1: platform.Config.currentState:type_name -> platform.Project
	4,  // 2: platform.Project.clusters:type_name -> platform.K8scluster
	5,  // 3: platform.Project.loadBalancerClusters:type_name -> platform.LBcluster
	6,  // 4: platform.K8scluster.clusterInfo:type_name -> platform.ClusterInfo
	6,  // 5: platform.LBcluster.clusterInfo:type_name -> platform.ClusterInfo
	7,  // 6: platform.LBcluster.roles:type_name -> platform.Role
	8,  // 7: platform.LBcluster.dns:type_name -> platform.DNS
	9,  // 8: platform.ClusterInfo.nodePools:type_name -> platform.NodePool
	0,  // 9: platform.Role.target:type_name -> platform.Target
	10, // 10: platform.NodePool.nodes:type_name -> platform.Node
	11, // 11: platform.NodePool.provider:type_name -> platform.Provider
	1,  // 12: platform.Node.nodeType:type_name -> platform.NodeType
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_proto_config_proto_init() }
func file_proto_config_proto_init() {
	if File_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_proto_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_proto_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*K8Scluster); i {
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
		file_proto_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LBcluster); i {
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
		file_proto_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterInfo); i {
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
		file_proto_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_proto_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNS); i {
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
		file_proto_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodePool); i {
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
		file_proto_config_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_proto_config_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provider); i {
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
			RawDescriptor: file_proto_config_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_config_proto_goTypes,
		DependencyIndexes: file_proto_config_proto_depIdxs,
		EnumInfos:         file_proto_config_proto_enumTypes,
		MessageInfos:      file_proto_config_proto_msgTypes,
	}.Build()
	File_proto_config_proto = out.File
	file_proto_config_proto_rawDesc = nil
	file_proto_config_proto_goTypes = nil
	file_proto_config_proto_depIdxs = nil
}

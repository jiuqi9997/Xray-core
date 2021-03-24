// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: app/dns/config.proto

package dns

import (
	proto "github.com/golang/protobuf/proto"
	domain "github.com/xtls/xray-core/common/matcher/domain"
	geoip "github.com/xtls/xray-core/common/matcher/geoip"
	net "github.com/xtls/xray-core/common/net"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type QueryStrategy int32

const (
	QueryStrategy_USE_IP  QueryStrategy = 0
	QueryStrategy_USE_IP4 QueryStrategy = 1
	QueryStrategy_USE_IP6 QueryStrategy = 2
)

// Enum value maps for QueryStrategy.
var (
	QueryStrategy_name = map[int32]string{
		0: "USE_IP",
		1: "USE_IP4",
		2: "USE_IP6",
	}
	QueryStrategy_value = map[string]int32{
		"USE_IP":  0,
		"USE_IP4": 1,
		"USE_IP6": 2,
	}
)

func (x QueryStrategy) Enum() *QueryStrategy {
	p := new(QueryStrategy)
	*p = x
	return p
}

func (x QueryStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_app_dns_config_proto_enumTypes[0].Descriptor()
}

func (QueryStrategy) Type() protoreflect.EnumType {
	return &file_app_dns_config_proto_enumTypes[0]
}

func (x QueryStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryStrategy.Descriptor instead.
func (QueryStrategy) EnumDescriptor() ([]byte, []int) {
	return file_app_dns_config_proto_rawDescGZIP(), []int{0}
}

type NameServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address           *net.Endpoint              `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ClientIp          []byte                     `protobuf:"bytes,5,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	PrioritizedDomain []*domain.Domain           `protobuf:"bytes,2,rep,name=prioritized_domain,json=prioritizedDomain,proto3" json:"prioritized_domain,omitempty"`
	Geoip             []*geoip.GeoIP             `protobuf:"bytes,3,rep,name=geoip,proto3" json:"geoip,omitempty"`
	OriginalRules     []*NameServer_OriginalRule `protobuf:"bytes,4,rep,name=original_rules,json=originalRules,proto3" json:"original_rules,omitempty"`
}

func (x *NameServer) Reset() {
	*x = NameServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_dns_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameServer) ProtoMessage() {}

func (x *NameServer) ProtoReflect() protoreflect.Message {
	mi := &file_app_dns_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameServer.ProtoReflect.Descriptor instead.
func (*NameServer) Descriptor() ([]byte, []int) {
	return file_app_dns_config_proto_rawDescGZIP(), []int{0}
}

func (x *NameServer) GetAddress() *net.Endpoint {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *NameServer) GetClientIp() []byte {
	if x != nil {
		return x.ClientIp
	}
	return nil
}

func (x *NameServer) GetPrioritizedDomain() []*domain.Domain {
	if x != nil {
		return x.PrioritizedDomain
	}
	return nil
}

func (x *NameServer) GetGeoip() []*geoip.GeoIP {
	if x != nil {
		return x.Geoip
	}
	return nil
}

func (x *NameServer) GetOriginalRules() []*NameServer_OriginalRule {
	if x != nil {
		return x.OriginalRules
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Nameservers used by this DNS. Only traditional UDP servers are support at
	// the moment. A special value 'localhost' as a domain address can be set to
	// use DNS on local system.
	//
	// Deprecated: Do not use.
	NameServers []*net.Endpoint `protobuf:"bytes,1,rep,name=NameServers,proto3" json:"NameServers,omitempty"`
	// NameServer list used by this DNS client.
	NameServer []*NameServer `protobuf:"bytes,5,rep,name=name_server,json=nameServer,proto3" json:"name_server,omitempty"`
	// Static hosts. Domain to IP.
	// Deprecated. Use static_hosts.
	//
	// Deprecated: Do not use.
	Hosts map[string]*net.IPOrDomain `protobuf:"bytes,2,rep,name=Hosts,proto3" json:"Hosts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Client IP for EDNS client subnet. Must be 4 bytes (IPv4) or 16 bytes
	// (IPv6).
	ClientIp    []byte                `protobuf:"bytes,3,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	StaticHosts []*Config_HostMapping `protobuf:"bytes,4,rep,name=static_hosts,json=staticHosts,proto3" json:"static_hosts,omitempty"`
	// Tag is the inbound tag of DNS client.
	Tag string `protobuf:"bytes,6,opt,name=tag,proto3" json:"tag,omitempty"`
	// DisableCache Disable DNS cache
	DisableCache  bool          `protobuf:"varint,8,opt,name=disableCache,proto3" json:"disableCache,omitempty"`
	QueryStrategy QueryStrategy `protobuf:"varint,9,opt,name=query_strategy,json=queryStrategy,proto3,enum=xray.app.dns.QueryStrategy" json:"query_strategy,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_dns_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_dns_config_proto_msgTypes[1]
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
	return file_app_dns_config_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Do not use.
func (x *Config) GetNameServers() []*net.Endpoint {
	if x != nil {
		return x.NameServers
	}
	return nil
}

func (x *Config) GetNameServer() []*NameServer {
	if x != nil {
		return x.NameServer
	}
	return nil
}

// Deprecated: Do not use.
func (x *Config) GetHosts() map[string]*net.IPOrDomain {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *Config) GetClientIp() []byte {
	if x != nil {
		return x.ClientIp
	}
	return nil
}

func (x *Config) GetStaticHosts() []*Config_HostMapping {
	if x != nil {
		return x.StaticHosts
	}
	return nil
}

func (x *Config) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Config) GetDisableCache() bool {
	if x != nil {
		return x.DisableCache
	}
	return false
}

func (x *Config) GetQueryStrategy() QueryStrategy {
	if x != nil {
		return x.QueryStrategy
	}
	return QueryStrategy_USE_IP
}

type NameServer_OriginalRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rule string `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
	Size uint32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *NameServer_OriginalRule) Reset() {
	*x = NameServer_OriginalRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_dns_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameServer_OriginalRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameServer_OriginalRule) ProtoMessage() {}

func (x *NameServer_OriginalRule) ProtoReflect() protoreflect.Message {
	mi := &file_app_dns_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameServer_OriginalRule.ProtoReflect.Descriptor instead.
func (*NameServer_OriginalRule) Descriptor() ([]byte, []int) {
	return file_app_dns_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *NameServer_OriginalRule) GetRule() string {
	if x != nil {
		return x.Rule
	}
	return ""
}

func (x *NameServer_OriginalRule) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type Config_HostMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   domain.MatchingType `protobuf:"varint,1,opt,name=type,proto3,enum=xray.common.matcher.domain.MatchingType" json:"type,omitempty"`
	Domain string              `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Ip     [][]byte            `protobuf:"bytes,3,rep,name=ip,proto3" json:"ip,omitempty"`
	// ProxiedDomain indicates the mapped domain has the same IP address on this
	// domain. Xray will use this domain for IP queries. This field is only
	// effective if ip is empty.
	ProxiedDomain string `protobuf:"bytes,4,opt,name=proxied_domain,json=proxiedDomain,proto3" json:"proxied_domain,omitempty"`
}

func (x *Config_HostMapping) Reset() {
	*x = Config_HostMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_dns_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config_HostMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config_HostMapping) ProtoMessage() {}

func (x *Config_HostMapping) ProtoReflect() protoreflect.Message {
	mi := &file_app_dns_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config_HostMapping.ProtoReflect.Descriptor instead.
func (*Config_HostMapping) Descriptor() ([]byte, []int) {
	return file_app_dns_config_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Config_HostMapping) GetType() domain.MatchingType {
	if x != nil {
		return x.Type
	}
	return domain.MatchingType_Full
}

func (x *Config_HostMapping) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Config_HostMapping) GetIp() [][]byte {
	if x != nil {
		return x.Ip
	}
	return nil
}

func (x *Config_HostMapping) GetProxiedDomain() string {
	if x != nil {
		return x.ProxiedDomain
	}
	return ""
}

var File_app_dns_config_proto protoreflect.FileDescriptor

var file_app_dns_config_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x64, 0x6e, 0x73, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x74,
	0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x6f, 0x69, 0x70, 0x2f, 0x67, 0x65, 0x6f, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xef, 0x02, 0x0a, 0x0a, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x33, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x70, 0x12, 0x51, 0x0a, 0x12, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x7a,
	0x65, 0x64, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x52, 0x11, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x64,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x36, 0x0a, 0x05, 0x67, 0x65, 0x6f, 0x69, 0x70, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x67, 0x65, 0x6f, 0x69,
	0x70, 0x2e, 0x47, 0x65, 0x6f, 0x49, 0x50, 0x52, 0x05, 0x67, 0x65, 0x6f, 0x69, 0x70, 0x12, 0x4c,
	0x0a, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0d, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0x36, 0x0a, 0x0c,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x22, 0x95, 0x05, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x3f, 0x0a, 0x0b, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x0b, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x12, 0x39, 0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x05, 0x48,
	0x6f, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x78, 0x72, 0x61,
	0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x05, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x70, 0x12, 0x43, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x68, 0x6f,
	0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x78, 0x72, 0x61, 0x79,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x42,
	0x0a, 0x0e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x52, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x1a, 0x55, 0x0a, 0x0a, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x6e, 0x65, 0x74, 0x2e, 0x49, 0x50, 0x4f, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x9a, 0x01, 0x0a, 0x0b, 0x48, 0x6f,
	0x73, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x70, 0x12,
	0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x65, 0x64, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x65, 0x64,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4a, 0x04, 0x08, 0x07, 0x10, 0x08, 0x2a, 0x35, 0x0a, 0x0d,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x0a, 0x0a,
	0x06, 0x55, 0x53, 0x45, 0x5f, 0x49, 0x50, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x53, 0x45,
	0x5f, 0x49, 0x50, 0x34, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x53, 0x45, 0x5f, 0x49, 0x50,
	0x36, 0x10, 0x02, 0x42, 0x46, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x64, 0x6e, 0x73, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x74, 0x6c, 0x73, 0x2f, 0x78, 0x72, 0x61, 0x79, 0x2d,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x6e, 0x73, 0xaa, 0x02, 0x0c, 0x58,
	0x72, 0x61, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x44, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_app_dns_config_proto_rawDescOnce sync.Once
	file_app_dns_config_proto_rawDescData = file_app_dns_config_proto_rawDesc
)

func file_app_dns_config_proto_rawDescGZIP() []byte {
	file_app_dns_config_proto_rawDescOnce.Do(func() {
		file_app_dns_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_dns_config_proto_rawDescData)
	})
	return file_app_dns_config_proto_rawDescData
}

var file_app_dns_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_app_dns_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_app_dns_config_proto_goTypes = []interface{}{
	(QueryStrategy)(0),              // 0: xray.app.dns.QueryStrategy
	(*NameServer)(nil),              // 1: xray.app.dns.NameServer
	(*Config)(nil),                  // 2: xray.app.dns.Config
	(*NameServer_OriginalRule)(nil), // 3: xray.app.dns.NameServer.OriginalRule
	nil,                             // 4: xray.app.dns.Config.HostsEntry
	(*Config_HostMapping)(nil),      // 5: xray.app.dns.Config.HostMapping
	(*net.Endpoint)(nil),            // 6: xray.common.net.Endpoint
	(*domain.Domain)(nil),           // 7: xray.common.matcher.domain.Domain
	(*geoip.GeoIP)(nil),             // 8: xray.common.matcher.geoip.GeoIP
	(*net.IPOrDomain)(nil),          // 9: xray.common.net.IPOrDomain
	(domain.MatchingType)(0),        // 10: xray.common.matcher.domain.MatchingType
}
var file_app_dns_config_proto_depIdxs = []int32{
	6,  // 0: xray.app.dns.NameServer.address:type_name -> xray.common.net.Endpoint
	7,  // 1: xray.app.dns.NameServer.prioritized_domain:type_name -> xray.common.matcher.domain.Domain
	8,  // 2: xray.app.dns.NameServer.geoip:type_name -> xray.common.matcher.geoip.GeoIP
	3,  // 3: xray.app.dns.NameServer.original_rules:type_name -> xray.app.dns.NameServer.OriginalRule
	6,  // 4: xray.app.dns.Config.NameServers:type_name -> xray.common.net.Endpoint
	1,  // 5: xray.app.dns.Config.name_server:type_name -> xray.app.dns.NameServer
	4,  // 6: xray.app.dns.Config.Hosts:type_name -> xray.app.dns.Config.HostsEntry
	5,  // 7: xray.app.dns.Config.static_hosts:type_name -> xray.app.dns.Config.HostMapping
	0,  // 8: xray.app.dns.Config.query_strategy:type_name -> xray.app.dns.QueryStrategy
	9,  // 9: xray.app.dns.Config.HostsEntry.value:type_name -> xray.common.net.IPOrDomain
	10, // 10: xray.app.dns.Config.HostMapping.type:type_name -> xray.common.matcher.domain.MatchingType
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_app_dns_config_proto_init() }
func file_app_dns_config_proto_init() {
	if File_app_dns_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_dns_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameServer); i {
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
		file_app_dns_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_app_dns_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameServer_OriginalRule); i {
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
		file_app_dns_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config_HostMapping); i {
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
			RawDescriptor: file_app_dns_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_dns_config_proto_goTypes,
		DependencyIndexes: file_app_dns_config_proto_depIdxs,
		EnumInfos:         file_app_dns_config_proto_enumTypes,
		MessageInfos:      file_app_dns_config_proto_msgTypes,
	}.Build()
	File_app_dns_config_proto = out.File
	file_app_dns_config_proto_rawDesc = nil
	file_app_dns_config_proto_goTypes = nil
	file_app_dns_config_proto_depIdxs = nil
}

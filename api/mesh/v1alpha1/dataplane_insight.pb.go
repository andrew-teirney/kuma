// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/dataplane_insight.proto

package v1alpha1

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// DataplaneInsight defines the observed state of a Dataplane.
type DataplaneInsight struct {
	// List of ADS subscriptions created by a given Dataplane.
	Subscriptions []*DiscoverySubscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
	// Insights about mTLS for Dataplane.
	MTLS                 *DataplaneInsight_MTLS `protobuf:"bytes,2,opt,name=mTLS,proto3" json:"mTLS,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DataplaneInsight) Reset()         { *m = DataplaneInsight{} }
func (m *DataplaneInsight) String() string { return proto.CompactTextString(m) }
func (*DataplaneInsight) ProtoMessage()    {}
func (*DataplaneInsight) Descriptor() ([]byte, []int) {
	return fileDescriptor_35794f05b529b342, []int{0}
}

func (m *DataplaneInsight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneInsight.Unmarshal(m, b)
}
func (m *DataplaneInsight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneInsight.Marshal(b, m, deterministic)
}
func (m *DataplaneInsight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneInsight.Merge(m, src)
}
func (m *DataplaneInsight) XXX_Size() int {
	return xxx_messageInfo_DataplaneInsight.Size(m)
}
func (m *DataplaneInsight) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneInsight.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneInsight proto.InternalMessageInfo

func (m *DataplaneInsight) GetSubscriptions() []*DiscoverySubscription {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

func (m *DataplaneInsight) GetMTLS() *DataplaneInsight_MTLS {
	if m != nil {
		return m.MTLS
	}
	return nil
}

// MTLS defines insights for mTLS
type DataplaneInsight_MTLS struct {
	// Expiration time of the last certificate that was generated for a
	// Dataplane.
	CertificateExpirationTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=certificate_expiration_time,json=certificateExpirationTime,proto3" json:"certificate_expiration_time,omitempty"`
	// Time on which the last certificate was generated.
	LastCertificateRegeneration *timestamp.Timestamp `protobuf:"bytes,2,opt,name=last_certificate_regeneration,json=lastCertificateRegeneration,proto3" json:"last_certificate_regeneration,omitempty"`
	// Number of certificate regenerations for a Dataplane.
	CertificateRegenerations uint32   `protobuf:"varint,3,opt,name=certificate_regenerations,json=certificateRegenerations,proto3" json:"certificate_regenerations,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *DataplaneInsight_MTLS) Reset()         { *m = DataplaneInsight_MTLS{} }
func (m *DataplaneInsight_MTLS) String() string { return proto.CompactTextString(m) }
func (*DataplaneInsight_MTLS) ProtoMessage()    {}
func (*DataplaneInsight_MTLS) Descriptor() ([]byte, []int) {
	return fileDescriptor_35794f05b529b342, []int{0, 0}
}

func (m *DataplaneInsight_MTLS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneInsight_MTLS.Unmarshal(m, b)
}
func (m *DataplaneInsight_MTLS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneInsight_MTLS.Marshal(b, m, deterministic)
}
func (m *DataplaneInsight_MTLS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneInsight_MTLS.Merge(m, src)
}
func (m *DataplaneInsight_MTLS) XXX_Size() int {
	return xxx_messageInfo_DataplaneInsight_MTLS.Size(m)
}
func (m *DataplaneInsight_MTLS) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneInsight_MTLS.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneInsight_MTLS proto.InternalMessageInfo

func (m *DataplaneInsight_MTLS) GetCertificateExpirationTime() *timestamp.Timestamp {
	if m != nil {
		return m.CertificateExpirationTime
	}
	return nil
}

func (m *DataplaneInsight_MTLS) GetLastCertificateRegeneration() *timestamp.Timestamp {
	if m != nil {
		return m.LastCertificateRegeneration
	}
	return nil
}

func (m *DataplaneInsight_MTLS) GetCertificateRegenerations() uint32 {
	if m != nil {
		return m.CertificateRegenerations
	}
	return 0
}

// DiscoverySubscription describes a single ADS subscription
// created by a Dataplane to the Control Plane.
// Ideally, there should be only one such subscription per Dataplane lifecycle.
// Presence of multiple subscriptions might indicate one of the following
// events:
// - transient loss of network connection between Dataplane and Control Plane
// - Dataplane restart (i.e. hot restart or crash)
// - Control Plane restart (i.e. rolling update or crash)
// - etc
type DiscoverySubscription struct {
	// Unique id per ADS subscription.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Control Plane instance that handled given subscription.
	ControlPlaneInstanceId string `protobuf:"bytes,2,opt,name=control_plane_instance_id,json=controlPlaneInstanceId,proto3" json:"control_plane_instance_id,omitempty"`
	// Time when a given Dataplane connected to the Control Plane.
	ConnectTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=connect_time,json=connectTime,proto3" json:"connect_time,omitempty"`
	// Time when a given Dataplane disconnected from the Control Plane.
	DisconnectTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=disconnect_time,json=disconnectTime,proto3" json:"disconnect_time,omitempty"`
	// Status of the ADS subscription.
	Status               *DiscoverySubscriptionStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *DiscoverySubscription) Reset()         { *m = DiscoverySubscription{} }
func (m *DiscoverySubscription) String() string { return proto.CompactTextString(m) }
func (*DiscoverySubscription) ProtoMessage()    {}
func (*DiscoverySubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_35794f05b529b342, []int{1}
}

func (m *DiscoverySubscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoverySubscription.Unmarshal(m, b)
}
func (m *DiscoverySubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoverySubscription.Marshal(b, m, deterministic)
}
func (m *DiscoverySubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoverySubscription.Merge(m, src)
}
func (m *DiscoverySubscription) XXX_Size() int {
	return xxx_messageInfo_DiscoverySubscription.Size(m)
}
func (m *DiscoverySubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoverySubscription.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoverySubscription proto.InternalMessageInfo

func (m *DiscoverySubscription) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DiscoverySubscription) GetControlPlaneInstanceId() string {
	if m != nil {
		return m.ControlPlaneInstanceId
	}
	return ""
}

func (m *DiscoverySubscription) GetConnectTime() *timestamp.Timestamp {
	if m != nil {
		return m.ConnectTime
	}
	return nil
}

func (m *DiscoverySubscription) GetDisconnectTime() *timestamp.Timestamp {
	if m != nil {
		return m.DisconnectTime
	}
	return nil
}

func (m *DiscoverySubscription) GetStatus() *DiscoverySubscriptionStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// DiscoverySubscriptionStatus defines status of an ADS subscription.
type DiscoverySubscriptionStatus struct {
	// Time when status of a given ADS subscription was most recently updated.
	LastUpdateTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
	// Total defines an aggregate over individual xDS stats.
	Total *DiscoveryServiceStats `protobuf:"bytes,2,opt,name=total,proto3" json:"total,omitempty"`
	// CDS defines all CDS stats.
	Cds *DiscoveryServiceStats `protobuf:"bytes,3,opt,name=cds,proto3" json:"cds,omitempty"`
	// EDS defines all EDS stats.
	Eds *DiscoveryServiceStats `protobuf:"bytes,4,opt,name=eds,proto3" json:"eds,omitempty"`
	// LDS defines all LDS stats.
	Lds *DiscoveryServiceStats `protobuf:"bytes,5,opt,name=lds,proto3" json:"lds,omitempty"`
	// RDS defines all RDS stats.
	Rds                  *DiscoveryServiceStats `protobuf:"bytes,6,opt,name=rds,proto3" json:"rds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DiscoverySubscriptionStatus) Reset()         { *m = DiscoverySubscriptionStatus{} }
func (m *DiscoverySubscriptionStatus) String() string { return proto.CompactTextString(m) }
func (*DiscoverySubscriptionStatus) ProtoMessage()    {}
func (*DiscoverySubscriptionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_35794f05b529b342, []int{2}
}

func (m *DiscoverySubscriptionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoverySubscriptionStatus.Unmarshal(m, b)
}
func (m *DiscoverySubscriptionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoverySubscriptionStatus.Marshal(b, m, deterministic)
}
func (m *DiscoverySubscriptionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoverySubscriptionStatus.Merge(m, src)
}
func (m *DiscoverySubscriptionStatus) XXX_Size() int {
	return xxx_messageInfo_DiscoverySubscriptionStatus.Size(m)
}
func (m *DiscoverySubscriptionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoverySubscriptionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoverySubscriptionStatus proto.InternalMessageInfo

func (m *DiscoverySubscriptionStatus) GetLastUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *DiscoverySubscriptionStatus) GetTotal() *DiscoveryServiceStats {
	if m != nil {
		return m.Total
	}
	return nil
}

func (m *DiscoverySubscriptionStatus) GetCds() *DiscoveryServiceStats {
	if m != nil {
		return m.Cds
	}
	return nil
}

func (m *DiscoverySubscriptionStatus) GetEds() *DiscoveryServiceStats {
	if m != nil {
		return m.Eds
	}
	return nil
}

func (m *DiscoverySubscriptionStatus) GetLds() *DiscoveryServiceStats {
	if m != nil {
		return m.Lds
	}
	return nil
}

func (m *DiscoverySubscriptionStatus) GetRds() *DiscoveryServiceStats {
	if m != nil {
		return m.Rds
	}
	return nil
}

// DiscoveryServiceStats defines all stats over a single xDS service.
type DiscoveryServiceStats struct {
	// Number of xDS responses sent to the Dataplane.
	ResponsesSent uint64 `protobuf:"varint,1,opt,name=responses_sent,json=responsesSent,proto3" json:"responses_sent,omitempty"`
	// Number of xDS responses ACKed by the Dataplane.
	ResponsesAcknowledged uint64 `protobuf:"varint,2,opt,name=responses_acknowledged,json=responsesAcknowledged,proto3" json:"responses_acknowledged,omitempty"`
	// Number of xDS responses NACKed by the Dataplane.
	ResponsesRejected    uint64   `protobuf:"varint,3,opt,name=responses_rejected,json=responsesRejected,proto3" json:"responses_rejected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoveryServiceStats) Reset()         { *m = DiscoveryServiceStats{} }
func (m *DiscoveryServiceStats) String() string { return proto.CompactTextString(m) }
func (*DiscoveryServiceStats) ProtoMessage()    {}
func (*DiscoveryServiceStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_35794f05b529b342, []int{3}
}

func (m *DiscoveryServiceStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryServiceStats.Unmarshal(m, b)
}
func (m *DiscoveryServiceStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryServiceStats.Marshal(b, m, deterministic)
}
func (m *DiscoveryServiceStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryServiceStats.Merge(m, src)
}
func (m *DiscoveryServiceStats) XXX_Size() int {
	return xxx_messageInfo_DiscoveryServiceStats.Size(m)
}
func (m *DiscoveryServiceStats) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryServiceStats.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryServiceStats proto.InternalMessageInfo

func (m *DiscoveryServiceStats) GetResponsesSent() uint64 {
	if m != nil {
		return m.ResponsesSent
	}
	return 0
}

func (m *DiscoveryServiceStats) GetResponsesAcknowledged() uint64 {
	if m != nil {
		return m.ResponsesAcknowledged
	}
	return 0
}

func (m *DiscoveryServiceStats) GetResponsesRejected() uint64 {
	if m != nil {
		return m.ResponsesRejected
	}
	return 0
}

func init() {
	proto.RegisterType((*DataplaneInsight)(nil), "kuma.mesh.v1alpha1.DataplaneInsight")
	proto.RegisterType((*DataplaneInsight_MTLS)(nil), "kuma.mesh.v1alpha1.DataplaneInsight.MTLS")
	proto.RegisterType((*DiscoverySubscription)(nil), "kuma.mesh.v1alpha1.DiscoverySubscription")
	proto.RegisterType((*DiscoverySubscriptionStatus)(nil), "kuma.mesh.v1alpha1.DiscoverySubscriptionStatus")
	proto.RegisterType((*DiscoveryServiceStats)(nil), "kuma.mesh.v1alpha1.DiscoveryServiceStats")
}

func init() {
	proto.RegisterFile("mesh/v1alpha1/dataplane_insight.proto", fileDescriptor_35794f05b529b342)
}

var fileDescriptor_35794f05b529b342 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xdf, 0x4e, 0x13, 0x41,
	0x14, 0xc6, 0xb3, 0xed, 0x82, 0x38, 0x08, 0xe2, 0x24, 0xc0, 0x52, 0x62, 0x24, 0x4d, 0x48, 0xf0,
	0xc2, 0x6d, 0xd0, 0x78, 0x45, 0x8c, 0x71, 0xc1, 0x18, 0x12, 0x8c, 0x3a, 0xc5, 0x1b, 0x2e, 0xdc,
	0x0c, 0x33, 0x87, 0x32, 0xb2, 0x9d, 0xd9, 0xcc, 0x4c, 0xab, 0xbe, 0x82, 0x4f, 0xe0, 0x03, 0xf0,
	0x04, 0xbe, 0x8b, 0x2f, 0xc3, 0x8d, 0x66, 0x66, 0xff, 0x74, 0x55, 0xfe, 0xf5, 0xae, 0xed, 0x39,
	0xbf, 0xef, 0x9c, 0xf3, 0xcd, 0x57, 0xb4, 0x39, 0x04, 0x73, 0xda, 0x1b, 0x6f, 0xd3, 0x2c, 0x3f,
	0xa5, 0xdb, 0x3d, 0x4e, 0x2d, 0xcd, 0x33, 0x2a, 0x21, 0x15, 0xd2, 0x88, 0xc1, 0xa9, 0x8d, 0x73,
	0xad, 0xac, 0xc2, 0xf8, 0x6c, 0x34, 0xa4, 0xb1, 0xeb, 0x8d, 0xab, 0xde, 0xce, 0xa3, 0x81, 0x52,
	0x83, 0x0c, 0x7a, 0xbe, 0xe3, 0x78, 0x74, 0xd2, 0xb3, 0x62, 0x08, 0xc6, 0xd2, 0x61, 0x5e, 0x40,
	0x9d, 0xd5, 0x31, 0xcd, 0x04, 0xa7, 0x16, 0x7a, 0xd5, 0x87, 0xa2, 0xd0, 0x3d, 0x6f, 0xa3, 0xa5,
	0xbd, 0x6a, 0xd2, 0x7e, 0x31, 0x08, 0xbf, 0x43, 0x0b, 0x66, 0x74, 0x6c, 0x98, 0x16, 0xb9, 0x15,
	0x4a, 0x9a, 0x28, 0xd8, 0x68, 0x6f, 0xcd, 0x3f, 0x7d, 0x1c, 0xff, 0x3f, 0x3a, 0xde, 0x13, 0x86,
	0xa9, 0x31, 0xe8, 0x6f, 0xfd, 0x06, 0x41, 0xfe, 0xe6, 0xf1, 0x0b, 0x14, 0x0e, 0x0f, 0x0f, 0xfa,
	0x51, 0x6b, 0x23, 0xb8, 0x52, 0xe7, 0x9f, 0x25, 0xe2, 0xb7, 0x87, 0x07, 0x7d, 0xe2, 0xb1, 0xce,
	0xef, 0x00, 0x85, 0xee, 0x2b, 0x3e, 0x42, 0xeb, 0x0c, 0xb4, 0x15, 0x27, 0x82, 0x51, 0x0b, 0x29,
	0x7c, 0xcd, 0x85, 0xa6, 0x6e, 0x44, 0xea, 0x0e, 0x8e, 0x02, 0x2f, 0xdf, 0x89, 0x0b, 0x37, 0xe2,
	0xca, 0x8d, 0xf8, 0xb0, 0x72, 0x83, 0xac, 0x35, 0xf0, 0xd7, 0x35, 0xed, 0xea, 0xf8, 0x13, 0x7a,
	0x98, 0x51, 0x63, 0xd3, 0xe6, 0x00, 0x0d, 0x03, 0x90, 0x50, 0x34, 0x95, 0xcb, 0x5f, 0xa7, 0xbe,
	0xee, 0x04, 0x76, 0x27, 0x3c, 0x69, 0xe0, 0x78, 0x07, 0xad, 0x5d, 0x25, 0x6d, 0xa2, 0xf6, 0x46,
	0xb0, 0xb5, 0x40, 0x22, 0x76, 0x39, 0x6b, 0xba, 0xbf, 0x5a, 0x68, 0xf9, 0x52, 0xa7, 0xf1, 0x2a,
	0x6a, 0x09, 0xee, 0x2f, 0xbf, 0x9b, 0xdc, 0xb9, 0x48, 0x42, 0xdd, 0x5a, 0x0a, 0x48, 0x4b, 0x70,
	0x9c, 0xa0, 0x35, 0xa6, 0xa4, 0xd5, 0x2a, 0x4b, 0xeb, 0x18, 0x59, 0x2a, 0x19, 0xa4, 0x82, 0xfb,
	0x5b, 0x1a, 0xfd, 0x2b, 0x65, 0xe7, 0xfb, 0xf2, 0x01, 0x7c, 0xdf, 0x3e, 0xc7, 0x6f, 0xd0, 0x3d,
	0xa6, 0xa4, 0x04, 0x66, 0x0b, 0x83, 0xdb, 0x37, 0x59, 0x90, 0xcc, 0x5d, 0x24, 0x33, 0x3f, 0x83,
	0xd6, 0x5c, 0x40, 0xe6, 0x4b, 0xd2, 0x9b, 0xbb, 0x8b, 0xee, 0x73, 0xb7, 0x7e, 0x43, 0x2b, 0xbc,
	0xd1, 0xce, 0xc5, 0x09, 0xe2, 0x45, 0x3e, 0xa0, 0x59, 0x63, 0xa9, 0x1d, 0x99, 0x68, 0xc6, 0xb3,
	0xbd, 0x5b, 0xe7, 0xb1, 0xef, 0x31, 0xbf, 0xdc, 0xf7, 0xc0, 0x1d, 0x5c, 0x0a, 0x75, 0x7f, 0xb4,
	0xd1, 0xfa, 0x35, 0x04, 0xde, 0x43, 0x4b, 0x3e, 0x14, 0xa3, 0xdc, 0xfd, 0x67, 0x6e, 0x9b, 0xb2,
	0x45, 0xc7, 0x7c, 0xf4, 0x88, 0x5f, 0xfc, 0x25, 0x9a, 0xb1, 0xca, 0xd2, 0xec, 0xda, 0xfc, 0xd7,
	0x5b, 0x80, 0x1e, 0x0b, 0x06, 0x6e, 0x01, 0x43, 0x0a, 0x0e, 0xef, 0xa0, 0x36, 0xe3, 0xa6, 0xb4,
	0x7f, 0x0a, 0xdc, 0x51, 0x0e, 0x06, 0x6e, 0x4a, 0xbf, 0xa7, 0x81, 0xa1, 0x80, 0x33, 0x5e, 0x19,
	0x3e, 0x0d, 0x9c, 0x15, 0xb0, 0xe6, 0x26, 0x9a, 0x9d, 0x1a, 0xd6, 0xdc, 0x74, 0xcf, 0x83, 0x66,
	0xe4, 0x1b, 0x65, 0xbc, 0x89, 0x16, 0x35, 0x98, 0x5c, 0x49, 0x03, 0x26, 0x35, 0x20, 0xad, 0x7f,
	0x92, 0x90, 0x2c, 0xd4, 0xbf, 0xf6, 0x41, 0x5a, 0xfc, 0x1c, 0xad, 0x4c, 0xda, 0x28, 0x3b, 0x93,
	0xea, 0x4b, 0x06, 0x7c, 0x00, 0x45, 0xfa, 0x43, 0xb2, 0x5c, 0x57, 0x5f, 0x35, 0x8a, 0xf8, 0x09,
	0xc2, 0x13, 0x4c, 0xc3, 0x67, 0x60, 0x16, 0xb8, 0xb7, 0x3e, 0x24, 0x0f, 0xea, 0x0a, 0x29, 0x0b,
	0x09, 0x3a, 0x9a, 0xab, 0xae, 0x39, 0x9e, 0xf5, 0x59, 0x78, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff,
	0xe1, 0x68, 0x8d, 0x9f, 0xca, 0x05, 0x00, 0x00,
}

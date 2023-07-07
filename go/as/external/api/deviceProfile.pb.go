// Code generated by protoc-gen-go. DO NOT EDIT.
// source: as/external/api/deviceProfile.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateDeviceProfileRequest struct {
	// Device-profile object to create.
	DeviceProfile        *DeviceProfile `protobuf:"bytes,1,opt,name=device_profile,json=deviceProfile,proto3" json:"device_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateDeviceProfileRequest) Reset()         { *m = CreateDeviceProfileRequest{} }
func (m *CreateDeviceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDeviceProfileRequest) ProtoMessage()    {}
func (*CreateDeviceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dcf865cca06987, []int{0}
}

func (m *CreateDeviceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDeviceProfileRequest.Unmarshal(m, b)
}
func (m *CreateDeviceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDeviceProfileRequest.Marshal(b, m, deterministic)
}
func (m *CreateDeviceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDeviceProfileRequest.Merge(m, src)
}
func (m *CreateDeviceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDeviceProfileRequest.Size(m)
}
func (m *CreateDeviceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDeviceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDeviceProfileRequest proto.InternalMessageInfo

func (m *CreateDeviceProfileRequest) GetDeviceProfile() *DeviceProfile {
	if m != nil {
		return m.DeviceProfile
	}
	return nil
}

type CreateDeviceProfileResponse struct {
	// Device-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDeviceProfileResponse) Reset()         { *m = CreateDeviceProfileResponse{} }
func (m *CreateDeviceProfileResponse) String() string { return proto.CompactTextString(m) }
func (*CreateDeviceProfileResponse) ProtoMessage()    {}
func (*CreateDeviceProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dcf865cca06987, []int{1}
}

func (m *CreateDeviceProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDeviceProfileResponse.Unmarshal(m, b)
}
func (m *CreateDeviceProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDeviceProfileResponse.Marshal(b, m, deterministic)
}
func (m *CreateDeviceProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDeviceProfileResponse.Merge(m, src)
}
func (m *CreateDeviceProfileResponse) XXX_Size() int {
	return xxx_messageInfo_CreateDeviceProfileResponse.Size(m)
}
func (m *CreateDeviceProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDeviceProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDeviceProfileResponse proto.InternalMessageInfo

func (m *CreateDeviceProfileResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetDeviceProfileRequest struct {
	// Device-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeviceProfileRequest) Reset()         { *m = GetDeviceProfileRequest{} }
func (m *GetDeviceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeviceProfileRequest) ProtoMessage()    {}
func (*GetDeviceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dcf865cca06987, []int{2}
}

func (m *GetDeviceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceProfileRequest.Unmarshal(m, b)
}
func (m *GetDeviceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceProfileRequest.Marshal(b, m, deterministic)
}
func (m *GetDeviceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceProfileRequest.Merge(m, src)
}
func (m *GetDeviceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeviceProfileRequest.Size(m)
}
func (m *GetDeviceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceProfileRequest proto.InternalMessageInfo

func (m *GetDeviceProfileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetDeviceProfileResponse struct {
	// Device-profile object.
	DeviceProfile *DeviceProfile `protobuf:"bytes,1,opt,name=device_profile,json=deviceProfile,proto3" json:"device_profile,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetDeviceProfileResponse) Reset()         { *m = GetDeviceProfileResponse{} }
func (m *GetDeviceProfileResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeviceProfileResponse) ProtoMessage()    {}
func (*GetDeviceProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dcf865cca06987, []int{3}
}

func (m *GetDeviceProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceProfileResponse.Unmarshal(m, b)
}
func (m *GetDeviceProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceProfileResponse.Marshal(b, m, deterministic)
}
func (m *GetDeviceProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceProfileResponse.Merge(m, src)
}
func (m *GetDeviceProfileResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeviceProfileResponse.Size(m)
}
func (m *GetDeviceProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceProfileResponse proto.InternalMessageInfo

func (m *GetDeviceProfileResponse) GetDeviceProfile() *DeviceProfile {
	if m != nil {
		return m.DeviceProfile
	}
	return nil
}

func (m *GetDeviceProfileResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GetDeviceProfileResponse) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type UpdateDeviceProfileRequest struct {
	// Device-profile object to update.
	DeviceProfile        *DeviceProfile `protobuf:"bytes,1,opt,name=device_profile,json=deviceProfile,proto3" json:"device_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateDeviceProfileRequest) Reset()         { *m = UpdateDeviceProfileRequest{} }
func (m *UpdateDeviceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateDeviceProfileRequest) ProtoMessage()    {}
func (*UpdateDeviceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dcf865cca06987, []int{4}
}

func (m *UpdateDeviceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDeviceProfileRequest.Unmarshal(m, b)
}
func (m *UpdateDeviceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDeviceProfileRequest.Marshal(b, m, deterministic)
}
func (m *UpdateDeviceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDeviceProfileRequest.Merge(m, src)
}
func (m *UpdateDeviceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateDeviceProfileRequest.Size(m)
}
func (m *UpdateDeviceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDeviceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDeviceProfileRequest proto.InternalMessageInfo

func (m *UpdateDeviceProfileRequest) GetDeviceProfile() *DeviceProfile {
	if m != nil {
		return m.DeviceProfile
	}
	return nil
}

type DeleteDeviceProfileRequest struct {
	// Device-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDeviceProfileRequest) Reset()         { *m = DeleteDeviceProfileRequest{} }
func (m *DeleteDeviceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDeviceProfileRequest) ProtoMessage()    {}
func (*DeleteDeviceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dcf865cca06987, []int{5}
}

func (m *DeleteDeviceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDeviceProfileRequest.Unmarshal(m, b)
}
func (m *DeleteDeviceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDeviceProfileRequest.Marshal(b, m, deterministic)
}
func (m *DeleteDeviceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDeviceProfileRequest.Merge(m, src)
}
func (m *DeleteDeviceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDeviceProfileRequest.Size(m)
}
func (m *DeleteDeviceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDeviceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDeviceProfileRequest proto.InternalMessageInfo

func (m *DeleteDeviceProfileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeviceProfileListItem struct {
	// Device-profile ID (UUID string).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Device-profile name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Organization ID.
	OrganizationId int64 `protobuf:"varint,3,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Network-server ID.
	NetworkServerId int64 `protobuf:"varint,4,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Network-server name.
	NetworkServerName    string   `protobuf:"bytes,7,opt,name=network_server_name,json=networkServerName,proto3" json:"network_server_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceProfileListItem) Reset()         { *m = DeviceProfileListItem{} }
func (m *DeviceProfileListItem) String() string { return proto.CompactTextString(m) }
func (*DeviceProfileListItem) ProtoMessage()    {}
func (*DeviceProfileListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dcf865cca06987, []int{6}
}

func (m *DeviceProfileListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceProfileListItem.Unmarshal(m, b)
}
func (m *DeviceProfileListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceProfileListItem.Marshal(b, m, deterministic)
}
func (m *DeviceProfileListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceProfileListItem.Merge(m, src)
}
func (m *DeviceProfileListItem) XXX_Size() int {
	return xxx_messageInfo_DeviceProfileListItem.Size(m)
}
func (m *DeviceProfileListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceProfileListItem.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceProfileListItem proto.InternalMessageInfo

func (m *DeviceProfileListItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeviceProfileListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceProfileListItem) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *DeviceProfileListItem) GetNetworkServerId() int64 {
	if m != nil {
		return m.NetworkServerId
	}
	return 0
}

func (m *DeviceProfileListItem) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *DeviceProfileListItem) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *DeviceProfileListItem) GetNetworkServerName() string {
	if m != nil {
		return m.NetworkServerName
	}
	return ""
}

type ListDeviceProfileRequest struct {
	// Max number of items to return.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Organization id to filter on.
	OrganizationId int64 `protobuf:"varint,3,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Application id to filter on.
	ApplicationId        int64    `protobuf:"varint,4,opt,name=application_id,json=applicationID,proto3" json:"application_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDeviceProfileRequest) Reset()         { *m = ListDeviceProfileRequest{} }
func (m *ListDeviceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*ListDeviceProfileRequest) ProtoMessage()    {}
func (*ListDeviceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dcf865cca06987, []int{7}
}

func (m *ListDeviceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDeviceProfileRequest.Unmarshal(m, b)
}
func (m *ListDeviceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDeviceProfileRequest.Marshal(b, m, deterministic)
}
func (m *ListDeviceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDeviceProfileRequest.Merge(m, src)
}
func (m *ListDeviceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_ListDeviceProfileRequest.Size(m)
}
func (m *ListDeviceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDeviceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDeviceProfileRequest proto.InternalMessageInfo

func (m *ListDeviceProfileRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListDeviceProfileRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListDeviceProfileRequest) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *ListDeviceProfileRequest) GetApplicationId() int64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

type ListDeviceProfileResponse struct {
	// Total number of device-profiles.
	TotalCount           int64                    `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Result               []*DeviceProfileListItem `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ListDeviceProfileResponse) Reset()         { *m = ListDeviceProfileResponse{} }
func (m *ListDeviceProfileResponse) String() string { return proto.CompactTextString(m) }
func (*ListDeviceProfileResponse) ProtoMessage()    {}
func (*ListDeviceProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dcf865cca06987, []int{8}
}

func (m *ListDeviceProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDeviceProfileResponse.Unmarshal(m, b)
}
func (m *ListDeviceProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDeviceProfileResponse.Marshal(b, m, deterministic)
}
func (m *ListDeviceProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDeviceProfileResponse.Merge(m, src)
}
func (m *ListDeviceProfileResponse) XXX_Size() int {
	return xxx_messageInfo_ListDeviceProfileResponse.Size(m)
}
func (m *ListDeviceProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDeviceProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDeviceProfileResponse proto.InternalMessageInfo

func (m *ListDeviceProfileResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListDeviceProfileResponse) GetResult() []*DeviceProfileListItem {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateDeviceProfileRequest)(nil), "api.CreateDeviceProfileRequest")
	proto.RegisterType((*CreateDeviceProfileResponse)(nil), "api.CreateDeviceProfileResponse")
	proto.RegisterType((*GetDeviceProfileRequest)(nil), "api.GetDeviceProfileRequest")
	proto.RegisterType((*GetDeviceProfileResponse)(nil), "api.GetDeviceProfileResponse")
	proto.RegisterType((*UpdateDeviceProfileRequest)(nil), "api.UpdateDeviceProfileRequest")
	proto.RegisterType((*DeleteDeviceProfileRequest)(nil), "api.DeleteDeviceProfileRequest")
	proto.RegisterType((*DeviceProfileListItem)(nil), "api.DeviceProfileListItem")
	proto.RegisterType((*ListDeviceProfileRequest)(nil), "api.ListDeviceProfileRequest")
	proto.RegisterType((*ListDeviceProfileResponse)(nil), "api.ListDeviceProfileResponse")
}

func init() {
	proto.RegisterFile("as/external/api/deviceProfile.proto", fileDescriptor_d0dcf865cca06987)
}

var fileDescriptor_d0dcf865cca06987 = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcb, 0x6e, 0xd4, 0x4a,
	0x10, 0xd5, 0x3c, 0xe2, 0xab, 0x54, 0x94, 0x89, 0xd2, 0x37, 0x37, 0x77, 0xe2, 0x84, 0x4c, 0x62,
	0x84, 0x08, 0x11, 0x63, 0x4b, 0xc9, 0x02, 0x85, 0x5d, 0x92, 0x41, 0xd1, 0x48, 0x08, 0x45, 0x06,
	0x84, 0xc4, 0x66, 0xd4, 0xb1, 0x6b, 0x26, 0xad, 0xd8, 0x6e, 0x63, 0xb7, 0x87, 0x97, 0xb2, 0xe1,
	0x17, 0xd8, 0xf0, 0x15, 0x2c, 0xf9, 0x09, 0x96, 0xfc, 0x02, 0x1f, 0x82, 0xdc, 0x6e, 0xc3, 0xd8,
	0xb1, 0x79, 0x44, 0xec, 0xdc, 0xdd, 0xa7, 0xaa, 0x4e, 0x9d, 0x7a, 0x18, 0x6e, 0xd2, 0xd8, 0xc2,
	0x57, 0x02, 0xa3, 0x80, 0x7a, 0x16, 0x0d, 0x99, 0xe5, 0xe2, 0x94, 0x39, 0x78, 0x1a, 0xf1, 0x31,
	0xf3, 0xd0, 0x0c, 0x23, 0x2e, 0x38, 0x69, 0xd1, 0x90, 0xe9, 0x1b, 0x13, 0xce, 0x27, 0x1e, 0x4a,
	0x10, 0x0d, 0x02, 0x2e, 0xa8, 0x60, 0x3c, 0x88, 0x33, 0x88, 0xde, 0x53, 0xaf, 0xf2, 0x74, 0x96,
	0x8c, 0x2d, 0xc1, 0x7c, 0x8c, 0x05, 0xf5, 0x43, 0x05, 0x58, 0x2f, 0x03, 0xd0, 0x0f, 0xc5, 0x6b,
	0xf5, 0xb8, 0x59, 0x66, 0x11, 0x66, 0xf1, 0x95, 0x77, 0xe3, 0x19, 0xe8, 0xc7, 0x11, 0x52, 0x81,
	0x83, 0x59, 0x76, 0x36, 0xbe, 0x48, 0x30, 0x16, 0xe4, 0x00, 0x3a, 0x19, 0xeb, 0x91, 0x32, 0xeb,
	0x36, 0xb6, 0x1a, 0x3b, 0x0b, 0x7b, 0xc4, 0xa4, 0x21, 0x33, 0x8b, 0x26, 0x8b, 0x85, 0xfc, 0x8c,
	0x3e, 0xac, 0x57, 0x3a, 0x8e, 0x43, 0x1e, 0xc4, 0x48, 0x3a, 0xd0, 0x64, 0xae, 0xf4, 0x36, 0x6f,
	0x37, 0x99, 0x6b, 0xdc, 0x81, 0xff, 0x4f, 0x50, 0x54, 0x92, 0x28, 0x43, 0x3f, 0x37, 0xa0, 0x7b,
	0x15, 0xab, 0xfc, 0x5e, 0x9f, 0x31, 0x39, 0x00, 0x70, 0x24, 0x63, 0x77, 0x44, 0x45, 0xb7, 0x29,
	0xcd, 0x74, 0x33, 0x13, 0xd7, 0xcc, 0xc5, 0x35, 0x9f, 0xe4, 0xea, 0xdb, 0xf3, 0x0a, 0x7d, 0x98,
	0xea, 0x04, 0x49, 0xe8, 0xe6, 0xa6, 0xad, 0x5f, 0x9b, 0x2a, 0xf4, 0xa1, 0x48, 0x0b, 0xf0, 0x54,
	0x1e, 0xfe, 0x76, 0x01, 0xee, 0x82, 0x3e, 0x40, 0x0f, 0x6b, 0x1c, 0x97, 0x45, 0xfd, 0xd4, 0x84,
	0xff, 0x0a, 0xc0, 0x87, 0x2c, 0x16, 0x43, 0x81, 0x7e, 0x19, 0x49, 0x08, 0xb4, 0x03, 0xea, 0xa3,
	0x14, 0x68, 0xde, 0x96, 0xdf, 0xe4, 0x36, 0x2c, 0xf1, 0x68, 0x42, 0x03, 0xf6, 0x46, 0xb6, 0xee,
	0x88, 0xb9, 0x52, 0x84, 0x96, 0xdd, 0x99, 0xbd, 0x1e, 0x0e, 0xc8, 0x2e, 0x2c, 0x07, 0x28, 0x5e,
	0xf2, 0xe8, 0x62, 0x14, 0x63, 0x34, 0xc5, 0x28, 0x85, 0xb6, 0x25, 0x74, 0x49, 0x3d, 0x3c, 0x96,
	0xf7, 0xc3, 0x41, 0xa9, 0x1e, 0x73, 0xd7, 0xaf, 0x87, 0xf6, 0x07, 0xf5, 0x20, 0x26, 0xfc, 0x5b,
	0x62, 0x28, 0xb3, 0xfd, 0x47, 0x66, 0xbb, 0x5c, 0xe0, 0xf8, 0x88, 0xfa, 0x68, 0x7c, 0x68, 0x40,
	0x37, 0xd5, 0xaa, 0x52, 0xe5, 0x15, 0x98, 0xf3, 0x98, 0xcf, 0x84, 0x94, 0xaf, 0x65, 0x67, 0x07,
	0xb2, 0x0a, 0x1a, 0x1f, 0x8f, 0x63, 0xcc, 0x9a, 0xac, 0x65, 0xab, 0xd3, 0xef, 0xab, 0x78, 0x0b,
	0x3a, 0x34, 0x0c, 0x3d, 0xe6, 0x7c, 0xc7, 0x65, 0x12, 0x2e, 0xce, 0xdc, 0x0e, 0x07, 0x46, 0x08,
	0x6b, 0x15, 0xcc, 0xd4, 0xa0, 0xf4, 0x60, 0x41, 0x70, 0x41, 0xbd, 0x91, 0xc3, 0x93, 0x20, 0x27,
	0x08, 0xf2, 0xea, 0x38, 0xbd, 0x21, 0x7b, 0xa0, 0x45, 0x18, 0x27, 0x5e, 0xca, 0xb2, 0x25, 0xf5,
	0xbb, 0xd2, 0x72, 0x79, 0x8f, 0xd8, 0x0a, 0xb9, 0xf7, 0xb1, 0x0d, 0x2b, 0x05, 0x44, 0x2a, 0x14,
	0x73, 0x90, 0x78, 0xa0, 0x65, 0xdb, 0x80, 0xf4, 0xa4, 0x9b, 0xfa, 0x9d, 0xa3, 0x6f, 0xd5, 0x03,
	0x32, 0xea, 0x46, 0xef, 0xdd, 0x97, 0xaf, 0xef, 0x9b, 0x6b, 0xc6, 0xca, 0xcc, 0x5a, 0xed, 0xe7,
	0x7b, 0xed, 0x7e, 0x63, 0x97, 0x20, 0xb4, 0x4e, 0x50, 0x90, 0x0d, 0xe9, 0xa9, 0x66, 0xad, 0xe8,
	0x37, 0x6a, 0x5e, 0x55, 0x90, 0x6d, 0x19, 0x64, 0x9d, 0xac, 0x55, 0x05, 0xb1, 0xde, 0x32, 0xf7,
	0x92, 0x4c, 0x41, 0xcb, 0x46, 0x57, 0x25, 0x55, 0x3f, 0xc7, 0xfa, 0xea, 0x95, 0xe6, 0x7b, 0x90,
	0x2e, 0x69, 0x63, 0x5f, 0x46, 0xe9, 0xeb, 0x3b, 0xd5, 0x51, 0x8a, 0xb3, 0x6f, 0x32, 0xf7, 0x32,
	0x4d, 0xcf, 0x05, 0x2d, 0x9b, 0x6c, 0x15, 0xb7, 0x7e, 0xcc, 0x6b, 0xe3, 0xaa, 0xec, 0x76, 0x7f,
	0x92, 0x9d, 0x03, 0xed, 0xb4, 0xbe, 0x24, 0xd3, 0xa9, 0xae, 0xc5, 0xf5, 0xcd, 0xba, 0x67, 0xa5,
	0xe3, 0x86, 0x8c, 0xb4, 0x4a, 0x2a, 0x8b, 0x75, 0x14, 0xc1, 0x36, 0xe3, 0xa6, 0x73, 0xce, 0xa2,
	0x30, 0x16, 0xd4, 0xb9, 0x90, 0xce, 0x68, 0x6c, 0xe6, 0xbf, 0xac, 0xf4, 0x7c, 0x44, 0x0a, 0x9e,
	0x4f, 0xd3, 0x4c, 0x4e, 0x1b, 0xcf, 0xef, 0x4d, 0x98, 0x38, 0x4f, 0xce, 0x4c, 0x87, 0xfb, 0xd6,
	0x05, 0xf5, 0x99, 0x93, 0x24, 0xd6, 0x0f, 0x47, 0xfd, 0x34, 0xda, 0x84, 0x5b, 0xd3, 0x7d, 0xab,
	0xf4, 0x07, 0x3c, 0xd3, 0xa4, 0x16, 0xfb, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x42, 0xe8,
	0x08, 0xa1, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeviceProfileServiceClient is the client API for DeviceProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceProfileServiceClient interface {
	// Create creates the given device-profile.
	Create(ctx context.Context, in *CreateDeviceProfileRequest, opts ...grpc.CallOption) (*CreateDeviceProfileResponse, error)
	// Get returns the device-profile matching the given id.
	Get(ctx context.Context, in *GetDeviceProfileRequest, opts ...grpc.CallOption) (*GetDeviceProfileResponse, error)
	// Update updates the given device-profile.
	Update(ctx context.Context, in *UpdateDeviceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete deletes the device-profile matching the given id.
	Delete(ctx context.Context, in *DeleteDeviceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// List lists the available device-profiles.
	List(ctx context.Context, in *ListDeviceProfileRequest, opts ...grpc.CallOption) (*ListDeviceProfileResponse, error)
}

type deviceProfileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceProfileServiceClient(cc grpc.ClientConnInterface) DeviceProfileServiceClient {
	return &deviceProfileServiceClient{cc}
}

func (c *deviceProfileServiceClient) Create(ctx context.Context, in *CreateDeviceProfileRequest, opts ...grpc.CallOption) (*CreateDeviceProfileResponse, error) {
	out := new(CreateDeviceProfileResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceProfileService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceProfileServiceClient) Get(ctx context.Context, in *GetDeviceProfileRequest, opts ...grpc.CallOption) (*GetDeviceProfileResponse, error) {
	out := new(GetDeviceProfileResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceProfileService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceProfileServiceClient) Update(ctx context.Context, in *UpdateDeviceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.DeviceProfileService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceProfileServiceClient) Delete(ctx context.Context, in *DeleteDeviceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.DeviceProfileService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceProfileServiceClient) List(ctx context.Context, in *ListDeviceProfileRequest, opts ...grpc.CallOption) (*ListDeviceProfileResponse, error) {
	out := new(ListDeviceProfileResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceProfileService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceProfileServiceServer is the server API for DeviceProfileService service.
type DeviceProfileServiceServer interface {
	// Create creates the given device-profile.
	Create(context.Context, *CreateDeviceProfileRequest) (*CreateDeviceProfileResponse, error)
	// Get returns the device-profile matching the given id.
	Get(context.Context, *GetDeviceProfileRequest) (*GetDeviceProfileResponse, error)
	// Update updates the given device-profile.
	Update(context.Context, *UpdateDeviceProfileRequest) (*empty.Empty, error)
	// Delete deletes the device-profile matching the given id.
	Delete(context.Context, *DeleteDeviceProfileRequest) (*empty.Empty, error)
	// List lists the available device-profiles.
	List(context.Context, *ListDeviceProfileRequest) (*ListDeviceProfileResponse, error)
}

// UnimplementedDeviceProfileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeviceProfileServiceServer struct {
}

func (*UnimplementedDeviceProfileServiceServer) Create(ctx context.Context, req *CreateDeviceProfileRequest) (*CreateDeviceProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedDeviceProfileServiceServer) Get(ctx context.Context, req *GetDeviceProfileRequest) (*GetDeviceProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedDeviceProfileServiceServer) Update(ctx context.Context, req *UpdateDeviceProfileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedDeviceProfileServiceServer) Delete(ctx context.Context, req *DeleteDeviceProfileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedDeviceProfileServiceServer) List(ctx context.Context, req *ListDeviceProfileRequest) (*ListDeviceProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterDeviceProfileServiceServer(s *grpc.Server, srv DeviceProfileServiceServer) {
	s.RegisterService(&_DeviceProfileService_serviceDesc, srv)
}

func _DeviceProfileService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceProfileService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).Create(ctx, req.(*CreateDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceProfileService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceProfileService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).Get(ctx, req.(*GetDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceProfileService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceProfileService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).Update(ctx, req.(*UpdateDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceProfileService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceProfileService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).Delete(ctx, req.(*DeleteDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceProfileService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceProfileService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).List(ctx, req.(*ListDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DeviceProfileService",
	HandlerType: (*DeviceProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DeviceProfileService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DeviceProfileService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DeviceProfileService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DeviceProfileService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DeviceProfileService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "as/external/api/deviceProfile.proto",
}

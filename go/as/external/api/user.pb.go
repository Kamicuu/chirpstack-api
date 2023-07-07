// Code generated by protoc-gen-go. DO NOT EDIT.
// source: as/external/api/user.proto

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

type User struct {
	// User ID.
	// Will be set automatically on create.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The session timeout, in minutes.
	SessionTtl int32 `protobuf:"varint,3,opt,name=session_ttl,json=sessionTTL,proto3" json:"session_ttl,omitempty"`
	// Set to true to make the user a global administrator.
	IsAdmin bool `protobuf:"varint,4,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	// Set to false to disable the user.
	IsActive bool `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// E-mail of the user.
	Email string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	// Optional note to store with the user.
	Note                 string   `protobuf:"bytes,7,opt,name=note,proto3" json:"note,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_213a97a4df8a40e1, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetSessionTtl() int32 {
	if m != nil {
		return m.SessionTtl
	}
	return 0
}

func (m *User) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *User) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

type UserListItem struct {
	// User ID.
	// Will be set automatically on create.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Email of the user.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// The session timeout, in minutes.
	SessionTtl int32 `protobuf:"varint,3,opt,name=session_ttl,json=sessionTTL,proto3" json:"session_ttl,omitempty"`
	// Set to true to make the user a global administrator.
	IsAdmin bool `protobuf:"varint,4,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	// Set to false to disable the user.
	IsActive bool `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserListItem) Reset()         { *m = UserListItem{} }
func (m *UserListItem) String() string { return proto.CompactTextString(m) }
func (*UserListItem) ProtoMessage()    {}
func (*UserListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_213a97a4df8a40e1, []int{1}
}

func (m *UserListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListItem.Unmarshal(m, b)
}
func (m *UserListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListItem.Marshal(b, m, deterministic)
}
func (m *UserListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListItem.Merge(m, src)
}
func (m *UserListItem) XXX_Size() int {
	return xxx_messageInfo_UserListItem.Size(m)
}
func (m *UserListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListItem.DiscardUnknown(m)
}

var xxx_messageInfo_UserListItem proto.InternalMessageInfo

func (m *UserListItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserListItem) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserListItem) GetSessionTtl() int32 {
	if m != nil {
		return m.SessionTtl
	}
	return 0
}

func (m *UserListItem) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *UserListItem) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *UserListItem) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *UserListItem) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type UserOrganization struct {
	// Organization ID.
	OrganizationId int64 `protobuf:"varint,1,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// User is admin within the context of the organization.
	// There is no need to set the is_device_admin and is_gateway_admin flags.
	IsAdmin bool `protobuf:"varint,2,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	// User is able to modify device related resources (applications,
	// device-profiles, devices, multicast-groups).
	IsDeviceAdmin bool `protobuf:"varint,3,opt,name=is_device_admin,json=isDeviceAdmin,proto3" json:"is_device_admin,omitempty"`
	// User is able to modify gateways.
	IsGatewayAdmin       bool     `protobuf:"varint,4,opt,name=is_gateway_admin,json=isGatewayAdmin,proto3" json:"is_gateway_admin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserOrganization) Reset()         { *m = UserOrganization{} }
func (m *UserOrganization) String() string { return proto.CompactTextString(m) }
func (*UserOrganization) ProtoMessage()    {}
func (*UserOrganization) Descriptor() ([]byte, []int) {
	return fileDescriptor_213a97a4df8a40e1, []int{2}
}

func (m *UserOrganization) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserOrganization.Unmarshal(m, b)
}
func (m *UserOrganization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserOrganization.Marshal(b, m, deterministic)
}
func (m *UserOrganization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserOrganization.Merge(m, src)
}
func (m *UserOrganization) XXX_Size() int {
	return xxx_messageInfo_UserOrganization.Size(m)
}
func (m *UserOrganization) XXX_DiscardUnknown() {
	xxx_messageInfo_UserOrganization.DiscardUnknown(m)
}

var xxx_messageInfo_UserOrganization proto.InternalMessageInfo

func (m *UserOrganization) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *UserOrganization) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *UserOrganization) GetIsDeviceAdmin() bool {
	if m != nil {
		return m.IsDeviceAdmin
	}
	return false
}

func (m *UserOrganization) GetIsGatewayAdmin() bool {
	if m != nil {
		return m.IsGatewayAdmin
	}
	return false
}

type CreateUserRequest struct {
	// User object to create.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Password of the user.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Add the user to the following organizations.
	Organizations        []*UserOrganization `protobuf:"bytes,3,rep,name=organizations,proto3" json:"organizations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_213a97a4df8a40e1, []int{3}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateUserRequest) GetOrganizations() []*UserOrganization {
	if m != nil {
		return m.Organizations
	}
	return nil
}

type CreateUserResponse struct {
	// User ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_213a97a4df8a40e1, []int{4}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUserRequest struct {
	// User ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_213a97a4df8a40e1, []int{5}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUserResponse struct {
	// User object.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_213a97a4df8a40e1, []int{6}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GetUserResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GetUserResponse) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type UpdateUserRequest struct {
	// User object to update.
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_213a97a4df8a40e1, []int{7}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type DeleteUserRequest struct {
	// User ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_213a97a4df8a40e1, []int{8}
}

func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(m, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListUserRequest struct {
	// Max number of user to return in the result-set.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUserRequest) Reset()         { *m = ListUserRequest{} }
func (m *ListUserRequest) String() string { return proto.CompactTextString(m) }
func (*ListUserRequest) ProtoMessage()    {}
func (*ListUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_213a97a4df8a40e1, []int{9}
}

func (m *ListUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserRequest.Unmarshal(m, b)
}
func (m *ListUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserRequest.Marshal(b, m, deterministic)
}
func (m *ListUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserRequest.Merge(m, src)
}
func (m *ListUserRequest) XXX_Size() int {
	return xxx_messageInfo_ListUserRequest.Size(m)
}
func (m *ListUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserRequest proto.InternalMessageInfo

func (m *ListUserRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListUserRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListUserResponse struct {
	// Total number of users.
	TotalCount int64 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// Result-set.
	Result               []*UserListItem `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListUserResponse) Reset()         { *m = ListUserResponse{} }
func (m *ListUserResponse) String() string { return proto.CompactTextString(m) }
func (*ListUserResponse) ProtoMessage()    {}
func (*ListUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_213a97a4df8a40e1, []int{10}
}

func (m *ListUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserResponse.Unmarshal(m, b)
}
func (m *ListUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserResponse.Marshal(b, m, deterministic)
}
func (m *ListUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserResponse.Merge(m, src)
}
func (m *ListUserResponse) XXX_Size() int {
	return xxx_messageInfo_ListUserResponse.Size(m)
}
func (m *ListUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserResponse proto.InternalMessageInfo

func (m *ListUserResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListUserResponse) GetResult() []*UserListItem {
	if m != nil {
		return m.Result
	}
	return nil
}

type UpdateUserPasswordRequest struct {
	// User ID.
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// New pasword.
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserPasswordRequest) Reset()         { *m = UpdateUserPasswordRequest{} }
func (m *UpdateUserPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserPasswordRequest) ProtoMessage()    {}
func (*UpdateUserPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_213a97a4df8a40e1, []int{11}
}

func (m *UpdateUserPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserPasswordRequest.Unmarshal(m, b)
}
func (m *UpdateUserPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserPasswordRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserPasswordRequest.Merge(m, src)
}
func (m *UpdateUserPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserPasswordRequest.Size(m)
}
func (m *UpdateUserPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserPasswordRequest proto.InternalMessageInfo

func (m *UpdateUserPasswordRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UpdateUserPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "api.User")
	proto.RegisterType((*UserListItem)(nil), "api.UserListItem")
	proto.RegisterType((*UserOrganization)(nil), "api.UserOrganization")
	proto.RegisterType((*CreateUserRequest)(nil), "api.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "api.CreateUserResponse")
	proto.RegisterType((*GetUserRequest)(nil), "api.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "api.GetUserResponse")
	proto.RegisterType((*UpdateUserRequest)(nil), "api.UpdateUserRequest")
	proto.RegisterType((*DeleteUserRequest)(nil), "api.DeleteUserRequest")
	proto.RegisterType((*ListUserRequest)(nil), "api.ListUserRequest")
	proto.RegisterType((*ListUserResponse)(nil), "api.ListUserResponse")
	proto.RegisterType((*UpdateUserPasswordRequest)(nil), "api.UpdateUserPasswordRequest")
}

func init() {
	proto.RegisterFile("as/external/api/user.proto", fileDescriptor_213a97a4df8a40e1)
}

var fileDescriptor_213a97a4df8a40e1 = []byte{
	// 855 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x41, 0x6f, 0xe3, 0x44,
	0x14, 0x96, 0xe3, 0x34, 0x4d, 0x5e, 0x76, 0x93, 0x66, 0x48, 0xdb, 0xd4, 0xcb, 0xd2, 0x60, 0x10,
	0x98, 0x4a, 0xd8, 0x52, 0xf6, 0x80, 0x58, 0x0e, 0xa8, 0xbb, 0x45, 0x55, 0xa5, 0x95, 0x28, 0xa6,
	0x7b, 0x80, 0x03, 0xd6, 0xd4, 0x9e, 0x66, 0x47, 0x6b, 0x7b, 0x8c, 0x67, 0xdc, 0x65, 0x41, 0xbd,
	0xc0, 0x91, 0x23, 0xe2, 0xc0, 0x1f, 0x40, 0xfc, 0x1f, 0xfe, 0x02, 0x3f, 0x04, 0xcd, 0x78, 0x9c,
	0xd8, 0x8e, 0x96, 0x02, 0xd2, 0xde, 0xfa, 0xde, 0x7c, 0xef, 0x7b, 0xdf, 0x7b, 0xf3, 0x4d, 0x5c,
	0xb0, 0x30, 0xf7, 0xc8, 0x77, 0x82, 0xe4, 0x29, 0x8e, 0x3d, 0x9c, 0x51, 0xaf, 0xe0, 0x24, 0x77,
	0xb3, 0x9c, 0x09, 0x86, 0x4c, 0x9c, 0x51, 0xeb, 0xcd, 0x25, 0x63, 0xcb, 0x98, 0xa8, 0x33, 0x9c,
	0xa6, 0x4c, 0x60, 0x41, 0x59, 0xca, 0x4b, 0x88, 0x75, 0xa8, 0x4f, 0x55, 0x74, 0x59, 0x5c, 0x79,
	0x82, 0x26, 0x84, 0x0b, 0x9c, 0x64, 0x1a, 0x70, 0xaf, 0x0d, 0x20, 0x49, 0x26, 0x5e, 0x96, 0x87,
	0xf6, 0x6f, 0x06, 0x74, 0x9f, 0x72, 0x92, 0xa3, 0x11, 0x74, 0x68, 0x34, 0x33, 0xe6, 0x86, 0x63,
	0xfa, 0x1d, 0x1a, 0xa1, 0x43, 0x18, 0x72, 0xc2, 0x39, 0x65, 0x69, 0x20, 0x44, 0x3c, 0x33, 0xe7,
	0x86, 0xb3, 0xe5, 0x83, 0x4e, 0x5d, 0x5c, 0x3c, 0x41, 0x07, 0xd0, 0xa7, 0x3c, 0xc0, 0x51, 0x42,
	0xd3, 0x59, 0x77, 0x6e, 0x38, 0x7d, 0x7f, 0x9b, 0xf2, 0x63, 0x19, 0xa2, 0x7b, 0x30, 0x90, 0x47,
	0xa1, 0xa0, 0xd7, 0x64, 0xb6, 0xa5, 0xce, 0xfa, 0x94, 0x1f, 0xab, 0x18, 0x4d, 0x61, 0x8b, 0x24,
	0x98, 0xc6, 0xb3, 0xde, 0xdc, 0x70, 0x06, 0x7e, 0x19, 0x20, 0x04, 0xdd, 0x94, 0x09, 0x32, 0xdb,
	0x56, 0x49, 0xf5, 0xb7, 0xfd, 0x53, 0x07, 0xee, 0x48, 0x6d, 0x4f, 0x28, 0x17, 0x67, 0x82, 0x24,
	0x1b, 0x1a, 0x57, 0x54, 0x9d, 0x3a, 0xd5, 0x6b, 0x53, 0xfe, 0x31, 0x40, 0x98, 0x13, 0x2c, 0x48,
	0x14, 0x60, 0x31, 0xeb, 0xcf, 0x0d, 0x67, 0xb8, 0xb0, 0xdc, 0x72, 0xbb, 0x6e, 0xb5, 0x5d, 0xf7,
	0xa2, 0x5a, 0xbf, 0x3f, 0xd0, 0xe8, 0x63, 0x21, 0x4b, 0x8b, 0x2c, 0xaa, 0x4a, 0x07, 0xb7, 0x97,
	0x6a, 0xf4, 0xb1, 0xb0, 0xff, 0x30, 0x60, 0x47, 0x6e, 0xe1, 0xf3, 0x7c, 0x89, 0x53, 0xfa, 0xbd,
	0xba, 0x7b, 0xf4, 0x3e, 0x8c, 0x59, 0x2d, 0x0e, 0x56, 0x6b, 0x19, 0xd5, 0xd3, 0x67, 0x27, 0x8d,
	0x59, 0x3b, 0xcd, 0x59, 0xdf, 0x83, 0x31, 0xe5, 0x41, 0x44, 0xae, 0x69, 0x48, 0x34, 0xc2, 0x54,
	0x88, 0xbb, 0x94, 0x9f, 0xa8, 0x6c, 0x89, 0x73, 0x60, 0x87, 0xf2, 0x60, 0x89, 0x05, 0x79, 0x81,
	0x5f, 0x36, 0xd6, 0x36, 0xa2, 0xfc, 0xb4, 0x4c, 0x2b, 0xa4, 0xfd, 0xb3, 0x01, 0x93, 0xc7, 0x6a,
	0x66, 0x29, 0xd8, 0x27, 0xdf, 0x16, 0x84, 0x0b, 0x74, 0x1f, 0xba, 0xd2, 0xd1, 0x4a, 0xe0, 0x70,
	0x31, 0x70, 0x71, 0x46, 0x5d, 0x75, 0xae, 0xd2, 0xc8, 0x82, 0x7e, 0x86, 0x39, 0x7f, 0xc1, 0xf2,
	0x48, 0xdf, 0xe3, 0x2a, 0x46, 0x9f, 0xc0, 0xdd, 0xfa, 0x3c, 0x7c, 0x66, 0xce, 0x4d, 0x67, 0xb8,
	0xd8, 0x5d, 0x71, 0xd4, 0x97, 0xe2, 0x37, 0xb1, 0xf6, 0xbb, 0x80, 0xea, 0x62, 0x78, 0xc6, 0x52,
	0x4e, 0xda, 0x1e, 0xb2, 0xe7, 0x30, 0x3a, 0x25, 0xa2, 0xae, 0xb7, 0x8d, 0xf8, 0xdd, 0x80, 0xf1,
	0x0a, 0xa2, 0x59, 0x6e, 0x99, 0xa9, 0xe9, 0x94, 0xce, 0xff, 0x77, 0x8a, 0xf9, 0x5f, 0x9c, 0xb2,
	0x80, 0xc9, 0x53, 0x15, 0xfc, 0xfb, 0xed, 0xdb, 0xef, 0xc0, 0xe4, 0x84, 0xc4, 0xa4, 0x59, 0xd3,
	0xde, 0xc0, 0xa7, 0x30, 0x96, 0x6f, 0xb0, 0x0e, 0x99, 0xc2, 0x56, 0x4c, 0x13, 0x2a, 0x34, 0xaa,
	0x0c, 0xd0, 0x1e, 0xf4, 0xd8, 0xd5, 0x15, 0x27, 0xe5, 0xcc, 0xa6, 0xaf, 0x23, 0xfb, 0x1b, 0xd8,
	0x59, 0x13, 0xe8, 0x15, 0x1e, 0xc2, 0x50, 0x30, 0x81, 0xe3, 0x20, 0x64, 0x45, 0x5a, 0xf1, 0x80,
	0x4a, 0x3d, 0x96, 0x19, 0xf4, 0x01, 0xf4, 0x72, 0xc2, 0x8b, 0x58, 0x92, 0xc9, 0x5b, 0x9f, 0xac,
	0xb4, 0x57, 0x3f, 0x08, 0xbe, 0x06, 0xd8, 0xe7, 0x70, 0xb0, 0x9e, 0xfc, 0x5c, 0xbb, 0xa7, 0x92,
	0xba, 0x0f, 0xdb, 0x72, 0xd4, 0xf5, 0x1b, 0xe9, 0xc9, 0xf0, 0x2c, 0xfa, 0x27, 0xe7, 0x2d, 0x7e,
	0xed, 0xc2, 0x50, 0x92, 0x7d, 0x49, 0x72, 0xf9, 0x12, 0xd0, 0x29, 0x74, 0x65, 0x57, 0x34, 0x55,
	0x22, 0x5a, 0xdb, 0xb0, 0x76, 0x5b, 0xd9, 0x72, 0x44, 0x1b, 0xfd, 0xf8, 0xe7, 0x5f, 0xbf, 0x74,
	0xee, 0x20, 0x58, 0xfd, 0xac, 0x73, 0x74, 0x06, 0xe6, 0x29, 0x11, 0xe8, 0x0d, 0x55, 0xd1, 0x74,
	0x9e, 0x35, 0x6d, 0x26, 0x35, 0xcb, 0xbe, 0x62, 0x99, 0xa0, 0xf1, 0x9a, 0xc5, 0xfb, 0x81, 0x46,
	0x37, 0xe8, 0x1c, 0x7a, 0xa5, 0xc1, 0xd1, 0x9e, 0x2a, 0xdc, 0x78, 0x7a, 0xd6, 0xfe, 0x46, 0x5e,
	0x73, 0xee, 0x2a, 0xce, 0xb1, 0x5d, 0x53, 0xf6, 0xd0, 0x38, 0x42, 0x5f, 0x41, 0xaf, 0xdc, 0xa3,
	0x66, 0xdc, 0xb0, 0x93, 0xb5, 0xb7, 0x61, 0xc5, 0xcf, 0xe4, 0xd7, 0xc4, 0x3e, 0x54, 0x84, 0x07,
	0xd6, 0xb4, 0x2e, 0x52, 0x7d, 0xc8, 0x68, 0x74, 0x23, 0xa9, 0xbf, 0x80, 0x5e, 0x69, 0x34, 0x4d,
	0xbd, 0xe1, 0xba, 0x57, 0x52, 0xeb, 0xf9, 0x8f, 0x36, 0xe6, 0xcf, 0x61, 0x54, 0x0a, 0xac, 0x6e,
	0x1c, 0xbd, 0xd5, 0x52, 0xdd, 0xb2, 0xc2, 0x2b, 0x5b, 0x38, 0xaa, 0x85, 0x6d, 0xdd, 0x6f, 0xab,
	0x0f, 0x68, 0x74, 0xe3, 0x55, 0xa6, 0x78, 0x68, 0x1c, 0x3d, 0xa2, 0xf0, 0x36, 0x65, 0x6e, 0xf8,
	0x8c, 0xe6, 0x19, 0x17, 0x38, 0x7c, 0xae, 0x1a, 0x62, 0xee, 0x56, 0x1f, 0x70, 0x19, 0x3f, 0x1a,
	0xa8, 0xde, 0xb2, 0xc5, 0xb9, 0xf1, 0xf5, 0x47, 0x4b, 0x2a, 0x9e, 0x15, 0x97, 0x6e, 0xc8, 0x12,
	0xef, 0x39, 0x4e, 0x68, 0x58, 0x14, 0xde, 0xba, 0xfe, 0x43, 0xd9, 0x70, 0xc9, 0xbc, 0xeb, 0x07,
	0x5e, 0xeb, 0xdf, 0x80, 0xcb, 0x9e, 0x12, 0xf9, 0xe0, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73,
	0x26, 0x05, 0x33, 0x20, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	// Get user list.
	List(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
	// Get data for a particular user.
	Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	// Create a new user.
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// Update an existing user.
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete a user.
	Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// UpdatePassword updates a password.
	UpdatePassword(ctx context.Context, in *UpdateUserPasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) List(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	out := new(ListUserResponse)
	err := c.cc.Invoke(ctx, "/api.UserService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/api.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/api.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.UserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePassword(ctx context.Context, in *UpdateUserPasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.UserService/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	// Get user list.
	List(context.Context, *ListUserRequest) (*ListUserResponse, error)
	// Get data for a particular user.
	Get(context.Context, *GetUserRequest) (*GetUserResponse, error)
	// Create a new user.
	Create(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// Update an existing user.
	Update(context.Context, *UpdateUserRequest) (*empty.Empty, error)
	// Delete a user.
	Delete(context.Context, *DeleteUserRequest) (*empty.Empty, error)
	// UpdatePassword updates a password.
	UpdatePassword(context.Context, *UpdateUserPasswordRequest) (*empty.Empty, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) List(ctx context.Context, req *ListUserRequest) (*ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedUserServiceServer) Get(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedUserServiceServer) Create(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedUserServiceServer) Update(ctx context.Context, req *UpdateUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedUserServiceServer) Delete(ctx context.Context, req *DeleteUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedUserServiceServer) UpdatePassword(ctx context.Context, req *UpdateUserPasswordRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).List(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePassword(ctx, req.(*UpdateUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _UserService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _UserService_UpdatePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "as/external/api/user.proto",
}

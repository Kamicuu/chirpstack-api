# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chirpstack-api/as_pb/external/api/user.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n,chirpstack-api/as_pb/external/api/user.proto\x12\x03\x61pi\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1bgoogle/protobuf/empty.proto\"u\n\x04User\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x1f\n\x0bsession_ttl\x18\x03 \x01(\x05R\nsessionTTL\x12\x10\n\x08is_admin\x18\x04 \x01(\x08\x12\x11\n\tis_active\x18\x05 \x01(\x08\x12\r\n\x05\x65mail\x18\x06 \x01(\t\x12\x0c\n\x04note\x18\x07 \x01(\t\"\xcf\x01\n\x0cUserListItem\x12\n\n\x02id\x18\x01 \x01(\x03\x12\r\n\x05\x65mail\x18\x02 \x01(\t\x12\x1f\n\x0bsession_ttl\x18\x03 \x01(\x05R\nsessionTTL\x12\x10\n\x08is_admin\x18\x04 \x01(\x08\x12\x11\n\tis_active\x18\x05 \x01(\x08\x12.\n\ncreated_at\x18\x08 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\t \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\x80\x01\n\x10UserOrganization\x12\'\n\x0forganization_id\x18\x01 \x01(\x03R\x0eorganizationID\x12\x10\n\x08is_admin\x18\x02 \x01(\x08\x12\x17\n\x0fis_device_admin\x18\x03 \x01(\x08\x12\x18\n\x10is_gateway_admin\x18\x04 \x01(\x08\"l\n\x11\x43reateUserRequest\x12\x17\n\x04user\x18\x01 \x01(\x0b\x32\t.api.User\x12\x10\n\x08password\x18\x02 \x01(\t\x12,\n\rorganizations\x18\x03 \x03(\x0b\x32\x15.api.UserOrganization\" \n\x12\x43reateUserResponse\x12\n\n\x02id\x18\x01 \x01(\x03\"\x1c\n\x0eGetUserRequest\x12\n\n\x02id\x18\x01 \x01(\x03\"\x8a\x01\n\x0fGetUserResponse\x12\x17\n\x04user\x18\x01 \x01(\x0b\x32\t.api.User\x12.\n\ncreated_at\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\",\n\x11UpdateUserRequest\x12\x17\n\x04user\x18\x01 \x01(\x0b\x32\t.api.User\"\x1f\n\x11\x44\x65leteUserRequest\x12\n\n\x02id\x18\x01 \x01(\x03\"0\n\x0fListUserRequest\x12\r\n\x05limit\x18\x01 \x01(\x03\x12\x0e\n\x06offset\x18\x02 \x01(\x03\"J\n\x10ListUserResponse\x12\x13\n\x0btotal_count\x18\x01 \x01(\x03\x12!\n\x06result\x18\x02 \x03(\x0b\x32\x11.api.UserListItem\">\n\x19UpdateUserPasswordRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\x03\x12\x10\n\x08password\x18\x02 \x01(\t2\x95\x04\n\x0bUserService\x12G\n\x04List\x12\x14.api.ListUserRequest\x1a\x15.api.ListUserResponse\"\x12\x82\xd3\xe4\x93\x02\x0c\x12\n/api/users\x12I\n\x03Get\x12\x13.api.GetUserRequest\x1a\x14.api.GetUserResponse\"\x17\x82\xd3\xe4\x93\x02\x11\x12\x0f/api/users/{id}\x12P\n\x06\x43reate\x12\x16.api.CreateUserRequest\x1a\x17.api.CreateUserResponse\"\x15\x82\xd3\xe4\x93\x02\x0f\"\n/api/users:\x01*\x12Y\n\x06Update\x12\x16.api.UpdateUserRequest\x1a\x16.google.protobuf.Empty\"\x1f\x82\xd3\xe4\x93\x02\x19\x1a\x14/api/users/{user.id}:\x01*\x12Q\n\x06\x44\x65lete\x12\x16.api.DeleteUserRequest\x1a\x16.google.protobuf.Empty\"\x17\x82\xd3\xe4\x93\x02\x11*\x0f/api/users/{id}\x12r\n\x0eUpdatePassword\x12\x1e.api.UpdateUserPasswordRequest\x1a\x16.google.protobuf.Empty\"(\x82\xd3\xe4\x93\x02\"\x1a\x1d/api/users/{user_id}/password:\x01*Bi\n!io.chirpstack.api.as.external.apiB\tUserProtoP\x01Z7github.com/kamicuu/chirpstack-api/go/v3/as/external/apib\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'chirpstack_api.as_pb.external.api.user_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n!io.chirpstack.api.as.external.apiB\tUserProtoP\001Z7github.com/kamicuu/chirpstack-api/go/v3/as/external/api'
  _USERSERVICE.methods_by_name['List']._options = None
  _USERSERVICE.methods_by_name['List']._serialized_options = b'\202\323\344\223\002\014\022\n/api/users'
  _USERSERVICE.methods_by_name['Get']._options = None
  _USERSERVICE.methods_by_name['Get']._serialized_options = b'\202\323\344\223\002\021\022\017/api/users/{id}'
  _USERSERVICE.methods_by_name['Create']._options = None
  _USERSERVICE.methods_by_name['Create']._serialized_options = b'\202\323\344\223\002\017\"\n/api/users:\001*'
  _USERSERVICE.methods_by_name['Update']._options = None
  _USERSERVICE.methods_by_name['Update']._serialized_options = b'\202\323\344\223\002\031\032\024/api/users/{user.id}:\001*'
  _USERSERVICE.methods_by_name['Delete']._options = None
  _USERSERVICE.methods_by_name['Delete']._serialized_options = b'\202\323\344\223\002\021*\017/api/users/{id}'
  _USERSERVICE.methods_by_name['UpdatePassword']._options = None
  _USERSERVICE.methods_by_name['UpdatePassword']._serialized_options = b'\202\323\344\223\002\"\032\035/api/users/{user_id}/password:\001*'
  _USER._serialized_start=145
  _USER._serialized_end=262
  _USERLISTITEM._serialized_start=265
  _USERLISTITEM._serialized_end=472
  _USERORGANIZATION._serialized_start=475
  _USERORGANIZATION._serialized_end=603
  _CREATEUSERREQUEST._serialized_start=605
  _CREATEUSERREQUEST._serialized_end=713
  _CREATEUSERRESPONSE._serialized_start=715
  _CREATEUSERRESPONSE._serialized_end=747
  _GETUSERREQUEST._serialized_start=749
  _GETUSERREQUEST._serialized_end=777
  _GETUSERRESPONSE._serialized_start=780
  _GETUSERRESPONSE._serialized_end=918
  _UPDATEUSERREQUEST._serialized_start=920
  _UPDATEUSERREQUEST._serialized_end=964
  _DELETEUSERREQUEST._serialized_start=966
  _DELETEUSERREQUEST._serialized_end=997
  _LISTUSERREQUEST._serialized_start=999
  _LISTUSERREQUEST._serialized_end=1047
  _LISTUSERRESPONSE._serialized_start=1049
  _LISTUSERRESPONSE._serialized_end=1123
  _UPDATEUSERPASSWORDREQUEST._serialized_start=1125
  _UPDATEUSERPASSWORDREQUEST._serialized_end=1187
  _USERSERVICE._serialized_start=1190
  _USERSERVICE._serialized_end=1723
# @@protoc_insertion_point(module_scope)

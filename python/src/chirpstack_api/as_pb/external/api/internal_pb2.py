# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chirpstack-api/as_pb/external/api/internal.proto
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
from chirpstack_api.as_pb.external.api import user_pb2 as chirpstack__api_dot_as__pb_dot_external_dot_api_dot_user__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n0chirpstack-api/as_pb/external/api/internal.proto\x12\x03\x61pi\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a,chirpstack-api/as_pb/external/api/user.proto\"\x84\x01\n\x06\x41PIKey\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x10\n\x08is_admin\x18\x03 \x01(\x08\x12\'\n\x0forganization_id\x18\x04 \x01(\x03R\x0eorganizationID\x12%\n\x0e\x61pplication_id\x18\x05 \x01(\x03R\rapplicationID\"3\n\x13\x43reateAPIKeyRequest\x12\x1c\n\x07\x61pi_key\x18\x01 \x01(\x0b\x32\x0b.api.APIKey\"5\n\x14\x43reateAPIKeyResponse\x12\n\n\x02id\x18\x01 \x01(\t\x12\x11\n\tjwt_token\x18\x02 \x01(\t\"!\n\x13\x44\x65leteAPIKeyRequest\x12\n\n\x02id\x18\x01 \x01(\t\"\x95\x01\n\x12ListAPIKeysRequest\x12\r\n\x05limit\x18\x01 \x01(\x03\x12\x0e\n\x06offset\x18\x02 \x01(\x03\x12\x10\n\x08is_admin\x18\x03 \x01(\x08\x12\'\n\x0forganization_id\x18\x04 \x01(\x03R\x0eorganizationID\x12%\n\x0e\x61pplication_id\x18\x05 \x01(\x03R\rapplicationID\"G\n\x13ListAPIKeysResponse\x12\x13\n\x0btotal_count\x18\x01 \x01(\x03\x12\x1b\n\x06result\x18\x02 \x03(\x0b\x32\x0b.api.APIKey\"\xfb\x01\n\x10OrganizationLink\x12\'\n\x0forganization_id\x18\x01 \x01(\x03R\x0eorganizationID\x12\x19\n\x11organization_name\x18\x02 \x01(\t\x12\x10\n\x08is_admin\x18\x03 \x01(\x08\x12\x17\n\x0fis_device_admin\x18\x06 \x01(\x08\x12\x18\n\x10is_gateway_admin\x18\x07 \x01(\x08\x12.\n\ncreated_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"/\n\x0cLoginRequest\x12\r\n\x05\x65mail\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\"\x1c\n\rLoginResponse\x12\x0b\n\x03jwt\x18\x01 \x01(\t\"X\n\x0fProfileResponse\x12\x17\n\x04user\x18\x01 \x01(\x0b\x32\t.api.User\x12,\n\rorganizations\x18\x03 \x03(\x0b\x32\x15.api.OrganizationLink\"D\n\x13GlobalSearchRequest\x12\x0e\n\x06search\x18\x01 \x01(\t\x12\r\n\x05limit\x18\x02 \x01(\x03\x12\x0e\n\x06offset\x18\x03 \x01(\x03\"?\n\x14GlobalSearchResponse\x12\'\n\x06result\x18\x01 \x03(\x0b\x32\x17.api.GlobalSearchResult\"\xa8\x02\n\x12GlobalSearchResult\x12\x0c\n\x04kind\x18\x01 \x01(\t\x12\r\n\x05score\x18\x02 \x01(\x02\x12\'\n\x0forganization_id\x18\x03 \x01(\x03R\x0eorganizationID\x12\x19\n\x11organization_name\x18\x04 \x01(\t\x12%\n\x0e\x61pplication_id\x18\x05 \x01(\x03R\rapplicationID\x12\x18\n\x10\x61pplication_name\x18\x06 \x01(\t\x12$\n\x0e\x64\x65vice_dev_eui\x18\x07 \x01(\tR\x0c\x64\x65viceDevEUI\x12\x13\n\x0b\x64\x65vice_name\x18\x08 \x01(\t\x12\x1f\n\x0bgateway_mac\x18\t \x01(\tR\ngatewayMAC\x12\x14\n\x0cgateway_name\x18\n \x01(\t\"_\n\x10SettingsResponse\x12\x1f\n\x08\x62randing\x18\x02 \x01(\x0b\x32\r.api.Branding\x12*\n\x0eopenid_connect\x18\x03 \x01(\x0b\x32\x12.api.OpenIDConnect\"0\n\x08\x42randing\x12\x14\n\x0cregistration\x18\x01 \x01(\t\x12\x0e\n\x06\x66ooter\x18\x02 \x01(\t\"q\n\rOpenIDConnect\x12\x0f\n\x07\x65nabled\x18\x01 \x01(\x08\x12\x1b\n\tlogin_url\x18\x02 \x01(\tR\x08loginURL\x12\x13\n\x0blogin_label\x18\x03 \x01(\t\x12\x1d\n\nlogout_url\x18\x04 \x01(\tR\tlogoutURL\"8\n\x19OpenIDConnectLoginRequest\x12\x0c\n\x04\x63ode\x18\x01 \x01(\t\x12\r\n\x05state\x18\x02 \x01(\t\"/\n\x1aOpenIDConnectLoginResponse\x12\x11\n\tjwt_token\x18\x01 \x01(\t\"C\n\x18GetDevicesSummaryRequest\x12\'\n\x0forganization_id\x18\x01 \x01(\x03R\x0eorganizationID\"\xd2\x01\n\x19GetDevicesSummaryResponse\x12\x14\n\x0c\x61\x63tive_count\x18\x01 \x01(\r\x12\x16\n\x0einactive_count\x18\x02 \x01(\r\x12=\n\x08\x64r_count\x18\x03 \x03(\x0b\x32+.api.GetDevicesSummaryResponse.DrCountEntry\x12\x18\n\x10never_seen_count\x18\x04 \x01(\r\x1a.\n\x0c\x44rCountEntry\x12\x0b\n\x03key\x18\x01 \x01(\r\x12\r\n\x05value\x18\x02 \x01(\r:\x02\x38\x01\"D\n\x19GetGatewaysSummaryRequest\x12\'\n\x0forganization_id\x18\x01 \x01(\x03R\x0eorganizationID\"d\n\x1aGetGatewaysSummaryResponse\x12\x14\n\x0c\x61\x63tive_count\x18\x01 \x01(\r\x12\x16\n\x0einactive_count\x18\x02 \x01(\r\x12\x18\n\x10never_seen_count\x18\x03 \x01(\r2\x9b\x08\n\x0fInternalService\x12N\n\x05Login\x12\x11.api.LoginRequest\x1a\x12.api.LoginResponse\"\x1e\x82\xd3\xe4\x93\x02\x18\"\x13/api/internal/login:\x01*\x12V\n\x07Profile\x12\x16.google.protobuf.Empty\x1a\x14.api.ProfileResponse\"\x1d\x82\xd3\xe4\x93\x02\x17\x12\x15/api/internal/profile\x12\x61\n\x0cGlobalSearch\x12\x18.api.GlobalSearchRequest\x1a\x19.api.GlobalSearchResponse\"\x1c\x82\xd3\xe4\x93\x02\x16\x12\x14/api/internal/search\x12\x66\n\x0c\x43reateAPIKey\x12\x18.api.CreateAPIKeyRequest\x1a\x19.api.CreateAPIKeyResponse\"!\x82\xd3\xe4\x93\x02\x1b\"\x16/api/internal/api-keys:\x01*\x12\x65\n\x0c\x44\x65leteAPIKey\x12\x18.api.DeleteAPIKeyRequest\x1a\x16.google.protobuf.Empty\"#\x82\xd3\xe4\x93\x02\x1d*\x1b/api/internal/api-keys/{id}\x12`\n\x0bListAPIKeys\x12\x17.api.ListAPIKeysRequest\x1a\x18.api.ListAPIKeysResponse\"\x1e\x82\xd3\xe4\x93\x02\x18\x12\x16/api/internal/api-keys\x12Y\n\x08Settings\x12\x16.google.protobuf.Empty\x1a\x15.api.SettingsResponse\"\x1e\x82\xd3\xe4\x93\x02\x18\x12\x16/api/internal/settings\x12w\n\x12OpenIDConnectLogin\x12\x1e.api.OpenIDConnectLoginRequest\x1a\x1f.api.OpenIDConnectLoginResponse\" \x82\xd3\xe4\x93\x02\x1a\x12\x18/api/internal/oidc/login\x12y\n\x11GetDevicesSummary\x12\x1d.api.GetDevicesSummaryRequest\x1a\x1e.api.GetDevicesSummaryResponse\"%\x82\xd3\xe4\x93\x02\x1f\x12\x1d/api/internal/devices/summary\x12}\n\x12GetGatewaysSummary\x12\x1e.api.GetGatewaysSummaryRequest\x1a\x1f.api.GetGatewaysSummaryResponse\"&\x82\xd3\xe4\x93\x02 \x12\x1e/api/internal/gateways/summaryBm\n!io.chirpstack.api.as.external.apiB\rInternalProtoP\x01Z7github.com/brocaar/chirpstack-api/go/v3/as/external/apib\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'chirpstack_api.as_pb.external.api.internal_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n!io.chirpstack.api.as.external.apiB\rInternalProtoP\001Z7github.com/brocaar/chirpstack-api/go/v3/as/external/api'
  _GETDEVICESSUMMARYRESPONSE_DRCOUNTENTRY._options = None
  _GETDEVICESSUMMARYRESPONSE_DRCOUNTENTRY._serialized_options = b'8\001'
  _INTERNALSERVICE.methods_by_name['Login']._options = None
  _INTERNALSERVICE.methods_by_name['Login']._serialized_options = b'\202\323\344\223\002\030\"\023/api/internal/login:\001*'
  _INTERNALSERVICE.methods_by_name['Profile']._options = None
  _INTERNALSERVICE.methods_by_name['Profile']._serialized_options = b'\202\323\344\223\002\027\022\025/api/internal/profile'
  _INTERNALSERVICE.methods_by_name['GlobalSearch']._options = None
  _INTERNALSERVICE.methods_by_name['GlobalSearch']._serialized_options = b'\202\323\344\223\002\026\022\024/api/internal/search'
  _INTERNALSERVICE.methods_by_name['CreateAPIKey']._options = None
  _INTERNALSERVICE.methods_by_name['CreateAPIKey']._serialized_options = b'\202\323\344\223\002\033\"\026/api/internal/api-keys:\001*'
  _INTERNALSERVICE.methods_by_name['DeleteAPIKey']._options = None
  _INTERNALSERVICE.methods_by_name['DeleteAPIKey']._serialized_options = b'\202\323\344\223\002\035*\033/api/internal/api-keys/{id}'
  _INTERNALSERVICE.methods_by_name['ListAPIKeys']._options = None
  _INTERNALSERVICE.methods_by_name['ListAPIKeys']._serialized_options = b'\202\323\344\223\002\030\022\026/api/internal/api-keys'
  _INTERNALSERVICE.methods_by_name['Settings']._options = None
  _INTERNALSERVICE.methods_by_name['Settings']._serialized_options = b'\202\323\344\223\002\030\022\026/api/internal/settings'
  _INTERNALSERVICE.methods_by_name['OpenIDConnectLogin']._options = None
  _INTERNALSERVICE.methods_by_name['OpenIDConnectLogin']._serialized_options = b'\202\323\344\223\002\032\022\030/api/internal/oidc/login'
  _INTERNALSERVICE.methods_by_name['GetDevicesSummary']._options = None
  _INTERNALSERVICE.methods_by_name['GetDevicesSummary']._serialized_options = b'\202\323\344\223\002\037\022\035/api/internal/devices/summary'
  _INTERNALSERVICE.methods_by_name['GetGatewaysSummary']._options = None
  _INTERNALSERVICE.methods_by_name['GetGatewaysSummary']._serialized_options = b'\202\323\344\223\002 \022\036/api/internal/gateways/summary'
  _APIKEY._serialized_start=196
  _APIKEY._serialized_end=328
  _CREATEAPIKEYREQUEST._serialized_start=330
  _CREATEAPIKEYREQUEST._serialized_end=381
  _CREATEAPIKEYRESPONSE._serialized_start=383
  _CREATEAPIKEYRESPONSE._serialized_end=436
  _DELETEAPIKEYREQUEST._serialized_start=438
  _DELETEAPIKEYREQUEST._serialized_end=471
  _LISTAPIKEYSREQUEST._serialized_start=474
  _LISTAPIKEYSREQUEST._serialized_end=623
  _LISTAPIKEYSRESPONSE._serialized_start=625
  _LISTAPIKEYSRESPONSE._serialized_end=696
  _ORGANIZATIONLINK._serialized_start=699
  _ORGANIZATIONLINK._serialized_end=950
  _LOGINREQUEST._serialized_start=952
  _LOGINREQUEST._serialized_end=999
  _LOGINRESPONSE._serialized_start=1001
  _LOGINRESPONSE._serialized_end=1029
  _PROFILERESPONSE._serialized_start=1031
  _PROFILERESPONSE._serialized_end=1119
  _GLOBALSEARCHREQUEST._serialized_start=1121
  _GLOBALSEARCHREQUEST._serialized_end=1189
  _GLOBALSEARCHRESPONSE._serialized_start=1191
  _GLOBALSEARCHRESPONSE._serialized_end=1254
  _GLOBALSEARCHRESULT._serialized_start=1257
  _GLOBALSEARCHRESULT._serialized_end=1553
  _SETTINGSRESPONSE._serialized_start=1555
  _SETTINGSRESPONSE._serialized_end=1650
  _BRANDING._serialized_start=1652
  _BRANDING._serialized_end=1700
  _OPENIDCONNECT._serialized_start=1702
  _OPENIDCONNECT._serialized_end=1815
  _OPENIDCONNECTLOGINREQUEST._serialized_start=1817
  _OPENIDCONNECTLOGINREQUEST._serialized_end=1873
  _OPENIDCONNECTLOGINRESPONSE._serialized_start=1875
  _OPENIDCONNECTLOGINRESPONSE._serialized_end=1922
  _GETDEVICESSUMMARYREQUEST._serialized_start=1924
  _GETDEVICESSUMMARYREQUEST._serialized_end=1991
  _GETDEVICESSUMMARYRESPONSE._serialized_start=1994
  _GETDEVICESSUMMARYRESPONSE._serialized_end=2204
  _GETDEVICESSUMMARYRESPONSE_DRCOUNTENTRY._serialized_start=2158
  _GETDEVICESSUMMARYRESPONSE_DRCOUNTENTRY._serialized_end=2204
  _GETGATEWAYSSUMMARYREQUEST._serialized_start=2206
  _GETGATEWAYSSUMMARYREQUEST._serialized_end=2274
  _GETGATEWAYSSUMMARYRESPONSE._serialized_start=2276
  _GETGATEWAYSSUMMARYRESPONSE._serialized_end=2376
  _INTERNALSERVICE._serialized_start=2379
  _INTERNALSERVICE._serialized_end=3430
# @@protoc_insertion_point(module_scope)

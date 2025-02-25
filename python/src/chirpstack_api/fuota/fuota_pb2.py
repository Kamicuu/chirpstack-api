# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chirpstack-api/fuota/fuota.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n chirpstack-api/fuota/fuota.proto\x12\x05\x66uota\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/duration.proto\"8\n\x10\x44\x65ploymentDevice\x12\x0f\n\x07\x64\x65v_eui\x18\x01 \x01(\x0c\x12\x13\n\x0bmc_root_key\x18\x02 \x01(\x0c\"\xa4\x05\n\nDeployment\x12\x16\n\x0e\x61pplication_id\x18\x01 \x01(\x03\x12(\n\x07\x64\x65vices\x18\x02 \x03(\x0b\x32\x17.fuota.DeploymentDevice\x12\x37\n\x14multicast_group_type\x18\x03 \x01(\x0e\x32\x19.fuota.MulticastGroupType\x12\x14\n\x0cmulticast_dr\x18\x04 \x01(\r\x12\"\n\x1amulticast_ping_slot_period\x18\x05 \x01(\r\x12\x1b\n\x13multicast_frequency\x18\x06 \x01(\r\x12\x1a\n\x12multicast_group_id\x18\x07 \x01(\r\x12\x19\n\x11multicast_timeout\x18\x08 \x01(\r\x12\x32\n\x0funicast_timeout\x18\t \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x1d\n\x15unicast_attempt_count\x18\n \x01(\r\x12#\n\x1b\x66ragmentation_fragment_size\x18\x0b \x01(\r\x12\x0f\n\x07payload\x18\x0c \x01(\x0c\x12 \n\x18\x66ragmentation_redundancy\x18\r \x01(\r\x12#\n\x1b\x66ragmentation_session_index\x18\x0e \x01(\r\x12\x1c\n\x14\x66ragmentation_matrix\x18\x0f \x01(\r\x12%\n\x1d\x66ragmentation_block_ack_delay\x18\x10 \x01(\r\x12 \n\x18\x66ragmentation_descriptor\x18\x11 \x01(\x0c\x12V\n$request_fragmentation_session_status\x18\x12 \x01(\x0e\x32(.fuota.RequestFragmentationSessionStatus\"@\n\x17\x43reateDeploymentRequest\x12%\n\ndeployment\x18\x01 \x01(\x0b\x32\x11.fuota.Deployment\"&\n\x18\x43reateDeploymentResponse\x12\n\n\x02id\x18\x01 \x01(\x0c\"(\n\x1aGetDeploymentStatusRequest\x12\n\n\x02id\x18\x01 \x01(\x0c\"\x8a\x03\n\x16\x44\x65ploymentDeviceStatus\x12\x0f\n\x07\x64\x65v_eui\x18\x01 \x01(\x0c\x12.\n\ncreated_at\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12?\n\x1bmc_group_setup_completed_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12;\n\x17mc_session_completed_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x43\n\x1f\x66rag_session_setup_completed_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12<\n\x18\x66rag_status_completed_at\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\xee\x03\n\x1bGetDeploymentStatusResponse\x12.\n\ncreated_at\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12?\n\x1bmc_group_setup_completed_at\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12;\n\x17mc_session_completed_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x43\n\x1f\x66rag_session_setup_completed_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x38\n\x14\x65nqueue_completed_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12<\n\x18\x66rag_status_completed_at\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x34\n\rdevice_status\x18\x08 \x03(\x0b\x32\x1d.fuota.DeploymentDeviceStatus\"H\n\x1eGetDeploymentDeviceLogsRequest\x12\x15\n\rdeployment_id\x18\x01 \x01(\x0c\x12\x0f\n\x07\x64\x65v_eui\x18\x02 \x01(\x0c\"\xcd\x01\n\x13\x44\x65ploymentDeviceLog\x12.\n\ncreated_at\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x0e\n\x06\x66_port\x18\x02 \x01(\r\x12\x0f\n\x07\x63ommand\x18\x03 \x01(\t\x12\x36\n\x06\x66ields\x18\x04 \x03(\x0b\x32&.fuota.DeploymentDeviceLog.FieldsEntry\x1a-\n\x0b\x46ieldsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"K\n\x1fGetDeploymentDeviceLogsResponse\x12(\n\x04logs\x18\x01 \x03(\x0b\x32\x1a.fuota.DeploymentDeviceLog*.\n\x12MulticastGroupType\x12\x0b\n\x07\x43LASS_B\x10\x00\x12\x0b\n\x07\x43LASS_C\x10\x01*j\n!RequestFragmentationSessionStatus\x12\x1a\n\x16\x41\x46TER_FRAGMENT_ENQUEUE\x10\x00\x12\x19\n\x15\x41\x46TER_SESSION_TIMEOUT\x10\x01\x12\x0e\n\nNO_REQUEST\x10\x02\x32\xb7\x02\n\x12\x46UOTAServerService\x12U\n\x10\x43reateDeployment\x12\x1e.fuota.CreateDeploymentRequest\x1a\x1f.fuota.CreateDeploymentResponse\"\x00\x12^\n\x13GetDeploymentStatus\x12!.fuota.GetDeploymentStatusRequest\x1a\".fuota.GetDeploymentStatusResponse\"\x00\x12j\n\x17GetDeploymentDeviceLogs\x12%.fuota.GetDeploymentDeviceLogsRequest\x1a&.fuota.GetDeploymentDeviceLogsResponse\"\x00\x42H\n\x17io.chirpstack.api.fuotaZ-github.com/kamicuu/chirpstack-api/go/v3/fuotab\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'chirpstack_api.fuota.fuota_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\027io.chirpstack.api.fuotaZ-github.com/kamicuu/chirpstack-api/go/v3/fuota'
  _DEPLOYMENTDEVICELOG_FIELDSENTRY._options = None
  _DEPLOYMENTDEVICELOG_FIELDSENTRY._serialized_options = b'8\001'
  _MULTICASTGROUPTYPE._serialized_start=2246
  _MULTICASTGROUPTYPE._serialized_end=2292
  _REQUESTFRAGMENTATIONSESSIONSTATUS._serialized_start=2294
  _REQUESTFRAGMENTATIONSESSIONSTATUS._serialized_end=2400
  _DEPLOYMENTDEVICE._serialized_start=108
  _DEPLOYMENTDEVICE._serialized_end=164
  _DEPLOYMENT._serialized_start=167
  _DEPLOYMENT._serialized_end=843
  _CREATEDEPLOYMENTREQUEST._serialized_start=845
  _CREATEDEPLOYMENTREQUEST._serialized_end=909
  _CREATEDEPLOYMENTRESPONSE._serialized_start=911
  _CREATEDEPLOYMENTRESPONSE._serialized_end=949
  _GETDEPLOYMENTSTATUSREQUEST._serialized_start=951
  _GETDEPLOYMENTSTATUSREQUEST._serialized_end=991
  _DEPLOYMENTDEVICESTATUS._serialized_start=994
  _DEPLOYMENTDEVICESTATUS._serialized_end=1388
  _GETDEPLOYMENTSTATUSRESPONSE._serialized_start=1391
  _GETDEPLOYMENTSTATUSRESPONSE._serialized_end=1885
  _GETDEPLOYMENTDEVICELOGSREQUEST._serialized_start=1887
  _GETDEPLOYMENTDEVICELOGSREQUEST._serialized_end=1959
  _DEPLOYMENTDEVICELOG._serialized_start=1962
  _DEPLOYMENTDEVICELOG._serialized_end=2167
  _DEPLOYMENTDEVICELOG_FIELDSENTRY._serialized_start=2122
  _DEPLOYMENTDEVICELOG_FIELDSENTRY._serialized_end=2167
  _GETDEPLOYMENTDEVICELOGSRESPONSE._serialized_start=2169
  _GETDEPLOYMENTDEVICELOGSRESPONSE._serialized_end=2244
  _FUOTASERVERSERVICE._serialized_start=2403
  _FUOTASERVERSERVICE._serialized_end=2714
# @@protoc_insertion_point(module_scope)

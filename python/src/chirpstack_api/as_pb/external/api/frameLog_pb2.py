# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chirpstack-api/as_pb/external/api/frameLog.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from chirpstack_api.common import common_pb2 as chirpstack__api_dot_common_dot_common__pb2
from chirpstack_api.gw import gw_pb2 as chirpstack__api_dot_gw_dot_gw__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n0chirpstack-api/as_pb/external/api/frameLog.proto\x12\x03\x61pi\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\"chirpstack-api/common/common.proto\x1a\x1a\x63hirpstack-api/gw/gw.proto\"\xb2\x01\n\x0eUplinkFrameLog\x12!\n\x07tx_info\x18\x01 \x01(\x0b\x32\x10.gw.UplinkTXInfo\x12!\n\x07rx_info\x18\x02 \x03(\x0b\x32\x10.gw.UplinkRXInfo\x12(\n\x10phy_payload_json\x18\x03 \x01(\tR\x0ephyPayloadJSON\x12\x30\n\x0cpublished_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\xb2\x01\n\x10\x44ownlinkFrameLog\x12#\n\x07tx_info\x18\x01 \x01(\x0b\x32\x12.gw.DownlinkTXInfo\x12(\n\x10phy_payload_json\x18\x02 \x01(\tR\x0ephyPayloadJSON\x12\x1d\n\ngateway_id\x18\x03 \x01(\tR\tgatewayID\x12\x30\n\x0cpublished_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp*\x1c\n\x08RXWindow\x12\x07\n\x03RX1\x10\x00\x12\x07\n\x03RX2\x10\x01\x42m\n!io.chirpstack.api.as.external.apiB\rFrameLogProtoP\x01Z7github.com/brocaar/chirpstack-api/go/v3/as/external/apib\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'chirpstack_api.as_pb.external.api.frameLog_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n!io.chirpstack.api.as.external.apiB\rFrameLogProtoP\001Z7github.com/brocaar/chirpstack-api/go/v3/as/external/api'
  _RXWINDOW._serialized_start=548
  _RXWINDOW._serialized_end=576
  _UPLINKFRAMELOG._serialized_start=187
  _UPLINKFRAMELOG._serialized_end=365
  _DOWNLINKFRAMELOG._serialized_start=368
  _DOWNLINKFRAMELOG._serialized_end=546
# @@protoc_insertion_point(module_scope)

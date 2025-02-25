# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chirpstack-api/nc/nc.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from chirpstack_api.gw import gw_pb2 as chirpstack__api_dot_gw_dot_gw__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1a\x63hirpstack-api/nc/nc.proto\x12\x02nc\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1a\x63hirpstack-api/gw/gw.proto\"\xfd\x01\n\x1bHandleUplinkMetaDataRequest\x12\x0f\n\x07\x64\x65v_eui\x18\x01 \x01(\x0c\x12!\n\x07tx_info\x18\x02 \x01(\x0b\x32\x10.gw.UplinkTXInfo\x12!\n\x07rx_info\x18\x03 \x03(\x0b\x32\x10.gw.UplinkRXInfo\x12\x1e\n\x16phy_payload_byte_count\x18\x04 \x01(\r\x12\x1e\n\x16mac_command_byte_count\x18\x05 \x01(\r\x12&\n\x1e\x61pplication_payload_byte_count\x18\x06 \x01(\r\x12\x1f\n\x0cmessage_type\x18\x07 \x01(\x0e\x32\t.nc.MType\"\x8e\x02\n\x1dHandleDownlinkMetaDataRequest\x12\x0f\n\x07\x64\x65v_eui\x18\x01 \x01(\x0c\x12\x1a\n\x12multicast_group_id\x18\x02 \x01(\x0c\x12#\n\x07tx_info\x18\x03 \x01(\x0b\x32\x12.gw.DownlinkTXInfo\x12\x1e\n\x16phy_payload_byte_count\x18\x04 \x01(\r\x12\x1e\n\x16mac_command_byte_count\x18\x05 \x01(\r\x12&\n\x1e\x61pplication_payload_byte_count\x18\x06 \x01(\r\x12\x1f\n\x0cmessage_type\x18\x07 \x01(\x0e\x32\t.nc.MType\x12\x12\n\ngateway_id\x18\x08 \x01(\x0c\"O\n\x1dHandleUplinkMACCommandRequest\x12\x0f\n\x07\x64\x65v_eui\x18\x01 \x01(\x0c\x12\x0b\n\x03\x63id\x18\x02 \x01(\r\x12\x10\n\x08\x63ommands\x18\x06 \x03(\x0c\"L\n#HandleRejectedUplinkFrameSetRequest\x12%\n\tframe_set\x18\x01 \x01(\x0b\x32\x12.gw.UplinkFrameSet*\xaf\x01\n\x05MType\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x10\n\x0cJOIN_REQUEST\x10\x01\x12\x0f\n\x0bJOIN_ACCEPT\x10\x02\x12\x17\n\x13UNCONFIRMED_DATA_UP\x10\x03\x12\x19\n\x15UNCONFIRMED_DATA_DOWN\x10\x04\x12\x15\n\x11\x43ONFIRMED_DATA_UP\x10\x05\x12\x17\n\x13\x43ONFIRMED_DATA_DOWN\x10\x06\x12\x12\n\x0eREJOIN_REQUEST\x10\x07\x32\xfe\x02\n\x18NetworkControllerService\x12Q\n\x14HandleUplinkMetaData\x12\x1f.nc.HandleUplinkMetaDataRequest\x1a\x16.google.protobuf.Empty\"\x00\x12U\n\x16HandleDownlinkMetaData\x12!.nc.HandleDownlinkMetaDataRequest\x1a\x16.google.protobuf.Empty\"\x00\x12U\n\x16HandleUplinkMACCommand\x12!.nc.HandleUplinkMACCommandRequest\x1a\x16.google.protobuf.Empty\"\x00\x12\x61\n\x1cHandleRejectedUplinkFrameSet\x12\'.nc.HandleRejectedUplinkFrameSetRequest\x1a\x16.google.protobuf.Empty\"\x00\x42\\\n\x14io.chirpstack.api.ncB\x16NetworkControllerProtoP\x01Z*github.com/kamicuu/chirpstack-api/go/v3/ncb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'chirpstack_api.nc.nc_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\024io.chirpstack.api.ncB\026NetworkControllerProtoP\001Z*github.com/kamicuu/chirpstack-api/go/v3/nc'
  _MTYPE._serialized_start=780
  _MTYPE._serialized_end=955
  _HANDLEUPLINKMETADATAREQUEST._serialized_start=92
  _HANDLEUPLINKMETADATAREQUEST._serialized_end=345
  _HANDLEDOWNLINKMETADATAREQUEST._serialized_start=348
  _HANDLEDOWNLINKMETADATAREQUEST._serialized_end=618
  _HANDLEUPLINKMACCOMMANDREQUEST._serialized_start=620
  _HANDLEUPLINKMACCOMMANDREQUEST._serialized_end=699
  _HANDLEREJECTEDUPLINKFRAMESETREQUEST._serialized_start=701
  _HANDLEREJECTEDUPLINKFRAMESETREQUEST._serialized_end=777
  _NETWORKCONTROLLERSERVICE._serialized_start=958
  _NETWORKCONTROLLERSERVICE._serialized_end=1340
# @@protoc_insertion_point(module_scope)

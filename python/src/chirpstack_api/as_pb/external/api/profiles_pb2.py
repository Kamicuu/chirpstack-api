# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chirpstack-api/as_pb/external/api/profiles.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n0chirpstack-api/as_pb/external/api/profiles.proto\x12\x03\x61pi\x1a\x1egoogle/protobuf/duration.proto\"\x9e\x05\n\x0eServiceProfile\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x15 \x01(\t\x12\'\n\x0forganization_id\x18\x16 \x01(\x03R\x0eorganizationID\x12*\n\x11network_server_id\x18\x17 \x01(\x03R\x0fnetworkServerID\x12\x0f\n\x07ul_rate\x18\x02 \x01(\r\x12\x16\n\x0eul_bucket_size\x18\x03 \x01(\r\x12\'\n\x0eul_rate_policy\x18\x04 \x01(\x0e\x32\x0f.api.RatePolicy\x12\x0f\n\x07\x64l_rate\x18\x05 \x01(\r\x12\x16\n\x0e\x64l_bucket_size\x18\x06 \x01(\r\x12\'\n\x0e\x64l_rate_policy\x18\x07 \x01(\x0e\x32\x0f.api.RatePolicy\x12&\n\x0f\x61\x64\x64_gw_metadata\x18\x08 \x01(\x08R\raddGWMetaData\x12\x1b\n\x13\x64\x65v_status_req_freq\x18\t \x01(\r\x12!\n\x19report_dev_status_battery\x18\n \x01(\x08\x12 \n\x18report_dev_status_margin\x18\x0b \x01(\x08\x12\x0e\n\x06\x64r_min\x18\x0c \x01(\r\x12\x0e\n\x06\x64r_max\x18\r \x01(\r\x12\x14\n\x0c\x63hannel_mask\x18\x0e \x01(\x0c\x12\x12\n\npr_allowed\x18\x0f \x01(\x08\x12\x12\n\nhr_allowed\x18\x10 \x01(\x08\x12\x12\n\nra_allowed\x18\x11 \x01(\x08\x12\x13\n\x0bnwk_geo_loc\x18\x12 \x01(\x08\x12\x1d\n\ntarget_per\x18\x13 \x01(\rR\ttargetPER\x12(\n\x10min_gw_diversity\x18\x14 \x01(\rR\x0eminGWDiversity\x12\x1f\n\x0bgws_private\x18\x18 \x01(\x08R\ngwsPrivate\"\xe0\x07\n\rDeviceProfile\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x15 \x01(\t\x12\'\n\x0forganization_id\x18\x16 \x01(\x03R\x0eorganizationID\x12*\n\x11network_server_id\x18\x17 \x01(\x03R\x0fnetworkServerID\x12\x18\n\x10supports_class_b\x18\x02 \x01(\x08\x12\x17\n\x0f\x63lass_b_timeout\x18\x03 \x01(\r\x12\x18\n\x10ping_slot_period\x18\x04 \x01(\r\x12 \n\x0cping_slot_dr\x18\x05 \x01(\rR\npingSlotDR\x12\x16\n\x0eping_slot_freq\x18\x06 \x01(\r\x12\x18\n\x10supports_class_c\x18\x07 \x01(\x08\x12\x17\n\x0f\x63lass_c_timeout\x18\x08 \x01(\r\x12\x13\n\x0bmac_version\x18\t \x01(\t\x12\x1b\n\x13reg_params_revision\x18\n \x01(\t\x12\x12\n\nrx_delay_1\x18\x0b \x01(\r\x12#\n\x0erx_dr_offset_1\x18\x0c \x01(\rR\x0brxDROffset1\x12\"\n\rrx_datarate_2\x18\r \x01(\rR\x0brxDataRate2\x12\x11\n\trx_freq_2\x18\x0e \x01(\r\x12\x1c\n\x14\x66\x61\x63tory_preset_freqs\x18\x0f \x03(\r\x12\x19\n\x08max_eirp\x18\x10 \x01(\rR\x07maxEIRP\x12\x16\n\x0emax_duty_cycle\x18\x11 \x01(\r\x12\x15\n\rsupports_join\x18\x12 \x01(\x08\x12\x11\n\trf_region\x18\x13 \x01(\t\x12/\n\x14supports_32bit_f_cnt\x18\x14 \x01(\x08R\x11supports32BitFCnt\x12\x15\n\rpayload_codec\x18\x18 \x01(\t\x12\x1e\n\x16payload_encoder_script\x18\x19 \x01(\t\x12\x1e\n\x16payload_decoder_script\x18\x1a \x01(\t\x12*\n\x11geoloc_buffer_ttl\x18\x1b \x01(\rR\x0fgeolocBufferTTL\x12\x1e\n\x16geoloc_min_buffer_size\x18\x1c \x01(\r\x12*\n\x04tags\x18\x1d \x03(\x0b\x32\x1c.api.DeviceProfile.TagsEntry\x12\x32\n\x0fuplink_interval\x18\x1e \x01(\x0b\x32\x19.google.protobuf.Duration\x12(\n\x10\x61\x64r_algorithm_id\x18\x1f \x01(\tR\x0e\x61\x64rAlgorithmID\x1a+\n\tTagsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01* \n\nRatePolicy\x12\x08\n\x04\x44ROP\x10\x00\x12\x08\n\x04MARK\x10\x01\x42m\n!io.chirpstack.api.as.external.apiB\rProfilesProtoP\x01Z7github.com/kamicuu/chirpstack-api/go/v3/as/external/apib\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'chirpstack_api.as_pb.external.api.profiles_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n!io.chirpstack.api.as.external.apiB\rProfilesProtoP\001Z7github.com/kamicuu/chirpstack-api/go/v3/as/external/api'
  _DEVICEPROFILE_TAGSENTRY._options = None
  _DEVICEPROFILE_TAGSENTRY._serialized_options = b'8\001'
  _RATEPOLICY._serialized_start=1757
  _RATEPOLICY._serialized_end=1789
  _SERVICEPROFILE._serialized_start=90
  _SERVICEPROFILE._serialized_end=760
  _DEVICEPROFILE._serialized_start=763
  _DEVICEPROFILE._serialized_end=1755
  _DEVICEPROFILE_TAGSENTRY._serialized_start=1712
  _DEVICEPROFILE_TAGSENTRY._serialized_end=1755
# @@protoc_insertion_point(module_scope)

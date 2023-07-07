# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chirpstack-api/as_pb/external/api/multicastGroup.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n6chirpstack-api/as_pb/external/api/multicastGroup.proto\x12\x03\x61pi\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1bgoogle/protobuf/empty.proto\"\x83\x02\n\x0eMulticastGroup\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0f\n\x07mc_addr\x18\x03 \x01(\t\x12\x14\n\x0cmc_nwk_s_key\x18\x04 \x01(\t\x12\x14\n\x0cmc_app_s_key\x18\x05 \x01(\t\x12\r\n\x05\x66_cnt\x18\x06 \x01(\r\x12+\n\ngroup_type\x18\x07 \x01(\x0e\x32\x17.api.MulticastGroupType\x12\n\n\x02\x64r\x18\x08 \x01(\r\x12\x11\n\tfrequency\x18\t \x01(\r\x12\x18\n\x10ping_slot_period\x18\n \x01(\r\x12%\n\x0e\x61pplication_id\x18\x0c \x01(\x03R\rapplicationID\"s\n\x16MulticastGroupListItem\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x0e\x61pplication_id\x18\x05 \x01(\x03R\rapplicationID\x12\x18\n\x10\x61pplication_name\x18\x06 \x01(\t\"K\n\x1b\x43reateMulticastGroupRequest\x12,\n\x0fmulticast_group\x18\x01 \x01(\x0b\x32\x13.api.MulticastGroup\"*\n\x1c\x43reateMulticastGroupResponse\x12\n\n\x02id\x18\x01 \x01(\t\"&\n\x18GetMulticastGroupRequest\x12\n\n\x02id\x18\x01 \x01(\t\"\xa9\x01\n\x19GetMulticastGroupResponse\x12,\n\x0fmulticast_group\x18\x01 \x01(\x0b\x32\x13.api.MulticastGroup\x12.\n\ncreated_at\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"K\n\x1bUpdateMulticastGroupRequest\x12,\n\x0fmulticast_group\x18\x01 \x01(\x0b\x32\x13.api.MulticastGroup\")\n\x1b\x44\x65leteMulticastGroupRequest\x12\n\n\x02id\x18\x01 \x01(\t\"i\n AddDeviceToMulticastGroupRequest\x12,\n\x12multicast_group_id\x18\x01 \x01(\tR\x10multicastGroupID\x12\x17\n\x07\x64\x65v_eui\x18\x02 \x01(\tR\x06\x64\x65vEUI\"n\n%RemoveDeviceFromMulticastGroupRequest\x12,\n\x12multicast_group_id\x18\x01 \x01(\tR\x10multicastGroupID\x12\x17\n\x07\x64\x65v_eui\x18\x02 \x01(\tR\x06\x64\x65vEUI\"\xb3\x01\n\x19ListMulticastGroupRequest\x12\r\n\x05limit\x18\x01 \x01(\x03\x12\x0e\n\x06offset\x18\x02 \x01(\x03\x12\'\n\x0forganization_id\x18\x03 \x01(\x03R\x0eorganizationID\x12\x17\n\x07\x64\x65v_eui\x18\x04 \x01(\tR\x06\x64\x65vEUI\x12\x0e\n\x06search\x18\x06 \x01(\t\x12%\n\x0e\x61pplication_id\x18\x07 \x01(\x03R\rapplicationID\"^\n\x1aListMulticastGroupResponse\x12\x13\n\x0btotal_count\x18\x01 \x01(\x03\x12+\n\x06result\x18\x02 \x03(\x0b\x32\x1b.api.MulticastGroupListItem\"o\n\x12MulticastQueueItem\x12,\n\x12multicast_group_id\x18\x01 \x01(\tR\x10multicastGroupID\x12\r\n\x05\x66_cnt\x18\x02 \x01(\r\x12\x0e\n\x06\x66_port\x18\x03 \x01(\r\x12\x0c\n\x04\x64\x61ta\x18\x04 \x01(\x0c\"Y\n EnqueueMulticastQueueItemRequest\x12\x35\n\x14multicast_queue_item\x18\x01 \x01(\x0b\x32\x17.api.MulticastQueueItem\"2\n!EnqueueMulticastQueueItemResponse\x12\r\n\x05\x66_cnt\x18\x01 \x01(\r\"T\n$FlushMulticastGroupQueueItemsRequest\x12,\n\x12multicast_group_id\x18\x01 \x01(\tR\x10multicastGroupID\"S\n#ListMulticastGroupQueueItemsRequest\x12,\n\x12multicast_group_id\x18\x01 \x01(\tR\x10multicastGroupID\"^\n$ListMulticastGroupQueueItemsResponse\x12\x36\n\x15multicast_queue_items\x18\x01 \x03(\x0b\x32\x17.api.MulticastQueueItem*.\n\x12MulticastGroupType\x12\x0b\n\x07\x43LASS_C\x10\x00\x12\x0b\n\x07\x43LASS_B\x10\x01\x32\xba\n\n\x15MulticastGroupService\x12o\n\x06\x43reate\x12 .api.CreateMulticastGroupRequest\x1a!.api.CreateMulticastGroupResponse\" \x82\xd3\xe4\x93\x02\x1a\"\x15/api/multicast-groups:\x01*\x12h\n\x03Get\x12\x1d.api.GetMulticastGroupRequest\x1a\x1e.api.GetMulticastGroupResponse\"\"\x82\xd3\xe4\x93\x02\x1c\x12\x1a/api/multicast-groups/{id}\x12y\n\x06Update\x12 .api.UpdateMulticastGroupRequest\x1a\x16.google.protobuf.Empty\"5\x82\xd3\xe4\x93\x02/\x1a*/api/multicast-groups/{multicast_group.id}:\x01*\x12\x66\n\x06\x44\x65lete\x12 .api.DeleteMulticastGroupRequest\x1a\x16.google.protobuf.Empty\"\"\x82\xd3\xe4\x93\x02\x1c*\x1a/api/multicast-groups/{id}\x12\x66\n\x04List\x12\x1e.api.ListMulticastGroupRequest\x1a\x1f.api.ListMulticastGroupResponse\"\x1d\x82\xd3\xe4\x93\x02\x17\x12\x15/api/multicast-groups\x12\x89\x01\n\tAddDevice\x12%.api.AddDeviceToMulticastGroupRequest\x1a\x16.google.protobuf.Empty\"=\x82\xd3\xe4\x93\x02\x37\"2/api/multicast-groups/{multicast_group_id}/devices:\x01*\x12\x98\x01\n\x0cRemoveDevice\x12*.api.RemoveDeviceFromMulticastGroupRequest\x1a\x16.google.protobuf.Empty\"D\x82\xd3\xe4\x93\x02>*</api/multicast-groups/{multicast_group_id}/devices/{dev_eui}\x12\xaa\x01\n\x07\x45nqueue\x12%.api.EnqueueMulticastQueueItemRequest\x1a&.api.EnqueueMulticastQueueItemResponse\"P\x82\xd3\xe4\x93\x02J\"E/api/multicast-groups/{multicast_queue_item.multicast_group_id}/queue:\x01*\x12\x89\x01\n\nFlushQueue\x12).api.FlushMulticastGroupQueueItemsRequest\x1a\x16.google.protobuf.Empty\"8\x82\xd3\xe4\x93\x02\x32*0/api/multicast-groups/{multicast_group_id}/queue\x12\x9a\x01\n\tListQueue\x12(.api.ListMulticastGroupQueueItemsRequest\x1a).api.ListMulticastGroupQueueItemsResponse\"8\x82\xd3\xe4\x93\x02\x32\x12\x30/api/multicast-groups/{multicast_group_id}/queueBs\n!io.chirpstack.api.as.external.apiB\x13MulticastGroupProtoP\x01Z7github.com/kamicuu/chirpstack-api/go/v3/as/external/apib\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'chirpstack_api.as_pb.external.api.multicastGroup_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n!io.chirpstack.api.as.external.apiB\023MulticastGroupProtoP\001Z7github.com/kamicuu/chirpstack-api/go/v3/as/external/api'
  _MULTICASTGROUPSERVICE.methods_by_name['Create']._options = None
  _MULTICASTGROUPSERVICE.methods_by_name['Create']._serialized_options = b'\202\323\344\223\002\032\"\025/api/multicast-groups:\001*'
  _MULTICASTGROUPSERVICE.methods_by_name['Get']._options = None
  _MULTICASTGROUPSERVICE.methods_by_name['Get']._serialized_options = b'\202\323\344\223\002\034\022\032/api/multicast-groups/{id}'
  _MULTICASTGROUPSERVICE.methods_by_name['Update']._options = None
  _MULTICASTGROUPSERVICE.methods_by_name['Update']._serialized_options = b'\202\323\344\223\002/\032*/api/multicast-groups/{multicast_group.id}:\001*'
  _MULTICASTGROUPSERVICE.methods_by_name['Delete']._options = None
  _MULTICASTGROUPSERVICE.methods_by_name['Delete']._serialized_options = b'\202\323\344\223\002\034*\032/api/multicast-groups/{id}'
  _MULTICASTGROUPSERVICE.methods_by_name['List']._options = None
  _MULTICASTGROUPSERVICE.methods_by_name['List']._serialized_options = b'\202\323\344\223\002\027\022\025/api/multicast-groups'
  _MULTICASTGROUPSERVICE.methods_by_name['AddDevice']._options = None
  _MULTICASTGROUPSERVICE.methods_by_name['AddDevice']._serialized_options = b'\202\323\344\223\0027\"2/api/multicast-groups/{multicast_group_id}/devices:\001*'
  _MULTICASTGROUPSERVICE.methods_by_name['RemoveDevice']._options = None
  _MULTICASTGROUPSERVICE.methods_by_name['RemoveDevice']._serialized_options = b'\202\323\344\223\002>*</api/multicast-groups/{multicast_group_id}/devices/{dev_eui}'
  _MULTICASTGROUPSERVICE.methods_by_name['Enqueue']._options = None
  _MULTICASTGROUPSERVICE.methods_by_name['Enqueue']._serialized_options = b'\202\323\344\223\002J\"E/api/multicast-groups/{multicast_queue_item.multicast_group_id}/queue:\001*'
  _MULTICASTGROUPSERVICE.methods_by_name['FlushQueue']._options = None
  _MULTICASTGROUPSERVICE.methods_by_name['FlushQueue']._serialized_options = b'\202\323\344\223\0022*0/api/multicast-groups/{multicast_group_id}/queue'
  _MULTICASTGROUPSERVICE.methods_by_name['ListQueue']._options = None
  _MULTICASTGROUPSERVICE.methods_by_name['ListQueue']._serialized_options = b'\202\323\344\223\0022\0220/api/multicast-groups/{multicast_group_id}/queue'
  _MULTICASTGROUPTYPE._serialized_start=2007
  _MULTICASTGROUPTYPE._serialized_end=2053
  _MULTICASTGROUP._serialized_start=156
  _MULTICASTGROUP._serialized_end=415
  _MULTICASTGROUPLISTITEM._serialized_start=417
  _MULTICASTGROUPLISTITEM._serialized_end=532
  _CREATEMULTICASTGROUPREQUEST._serialized_start=534
  _CREATEMULTICASTGROUPREQUEST._serialized_end=609
  _CREATEMULTICASTGROUPRESPONSE._serialized_start=611
  _CREATEMULTICASTGROUPRESPONSE._serialized_end=653
  _GETMULTICASTGROUPREQUEST._serialized_start=655
  _GETMULTICASTGROUPREQUEST._serialized_end=693
  _GETMULTICASTGROUPRESPONSE._serialized_start=696
  _GETMULTICASTGROUPRESPONSE._serialized_end=865
  _UPDATEMULTICASTGROUPREQUEST._serialized_start=867
  _UPDATEMULTICASTGROUPREQUEST._serialized_end=942
  _DELETEMULTICASTGROUPREQUEST._serialized_start=944
  _DELETEMULTICASTGROUPREQUEST._serialized_end=985
  _ADDDEVICETOMULTICASTGROUPREQUEST._serialized_start=987
  _ADDDEVICETOMULTICASTGROUPREQUEST._serialized_end=1092
  _REMOVEDEVICEFROMMULTICASTGROUPREQUEST._serialized_start=1094
  _REMOVEDEVICEFROMMULTICASTGROUPREQUEST._serialized_end=1204
  _LISTMULTICASTGROUPREQUEST._serialized_start=1207
  _LISTMULTICASTGROUPREQUEST._serialized_end=1386
  _LISTMULTICASTGROUPRESPONSE._serialized_start=1388
  _LISTMULTICASTGROUPRESPONSE._serialized_end=1482
  _MULTICASTQUEUEITEM._serialized_start=1484
  _MULTICASTQUEUEITEM._serialized_end=1595
  _ENQUEUEMULTICASTQUEUEITEMREQUEST._serialized_start=1597
  _ENQUEUEMULTICASTQUEUEITEMREQUEST._serialized_end=1686
  _ENQUEUEMULTICASTQUEUEITEMRESPONSE._serialized_start=1688
  _ENQUEUEMULTICASTQUEUEITEMRESPONSE._serialized_end=1738
  _FLUSHMULTICASTGROUPQUEUEITEMSREQUEST._serialized_start=1740
  _FLUSHMULTICASTGROUPQUEUEITEMSREQUEST._serialized_end=1824
  _LISTMULTICASTGROUPQUEUEITEMSREQUEST._serialized_start=1826
  _LISTMULTICASTGROUPQUEUEITEMSREQUEST._serialized_end=1909
  _LISTMULTICASTGROUPQUEUEITEMSRESPONSE._serialized_start=1911
  _LISTMULTICASTGROUPQUEUEITEMSRESPONSE._serialized_end=2005
  _MULTICASTGROUPSERVICE._serialized_start=2056
  _MULTICASTGROUPSERVICE._serialized_end=3394
# @@protoc_insertion_point(module_scope)

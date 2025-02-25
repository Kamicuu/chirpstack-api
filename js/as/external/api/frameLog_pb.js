/**
 * @fileoverview
 * @enhanceable
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js');
var common_common_pb = require('../../../common/common_pb.js');
var gw_gw_pb = require('../../../gw/gw_pb.js');
goog.exportSymbol('proto.api.DownlinkFrameLog', null, global);
goog.exportSymbol('proto.api.RXWindow', null, global);
goog.exportSymbol('proto.api.UplinkFrameLog', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.api.UplinkFrameLog = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.api.UplinkFrameLog.repeatedFields_, null);
};
goog.inherits(proto.api.UplinkFrameLog, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.api.UplinkFrameLog.displayName = 'proto.api.UplinkFrameLog';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.api.UplinkFrameLog.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.api.UplinkFrameLog.prototype.toObject = function(opt_includeInstance) {
  return proto.api.UplinkFrameLog.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.api.UplinkFrameLog} msg The msg instance to transform.
 * @return {!Object}
 */
proto.api.UplinkFrameLog.toObject = function(includeInstance, msg) {
  var f, obj = {
    txInfo: (f = msg.getTxInfo()) && gw_gw_pb.UplinkTXInfo.toObject(includeInstance, f),
    rxInfoList: jspb.Message.toObjectList(msg.getRxInfoList(),
    gw_gw_pb.UplinkRXInfo.toObject, includeInstance),
    phyPayloadJson: msg.getPhyPayloadJson(),
    publishedAt: (f = msg.getPublishedAt()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.api.UplinkFrameLog}
 */
proto.api.UplinkFrameLog.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.api.UplinkFrameLog;
  return proto.api.UplinkFrameLog.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.api.UplinkFrameLog} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.api.UplinkFrameLog}
 */
proto.api.UplinkFrameLog.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new gw_gw_pb.UplinkTXInfo;
      reader.readMessage(value,gw_gw_pb.UplinkTXInfo.deserializeBinaryFromReader);
      msg.setTxInfo(value);
      break;
    case 2:
      var value = new gw_gw_pb.UplinkRXInfo;
      reader.readMessage(value,gw_gw_pb.UplinkRXInfo.deserializeBinaryFromReader);
      msg.getRxInfoList().push(value);
      msg.setRxInfoList(msg.getRxInfoList());
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setPhyPayloadJson(value);
      break;
    case 4:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setPublishedAt(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.api.UplinkFrameLog} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.api.UplinkFrameLog.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.api.UplinkFrameLog.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.api.UplinkFrameLog.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getTxInfo();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      gw_gw_pb.UplinkTXInfo.serializeBinaryToWriter
    );
  }
  f = this.getRxInfoList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      gw_gw_pb.UplinkRXInfo.serializeBinaryToWriter
    );
  }
  f = this.getPhyPayloadJson();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = this.getPublishedAt();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
};


/**
 * Creates a deep clone of this proto. No data is shared with the original.
 * @return {!proto.api.UplinkFrameLog} The clone.
 */
proto.api.UplinkFrameLog.prototype.cloneMessage = function() {
  return /** @type {!proto.api.UplinkFrameLog} */ (jspb.Message.cloneMessage(this));
};


/**
 * optional gw.UplinkTXInfo tx_info = 1;
 * @return {proto.gw.UplinkTXInfo}
 */
proto.api.UplinkFrameLog.prototype.getTxInfo = function() {
  return /** @type{proto.gw.UplinkTXInfo} */ (
    jspb.Message.getWrapperField(this, gw_gw_pb.UplinkTXInfo, 1));
};


/** @param {proto.gw.UplinkTXInfo|undefined} value  */
proto.api.UplinkFrameLog.prototype.setTxInfo = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.api.UplinkFrameLog.prototype.clearTxInfo = function() {
  this.setTxInfo(undefined);
};


/**
 * Returns whether this field is set.
 * @return{!boolean}
 */
proto.api.UplinkFrameLog.prototype.hasTxInfo = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated gw.UplinkRXInfo rx_info = 2;
 * If you change this array by adding, removing or replacing elements, or if you
 * replace the array itself, then you must call the setter to update it.
 * @return {!Array.<!proto.gw.UplinkRXInfo>}
 */
proto.api.UplinkFrameLog.prototype.getRxInfoList = function() {
  return /** @type{!Array.<!proto.gw.UplinkRXInfo>} */ (
    jspb.Message.getRepeatedWrapperField(this, gw_gw_pb.UplinkRXInfo, 2));
};


/** @param {Array.<!proto.gw.UplinkRXInfo>} value  */
proto.api.UplinkFrameLog.prototype.setRxInfoList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 2, value);
};


proto.api.UplinkFrameLog.prototype.clearRxInfoList = function() {
  this.setRxInfoList([]);
};


/**
 * optional string phy_payload_json = 3;
 * @return {string}
 */
proto.api.UplinkFrameLog.prototype.getPhyPayloadJson = function() {
  return /** @type {string} */ (jspb.Message.getFieldProto3(this, 3, ""));
};


/** @param {string} value  */
proto.api.UplinkFrameLog.prototype.setPhyPayloadJson = function(value) {
  jspb.Message.setField(this, 3, value);
};


/**
 * optional google.protobuf.Timestamp published_at = 4;
 * @return {proto.google.protobuf.Timestamp}
 */
proto.api.UplinkFrameLog.prototype.getPublishedAt = function() {
  return /** @type{proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 4));
};


/** @param {proto.google.protobuf.Timestamp|undefined} value  */
proto.api.UplinkFrameLog.prototype.setPublishedAt = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


proto.api.UplinkFrameLog.prototype.clearPublishedAt = function() {
  this.setPublishedAt(undefined);
};


/**
 * Returns whether this field is set.
 * @return{!boolean}
 */
proto.api.UplinkFrameLog.prototype.hasPublishedAt = function() {
  return jspb.Message.getField(this, 4) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.api.DownlinkFrameLog = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.api.DownlinkFrameLog, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.api.DownlinkFrameLog.displayName = 'proto.api.DownlinkFrameLog';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.api.DownlinkFrameLog.prototype.toObject = function(opt_includeInstance) {
  return proto.api.DownlinkFrameLog.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.api.DownlinkFrameLog} msg The msg instance to transform.
 * @return {!Object}
 */
proto.api.DownlinkFrameLog.toObject = function(includeInstance, msg) {
  var f, obj = {
    txInfo: (f = msg.getTxInfo()) && gw_gw_pb.DownlinkTXInfo.toObject(includeInstance, f),
    phyPayloadJson: msg.getPhyPayloadJson(),
    gatewayId: msg.getGatewayId(),
    publishedAt: (f = msg.getPublishedAt()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.api.DownlinkFrameLog}
 */
proto.api.DownlinkFrameLog.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.api.DownlinkFrameLog;
  return proto.api.DownlinkFrameLog.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.api.DownlinkFrameLog} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.api.DownlinkFrameLog}
 */
proto.api.DownlinkFrameLog.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new gw_gw_pb.DownlinkTXInfo;
      reader.readMessage(value,gw_gw_pb.DownlinkTXInfo.deserializeBinaryFromReader);
      msg.setTxInfo(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPhyPayloadJson(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setGatewayId(value);
      break;
    case 4:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setPublishedAt(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.api.DownlinkFrameLog} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.api.DownlinkFrameLog.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.api.DownlinkFrameLog.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.api.DownlinkFrameLog.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getTxInfo();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      gw_gw_pb.DownlinkTXInfo.serializeBinaryToWriter
    );
  }
  f = this.getPhyPayloadJson();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = this.getGatewayId();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = this.getPublishedAt();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
};


/**
 * Creates a deep clone of this proto. No data is shared with the original.
 * @return {!proto.api.DownlinkFrameLog} The clone.
 */
proto.api.DownlinkFrameLog.prototype.cloneMessage = function() {
  return /** @type {!proto.api.DownlinkFrameLog} */ (jspb.Message.cloneMessage(this));
};


/**
 * optional gw.DownlinkTXInfo tx_info = 1;
 * @return {proto.gw.DownlinkTXInfo}
 */
proto.api.DownlinkFrameLog.prototype.getTxInfo = function() {
  return /** @type{proto.gw.DownlinkTXInfo} */ (
    jspb.Message.getWrapperField(this, gw_gw_pb.DownlinkTXInfo, 1));
};


/** @param {proto.gw.DownlinkTXInfo|undefined} value  */
proto.api.DownlinkFrameLog.prototype.setTxInfo = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.api.DownlinkFrameLog.prototype.clearTxInfo = function() {
  this.setTxInfo(undefined);
};


/**
 * Returns whether this field is set.
 * @return{!boolean}
 */
proto.api.DownlinkFrameLog.prototype.hasTxInfo = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string phy_payload_json = 2;
 * @return {string}
 */
proto.api.DownlinkFrameLog.prototype.getPhyPayloadJson = function() {
  return /** @type {string} */ (jspb.Message.getFieldProto3(this, 2, ""));
};


/** @param {string} value  */
proto.api.DownlinkFrameLog.prototype.setPhyPayloadJson = function(value) {
  jspb.Message.setField(this, 2, value);
};


/**
 * optional string gateway_id = 3;
 * @return {string}
 */
proto.api.DownlinkFrameLog.prototype.getGatewayId = function() {
  return /** @type {string} */ (jspb.Message.getFieldProto3(this, 3, ""));
};


/** @param {string} value  */
proto.api.DownlinkFrameLog.prototype.setGatewayId = function(value) {
  jspb.Message.setField(this, 3, value);
};


/**
 * optional google.protobuf.Timestamp published_at = 4;
 * @return {proto.google.protobuf.Timestamp}
 */
proto.api.DownlinkFrameLog.prototype.getPublishedAt = function() {
  return /** @type{proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 4));
};


/** @param {proto.google.protobuf.Timestamp|undefined} value  */
proto.api.DownlinkFrameLog.prototype.setPublishedAt = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


proto.api.DownlinkFrameLog.prototype.clearPublishedAt = function() {
  this.setPublishedAt(undefined);
};


/**
 * Returns whether this field is set.
 * @return{!boolean}
 */
proto.api.DownlinkFrameLog.prototype.hasPublishedAt = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * @enum {number}
 */
proto.api.RXWindow = {
  RX1: 0,
  RX2: 1
};

goog.object.extend(exports, proto.api);

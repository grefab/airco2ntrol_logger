/**
 * @fileoverview gRPC-Web generated client stub for airco2ntrol
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.airco2ntrol = require('./api_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.airco2ntrol.StorageClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.airco2ntrol.StoragePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.google.protobuf.Timestamp,
 *   !proto.airco2ntrol.AirQuality>}
 */
const methodDescriptor_Storage_GetSince = new grpc.web.MethodDescriptor(
  '/airco2ntrol.Storage/GetSince',
  grpc.web.MethodType.SERVER_STREAMING,
  google_protobuf_timestamp_pb.Timestamp,
  proto.airco2ntrol.AirQuality,
  /**
   * @param {!proto.google.protobuf.Timestamp} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.airco2ntrol.AirQuality.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.google.protobuf.Timestamp,
 *   !proto.airco2ntrol.AirQuality>}
 */
const methodInfo_Storage_GetSince = new grpc.web.AbstractClientBase.MethodInfo(
  proto.airco2ntrol.AirQuality,
  /**
   * @param {!proto.google.protobuf.Timestamp} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.airco2ntrol.AirQuality.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Timestamp} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.airco2ntrol.AirQuality>}
 *     The XHR Node Readable Stream
 */
proto.airco2ntrol.StorageClient.prototype.getSince =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/airco2ntrol.Storage/GetSince',
      request,
      metadata || {},
      methodDescriptor_Storage_GetSince);
};


/**
 * @param {!proto.google.protobuf.Timestamp} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.airco2ntrol.AirQuality>}
 *     The XHR Node Readable Stream
 */
proto.airco2ntrol.StoragePromiseClient.prototype.getSince =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/airco2ntrol.Storage/GetSince',
      request,
      metadata || {},
      methodDescriptor_Storage_GetSince);
};


module.exports = proto.airco2ntrol;


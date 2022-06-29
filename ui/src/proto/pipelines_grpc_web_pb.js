/**
 * @fileoverview gRPC-Web generated client stub for pipelines
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.pipelines = require('./pipelines_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.pipelines.PipelineStatusClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

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
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.pipelines.PipelineStatusPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

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
 *   !proto.pipelines.GetStatusRequest,
 *   !proto.pipelines.GetStatusResponse>}
 */
const methodDescriptor_PipelineStatus_GetStatus = new grpc.web.MethodDescriptor(
  '/pipelines.PipelineStatus/GetStatus',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.pipelines.GetStatusRequest,
  proto.pipelines.GetStatusResponse,
  /**
   * @param {!proto.pipelines.GetStatusRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pipelines.GetStatusResponse.deserializeBinary
);


/**
 * @param {!proto.pipelines.GetStatusRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.pipelines.GetStatusResponse>}
 *     The XHR Node Readable Stream
 */
proto.pipelines.PipelineStatusClient.prototype.getStatus =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/pipelines.PipelineStatus/GetStatus',
      request,
      metadata || {},
      methodDescriptor_PipelineStatus_GetStatus);
};


/**
 * @param {!proto.pipelines.GetStatusRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.pipelines.GetStatusResponse>}
 *     The XHR Node Readable Stream
 */
proto.pipelines.PipelineStatusPromiseClient.prototype.getStatus =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/pipelines.PipelineStatus/GetStatus',
      request,
      metadata || {},
      methodDescriptor_PipelineStatus_GetStatus);
};


module.exports = proto.pipelines;


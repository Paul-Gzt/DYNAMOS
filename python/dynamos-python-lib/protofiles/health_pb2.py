# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: health.proto
# Protobuf Python Version: 5.29.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    0,
    '',
    'health.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0chealth.proto\x12\x07\x64ynamos\"%\n\x12HealthCheckRequest\x12\x0f\n\x07service\x18\x01 \x01(\t\"\xa2\x01\n\x13HealthCheckResponse\x12:\n\x06status\x18\x01 \x01(\x0e\x32*.dynamos.HealthCheckResponse.ServingStatus\"O\n\rServingStatus\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0b\n\x07SERVING\x10\x01\x12\x0f\n\x0bNOT_SERVING\x10\x02\x12\x13\n\x0fSERVICE_UNKNOWN\x10\x03\x32\x92\x01\n\x06Health\x12\x42\n\x05\x43heck\x12\x1b.dynamos.HealthCheckRequest\x1a\x1c.dynamos.HealthCheckResponse\x12\x44\n\x05Watch\x12\x1b.dynamos.HealthCheckRequest\x1a\x1c.dynamos.HealthCheckResponse0\x01\x42\'Z%github.com/Jorrit05/DYNAMOS/pkg/protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'health_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z%github.com/Jorrit05/DYNAMOS/pkg/proto'
  _globals['_HEALTHCHECKREQUEST']._serialized_start=25
  _globals['_HEALTHCHECKREQUEST']._serialized_end=62
  _globals['_HEALTHCHECKRESPONSE']._serialized_start=65
  _globals['_HEALTHCHECKRESPONSE']._serialized_end=227
  _globals['_HEALTHCHECKRESPONSE_SERVINGSTATUS']._serialized_start=148
  _globals['_HEALTHCHECKRESPONSE_SERVINGSTATUS']._serialized_end=227
  _globals['_HEALTH']._serialized_start=230
  _globals['_HEALTH']._serialized_end=376
# @@protoc_insertion_point(module_scope)

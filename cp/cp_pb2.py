# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: cp.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x08\x63p.proto\x12\x02\x63p\"3\n\x02\x43p\x12\r\n\x05level\x18\x01 \x01(\x05\x12\r\n\x05stone\x18\x02 \x01(\t\x12\x0f\n\x07squares\x18\x03 \x01(\t\"\x1b\n\x03Res\x12\t\n\x01x\x18\x01 \x01(\x05\x12\t\n\x01y\x18\x02 \x01(\x05\x32$\n\x05\x43pApi\x12\x1b\n\x06\x41ttack\x12\x06.cp.Cp\x1a\x07.cp.Res\"\x00\x62\x06proto3')



_CP = DESCRIPTOR.message_types_by_name['Cp']
_RES = DESCRIPTOR.message_types_by_name['Res']
Cp = _reflection.GeneratedProtocolMessageType('Cp', (_message.Message,), {
  'DESCRIPTOR' : _CP,
  '__module__' : 'cp_pb2'
  # @@protoc_insertion_point(class_scope:cp.Cp)
  })
_sym_db.RegisterMessage(Cp)

Res = _reflection.GeneratedProtocolMessageType('Res', (_message.Message,), {
  'DESCRIPTOR' : _RES,
  '__module__' : 'cp_pb2'
  # @@protoc_insertion_point(class_scope:cp.Res)
  })
_sym_db.RegisterMessage(Res)

_CPAPI = DESCRIPTOR.services_by_name['CpApi']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _CP._serialized_start=16
  _CP._serialized_end=67
  _RES._serialized_start=69
  _RES._serialized_end=96
  _CPAPI._serialized_start=98
  _CPAPI._serialized_end=134
# @@protoc_insertion_point(module_scope)

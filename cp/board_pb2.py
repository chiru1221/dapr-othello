# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: board.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0b\x62oard.proto\x12\x05\x62oard\"=\n\x05\x42oard\x12\r\n\x05stone\x18\x01 \x01(\t\x12\t\n\x01x\x18\x02 \x01(\x05\x12\t\n\x01y\x18\x03 \x01(\x05\x12\x0f\n\x07squares\x18\x04 \x01(\t\"\x16\n\x03Res\x12\x0f\n\x07squares\x18\x01 \x01(\t2X\n\x08\x42oardApi\x12%\n\x07Putable\x12\x0c.board.Board\x1a\n.board.Res\"\x00\x12%\n\x07Reverse\x12\x0c.board.Board\x1a\n.board.Res\"\x00\x42\x1bZ\x19\x65xample.com/othello/boardb\x06proto3')



_BOARD = DESCRIPTOR.message_types_by_name['Board']
_RES = DESCRIPTOR.message_types_by_name['Res']
Board = _reflection.GeneratedProtocolMessageType('Board', (_message.Message,), {
  'DESCRIPTOR' : _BOARD,
  '__module__' : 'board_pb2'
  # @@protoc_insertion_point(class_scope:board.Board)
  })
_sym_db.RegisterMessage(Board)

Res = _reflection.GeneratedProtocolMessageType('Res', (_message.Message,), {
  'DESCRIPTOR' : _RES,
  '__module__' : 'board_pb2'
  # @@protoc_insertion_point(class_scope:board.Res)
  })
_sym_db.RegisterMessage(Res)

_BOARDAPI = DESCRIPTOR.services_by_name['BoardApi']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\031example.com/othello/board'
  _BOARD._serialized_start=22
  _BOARD._serialized_end=83
  _RES._serialized_start=85
  _RES._serialized_end=107
  _BOARDAPI._serialized_start=109
  _BOARDAPI._serialized_end=197
# @@protoc_insertion_point(module_scope)

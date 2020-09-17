# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user_matching.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='user_matching.proto',
  package='',
  syntax='proto3',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x13user_matching.proto\"$\n\x11GetMatchesRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\"#\n\x0fGetMatchesReply\x12\x10\n\x08user_ids\x18\x01 \x03(\t\"\x07\n\x05\x45mpty2`\n\x07Matcher\x12\x34\n\nGetMatches\x12\x12.GetMatchesRequest\x1a\x10.GetMatchesReply\"\x00\x12\x1f\n\x0bRecalcScaNN\x12\x06.Empty\x1a\x06.Empty\"\x00\x62\x06proto3'
)




_GETMATCHESREQUEST = _descriptor.Descriptor(
  name='GetMatchesRequest',
  full_name='GetMatchesRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='GetMatchesRequest.user_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=23,
  serialized_end=59,
)


_GETMATCHESREPLY = _descriptor.Descriptor(
  name='GetMatchesReply',
  full_name='GetMatchesReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_ids', full_name='GetMatchesReply.user_ids', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=61,
  serialized_end=96,
)


_EMPTY = _descriptor.Descriptor(
  name='Empty',
  full_name='Empty',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=98,
  serialized_end=105,
)

DESCRIPTOR.message_types_by_name['GetMatchesRequest'] = _GETMATCHESREQUEST
DESCRIPTOR.message_types_by_name['GetMatchesReply'] = _GETMATCHESREPLY
DESCRIPTOR.message_types_by_name['Empty'] = _EMPTY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

GetMatchesRequest = _reflection.GeneratedProtocolMessageType('GetMatchesRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETMATCHESREQUEST,
  '__module__' : 'user_matching_pb2'
  # @@protoc_insertion_point(class_scope:GetMatchesRequest)
  })
_sym_db.RegisterMessage(GetMatchesRequest)

GetMatchesReply = _reflection.GeneratedProtocolMessageType('GetMatchesReply', (_message.Message,), {
  'DESCRIPTOR' : _GETMATCHESREPLY,
  '__module__' : 'user_matching_pb2'
  # @@protoc_insertion_point(class_scope:GetMatchesReply)
  })
_sym_db.RegisterMessage(GetMatchesReply)

Empty = _reflection.GeneratedProtocolMessageType('Empty', (_message.Message,), {
  'DESCRIPTOR' : _EMPTY,
  '__module__' : 'user_matching_pb2'
  # @@protoc_insertion_point(class_scope:Empty)
  })
_sym_db.RegisterMessage(Empty)



_MATCHER = _descriptor.ServiceDescriptor(
  name='Matcher',
  full_name='Matcher',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=107,
  serialized_end=203,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetMatches',
    full_name='Matcher.GetMatches',
    index=0,
    containing_service=None,
    input_type=_GETMATCHESREQUEST,
    output_type=_GETMATCHESREPLY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='RecalcScaNN',
    full_name='Matcher.RecalcScaNN',
    index=1,
    containing_service=None,
    input_type=_EMPTY,
    output_type=_EMPTY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_MATCHER)

DESCRIPTOR.services_by_name['Matcher'] = _MATCHER

# @@protoc_insertion_point(module_scope)

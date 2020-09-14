# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import user_matching_pb2 as user__matching__pb2


class MatcherStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetMatches = channel.unary_unary(
                '/Matcher/GetMatches',
                request_serializer=user__matching__pb2.GetMatchesRequest.SerializeToString,
                response_deserializer=user__matching__pb2.GetMatchesReply.FromString,
                )
        self.RecalcScaNN = channel.unary_unary(
                '/Matcher/RecalcScaNN',
                request_serializer=user__matching__pb2.Empty.SerializeToString,
                response_deserializer=user__matching__pb2.Empty.FromString,
                )


class MatcherServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetMatches(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RecalcScaNN(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MatcherServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetMatches': grpc.unary_unary_rpc_method_handler(
                    servicer.GetMatches,
                    request_deserializer=user__matching__pb2.GetMatchesRequest.FromString,
                    response_serializer=user__matching__pb2.GetMatchesReply.SerializeToString,
            ),
            'RecalcScaNN': grpc.unary_unary_rpc_method_handler(
                    servicer.RecalcScaNN,
                    request_deserializer=user__matching__pb2.Empty.FromString,
                    response_serializer=user__matching__pb2.Empty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'Matcher', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Matcher(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetMatches(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Matcher/GetMatches',
            user__matching__pb2.GetMatchesRequest.SerializeToString,
            user__matching__pb2.GetMatchesReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RecalcScaNN(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Matcher/RecalcScaNN',
            user__matching__pb2.Empty.SerializeToString,
            user__matching__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

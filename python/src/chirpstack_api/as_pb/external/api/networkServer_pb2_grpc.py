# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from chirpstack_api.as_pb.external.api import networkServer_pb2 as chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


class NetworkServerServiceStub(object):
    """NetworkServerService is the service managing network-servers.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Create = channel.unary_unary(
                '/api.NetworkServerService/Create',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.CreateNetworkServerRequest.SerializeToString,
                response_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.CreateNetworkServerResponse.FromString,
                )
        self.Get = channel.unary_unary(
                '/api.NetworkServerService/Get',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.GetNetworkServerRequest.SerializeToString,
                response_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.GetNetworkServerResponse.FromString,
                )
        self.Update = channel.unary_unary(
                '/api.NetworkServerService/Update',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.UpdateNetworkServerRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.Delete = channel.unary_unary(
                '/api.NetworkServerService/Delete',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.DeleteNetworkServerRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.List = channel.unary_unary(
                '/api.NetworkServerService/List',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.ListNetworkServerRequest.SerializeToString,
                response_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.ListNetworkServerResponse.FromString,
                )
        self.GetADRAlgorithms = channel.unary_unary(
                '/api.NetworkServerService/GetADRAlgorithms',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.GetADRAlgorithmsRequest.SerializeToString,
                response_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.GetADRAlgorithmsResponse.FromString,
                )


class NetworkServerServiceServicer(object):
    """NetworkServerService is the service managing network-servers.
    """

    def Create(self, request, context):
        """Create creates the given network-server.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Get(self, request, context):
        """Get returns the network-server matching the given id.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Update(self, request, context):
        """Update updates the given network-server.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Delete(self, request, context):
        """Delete deletes the network-server matching the given id.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def List(self, request, context):
        """List lists the available network-servers.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetADRAlgorithms(self, request, context):
        """GetADRAlgorithms returns the available ADR algorithms.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_NetworkServerServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Create': grpc.unary_unary_rpc_method_handler(
                    servicer.Create,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.CreateNetworkServerRequest.FromString,
                    response_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.CreateNetworkServerResponse.SerializeToString,
            ),
            'Get': grpc.unary_unary_rpc_method_handler(
                    servicer.Get,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.GetNetworkServerRequest.FromString,
                    response_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.GetNetworkServerResponse.SerializeToString,
            ),
            'Update': grpc.unary_unary_rpc_method_handler(
                    servicer.Update,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.UpdateNetworkServerRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'Delete': grpc.unary_unary_rpc_method_handler(
                    servicer.Delete,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.DeleteNetworkServerRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'List': grpc.unary_unary_rpc_method_handler(
                    servicer.List,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.ListNetworkServerRequest.FromString,
                    response_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.ListNetworkServerResponse.SerializeToString,
            ),
            'GetADRAlgorithms': grpc.unary_unary_rpc_method_handler(
                    servicer.GetADRAlgorithms,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.GetADRAlgorithmsRequest.FromString,
                    response_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.GetADRAlgorithmsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'api.NetworkServerService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class NetworkServerService(object):
    """NetworkServerService is the service managing network-servers.
    """

    @staticmethod
    def Create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.NetworkServerService/Create',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.CreateNetworkServerRequest.SerializeToString,
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.CreateNetworkServerResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.NetworkServerService/Get',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.GetNetworkServerRequest.SerializeToString,
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.GetNetworkServerResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Update(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.NetworkServerService/Update',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.UpdateNetworkServerRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Delete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.NetworkServerService/Delete',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.DeleteNetworkServerRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def List(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.NetworkServerService/List',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.ListNetworkServerRequest.SerializeToString,
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.ListNetworkServerResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetADRAlgorithms(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.NetworkServerService/GetADRAlgorithms',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.GetADRAlgorithmsRequest.SerializeToString,
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_networkServer__pb2.GetADRAlgorithmsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

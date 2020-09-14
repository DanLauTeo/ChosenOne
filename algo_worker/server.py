import asyncio
from concurrent import futures

import argparse

import grpc

import user_matching_pb2
import user_matching_pb2_grpc

from matching_service import MatchingService


class MatcherServicer(user_matching_pb2_grpc.MatcherServicer):

    def __init__(self, matching_service):
        self.matching_service = matching_service
        self.recalc_loop = asyncio.new_event_loop()

    def GetMatches(self, request, context):
        return user_matching_pb2.GetMatchesReply(
            user_ids=self.matching_service.get_matches(request.user_id)
        )

    def RecalcScaNN(self, request, context):
        self.recalc_loop.create_task(self.matching_service.recalc_scann())
        return user_matching_pb2.Empty()


async def serve(port):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

    matching_service = await MatchingService.create()

    servicer = MatcherServicer(matching_service)

    user_matching_pb2_grpc.add_MatcherServicer_to_server(servicer, server)

    server.add_insecure_port(f"[::]:{port}")
    server.start()

    print("Server started")
    print(f"Listening on port {port}")

    server.wait_for_termination()


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="User mathing service")
    parser.add_argument("--port", dest="port", metavar="PORT", type=int, required=True, help="server port")

    args = parser.parse_args()

    asyncio.run(serve(args.port))

from concurrent import futures
import logging

import grpc
import myproto_pb2
import myproto_pb2_grpc
import spacy


class ProtoTest(myproto_pb2_grpc.MyProto):

    def ReturnVector(self, request, context):
        nlp = spacy.load('ru_core_news_md')
        doc = nlp(request.description)
        return myproto_pb2.VectorReply(vector=doc.vector.tolist())


def serve():
    port = '50051'
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=100))
    myproto_pb2_grpc.add_MyProtoServicer_to_server(ProtoTest(), server)
    server.add_insecure_port('[::]:' + port)
    server.start()
    print("Server started, listening on " + port)
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    print("PyServer is running")
    serve()

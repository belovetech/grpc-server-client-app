#!/usr/bin/env python3
import grpc
from gen.hello import hello_pb2
from gen.hello import hello_pb2_grpc


def run(name: str = 'Abeeb'):
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = hello_pb2_grpc.GreeterStub(channel)
        # make rpc call
        response = stub.SayHello(hello_pb2.HelloRequest(name=name))
        print("Server response: " + response.message)


if __name__ == '__main__':
    names = ['Abeeb', 'John', 'Doe', 'Jane']
    for name in names:
        print(f"Calling server with name: {name}")
        run(name)
    print("Client finished calling server")

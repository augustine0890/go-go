#!/bin/bash

protoc --proto_path=greetpb greetpb/*.proto --go_out=greetpb --go-grpc_out=greetpb
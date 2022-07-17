- Install protocol buffer compiler, `protoc`
  - `brew install protobuf`
  - `protoc --version`

## Greet
- Generate the stubs:
  - `protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.`
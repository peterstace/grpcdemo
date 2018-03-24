# gRPC Example

Contains a gRPC server and client for a simple "add two numbers together" RPC
call.

To regenerate the `.pb.go` file after changing the IDL, run the following from
the root directory of the repo: 

```
protoc calc/calc.proto --go_out=plugins=grpc:.
```

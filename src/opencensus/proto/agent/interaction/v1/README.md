# Interaction Service

How to regenerate the interaction grpc-gateway protos and OpenApi config:
```bash
cd $GOPATH/src && protoc -I/usr/local/include -I. \
  -I./github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I./github.com/census-instrumentation/opencensus-proto/src \
  --plugin=protoc-gen-grpc-gateway=${GOPATH}/bin/protoc-gen-grpc-gateway \
  --go_out=plugins=grpc:. \
  --grpc-gateway_out=logtostderr=true,grpc_api_configuration=./github.com/census-instrumentation/opencensus-proto/src/opencensus/proto/agent/interaction/v1/interaction_http.yaml:. \
  --swagger_out=logtostderr=true,grpc_api_configuration=./github.com/census-instrumentation/opencensus-proto/src/opencensus/proto/agent/interaction/v1/interaction_http.yaml:. \
  github.com/census-instrumentation/opencensus-proto/src/opencensus/proto/agent/interaction/v1/interaction_service.proto
```

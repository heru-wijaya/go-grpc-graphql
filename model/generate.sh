#!/bin/bash

echo "starting generate proto"
protoc --go_out=./model/pb --go_opt=module=github.com/heru-wijaya/go-grpc-skeleton/model/pb  ./model/*.proto
protoc --go-grpc_out=./model/pb --go-grpc_opt=module=github.com/heru-wijaya/go-grpc-skeleton/model/pb ./model/*.proto
echo "successfully generate proto"
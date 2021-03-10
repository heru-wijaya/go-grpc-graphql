# GO-GRPC-SKELETON

this is a simple project skeleton using golang, grpc, postgreSQL, and protobuf

## how to test
1. install postgreSQL
2. create database accounts
3. run up.sql to create table
4. install docker
5. build the docker image using this command
> docker build -t account .
6. run the image (you can edit the port if you want, network host is for connecting to localhost database)
> docker run --network="host" account
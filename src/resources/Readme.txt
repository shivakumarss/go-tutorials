
export PATH=$PATH:/usr/local/go/bin:/home/hduser/go/bin

protoc -I=/home/hduser/Downloads/protobuf/src --go_out=/home/hduser/Downloads/protobuf/dest/ /home/hduser/Downloads/protobuf/src/SearchRequest.proto

protoc -I=/home/hduser/Downloads/protobuf/src --go_out=/home/hduser/Downloads/protobuf/dest/ /home/hduser/Downloads/protobuf/src/SearchRequest.proto







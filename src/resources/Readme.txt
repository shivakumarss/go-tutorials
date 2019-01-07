
export PATH=$PATH:/usr/local/go/bin:/home/hduser/go/bin

protoc -I=/home/hduser/Downloads/protobuf/src --go_out=/home/hduser/Downloads/protobuf/dest/ /home/hduser/Downloads/protobuf/src/SearchRequest.proto

protoc -I=/home/hduser/Downloads/protobuf/src --go_out=/home/hduser/Downloads/protobuf/dest/ /home/hduser/Downloads/protobuf/src/book.proto

in GOPATH these *.pb.go files needs to be present for referencing




export PATH=$PATH:/usr/local/go/bin:/home/hduser/go/bin

export PATH=$PATH:/usr/local/lib/python3.5/dist-packages

export PYTHONPATH=$HOME/dirWithScripts/:$PYTHONPATH


/usr/local/lib/python3.5/dist-packages

export PYTHONPATH=/usr/local/lib/python3.5/dist-packages

sys.path.append('/usr/local/lib/python3.5/dist-packages')
sys.path.append('/usr/local/lib/python3.5/dist-packages/google')




FROM golang:1.17
EXPOSE 8080
WORKDIR /go/src
RUN apt update
RUN apt install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
ENV GOPATH=/go/bin

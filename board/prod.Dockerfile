FROM golang:1.17 as builder
WORKDIR /go/src
COPY ./board /go/src
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go install
RUN go build -o server

FROM alpine:latest
COPY --from=builder /go/src/server ./server
EXPOSE 8080
ENTRYPOINT ["./server"]

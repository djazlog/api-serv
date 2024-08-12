FROM golang:1.23rc1-alpine AS builder

COPY . /grpc/source/
WORKDIR /grpc/source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /grpc/source/bin/crud_server .

CMD ["./crud_server"]
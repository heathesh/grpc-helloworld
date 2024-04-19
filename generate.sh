protoc --go_out=healthcheck --go_opt=paths=source_relative \
    --go-grpc_out=healthcheck --go-grpc_opt=paths=source_relative \
    *.proto
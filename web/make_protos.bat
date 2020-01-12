@protoc --proto_path=../proto --js_out=import_style=commonjs,binary:src/pb --grpc-web_out=import_style=commonjs,mode=grpcwebtext:src/pb ../proto/*.proto

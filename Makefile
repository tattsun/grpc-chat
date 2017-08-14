grpc:
	protoc -I proto/ proto/chat.proto \
	--go_out=plugins=grpc,import_path=main:server \
	--ts_out=service=true:client/src/api \
    --js_out=import_style=commonjs,binary:client/src/api

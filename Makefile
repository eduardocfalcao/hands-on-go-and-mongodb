protoc:
	protoc -I. src/proto/sweatmgr/*.proto --go_out=plugins=grpc:.
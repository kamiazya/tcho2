
proto:
	protoc -I ./contexts/tag-manager/interfacses/grpc ./contexts/tag-manager/interfacses/grpc/*.proto --go_out=plugins=grpc:./contexts/tag-manager/interfacses/grpc/

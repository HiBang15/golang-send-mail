genpb:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:proto/pb

sqlc:
	cd db/mail; sqlc generate; cd ../..;
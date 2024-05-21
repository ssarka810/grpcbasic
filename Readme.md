go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

#run tthe protoc commands to generate proto files
protoc --go_out=. --go_opt=module=<go_module> --go-grpc_out=. --go-grpc_opt=module=<go_module> greet/proto/\*.proto
protoc --go_out=. --go_opt=module=github.com/ss810n/go-grpc --go-grpc_out=. --go-grpc_opt=module=github.com/ss810n/go-grpc greet/proto/\*.proto

#if the above command does not work properly and you are in mac then try these commands
brew install protoc-gen-go
brew install protoc-gen-go-grpc
protoc-gen-go --version

#run
protoc --go_out=. --go_opt=module=github.com/ss810n/go-grpc --go-grpc_out=. --go-grpc_opt=module=github.com/ss810n/go-grpc greet/proto/\*.proto

#Generate CA private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -nodes -days 365 -keyout ca-key.pem -out ca-cert.pem -subj "/C=IN/ST=WB/L=KOL/O=DEV/OU=ENG/CN=\*.testssl/emailAddress=test@gmail.com"

#Generate Web Server’s Private Key and CSR (Certificate Signing Request)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=IN/ST=WB/L=KOL/O=DEV/OU=ENG/CN=\*.testssl/email=test@gmail.com"

# Generate a private key and a certificate

openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

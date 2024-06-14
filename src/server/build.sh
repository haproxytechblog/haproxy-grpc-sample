# Install Go
wget https://go.dev/dl/go1.22.4.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# Install protoc
sudo apt update && sudo apt install -y protobuf-compiler

# Install protoc to go generator
export GOPATH=/home/vagrant/go
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"

mkdir -p /home/vagrant/go/sample
cd /home/vagrant/go/sample
cp -r /vagrant/src/* ./

cd server/codenamecreator
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative ./*.proto

cd ../
go build -o ../output/server server.go
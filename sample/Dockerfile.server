FROM golang:alpine AS build

RUN apk add git protobuf
RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go

# Copy files to container
WORKDIR /go/src/app
COPY . .

# Build proto file
WORKDIR /go/src/app/codenamecreator
RUN protoc --go_out=plugins=grpc:. *.proto

# Build app
WORKDIR /go/src/app/
RUN go build -o /output/server ./server/server.go




FROM golang:alpine
EXPOSE 3000
WORKDIR /app
COPY --from=build /output/server .
COPY ./creds/*.crt ./creds/*.key ./
ENTRYPOINT ["./server"]
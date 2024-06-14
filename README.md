# HAProxy gRPC Sample

Demonstrates proxying gRPC traffic with HAProxy.

## Set Up

Run:

```
docker compose build
docker compose up
```

You should see the client connect through HAProxy to the gRPC server and get a stream of "codenames".

## Other info

The `haproxy.crt` and `haproxy.key` are generated with a SAN of "haproxy", which is required since CN allows only domain names.

```
openssl req -newkey rsa:2048 -nodes -x509 -days 3650 -keyout haproxy.key -out haproxy.crt -subj /CN=haproxy -addext "subjectAltName=DNS:haproxy"
```
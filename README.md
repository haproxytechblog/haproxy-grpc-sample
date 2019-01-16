# HAProxy gRPC Sample

Demonstrates proxying gRPC traffic with HAProxy.

## Set Up

Be sure that you have [Docker](https://docs.docker.com/v17.12/install/) and [Docker Compose](https://docs.docker.com/compose/install/) (version 3 or newer) installed. Then, run:

```
docker-compose build
docker-compose up
```

You should see the client connect through HAProxy to the gRPC server and get a stream of "codenames".
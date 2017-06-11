# Mailer Microservice

GRPC Micro Container Mail Service developed in GO. This is a small tutorial on how to create Micro containers (< 12 mb) running a GRPC powered server-client in Go language.

## Why GO?

* Very simple, elegant and performant language
* Building process generates small binaries ideal for micro containers
* Fun to learn

## Why Docker EDGE?

* Supports multi stage builds. This means in the same Dockerfile multiple images can be generated, this is ideal for creating first a builder container for your sourcecode which then generates an artifact (binary) that will be the content of our final "production ready" image, very lean.

## Prerrequisites

* Download and install [Docker](https://docs.docker.com/docker-for-mac/install/) (Multi Stage Dockerfile support needed, Docker EDGE recommended)

## Getting Started

These instructions will get you running the code from go and building the images from Docker.

Download the source
```
go get github.com/gianksp/mailer
```

Run server

```
cd ~/go/src/github.com/gianksp/mailer/server
go run main.go
```

Run client

```
cd ~/go/src/github.com/gianksp/mailer/client
go run main.go
```

Client will start sending loop requests to the server via GRPC and will receive the request object as the response by the server.

### Dockerizing

Now it's time to generate our Micro containers

From the root of the mailer project
```
cd ~/go/src/github.com/gianksp/mailer
```

Build the docker server images
```
docker build . -f Dockerfile.server --rm=true -t mailer/service:1.0 ; docker rmi $(docker images -f "dangling=true" -q)
```
Build the docker client images
```
docker build . -f Dockerfile.client --rm=true -t mailer/client:1.0 ; docker rmi $(docker images -f "dangling=true" -q)
```
The commands above generate the following images. A couple of beautiful 10Mb images for the server and the client
![Micro containers](http://i.imgur.com/nUmqH1A.png)

Before we can see them in action within their docker containers we must make sure that they have clear connectivity. We generate
a Docker network by the name of mailernet
```
docker network create --subnet=172.18.0.0/16 mailernet
```

Run the docker server image assigning a static ip within the previous network
```
docker run -i --net mailernet --ip 172.18.0.22 --name mailer_service mailer/service:1.0
```
Run the docker client image, set the env var SRV for the client to connect to the server
```
docker run -i --net mailernet --name mailer_client -e SRV='172.18.0.22' mailer/client:1.0
```

This will trigger a loop through GRPC againt the server for you to observe the flow of information.

## Additional Commands

Generate the proto go class for grpc if not available
```
protoc -I mail/ mail/proto/mail.proto --go_out=plugins=grpc:mail
```

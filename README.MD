# Project Title

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system. The following class has already been generated but knowing how to do it does not harm

Generate the proto go class for grpc if not available
```
protoc -I mail/ mail/proto/mail.proto --go_out=plugins=grpc:mail
```

### Running

Build the docker server images
```
docker build . -f Dockerfile.server --rm=true -t mailer/service:1.0 ; docker rmi $(docker images -f "dangling=true" -q)
```
Build the docker client images
```
docker build . -f Dockerfile.client --rm=true -t mailer/client:1.0 ; docker rmi $(docker images -f "dangling=true" -q)
```

Create network
```
docker network create --subnet=172.18.0.0/16 mailernet
```

Run the docker server image
```
docker run -i --net mailernet --ip 172.18.0.22 --name mailer_service mailer/service:1.0
```
Run the docker client image
```
docker run -i --net mailernet --name mailer_client -e SRV='172.18.0.22' mailer/client:1.0
```

## Built With

* [Dropwizard](http://www.dropwizard.io/1.0.2/docs/) - The web framework used
* [Maven](https://maven.apache.org/) - Dependency Management
* [ROME](https://rometools.github.io/rome/) - Used to generate RSS Feeds

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).

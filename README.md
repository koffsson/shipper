# Microservices in Golang

This is my implementation of the microservices-in-golang blog. It is a
10-step guide on how to create a shipping container management
platform. The first blog post can be found [here](https://ewanvalentine.io/microservices-in-golang-part-1/).

## Getting Started (IN NEED OF UPDATE(S))

These instructions will get you a copy of the project up and running on your
local machine for development and testing purposes. See deployment for notes
on how to deploy the project on a live system.

## Prerequisites

### Part 1

* [gRPC/protobuf](https://grpc.io/docs/quickstart/go.html) - Googles remote procedure call
and the accompanying protocol buffer. Follow the instructions to get gRPC up and running.

### Part 2

* [go-micro](https://github.com/micro/go-micro) - A pluggable gRPC framework. After gRPC/protobuf is installed, run the following command:

```
go get -v -u github.com/micro/protobuf/{proto,protoc-gen-go}
```
### Part 3

* [docker-compose](https://docs.docker.com/compose/overview/) - Compose is a tool for defining and running multi-container Docker applications. Run this command to download the latest version of Docker Compose:

```
sudo curl -L https://github.com/docker/compose/releases/download/1.21.2/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
```
* Apply executable permissions to the binary:

```
sudo chmod +x /usr/local/bin/docker-compose
```
* Optionally, we can install (bash) command completion:

```
sudo curl -L https://raw.githubusercontent.com/docker/compose/1.21.2/contrib/completion/bash/docker-compose -o /etc/bash_completion.d/docker-compose
```
* To test the installation, run:

```
docker-compose --version
```
* [gopkg.in/mgo.v2](https://gopkg.in/mgo.v2) - Rich MongoDB driver for Go. Install by simply running following command:

```
go get gopkg.in/mgo.v2
```
* [bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt) - Package bcrypt implements Provos and Mazières's bcrypt adaptive hashing algorithm. Install by running:

```
go get golang.org/x/crypto/bcrypt
```
* [GORM](http://gorm.io) - Golang Object Relational Mapper (GORM). Install by running:

```
go get github.com/jinzhu/gorm
```
* <del>[go.uuid](https://github.com/satori/go.uuid) - A pure Golang implementation of Universal Unique Identifier [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier). Install by running:</del>
<del>

```
go get github.com/satori/go.uuid
```
</del>

* [gouuid](https://github.com/nu7hatch/gouuid) - A pure Golang implementation of Universal Unique Identifier [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier). Install by running:

```
go get go get github.com/nu7hatch/gouuid
```

* [jwt-go](https://github.com/dgrijalva/jwt-go) - Golang implementation of [JSON Web Tokens](https://jwt.io/introduction). Install by running:

```
go get github.com/dgrijalva/jwt-go
```
### Installing (IN NEED OF UPDATE(S))

A step by step series of examples that tell you have to get a development env running

Say what the step will be

```
Give the example
```

And repeat

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo

## Running the tests (IN NEED OF UPDATE(S))

Explain how to run the automated tests for this system

### Break down into end to end tests (IN NEED OF UPDATE(S))

Explain what these tests test and why

```
Give an example
```

### And coding style tests (IN NEED OF UPDATE(S))

Explain what these tests test and why

```
Give an example
```

## Deployment (IN NEED OF UPDATE(S))

Add additional notes about how to deploy this on a live system

## Built With (IN NEED OF UPDATE(S))

* [Golang](https://golang.org/)
* [gRPC/protobuf](https://grpc.io/docs/quickstart/go.html)
* [MongoDB](https://www.mongodb.com/)
* [Docker](https://www.docker.com/)
* [Google Cloud](https://cloud.google.com/)
* [Kubernetes](https://kubernetes.io/)
* [NATS](https://nats.io/)
* [CircleCI](https://circleci.com/)
* [Terraform](https://www.terraform.io/)
* [go-micro](https://github.com/micro/go-micro)



## Contributing (IN NEED OF UPDATE(S))

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning (IN NEED OF UPDATE(S))

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).

## Authors (IN NEED OF UPDATE(S))

* **Billie Thompson** - *Initial work* - [PurpleBooth](https://github.com/PurpleBooth)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License (IN NEED OF UPDATE(S))

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments (IN NEED OF UPDATE(S))

* Hat tip to anyone who's code was used
* Inspiration
* etc

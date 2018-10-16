# password manager service
go server for safely storing and managing encrypted passwords

![architecture](docs/diagrams/architecture.png)

# Development

### Requirements

The following software must be installed before loading up the project.

- [go](https://golang.org/doc/install). Please also [setup GOBIN and GOPATH](https://github.com/golang/go/wiki/SettingGOPATH)
- [fabric](https://github.com/DecipherNow/gm-fabric-go#installation)
- [docker cli](https://docs.docker.com/install)
- [node](https://nodejs.org/en/download/)
- [aglio](https://github.com/danielgtaylor/aglio)

Currently the following operating systems are tested and supported

- OS X 10.13 High Sierra
- Ubuntu 18
- Pop!OS 17

### Setting up the Project

Clone project

```sh
# create project path
mkdir -p ${GOPATH}/src/github.com/dgoldstein1
cd ${GOPATH}/src/github.com/dgoldstein1
# clone using ssh. Make sure you've added your public key to gitlab
git clone git@github.com:dgoldstein1/passwordmanagerservice.git
```

#### Launch the Project

```sh
# start docker
docker-compose up -d
# see if all containers are up
docker-compose ps
```

### Developemnt

Any changes made in the `./cmd/server` directors are automatically applied to the running container and the service is restarted. If you need to change packages outside this directory (e.g. protobuf, docs, vendor), the following can be done:

```sh
# generate execuatbe artifact to $GOPATH/bin
./build_server.sh
# generate new protobuf definitions for entire app
./generate_proto.sh
# create rpm
./rpm/build.sh
# regenerate documentation from api blueprint files
./create_documentation.sh
```

### Deployment

Develop and Master are automatically deployed on merge.

# Testing

```sh
cd ./cmd/server/methods/
go test -v -cover -covermode=count -coverprofile=coverage.out
```

## Authors

* **David Goldstein** - [DavidCharlesGoldstein.com](http://www.davidcharlesgoldstein.com/) - [Decipher Technology Studios](http://deciphernow.com/)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
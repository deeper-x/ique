# iQue

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/cb8ee4ca23764801b9274e5d18755f38)](https://app.codacy.com/manual/deeper-x/ique?utm_source=github.com&utm_medium=referral&utm_content=deeper-x/ique&utm_campaign=Badge_Grade_Settings)

RabbitMQ with Golang. Producer->Consumer model, with queue 

![Go](https://github.com/deeper-x/ique/workflows/Go/badge.svg)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/7b9c3fd94126499098ace12437471384)](https://www.codacy.com/manual/deeper-x/ique?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=deeper-x/ique&amp;utm_campaign=Badge_Grade)


### Prerequisites

You need go, and docker

```bash
$ go version
go version go1.13.8 linux/amd64
$ docker version
Client: Docker Engine - Community
 Version:           19.03.5
 API version:       1.40
 Go version:        go1.12.12
 Git commit:        633a0ea838
 Built:             Wed Nov 13 07:29:52 2019
 OS/Arch:           linux/amd64
 Experimental:      false

Server: Docker Engine - Community
 Engine:
  Version:          19.03.5
  API version:      1.40 (minimum version 1.12)
  Go version:       go1.12.12
  Git commit:       633a0ea838
  Built:            Wed Nov 13 07:28:22 2019
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.2.10
  GitCommit:        b34a5c8af56e510852c35414db4c1f4fa6172339
 runc:
  Version:          1.0.0-rc8+dev
  GitCommit:        3e425f80a8c931f88e6d94a8c831b9d5aa481657
 docker-init:
  Version:          0.18.0
  GitCommit:        fec3683

```

Go setup is:

```bash
$ export GOPATH=${HOME}/go
$ export GOBIN=${GOPATH}/bin
$ export PATH=${PATH}:${GOBIN}
$ export GO111MODULE=on
$ go env
[...]
```

### Installing

Get the development env running
[TODO] Makefile, with build/install, run and test

```bash
$ git clone https://github.com/deeper-x/ique.git
$ cd ique
$ go get -d ./...
$ go build

```

### Deployment and usage (development)

[WIP] In this 1st release, sender push a default message (demo text) to a default queue (msg-qu). Best yet to come.

RabbitMQ server:
```bash
$ bash ./rabbitmq.sh
2020-02-26 11:02:40.246 [info] <0.277.0> 
 Starting RabbitMQ 3.8.2 on Erlang 22.2.7
 Copyright (c) 2007-2019 Pivotal Software, Inc.
 Licensed under the MPL 1.1. Website: https://rabbitmq.com

  ##  ##      RabbitMQ 3.8.2
  ##  ##
  ##########  Copyright (c) 2007-2019 Pivotal Software, Inc.
  ######  ##
  ##########  Licensed under the MPL 1.1. Website: https://rabbitmq.com

  Doc guides: https://rabbitmq.com/documentation.html
  Support:    https://rabbitmq.com/contact.html
  Tutorials:  https://rabbitmq.com/getstarted.html
  Monitoring: https://rabbitmq.com/monitoring.html

  Logs: <stdout>

  Config file(s): /etc/rabbitmq/rabbitmq.conf
```

Run consumer and sender:

```bash
# shell 1
$ go run main.go 
Please insert runner [sender/receiver]:sender
2020/02/26 18:11:27 Message sent on msg-qu: demo text

# shell 2
$ go run main.go 
Please insert runner [sender/receiver]:receiver
2020/02/26 18:11:00 Waiting for messages....
2020/02/26 18:11:27 Received: demo text

```

## Unit test

Run the automated tests for this system:

```bash
$ go test -v -cover ./...
```



## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/deeper-x/ique/tags). 

## Authors

* **Alberto de Prezzo** *


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details


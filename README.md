# docker-deploy-client
[![License MIT](https://img.shields.io/npm/l/express.svg)](http://opensource.org/licenses/MIT)
[![Build Status](https://travis-ci.org/composer22/docker-deploy-client.svg?branch=master)](http://travis-ci.org/composer22/docker-deploy-client)
[![Current Release](https://img.shields.io/badge/release-v0.0.1-brightgreen.svg)](https://github.com/composer22/docker-deploy-client/releases/tag/v0.0.1)
[![Coverage Status](https://coveralls.io/repos/composer22/docker-deploy-client/badge.svg?branch=master)](https://coveralls.io/r/composer22/docker-deploy-client?branch=master)

A CLI for interacting with docker-deploy-server written in [Go.](http://golang.org)

## About

This client application can be used to submit deploy requests to docker-deploy-server
or check the status of a previous deploy request to the server.

## Requirements

If you compile, you need the server source:

go get github.com/composer22/docker-deploy-server

You will also need a valid API token defined in the server to access the server.
See that documentation for more information.


## Usage

This command performs two functions:

* Submitting a deploy request to the docker-deploy-server.
* Retrieving the result of a previous deploy request.

When submitting a deploy request, a unique deploy ID is returned as reference for future
result queries.

```
$ ./docker-deploy-client --help
Client for making requests to a docker-deploy-server to deploy
 Docker containers into one or more machines (swarm cluster) and check the
status of a previous deploy request,

Usage:
  docker-deploy-client [command]

Available Commands:
  deploy      Deploy a Docker image to one or more machines.
  status      Retrieve status of a previous deploy

Flags:
      --config string          config file (default is $HOME/.docker-deploy-client.yaml)
  -f, --formatted              JSON indented status results (default true)
  -i, --poll_interval string   Polling interval for status check (default "5")
  -o, --token string           API Token
  -u, --url string             docker-deploy-server endpoint

Use "docker-deploy-client [command] --help" for more information about a command.

```

see command help for examples.

## Config file

An example config layout with comments is included in /examples.

rename your copy as .docker-deploy-client.yml

This can be placed in:
* the root directory of the user.
* the directory where you keep the docker-deploy-client binary.

## Building

This code was tested with version 1.6.2 or higher of Go.

Information on Golang installation, including pre-built binaries, is available at
<http://golang.org/doc/install>.

Run `go version` to see the version of Go which you have installed.

Run `go build` inside the directory to build.

Run `go test ./...` to run the unit regression tests.

A successful build run produces no messages and creates an executable called `docker-deploy-client` in this
directory.

Run `go help` for more guidance, and visit <http://golang.org/> for tutorials, presentations, references and more.

Run `./build.sh` for multiple platforms.

## License

(The MIT License)

Copyright (c) 2015 Pyxxel Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to
deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.

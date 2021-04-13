# pinger/Golang

A simple program written in Golang named `pinger`. This service reponds with `"hello world"` at the root path and can be configured to ping another server through the environment variables (see the file at `./cmd/pinger/config.go`). By default, running it will make it start a server and ping itself.

A basic `Makefile` is provided that allows you to:

- pull in dependencies - `make dep`
- builds the binaries - `make build`
- test runs - `make run`
- run tests - `make test`


## Pre-requisites

You will need the following installed:

- `go` to run the application (check with `go version`)
- `docker` for image building/publishing (check with `docker version`)
- `docker-compose` for environment provisioning (check with `docker-compose version`)
- `git` for source control (check with `git -v`)
- `make` for simple convenience scripts (check with `make -v`)


## Directory structure

| Directory | Description |
| --- | --- |
| `/bin` | Contains binaries |
| `/cmd` | Contains source code for CLI interfaces |
| `/deployments` | Contains image files and manifests for deployments |
| `/docs` | Contains documentation |
| `/vendor` | Contains dependencies (use `make dep` to populate it) |


## Get Started

1. Clone this repository
2. Create your own repository on GitLab
3. Set your local repository's remote to point to your GitLab repository
4. Make your changes locally
5. Push to your GitLab repository

## Build Docker image

 docker build -f ./deployments/build/Dockerfile -t devops/pinger:latest .

## Running pinger

 docker run -it -p 8000:8000 devops/pinger:latest

# Expected output:



## Exports the docker image as a tarball

 docker save -o ./build/pinger.tar devops/pinger:latest 

## Imports the docker image from a tarball

 docker load -i ./build/pinger.tar

## Creation of docker-compose.yml in the ./deployments to demonstrate two pinger services that ping each other

 docker-compose up  

# Expected output:




## Pipeline to automate building the docker image, exporting as tarball and loading the image from tarball as an artifacts

 cat ./.gitlab-ci.yml



## Done !



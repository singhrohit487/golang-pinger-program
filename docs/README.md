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

--------------------------------------
### Expected output:
service|2021/04/13 10:08:30 initialising service...
service|2021/04/13 10:08:30 attempting to listen on '0.0.0.0:8000'...
 server|2021/04/13 10:08:31 < localhost:8000 <- [::1]:42104 | HTTP/1.1 GET / 
service|2021/04/13 10:08:31 > http://localhost:8000/ -> '200 OK'
 server|2021/04/13 10:08:32 < localhost:8000 <- [::1]:42106 | HTTP/1.1 GET / 
service|2021/04/13 10:08:32 > http://localhost:8000/ -> '200 OK'
 server|2021/04/13 10:08:33 < localhost:8000 <- [::1]:42108 | HTTP/1.1 GET / 
service|2021/04/13 10:08:33 > http://localhost:8000/ -> '200 OK'
 server|2021/04/13 10:08:34 < localhost:8000 <- [::1]:42110 | HTTP/1.1 GET / 
service|2021/04/13 10:08:34 > http://localhost:8000/ -> '200 OK'
 server|2021/04/13 10:08:35 < localhost:8000 <- [::1]:42112 | HTTP/1.1 GET / 
service|2021/04/13 10:08:35 > http://localhost:8000/ -> '200 OK'
 server|2021/04/13 10:08:36 < localhost:8000 <- [::1]:42114 | HTTP/1.1 GET / 
service|2021/04/13 10:08:36 > http://localhost:8000/ -> '200 OK'


## Exports the docker image as a tarball

 docker save -o ./build/pinger.tar devops/pinger:latest 

## Creation of docker-compose.yml in the ./deployments to demonstrate two pinger services that ping each other

 docker-compose -f ./deployments/docker-compose.yml up  

-------------------------------------------
### Expected output:
Creating deployments_pinger1_1 ... done
Creating deployments_pinger2_1 ... done
Attaching to deployments_pinger2_1, deployments_pinger1_1
pinger2_1  | service|2021/04/13 14:39:26 initialising service...
pinger2_1  | service|2021/04/13 14:39:26 attempting to listen on '0.0.0.0:8000'...
pinger1_1  | service|2021/04/13 14:39:26 initialising service...
pinger1_1  | service|2021/04/13 14:39:26 attempting to listen on '0.0.0.0:8000'...
pinger2_1  |  server|2021/04/13 14:39:27 < localhost:8000 <- [::1]:42312 | HTTP/1.1 GET / 
pinger2_1  | service|2021/04/13 14:39:27 > http://localhost:8000/ -> '200 OK'
pinger1_1  |  server|2021/04/13 14:39:27 < localhost:8000 <- [::1]:42314 | HTTP/1.1 GET / 
pinger1_1  | service|2021/04/13 14:39:27 > http://localhost:8000/ -> '200 OK'
pinger2_1  |  server|2021/04/13 14:39:28 < localhost:8000 <- [::1]:42316 | HTTP/1.1 GET / 
pinger2_1  | service|2021/04/13 14:39:28 > http://localhost:8000/ -> '200 OK'
pinger1_1  |  server|2021/04/13 14:39:28 < localhost:8000 <- [::1]:42318 | HTTP/1.1 GET / 
pinger1_1  | service|2021/04/13 14:39:28 > http://localhost:8000/ -> '200 OK'


## Pipeline to automate building the binary and docker image, exporting as tarball as an artifacts

 cat ./.gitlab-ci.yml

# That's it !

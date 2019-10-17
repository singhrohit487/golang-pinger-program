dep: # installs dependencies
	GO111MODULE=on go mod vendor

build: # builds the binary
	GO111MODULE=on go build -mod vendor -o ./bin/pinger ./cmd/pinger

run: # runs the code in development
	GO111MODULE=on go run ./cmd/pinger

test: # runs tests
	GO111MODULE=on go test ./...

docker_image: # builds the docker image
	# DO NOT CHANGE THIS
	docker build -f ./deployments/build/Dockerfile -t devops/pinger:latest .

docker_testrun: # runs the docker container
	# DO NOT CHANGE THIS
	docker run -it -p 8000:8000 devops/pinger:latest

docker_tar: docker_image # exports the docker image as a tarball
	# DO NOT CHANGE THIS
	docker save -o ./build/pinger.tar devops/pinger:latest 

docker_untar: # imports the docker image from a tarball
	# DO NOT CHANGE THIS
	docker load -i ./build/pinger.tar

testenv: # runs the orchestrated containers
	# DO NOT CHANGE THIS
	docker-compose up -f ./deployments/docker-compose.yml -V

verify_pipeline: # verifies pipeline files are there
	# DO NOT CHANGE THIS
	cat ./.gitlab-ci.yml

verify_readme: # verifies readme file is there
	# DO NOT CHANGE THIS
	cat ./docs/README.md

done?:
	# checking documentation section...
	@$(MAKE) verify_readme
	# checking pipelining section...
	@$(MAKE) verify_pipeline
	# checking containerisation section...
	@$(MAKE) docker_image
	# checking environment section...
	@$(MAKE) docker_testrun
	@$(MAKE) testenv
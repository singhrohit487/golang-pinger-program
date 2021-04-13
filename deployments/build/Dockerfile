FROM golang:1.15.6-alpine3.12
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# workspace directory
WORKDIR /app

# copy `go.mod` and `go.sum`
ADD go.mod go.sum ./

# install dependencies
RUN go mod download

# copy source code
COPY ./cmd/pinger/ .

# build executable
RUN go build -o ./cmd/pinger .

# set entrypoint
ENTRYPOINT [ "./cmd/pinger" ]

FROM golang:alpine AS builder

ENV GO111MODULE=on

# Set the Current Working Directory inside the container
WORKDIR /app

RUN apk add git

#  Copying go.mod / go.sum files into docker first
COPY /go.mod .
COPY /go.sum .

# Download dependencies
RUN go mod download

# Copy everything from the current directory to the Working Directory inside the container
COPY . /app

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -o ./app/auth ./cmd/main.go

EXPOSE ${PORT_AUTH}

CMD ["./app/auth"]

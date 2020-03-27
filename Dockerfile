FROM golang:latest
LABEL maintainer="Nick Pocock"

# Set the Current Working Directory inside the container
WORKDIR /app

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o recs-api

ENTRYPOINT ["/recs-api"]
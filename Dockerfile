# syntax=docker/dockerfile:1
FROM golang:1.23.0 AS build
# FROM ubuntu:latest
# FROM alpine:3.14

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in. Keep slash at the end.
#COPY *.go ./
COPY . ./

# Download Go modules
RUN go mod download


# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/Application


#FROM scratch
FROM alpine:3.14

WORKDIR /app
RUN mkdir -p static
RUN mkdir -p logs
COPY ./web/*   /app/web/
COPY --from=build /app/Application /app/Application

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose

EXPOSE 2025

# Run
CMD ["/app/Application"]
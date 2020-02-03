ARG GO_VERSION=1.13.1
# First stage: build the executable.
FROM golang:${GO_VERSION} AS builder

MAINTAINER https://github.com/OrkhanHuseynli

WORKDIR /app
RUN ls
COPY ./ ./
WORKDIR /app/src
RUN echo "******** LIST DIR ************"
RUN ls
RUN echo "******** RUN BUILD ************"
RUN go build main.go
RUN ls
EXPOSE 8080
ENTRYPOINT ["/main"]
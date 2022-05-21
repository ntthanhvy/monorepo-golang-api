
# Example project for monorepo golang services

Project name: monorepo-golang-api


## Services

### Core-api

service for external request call into

### Process-data

Service for process data

## Build steps

1. Dockerfile

```dockerfile

FROM golang:1.17-alpine as build

WORKDIR /app
RUN mkdir -p ./services/<service-name>
RUN mkdir -p ./lib


COPY services/<service-name>/. ./services/<service-name>/
COPY lib ./lib

WORKDIR /app/services/<service-name>
RUN go mod tidy

RUN go build -o out

FROM alpine:3.10 as runner

WORKDIR /app

COPY --from=build /app/services/<service-name>/out ./

RUN chmod +x ./out

CMD [ "./out" ]

```
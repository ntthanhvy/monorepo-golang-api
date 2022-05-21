FROM golang:1.17-alpine as build

WORKDIR /app
RUN mkdir -p ./services/core-api
RUN mkdir -p ./lib


COPY services/core-api/. ./services/core-api/
COPY lib ./lib

WORKDIR /app/services/core-api
RUN go mod tidy

RUN go build -o out

FROM alpine:3.10 as runner

WORKDIR /app

COPY --from=build /app/services/core-api/out ./

RUN chmod +x ./out

CMD [ "./out" ]

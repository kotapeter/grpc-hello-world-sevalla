FROM golang:1.22.1-alpine as build

WORKDIR /build

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o /hello-client

FROM alpine as release

COPY --from=build /hello-client /hello-client

EXPOSE 50051
ENTRYPOINT ["/hello-client"]

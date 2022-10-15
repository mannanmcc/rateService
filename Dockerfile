# syntax=docker/dockerfile:1

FROM golang:1.16-alpine as build

WORKDIR /build

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod vendor -o /app ./cmd/server

ENTRYPOINT ["/app"]

FROM scratch

COPY --from=build /app /

ENTRYPOINT [ "/app" ]
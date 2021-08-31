FROM golang:1.17.0-alpine3.14 as builder

COPY . /src
WORKDIR /src

RUN go build -o app cmd/go-clean-boilerplate/main.go

FROM alpine:3.14.2 as runner

COPY --from=builder /src/app /opt/app/
WORKDIR /opt/app

ENTRYPOINT ["./app"]


# syntax=docker/dockerfile:1
FROM golang:1.18-alpine AS build

LABEL MAINTAINER = 'Raúl Fernández (info@raulfernandez.dev)'

RUN apk add --update git \
	&& apk add ca-certificates
WORKDIR /go/src/github.com/rfdez/diade
COPY . .
RUN go mod tidy && TAG=$(git describe --tags --always) \
    && LDFLAGS=$(echo "-s -w -X main.version="$TAG) \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/diade-api -ldflags "$LDFLAGS" cmd/diade-api/main.go

FROM scratch
COPY --from=build /go/bin/diade-api /go/bin/diade-api
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
ENTRYPOINT ["/go/bin/diade-api"]

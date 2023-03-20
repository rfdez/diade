# syntax=docker/dockerfile:1.2
FROM golang:1.20-alpine AS build

RUN apk --no-cache --no-progress add git ca-certificates tzdata make \
    && update-ca-certificates \
    && rm -rf /var/cache/apk/*

# syntax=docker/dockerfile:1.2
FROM scratch

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY diade /

ENTRYPOINT ["/diade"]

EXPOSE 8080

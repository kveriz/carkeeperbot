ARG GO_VERSION="1.20.5"
ARG PLATFORM="linux"
ARG OS="arm64"
FROM --platform=${PLATFORM} golang:${GO_VERSION}-alpine AS build
WORKDIR /app

COPY ../.. ./

# build app
RUN apk add --no-cache ca-certificates

RUN go mod download

RUN go env -w GO111MODULE=on && \
    CGO_ENABLED=0 GOARCH=${PLATFORM} GOOS=${OS} go build -o /bot ./cmd/bot/main.go

##
## Deploy
##
FROM --platform=${PLATFORM} scratch

WORKDIR /

COPY --from=build /bot /bot
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/database/migrations /migrations

# CMD ["/bot"]
ENTRYPOINT ["/bot", "-conf", "/config.yaml"]
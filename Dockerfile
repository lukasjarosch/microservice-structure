# Accept the Go version for the image to be set as a build argument.
# Default to Go 1.12
ARG GO_VERSION=1.12

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine AS builder

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
RUN apk add --no-cache ca-certificates

WORKDIR /build
COPY . .

# build binary
RUN CGO_ENABLED=0 GOFLAGS=-mod=vendor GOOS=linux GOARCH=amd64 go build  -ldflags="-w -s" -o /bin/service ./cmd/service/main.go

# runtime container image
FROM alpine

# Load binary from builder
COPY --from=builder /bin/service /service
#COPY ./migrations /migrations

# ensure ca-certs
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true


ENTRYPOINT ["/service"]

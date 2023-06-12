# -----------------------------------------------------------------------------
#  Build Stage
# -----------------------------------------------------------------------------
FROM golang:alpine AS build

ENV CGO_ENABLED=1

RUN apk add --no-cache \
    gcc \
    musl-dev

WORKDIR /workspace

COPY . /workspace/

RUN \
    go mod init github.com/mattn/sample && \
    go mod tidy && \
    go install -ldflags='-s -w -extldflags "-static"' ./simple.go

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM scratch

COPY --from=build /go/bin/secretary /usr/local/bin/secretary

ENTRYPOINT [ "/usr/local/bin/secretary" ]

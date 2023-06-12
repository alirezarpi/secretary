# -----------------------------------------------------------------------------
#  Build Stage
# -----------------------------------------------------------------------------
FROM golang:alpine AS build

ENV CGO_ENABLED=1

WORKDIR /workspace

RUN apk add --no-cache \
    gcc \
    musl-dev

COPY . .

RUN go mod init github.com/alirezarpi/secretary && \
    go mod tidy && \
    go build ./gateway/main.go -o ./secretary

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM scratch

COPY --from=build /go/bin/secretary /secretary/

VOLUME /secretary/

ENTRYPOINT [ "/secretary/secretary" ]

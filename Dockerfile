# -----------------------------------------------------------------------------
#  Build Stage
# -----------------------------------------------------------------------------
FROM golang:alpine AS build

ENV CGO_ENABLED=1
ENV CGO_CFLAGS="-D_LARGEFILE64_SOURCE"

WORKDIR /workspace

RUN apk add --no-cache \
    gcc \
    musl-dev

COPY . .

RUN go mod tidy && \
    go build -o ./secretary ./gateway/main.go 

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM scratch

WORKDIR /secretary/

COPY --from=build /workspace/secretary /secretary/

ENTRYPOINT [ "/secretary/secretary" ]

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

RUN go mod tidy && \
    go build -o ./secretary ./gateway/main.go 

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM scratch

COPY --from=build /workspace/secretary /secretary/

VOLUME /secretary/

ENTRYPOINT [ "/secretary/secretary" ]

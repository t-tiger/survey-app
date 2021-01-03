# syntax = docker/dockerfile:experimental
# =====================
FROM golang:1.15-alpine3.12 as base
RUN apk add --no-cache make gcc musl-dev
WORKDIR /go/src/github.com/t-tiger/survey/server/

# =====================
FROM base as builder

COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile

RUN --mount=type=cache,id=surver-server-pkg,target=/go/pkg \
  make mod-download

COPY . .
RUN --mount=type=cache,id=surver-server-pkg,target=/go/pkg \
    --mount=type=cache,id=surver-server-go-build,target=/root/.cache/go-build \
  make && make install

# =====================
FROM alpine:3.12 as app
WORKDIR /app
COPY --from=builder /go/src/github.com/t-tiger/survey/server/bin/survey-server /app/bin/survey-server

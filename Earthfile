VERSION 0.8
FROM golang:1.22-alpine
WORKDIR /go-workdir

unit-test:
  BUILD ./libs/db-package+unit-test
  # BUILD ./apps/merch-rack+unit-test

compile:
  BUILD ./apps/merch-rack+compile

image:
  BUILD ./apps/merch-rack+image
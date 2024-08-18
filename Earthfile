VERSION 0.8
FROM golang:1.22-alpine

test:
  BUILD ./libs/db-package+test
  # BUILD ./apps/merch-rack+test

build:
  BUILD ./apps/merch-rack+build

image:
  ARG buildenv
  BUILD ./apps/merch-rack+image --buildenv=$buildenv
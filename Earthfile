VERSION 0.8
FROM golang:1.22-alpine

unit-test:
  BUILD ./libs/db-package+unit-test
  # BUILD ./apps/merch-rack+unit-test

build:
  BUILD ./apps/merch-rack+build

image:
  ARG buildenv
  BUILD ./apps/merch-rack+image --buildenv=$buildenv
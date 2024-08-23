VERSION 0.8

test:
  BUILD ./libs/db-package+test
  # BUILD ./apps/merch-rack+test

build:
  BUILD +test
  BUILD ./apps/merch-rack+build

image:
  ARG buildenv
  BUILD ./apps/merch-rack+image --buildenv=$buildenv
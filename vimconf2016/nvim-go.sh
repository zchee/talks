#!/bin/bash
set -e

# fmt
# run
# rename
# test
# build
# testswitch
# guru
# delve

cd vimconf2016/go/fmt \
  && nvim gofmt.go \
  \
  && cd ../run \
  && nvim gorun.go \
  \
  && cd ../rename \
  && nvim gorename.go \
  \
  && cd $GOPATH/src/github.com/mattn/go-scan \
  && nvim scan.go \
  \
  && cd $GOPATH/src/github.com/zchee/nvim-go/src/nvim-go/commands \
  && nvim build.go \
  && nvim build_test.go \
  && nvim guru.go \
  \
  && cd $GOPATH/src/github.com/zchee/dlvtest \
  && nvim dlvtest.go

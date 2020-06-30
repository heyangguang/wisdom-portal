#!/usr/bin/env bash
export GO111MODULE=off

GOOS=linux GOARCH=amd64 go build  \
	  -o wisdoms-ctl
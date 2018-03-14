#!/bin/bash
packr -z
GOOS=linux GOARCH=amd64 go build
packr clean

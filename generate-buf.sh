#!/bin/bash

buf generate
echo "Generated protobuf files finished"
go run ./cmd/process_id_inject/main.go

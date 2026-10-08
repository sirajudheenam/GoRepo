#!/bin/bash

set -e

go run rest/rest_server.go
go run grpc/grpc_server.go



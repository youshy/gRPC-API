#!/bin/bash

protoc calculate.proto --go_out=plugins=grpc:.

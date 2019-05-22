#!/bin/bash

protoc greet/greetpd/greet.proto --go_out=plugins=grpc:. 
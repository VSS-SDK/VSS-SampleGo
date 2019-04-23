#!/bin/bash

COMPILE_PROTOS() {
    protoc -I=. --go_out=. protos/state.proto
    protoc -I=. --go_out=. protos/command.proto
    protoc -I=. --go_out=. protos/debug.proto
    protoc -I=. --go_out=. protos/control.proto
}

INSTALL () {
    COMPILE_PROTOS;
 }

INSTALL;
#!/bin/bash

COMPILE_PROTOS() {
    protoc -I=. --go_out=. protos/vss_protos.proto
}

INSTALL () {
    COMPILE_PROTOS;
 }

INSTALL;
#!/usr/bin/env bash

cd canary; ./make_protos.sh; cd ..
cd storage; ./make_protos.sh; cd ..

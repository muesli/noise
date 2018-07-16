#!/bin/sh

cd `dirname $0`
vgo build -buildmode c-shared -o target/libnoise.so github.com/perlin-network/noise/libnoise

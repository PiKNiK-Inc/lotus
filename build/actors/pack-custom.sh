#!/bin/bash

VERSION=$1
rm -f "${VERSION}.tar.zst"
tar -cf "${VERSION}.tar.zst" --use-compress-program "zstd -19" -- *.car

make -C ../../ bundle-gen

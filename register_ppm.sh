#!/bin/sh

IMG_DIR="${PWD}/images"
BINARY_NAME=ppm-search-engine

if [ $# -ge 2 ] ; then
  BINARY_NAME=$1
fi

for CATEGORY in `ls ${IMG_DIR}` ; do
  for IMAGE in `ls ${IMG_DIR}/${CATEGORY}` ; do
    cd ${PWD} && ./${BINARY_NAME} register images/${CATEGORY}/${IMAGE}
  done
done

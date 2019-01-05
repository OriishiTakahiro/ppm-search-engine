#!/bin/sh

SRC_DIR="${PWD}/tmp/101_ObjectCategories"   # Default value
DST_DIR="${PWD}/images/"                    # Default value

if [ `$# -lt 2` ] ;  then
  SRC_DIR=$1
  SRC_DIR=$2
fi

echo "convert ${SRC_DIR}/*.jpg to ${DST_DIR}/*.ppm"

for CATEGORY in ${SRC_DIR}/* ; do
  mkdir -p ${DST_DIR}/${CATEGORY}/${CATEGORY}
  for IMAGE in ${SRC_DIR}/${CATEGORY}/* ; do
    convert ${SRC_DIR}/${CATEGORY}/${IMAGE}.jpg -quality 100 ${DST_DIR}/${CATEGORY}/${CATEGORY}/${IMAGE}.ppm
  done
done

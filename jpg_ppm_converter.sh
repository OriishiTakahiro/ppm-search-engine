#!/bin/sh

SRC_DIR="${PWD}/tmp/101_ObjectCategories"   # Default value
DST_DIR="${PWD}/images"                    # Default value

if [ $# -ge 3 ] ;  then
  SRC_DIR=$1
  DST_DIR=$2
fi

echo "convert ${SRC_DIR}/*.jpg to ${DST_DIR}/*.ppm"

for CATEGORY in `ls ${SRC_DIR}` ; do
  mkdir ${DST_DIR}/${CATEGORY}
  for IMAGE in `ls ${SRC_DIR}/${CATEGORY}` ; do
    FILE_NAME=`basename ${SRC_DIR}/${IMAGE} .jpg`
    echo "convert ${CATEGORY}/${FILE_NAME}.jpg to ${CATEGORY}/${FILE_NAME}.ppm"
    convert ${SRC_DIR}/${CATEGORY}/${FILE_NAME}.jpg -quality 100 ${DST_DIR}/${CATEGORY}/${FILE_NAME}.ppm
  done
done

#!/bin/sh

cd appengine

if [ -e vendor/lib ]; then
  rm vendor/lib
fi

dep ensure $*

if [ ! -e vendor/lib ]; then
  ln -s ../src/lib vendor/lib
fi

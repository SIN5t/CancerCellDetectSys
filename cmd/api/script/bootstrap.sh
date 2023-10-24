#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=upload
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}
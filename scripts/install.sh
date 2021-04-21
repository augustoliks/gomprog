#!/bin/bash

VERSION="0.1.8"
TAG_VERSION="v0.1.8"

PATH_BIN=/usr/local/bin/
PATH_DOC=/usr/share/doc/gomprog/

PACKAGE=gomprog_$(echo $VERSION)_linux_$(uname -i).zip
PACKAGE_LINK=https://github.com/augustoliks/gomprog/releases/download/$TAG_VERSION/$PACKAGE

TEMP_DIRECTORY=$(mktemp -d)

echo "
>>> Download Package
>>>     Package: $PACKAGE
>>>     Package Link: $PACKAGE_LINK

    + wget -q $PACKAGE_LINK -O "$TEMP_DIRECTORY/$PACKAGE"
"
wget -q $PACKAGE_LINK -O "$TEMP_DIRECTORY/$PACKAGE"

echo "
>>> Unzip Package:

    + unzip $TEMP_DIRECTORY/$PACKAGE -d $TEMP_DIRECTORY/
"
unzip -qq $TEMP_DIRECTORY/$PACKAGE -d $TEMP_DIRECTORY/

echo "
>>> Copy Binary to $PATH_BIN:

    + cp $TEMP_DIRECTORY/gomprog-* $PATH_BIN
"
cp $TEMP_DIRECTORY/gomprog-* $PATH_BIN

echo "
>>> Copy Doc file to $PATH_DOC:

    + cp $TEMP_DIRECTORY/README.md $PATH_DOC
"
mkdir -p $PATH_DOC
cp $TEMP_DIRECTORY/README.md $PATH_DOC

echo "
* remove $TEMP_DIRECTORY

    + rm -rf $TEMP_DIRECTORY
"
rm -rf $TEMP_DIRECTORY

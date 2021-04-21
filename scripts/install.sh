#!/bin/bash

# https://github.com/augustoliks/gomprog/releases/download/0.1.5/gomprog_0.1.5_linux_64-bit.zip

VERSION="v0.1.6"

PATH_BIN=/usr/local/bin/
PATH_DOC=/usr/share/doc/gomprog/

PACKAGE=gomprog_$VERSION_linux_$(uname -i).zip
PACKAGE_LINK=https://github.com/augustoliks/gomprog/releases/download/$VERSION/$PACKAGE

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
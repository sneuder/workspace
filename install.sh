#!/bin/bash

VERSION=1.0.0

wget "https://github.com/sneuder/workspace/releases/download/v$VERSION/wkspace-$VERSION.tar.gz" -O "wkspace-$VERSION.tar.gz"

tar -xzf "wkspace-$VERSION.tar.gz"

sudo mv ./wkspace /usr/local/bin/wkspace

rm "wkspace-$VERSION.tar.gz"

echo "Installation completed for wkspace version $VERSION."
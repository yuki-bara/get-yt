#!/usr/bin/bash

mkdir -p ~/get-yt-install/dev || exit 1

cleanup() {
    trap '' SIGINT SIGHUP SIGTERM EXIT
    rm -r ~/get-yt-install
}

trap "cleanup > /dev/null 2>&1" SIGINT SIGHUP SIGTERM EXIT

cd ~/get-yt-install || exit 1

read -p "version: " -r version

if ! wget "https://github.com/yuki-bara/get-yt/archive/refs/tags/v$version.tar.gz" -P dev; then
    exit 1
fi

tar -xzf "dev/v$version.tar.gz" -C .

cd "get-yt-$version" || exit 1

make

echo "move file to /usr/bin"

sudo mv bin/get-yt /usr/bin

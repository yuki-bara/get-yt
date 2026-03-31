#!/usr/bin/bash

version="1.5"

mkdir -p ~/.cache/get-yt-install/dev || exit 1

cleanup() {
    trap '' SIGINT SIGHUP SIGTERM EXIT
    rm -r ~/.cache/get-yt-install/dev
}

trap "cleanup > /dev/null 2>&1" SIGINT SIGHUP SIGTERM EXIT

cd ~/.cache/get-yt-install/dev || exit 1

if ! wget "https://github.com/yuki-bara/get-yt/archive/refs/tags/v$version.tar.gz" -P dev; then
    exit 1
fi

tar -xzf "dev/v$version.tar.gz" -C .

cd "get-yt-$version" || exit 1

make

echo "Moving to /usr/bin..."

if ! sudo mv bin/get-yt /usr/bin/ 2>/dev/null; then
    echo "Error: /usr/bin is read-only. Moving to /usr/local/bin instead..."
    sudo mkdir -p /usr/local/bin
    sudo mv bin/get-yt /usr/local/bin/
    sudo chmod +x /usr/local/bin/get-yt
fi

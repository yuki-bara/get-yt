# SPDX-License-Identifier: 0BSD
# Author: Makkhawan Sardlah

CMD = go
ACT = build
NAM = get-yt
VERSION=1.1

build :
	mkdir -p bin
	$(CMD) $(ACT) -ldflags="-X 'main.version=$(VERSION)'" -o bin/$(NAM) src/$(NAM).go

clean :
	rm -r bin

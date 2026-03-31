# SPDX-License-Identifier: 0BSD
# Author: Makkhawan Sardlah

CMD = go
ACT = build
NAM = get-yt
VERSION=1.5

build :
	mkdir -p bin
	$(CMD) $(ACT) -ldflags="-X 'main.version=$(VERSION)'" -o bin/$(NAM) ./src/unix

build-win :
	mkdir -p bin/win
	GOOS=windows GOARCH=amd64 go build -o bin/win/$(NAM)-64.exe ./src/win
	GOOS=windows GOARCH=arm64 go build -o bin/win/$(NAM)-arm.exe ./src/win
	GOOS=windows GOARCH=386 go build -o bin/win/$(NAM).exe ./src/win

build-root :
	mkdir bin
	$(CMD) $(ACT) -ldflags="-X 'main.version=$(VERSION)'" -o bin\$(NAM).exe .\src\win

clean :
	rm -r bin

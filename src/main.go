// SPDX-License-Identifier: 0BSD
// Author: Makkhawan Sardlah

package main

import (
	"fmt"
	net "get-yt/src/module"
	"os"
)

var version = "version"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Args >> error: No Action")
		return
	}

	if os.Args[1] == "-V" {
		fmt.Printf("\033[33m Version:\033[36m %s \033[0m\n", version)
		return
	}
	if len(os.Args) > 2 {
		if os.Args[1] == "-mp4" {
			net.Download_mp4(os.Args[2:])
		} else if os.Args[1] == "-m4a" {
			net.Download_m4a(os.Args[2:])
		}
	}
}

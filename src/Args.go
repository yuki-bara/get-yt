// SPDX-License-Identifier: 0BSD
// Author: Makkhawan Sardlah

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func Step_1(Args []string) int {
	if len(os.Args) < 2 {
		fmt.Println("Args >> error: No Action")
		return 1
	}
	if os.Args[1] == "-V" {
		fmt.Printf("\033[33m Version:\033[36m %s \033[0m\n", version)
		return -1
	}
	if os.Args[1] == "-C" {
		cacheDir, err := os.UserCacheDir()
		if err == nil {
			os.RemoveAll(filepath.Join(cacheDir, "get-yt"))
		}
	}
	return 0
}

func Step_2(Args []string) {
	if len(Args) > 2 {
		var Action = Args[1]
		switch Action {
		case "-mp4":
			Download_mp4(Args[2:])
		case "-m4a":
			Download_m4a(Args[2:])
		case "-m4v":
			Download_m4v(Args[2:])
		case "-mp4u":
			Download_mp4u(Args[2:])
		}
	}
}

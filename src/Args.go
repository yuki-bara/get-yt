// SPDX-License-Identifier: 0BSD
// Author: Makkhawan Sardlah

package main

import (
	"fmt"
	"os"
)

func Step_1(Args []string) int {
	if len(os.Args) < 2 {
		fmt.Println("Args >> error: No Action")
		return 1
	}
	if os.Args[1] == "-version" {
		fmt.Printf("\033[33m Version:\033[36m %s \033[0m\n", version)
		return -1
	}
	return 0
}

func Step_2(Args []string) {
	if len(Args) > 2 {
		var Action = Args[1]
		switch Action {
		case "mp4":
			Download_mp4(Args[2:])
		case "m4a":
			Download_m4a(Args[2:])
		}
	}
}

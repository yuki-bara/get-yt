// SPDX-License-Identifier: 0BSD
// Author: Makkhawan Sardlah

package main

import (
	"bufio"
	"fmt"
	"os"
)

func examine(mode string, url string) {
	if len(mode) > 2 {
		switch mode {
		case "mp4":
			Download_mp4(url)
		case "m4a":
			Download_m4a(url)
		}
	}
}

func main() {
	var url string
	var mode string
	fmt.Print("url / id: ")
	fmt.Scanln(&url)
	fmt.Print("file (mp4 m4a): ")
	fmt.Scanln(&mode)
	examine(mode, url)
	fmt.Println("\n------------------------------")
	fmt.Println("Done! Press Enter to exit...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}

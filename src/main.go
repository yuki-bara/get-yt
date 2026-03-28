// SPDX-License-Identifier: 0BSD
// Author: Makkhawan Sardlah

package main

import (
	"os"
)

var version = "version"

func main() {
	err := Step_1(os.Args)
	if err != 0 {
		return
	}
	Step_2(os.Args)
}

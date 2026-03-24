// SPDX-License-Identifier: 0BSD
// Author: Makkhawan Sardlah

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

var version = "version"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Args >> error: No URL or No Action")
		return
	}

	if os.Args[1] == "-V" {
		fmt.Println("Version: ", version)
		return
	}

	if len(os.Args[1]) < 11 {
		fmt.Println("Id >> error: No id")
		return
	}
	videoID := os.Args[1][len(os.Args[1])-11:]

	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		fmt.Println("pull >> error: ", err)
		return
	}

	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		fmt.Println("Stream >> error: ", err)
		return
	}
	defer stream.Close()

	file, err := os.Create(video.Title + ".mp4")
	if err != nil {
		fmt.Println("file >> error:", err)
		return
	}
	defer file.Close()

	fmt.Printf("download: %s...\n", video.Title)
	_, err = io.Copy(file, stream)
	if err != nil {
		fmt.Println("download >> error: ", err)
		return
	}

	fmt.Println("succeed!")
}

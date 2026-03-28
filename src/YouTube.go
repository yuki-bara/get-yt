// SPDX-License-Identifier: 0BSD
// Author: Makkhawan Sardlah

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

func Check_id(id string) string {
	if len(id) < 11 {
		fmt.Println("Id >> error: No url or id")
		return "nil"
	}
	return id[len(id)-11:]
}

func Create_file(file_extension string, video *youtube.Video, formats *youtube.Format, client *youtube.Client) {
	stream, _, err := client.GetStream(video, formats)
	if err != nil {
		fmt.Println("Stream >> error: ", err)
		return
	}
	defer stream.Close()
	file, err := os.Create(video.Title + file_extension)
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

func Getvideo(id string) (youtube.Client, *youtube.Video, youtube.FormatList, error) {
	client := youtube.Client{}

	video, err := client.GetVideo(id)
	if err != nil {
		fmt.Println("pull >> error: ", err)
		return client, nil, nil, err
	}

	formats := video.Formats.WithAudioChannels()
	return client, video, formats, nil
}

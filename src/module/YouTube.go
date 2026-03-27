// SPDX-License-Identifier: 0BSD
// Author: Makkhawan Sardlah

package net

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/kkdai/youtube/v2"
)

func Download_mp4(ID []string) {
	var wg sync.WaitGroup
	for i := 0; i < len(ID); i++ {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			if len(id) < 11 {
				fmt.Println("Id >> error: No url or id")
				return
			}
			videoID := id[len(id)-11:]

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
		}(ID[i])
	}
	wg.Wait()
}

func Download_m4a(ID []string) {
	var wg sync.WaitGroup
	for i := 0; i < len(ID); i++ {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			if len(id) < 11 {
				fmt.Println("Id >> error: No url or id")
				return
			}
			videoID := id[len(id)-11:]

			client := youtube.Client{}

			video, err := client.GetVideo(videoID)
			if err != nil {
				fmt.Println("pull >> error: ", err)
				return
			}

			formats := video.Formats.WithAudioChannels()
			var targetFormat *youtube.Format
			for _, f := range formats {
				if f.FPS == 0 {
					targetFormat = &f
					break
				}
			}
			stream, _, err := client.GetStream(video, targetFormat)
			if err != nil {
				fmt.Println("Stream >> error: ", err)
				return
			}
			defer stream.Close()

			file, err := os.Create(video.Title + ".m4a")
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
		}(ID[i])
	}
	wg.Wait()
}

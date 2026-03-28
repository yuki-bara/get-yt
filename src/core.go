package main

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/kkdai/youtube/v2"
)

func Download_mp4(ID []string) {
	var wg sync.WaitGroup
	for i := 0; i < len(ID); i++ {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()

			videoID := Check_id(id)

			client, video, err := Create_client(videoID)

			if err != nil {
				return
			}

			formats := video.Formats.WithAudioChannels()

			Create_file(".mp4", video, &formats[0], &client)

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

			videoID := Check_id(id)

			if videoID == "nil" {
				return
			}

			client, video, err := Create_client(videoID)

			if err != nil {
				return
			}

			formats := video.Formats.Type("audio")

			var targetFormat *youtube.Format
			for i := range formats {
				f := &formats[i]
				if f.FPS == 0 && f.AudioChannels > 0 {
					targetFormat = f
					break
				}
			}
			Create_file(".m4a", video, targetFormat, &client)
		}(ID[i])
	}
	wg.Wait()
}

func Download_m4v(ID []string) {
	var wg sync.WaitGroup
	for i := 0; i < len(ID); i++ {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()

			videoID := Check_id(id)

			if videoID == "nil" {
				return
			}

			client, video, err := Create_client(videoID)

			if err != nil {
				return
			}

			formats := video.Formats.Type("video")
			var targetFormat *youtube.Format

			for i := range formats {
				f := &formats[i]
				if f.FPS > 0 && f.AudioChannels == 0 {
					targetFormat = f
					break
				}
			}

			Create_file(".m4v", video, targetFormat, &client)
		}(ID[i])
	}
	wg.Wait()
}

func Download_mp4u(ID []string) {
	var wg sync.WaitGroup
	for i := 0; i < len(ID); i++ {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()

			videoID := Check_id(id)

			if videoID == "nil" {
				return
			}

			client, video, err := Create_client(videoID)

			if err != nil {
				return
			}

			formats := video.Formats.Type("video")
			var targetFormat *youtube.Format

			for i := range formats {
				f := &formats[i]
				if f.FPS > 0 && f.AudioChannels == 0 {
					targetFormat = f
					break
				}
			}

			Create_file(".m4v", video, targetFormat, &client)

			formats = video.Formats.Type("audio")

			for i := range formats {
				f := &formats[i]
				if f.FPS == 0 && f.AudioChannels > 0 {
					targetFormat = f
					break
				}
			}

			Create_file(".m4a", video, targetFormat, &client)

			cmd := exec.Command("ffmpeg", "-i", video.Title+".m4v", "-i", video.Title+".m4a", "-c", "copy", "-y", video.Title+".mp4")
			err = cmd.Run()
			if err != nil {
				fmt.Println("Error merging:", err)
				return
			}
			fmt.Println("succeed combined to mp4")

		}(ID[i])
	}
	wg.Wait()
}

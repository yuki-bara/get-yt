package main

import (
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

			client, video, formats, err := Getvideo(videoID)

			if err != nil {
				return
			}

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

			client, video, formats, err := Getvideo(videoID)

			if err != nil {
				return
			}

			var targetFormat *youtube.Format
			for _, f := range formats {
				if f.FPS == 0 {
					targetFormat = &f
					break
				}
			}
			Create_file(".m4a", video, targetFormat, &client)
		}(ID[i])
	}
	wg.Wait()
}

// SPDX-License-Identifier: 0BSD
// Author: Makkhawan Sardlah

package main

import (
	"github.com/kkdai/youtube/v2"
)

func Download_mp4(URL string) {
	videoID := Check_id(URL)

	client, video, err := Create_client(videoID)

	if err != nil {
		return
	}

	formats := video.Formats.WithAudioChannels()

	Create_file(".mp4", video, &formats[0], &client)
}

func Download_m4a(URL string) {

	videoID := Check_id(URL)

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
}

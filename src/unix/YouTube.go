// SPDX-License-Identifier: 0BSD
// Author: Makkhawan Sardlah

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/kkdai/youtube/v2"
	"github.com/klauspost/compress/zstd"
)

func Check_id(id string) string {
	if len(id) < 11 {
		fmt.Println("Id >> error: No url or id")
		return "nil"
	}
	return id[len(id)-11:]
}

func Create_file(file_extension string, video *youtube.Video, format *youtube.Format, client *youtube.Client) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println("cache >> error: ", err)
		return
	}
	dircache := filepath.Join(cacheDir, "get-yt")
	cachefile := filepath.Join(dircache, video.ID+file_extension+".zst")
	_, err = os.Stat(cachefile)
	filename := video.Title + file_extension
	if err == nil {
		fmt.Printf("cache >> copyfile: cache/%s >> %s", cachefile, filename)
		cache, _ := os.Open(cachefile)
		defer cache.Close()
		decoder, _ := zstd.NewReader(cache)
		defer decoder.Close()
		file, _ := os.Create(filename)
		defer file.Close()
		_, _ = io.Copy(file, decoder)
		fmt.Println(filename + " >> copy succeed!")
	} else if os.IsNotExist(err) {
		stream, _, err := client.GetStream(video, format)
		if err != nil {
			fmt.Println("Stream >> error: ", err)
			return
		}
		defer stream.Close()
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("file >> error:", err)
			return
		}
		defer file.Close()

		fmt.Printf("download: %s...\n", video.Title+file_extension)
		_, err = io.Copy(file, stream)
		if err != nil {
			fmt.Println("download >> error: ", err)
			return
		}
		err = os.MkdirAll(dircache, 0755)
		if err != nil {
			fmt.Println("dir >> error: ", err)
		}
		cache, err := os.Create(cachefile)
		if err != nil {
			fmt.Println("file >> error: ", err)
			return
		}
		defer cache.Close()
		file.Seek(0, 0)
		encoder, _ := zstd.NewWriter(cache)
		defer encoder.Close()
		io.Copy(encoder, file)
		fmt.Println(filename + " >> download succeed!")
	} else {
		fmt.Print("cache >> error: ", err)
	}
}

func Create_client(id string) (youtube.Client, *youtube.Video, error) {
	client := youtube.Client{}

	video, err := client.GetVideo(id)
	if err != nil {
		fmt.Println("pull >> error: ", err)
		return client, nil, err
	}
	return client, video, nil
}

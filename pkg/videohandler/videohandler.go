package videohandler

import (
	"Zyzz-Motivation-Booster/pkg/storage"
	"github.com/kkdai/youtube/v2"
	"io"
	"math/rand"
	"os"
)

const (
	balkan_gains_all_videos_playlist       = "https://www.youtube.com/watch?v=sOAi9JJhATE&list=UUPXRHu3iYggeluGLgz4QTOQ"
	zyzz_inspirational_all_videos_playlist = "https://www.youtube.com/watch?v=QwIrRS3N1Zw&list=UUfgzNJDxhp3zZ3uNtfmKHTQ"
	//arcdelio_all_videos_playlist           = "https://www.youtube.com/watch?v=AawRADgo0fU&list=UUCCg8Fh5Avkt0lqD4wnAL6A"
	//daily_zyzz_all_videos_playlist         = "https://www.youtube.com/watch?v=WRQrQb4Xu2I&list=UU5mTOjypsz4bHilrl2HKT9g"
)

func GetVideoList() ([]*youtube.PlaylistEntry, error) {
	client := youtube.Client{}
	playlist1, err := client.GetPlaylist(balkan_gains_all_videos_playlist)
	if err != nil {
		return nil, err
	}
	playlist2, err := client.GetPlaylist(zyzz_inspirational_all_videos_playlist)
	if err != nil {
		return nil, err
	}
	/*playlist3, err := client.GetPlaylist(arcdelio_all_videos_playlist)
	if err != nil {
		return nil, err
	}
	playlist4, err := client.GetPlaylist(daily_zyzz_all_videos_playlist)
	if err != nil {
		return nil, err
	}*/
	var videoList []*youtube.PlaylistEntry
	videoList = append(videoList, playlist1.Videos...)
	videoList = append(videoList, playlist2.Videos...)
	/*videoList = append(videoList, playlist3.Videos...)
	videoList = append(videoList, playlist4.Videos...)*/
	return videoList, nil
}

// SelectVideo This function searches for videos with the least viewing count (viewing count = how often that video has been sent to you) and picks a random one and returns the url
func SelectVideo() (string, string, error) {
	videoList, err := GetVideoList()
	if err != nil {
		return "", "", err
	}

	minViewingCount := 999999999
	for _, v := range videoList {
		url := BuildYouTubeUrl(v.ID)
		viewingCount := storage.GetViewingCount(url)
		if viewingCount < minViewingCount {
			minViewingCount = viewingCount
		}
	}

	var videoUrls []string
	var videoNames []string
	// Get Videos with smallest viewing count
	for _, v := range videoList {
		url := BuildYouTubeUrl(v.ID)
		if storage.GetViewingCount(url) == minViewingCount {
			videoUrls = append(videoUrls, url)
			videoNames = append(videoNames, v.Title)
		}
	}

	// Pick a random video
	number := rand.Intn(len(videoUrls))
	video := videoUrls[number]

	// Increase viewing count
	storage.IncreaseViewingCount(video)
	return video, videoNames[number], nil
}

func BuildYouTubeUrl(url string) string {
	return "https://www.youtube.com/watch?v=" + url
}

func DownloadVideo(url string) (string, error) {
	client := youtube.Client{}
	video, err := client.GetVideo(url)
	if err != nil {
		return "", err
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		return "", err
	}

	fileName := "video.mp4"
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

package cron

import (
	"Zyzz-Motivation-Booster/pkg/storage"
	"Zyzz-Motivation-Booster/pkg/telegram"
	"Zyzz-Motivation-Booster/pkg/videohandler"
	"github.com/jasonlvhit/gocron"
	"log"
	"time"
)

func SendVideo() {
	log.Println("Selecting video...")
	video, title, err := videohandler.SelectVideo()
	if err != nil {
		log.Println("Bruh, I couldn't get the videos. Trying again in one minute.")
		time.Sleep(time.Minute)
		SendVideo()
		return
	}
	storage.Save()
	log.Println("Downloading video " + video + "...")
	fileName, err := videohandler.DownloadVideo(video)
	if err != nil {
		log.Println("Bruh, I couldn't download the video with the url " + video + ". Trying again in one minute.")
		time.Sleep(time.Minute)
		SendVideo()
		return
	}
	log.Println("Sending video " + fileName + "...")
	err = telegram.SendVideo(fileName, title)
	if err != nil {
		log.Println("Bruh, I couldn't send the video with the url " + video + ". Trying again in one minute.")
		time.Sleep(time.Minute)
		SendVideo()
		return
	}
}

func Setup() {
	gocron.ChangeLoc(time.UTC)
	gocron.Every(1).Day().At("07:00").Do(SendVideo)
	<-gocron.Start()
}

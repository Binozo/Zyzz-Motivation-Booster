package telegram

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func SendMessage(message string) error {
	formData := url.Values{
		"chat_id": {os.Getenv("telegramchatid")},
		"text":    {message},
	}

	client := &http.Client{}

	//Not working, the post data is not a form
	req, err := http.NewRequest("POST", "https://api.telegram.org/bot"+os.Getenv("telegrambottoken")+"/sendMessage", strings.NewReader(formData.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func SendVideo(filename string, title string) error {
	SendMessage("We're all gonna make it Brah!\n" + title)

	// Create buffer
	buf := new(bytes.Buffer) // caveat IMO dont use this for large files, \
	// create a tmpfile and assemble your multipart from there (not tested)
	w := multipart.NewWriter(buf)
	fw, err := w.CreateFormFile("document", filename)
	if err != nil {
		return err
	}
	fd, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fd.Close()
	// Write file field from file to upload
	_, err = io.Copy(fw, fd)
	if err != nil {
		return err
	}
	// Important if you do not close the multipart writer you will not have a
	// terminating boundry
	w.Close()
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://api.telegram.org/bot"+os.Getenv("telegrambottoken")+"/sendDocument?chat_id="+os.Getenv("telegramchatid"), buf)
	req.Header.Set("Content-Type", w.FormDataContentType())
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//log.Println(string(content))
	return nil
}

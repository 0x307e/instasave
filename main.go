package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/makotia/instasave/models"
)

func main() {
	var (
		url       string
		err       error
		resp      *http.Response
		config    models.Config
		sessionID string
	)

	if _, err = toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sessionID = config.InstagramDefaultSetting.SessionID

	client := new(http.Client)

	for i := 0; i < len(config.InstagramSetting); i++ {
		fmt.Printf("%#v\n", config.InstagramSetting[i])
		if config.InstagramSetting[i].SessionID != "" {
			sessionID = config.InstagramSetting[i].SessionID
		}

		url = fmt.Sprintf("https://i.instagram.com/api/v1/feed/user/%d/reel_media/", config.InstagramSetting[i].UserID)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Cookie", fmt.Sprintf("sessionid=%s", sessionID))
		req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0.1; SM-G935T Build/MMB29M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/51.0.2704.81 Mobile Safari/537.36 Instagram 8.4.0 Android (23/6.0.1; 560dpi; 1440x2560; samsung; SM-G935T; hero2qltetmo; qcom; en_US")

		if resp, err = client.Do(req); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		res, _ := models.UnmarshalStory(byteArray)
		fmt.Printf("%#v\n", res)
	}
}

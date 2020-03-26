package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/makotia/instasave/models"
	"github.com/makotia/instasave/utils"
)

func main() {
	var (
		url       string
		err       error
		req       *http.Request
		byteArray []byte
		resp      *http.Response
		story     models.Story
		feedList  models.FeedList
		config    models.Config
		sessionID string
		dlDir     string
	)

	if _, err = toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sessionID = config.DefaultSetting.SessionID

	if config.DefaultSetting.DownloadDir != "" {
		dlDir = config.DefaultSetting.DownloadDir
	} else {
		dlDir = "./dl"
	}

	client := new(http.Client)

	for i := 0; i < len(config.InstagramSetting); i++ {
		if config.InstagramSetting[i].SessionID != "" {
			sessionID = config.InstagramSetting[i].SessionID
		}

		// Story
		url = fmt.Sprintf("https://i.instagram.com/api/v1/feed/user/%d/reel_media/", config.InstagramSetting[i].UserID)
		req, _ = http.NewRequest("GET", url, nil)
		req.Header.Set("Cookie", fmt.Sprintf("sessionid=%s", sessionID))
		req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0.1; SM-G935T Build/MMB29M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/51.0.2704.81 Mobile Safari/537.36 Instagram 8.4.0 Android (23/6.0.1; 560dpi; 1440x2560; samsung; SM-G935T; hero2qltetmo; qcom; en_US")

		if resp, err = client.Do(req); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		byteArray, _ = ioutil.ReadAll(resp.Body)
		story, _ = models.UnmarshalStory(byteArray)
		if story.Status == "ok" {
			for _, item := range story.Items {
				var (
					height = 0
					width  = 0
					dlurl  = ""
					ext    = ""
					path   = ""
				)

				if len(item.VideoVersions) != 0 {
					ext = "mp4"
					for _, v := range item.VideoVersions {
						if height < v.Height && width < v.Width {
							dlurl = v.URL
							height = v.Height
							width = v.Width
						}
					}
				} else if len(item.ImageVersions.Images) != 0 {
					ext = "jpg"
					for _, image := range item.ImageVersions.Images {
						if height < image.Height && width < image.Width {
							dlurl = image.URL
							height = image.Height
							width = image.Width
						}
					}
				}
				savePath := fmt.Sprintf("%s/%s", dlDir, story.User.UserName)
				if path, err = utils.Download(dlurl, savePath, item.ID, ext); err != nil {
					fmt.Println(err)
				}
				fmt.Println(path)
			}
		}
		// Feed
		url = fmt.Sprintf("https://instagram.com/%s/?__a=1", story.User.UserName)
		req, _ = http.NewRequest("GET", url, nil)
		req.Header.Set("Cookie", fmt.Sprintf("sessionid=%s", sessionID))
		req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0.1; SM-G935T Build/MMB29M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/51.0.2704.81 Mobile Safari/537.36 Instagram 8.4.0 Android (23/6.0.1; 560dpi; 1440x2560; samsung; SM-G935T; hero2qltetmo; qcom; en_US")

		if resp, err = client.Do(req); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		byteArray, _ = ioutil.ReadAll(resp.Body)
		feedList, _ = models.UnmarshalFeedList(byteArray)
		for _, edge := range feedList.GraphQL.User.EdgeOwnerToTimelineMedia.Edges {
			fmt.Println(fmt.Sprintf("https://instagram.com/%s/p/%s/?__a=1", feedList.GraphQL.User.UserName, edge.Node.ShortCode))
		}
	}
}

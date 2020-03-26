package models

import "encoding/json"

// Story reel_media で返ってくるモデル
type Story struct {
	Items  []Item    `json:"items"`
	User   StoryUser `json:"user"`
	Status string    `json:"status"`
}

// Item item
type Item struct {
	ID            string     `json:"id"`
	TimeStamp     int        `json:"device_timestamp"`
	MediaType     int        `json:"media_type"`
	VideoVersions []Video    `json:"video_versions"`
	ImageVersions Candidates `json:"image_versions2"`
}

// StoryUser user
type StoryUser struct {
	UserName string `json:"username"`
	FullName string `json:"full_name"`
}

// Video video
type Video struct {
	URL    string `json:"url"`
	ID     string `json:"id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Image image
type Image struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

// Candidates candidates
type Candidates struct {
	Images []Image `json:"candidates"`
}

// UnmarshalStory Story のレスポンスを struct に
func UnmarshalStory(data []byte) (s Story, err error) {
	err = json.Unmarshal(data, &s)
	return s, err
}

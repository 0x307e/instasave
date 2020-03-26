package models

import "encoding/json"

// Feed 投稿ページに ?__a=1 したときのレスポンス
type Feed struct {
	GraphQL FeedGraphQL `json:"graphql"`
}

// FeedGraphQL graphql
type FeedGraphQL struct {
	ShortCodeMedia ShortCodeMedia `json:"shortcode_media"`
}

// ShortCodeMedia ShortCodeMedia
type ShortCodeMedia struct {
	ID                    string                `json:"id"`
	ShortCode             string                `json:"shortcode"`
	Dimensions            Dimensions            `json:"dimensions"`
	IsVideo               bool                  `json:"is_video"`
	EdgeMediaToCaption    EdgeMediaToCaption    `json:"edge_media_to_caption"`
	TimeStamp             int                   `json:"taken_at_timestamp"`
	Owner                 StoryUser             `json:"owner"`
	EdgeSidecarToChildren EdgeSidecarToChildren `json:"edge_sidecar_to_children"`
}

// Dimensions Dimensions
type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// EdgeMediaToCaption EdgeMediaToCaption
type EdgeMediaToCaption struct {
	Edges []FeedEdge `json:"edges"`
}

// FeedEdge edge
type FeedEdge struct {
	Node FeedNode `json:"node"`
}

// FeedNode node
type FeedNode struct {
	Text string `json:"text"`
}

// EdgeSidecarToChildren EdgeSidecarToChildren
type EdgeSidecarToChildren struct {
	Edges []FeedChildrenEdge `json:"edges"`
}

// FeedChildrenEdge FeedChildrenEdge
type FeedChildrenEdge struct {
	Node FeedChildrenNode `json:"node"`
}

// FeedChildrenNode FeedChildrenNode
type FeedChildrenNode struct {
	ID               string      `json:"id"`
	ShortCode        string      `json:"shortcode"`
	Dimensions       Dimensions  `json:"dimensions"`
	DisplayResources []FeedImage `json:"display_resources"`
	VideoURL         string      `json:"video_url"`
	IsVideo          bool        `json:"is_video"`
}

// FeedImage image
// Image image
type FeedImage struct {
	Width  int    `json:"config_width"`
	Height int    `json:"config_height"`
	URL    string `json:"src"`
}

// UnmarshalFeed Feed のレスポンスを struct に
func UnmarshalFeed(data []byte) (f Feed, err error) {
	err = json.Unmarshal(data, &f)
	return f, err
}

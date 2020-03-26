package models

import "encoding/json"

// FeedList ユーザーページに ?__a=1 したときのレスポンス
type FeedList struct {
	GraphQL GraphQL `json:"graphql"`
}

// GraphQL GraphQL
type GraphQL struct {
	User FeedListUser `json:"user"`
}

// FeedListUser user
type FeedListUser struct {
	UserName                 string                   `json:"username"`
	EdgeOwnerToTimelineMedia EdgeOwnerToTimelineMedia `json:"edge_owner_to_timeline_media"`
}

// EdgeOwnerToTimelineMedia EdgeOwnerToTimelineMedia
type EdgeOwnerToTimelineMedia struct {
	Edges []Edge `json:"edges"`
}

// Edge edge
type Edge struct {
	Node Node `json:"node"`
}

// Node node
type Node struct {
	ShortCode string `json:"shortcode"`
}

// UnmarshalFeedList FeedList のレスポンスを struct に
func UnmarshalFeedList(data []byte) (f FeedList, err error) {
	err = json.Unmarshal(data, &f)
	return f, err
}

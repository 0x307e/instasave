package models

import "encoding/json"

// FeedList ユーザーページに ?__a=1 したときのレスポンス
type FeedList struct {
	GraphQL FeedListGraphQL `json:"graphql"`
}

// FeedListGraphQL GraphQL
type FeedListGraphQL struct {
	User FeedListUser `json:"user"`
}

// FeedListUser user
type FeedListUser struct {
	UserName                 string                   `json:"username"`
	EdgeOwnerToTimelineMedia EdgeOwnerToTimelineMedia `json:"edge_owner_to_timeline_media"`
}

// EdgeOwnerToTimelineMedia EdgeOwnerToTimelineMedia
type EdgeOwnerToTimelineMedia struct {
	Edges []FeedListEdge `json:"edges"`
}

// FeedListEdge edge
type FeedListEdge struct {
	Node FeedListNode `json:"node"`
}

// FeedListNode node
type FeedListNode struct {
	ShortCode string `json:"shortcode"`
	TimeStamp int    `json:"taken_at_timestamp"`
}

// UnmarshalFeedList FeedList のレスポンスを struct に
func UnmarshalFeedList(data []byte) (f FeedList, err error) {
	err = json.Unmarshal(data, &f)
	return f, err
}

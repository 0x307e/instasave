package models

// Config 全ての設定を格納
type Config struct {
	InstagramDefaultSetting InstagramDefaultSetting `toml:"InstagramDefaultSetting"`
	InstagramSetting        []InstagramSetting      `toml:"InstagramSetting"`
}

// InstagramDefaultSetting インスタの SessionID など
type InstagramDefaultSetting struct {
	SessionID string `toml:"SessionID"`
}

// InstagramSetting インスタの SessionID など
type InstagramSetting struct {
	SessionID string `toml:"SessionID"`
	UserID    int    `toml:"UserID"`
}

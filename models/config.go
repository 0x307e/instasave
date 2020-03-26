package models

// Config 全ての設定を格納
type Config struct {
	DefaultSetting   DefaultSetting     `toml:"DefaultSetting"`
	InstagramSetting []InstagramSetting `toml:"InstagramSetting"`
}

// DefaultSetting デフォルト設定
type DefaultSetting struct {
	SessionID   string `toml:"SessionID"`
	DownloadDir string `toml:"DownloadDir"`
}

// InstagramSetting インスタの SessionID など
type InstagramSetting struct {
	SessionID string `toml:"SessionID"`
	UserID    int    `toml:"UserID"`
}

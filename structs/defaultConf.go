package structs

type Config struct {
	DefaultBranch string `json:"defaultBranch"`
	UseEmoji      bool   `json:"useEmoji"`
	UseAi         bool   `json:"useAi"`
	Url           string `json:"url"`
}

func DefaultConfigs() Config {
	return Config{
		DefaultBranch: "main",
		UseEmoji:      true,
		UseAi:         false,
		Url:           "",
	}
}

package structs

type Config struct {
	DefaultBranch string `json:"defaultBranch"`
	UseEmoji      bool   `json:"useEmoji"`
	UseAi         bool   `json:"useAi"`
	Url           string `json:"url"`
	Login         Login  `json:"login"`
}

type Login struct {
	User      string `json:"user"`
	Email     string `json:"email"`
	Abundance string `json:"abundance"`
}

func DefaultConfigs() Config {
	return Config{
		DefaultBranch: "main",
		UseEmoji:      true,
		UseAi:         false,
		Url:           "",
		Login:         Login{User: "User", Email: "", Abundance: "--local"},
	}
}

package utils

import "fmt"

var emojis = map[string]string{
	"feat":     "✨",
	"fix":      "🐛",
	"docs":     "📄",
	"chore":    "🎨",
	"refactor": "♻️",
	"remove":   "🔥",
}

func Parse(prefix string, msg string, wEmoji bool) string {
	if !wEmoji {
		return fmt.Sprintf("%s: %s", prefix, msg)
	}

	if prefix == "" {
		return fmt.Sprint(msg)
	}

	return fmt.Sprintf("%s %s: %s", emojis[prefix], prefix, msg)
}

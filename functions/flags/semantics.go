package functions

func CommitFeat(argM string) string {
	return "✨ feat: " + argM
}

func CommitFix(argM string) string {
	return "🐛 fix: " + argM
}

func CommitChore(argM string) string {
	return "🎨 chore: " + argM
}

func CommitRefactor(argM string) string {
	return "♻️ refactor: " + argM
}

func CommitDocs(argM string) string {
	return "📄 doc: " + argM
}

func CommitRemove(argM string) string {
	return "🔥 remove: " + argM
}

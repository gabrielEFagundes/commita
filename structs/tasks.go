package structs

type Task struct {
	Label string
	Cmd   []string
}

func MountTask(msg string, b string) []Task {
	return []Task{
		{Label: "Initializing your repository...", Cmd: []string{"init"}},
		{Label: "Adding your remote to the origin...", Cmd: []string{"remote"}},
		{Label: "Staging your changes...", Cmd: []string{"add", "."}},
		{Label: "Commiting your changes...", Cmd: []string{"commit", "-m", msg}},
		{Label: "Changing to your branch...", Cmd: []string{"checkout"}},
		{Label: "Pushing your changes...", Cmd: []string{"push", "-u", "origin", b}},
	}
}

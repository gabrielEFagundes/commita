package structs

type Task struct {
	Label string
	Cmd   []string
}

func MountTask(msg string, conf Config) []Task {
	return []Task{
		{Label: "Initializing your repository...", Cmd: []string{"init"}},
		{Label: "Adding your remote to the origin...", Cmd: []string{"remote", "add", "origin", conf.Url}}, // by default, it'll try to add the remote
		{Label: "Staging your changes...", Cmd: []string{"add", "."}},
		{Label: "Commiting your changes...", Cmd: []string{"commit", "-m", msg}},
		{Label: "Changing to your branch...", Cmd: []string{"checkout", "-b", conf.DefaultBranch}}, // by default, it'll try to create the branch
		{Label: "Pushing your changes...", Cmd: []string{"push", "-u", "origin", conf.DefaultBranch}},
	}
}

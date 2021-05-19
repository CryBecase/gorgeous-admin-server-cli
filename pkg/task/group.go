package task

import (
	"github.com/schollz/progressbar/v3"

	"gorgeous-admin-server-cli/pkg/pb"
)

type Group struct {
	theme string
	tasks []Task

	progressBar *progressbar.ProgressBar
}

func NewGroup(theme string, tasks []Task) *Group {
	return &Group{
		theme:       theme,
		tasks:       tasks,
		progressBar: pb.NewProgressBar(len(tasks), theme),
	}
}

func (g *Group) Exec() error {
	for _, t := range g.tasks {
		err := t.Exec()
		if err != nil {
			return err
		}
		_ = g.progressBar.Add(1)
	}
	_ = g.progressBar.Finish()

	return nil
}

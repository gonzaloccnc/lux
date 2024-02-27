package singlechoice

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
)

type SingleChoice struct {
	items   []string
	label   string
	PerPage int
}

func New(choices []string, label string) SingleChoice {
	return SingleChoice{
		items:   choices,
		label:   label,
		PerPage: 5,
	}
}

func (sc *SingleChoice) Run() (string, error) {
	pgm := tea.NewProgram(newModel(*sc))
	mp, err := pgm.Run()

	if err != nil {
		return "", errors.New("error running prompt " + err.Error())
	}

	return mp.(model).choice, mp.(model).err
}

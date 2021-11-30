package model

import (
	"fmt"
	"log"

	"github.com/ahui2016/gof/recipes"
)

type Task struct {
	Recipe  string
	Options map[string]string
	Names   []string
}

type Tasks struct {
	AllTasks []Task `yaml:"all-tasks"`
}

func (all Tasks) ExecAll(useDefault bool) error {
	if len(all.AllTasks) == 0 {
		return fmt.Errorf("no task")
	}
	for _, task := range all.AllTasks {
		recipe, ok := recipes.Get[task.Recipe]
		if !ok {
			return fmt.Errorf("not found recipe: %s", task.Recipe)
		}
		recipe.Refresh()
		recipe.Prepare(task.Names, task.Options, useDefault)
		if err := recipe.Validate(); err != nil {
			return err
		}
		if err := recipe.Exec(); err != nil {
			return err
		}
	}
	log.Print("all tasks are finished.")
	return nil
}

package app

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type ClassFieldConfig struct {
	Type   string                 `yaml:"type"`
	Min    *int                   `yaml:"min,omitempty"`
	Max    *int                   `yaml:"max,omitempty"`
	Values map[string]interface{} `yaml:"values,omitempty"`
}

type ElementsConfig struct {
	Name     string
	Addr     interface{}
	Register interface{}
	Class    interface{}
	task     *Task
}

type TaskConfig struct {
	Addr     *uint8
	Elements map[string]*ElementsConfig
}

type Task struct {
	Name   string
	config *TaskConfig
}

type MainConfig struct {
	Tasks   map[string]*TaskConfig                  `yaml:"tasks"`
	Classes map[string]map[string]*ClassFieldConfig `yaml:"classes"`
}

func (app *App) InitTasks() error {
	config := &MainConfig{}

	rawYaml, err := os.ReadFile(app.Config.TasksConfig)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(rawYaml, config)
	if err != nil {
		return err
	}

	app.Classes = config.Classes

	for k, v := range config.Tasks {
		task := &Task{Name: k, config: v}
		for _, c := range v.Elements {
			c.task = task
		}
		app.Tasks = append(app.Tasks, task)
	}

	return nil
}

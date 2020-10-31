package config

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const (
	goalsYAMLFileName = "goals.yaml"
)

type GoalsConfig struct {
	Goals map[string]*Goal
}

func (config *GoalsConfig) RunGoals(goalsToRun ...string) error {
	for _, goalToRun := range goalsToRun {
		_, ok := config.Goals[goalToRun]
		if !ok {
			return errors.Errorf("Couldn't find goal declared: %s", goalToRun)
		}
	}

	for _, goalToRun := range goalsToRun {
		err := config.Goals[goalToRun].RunFunctions()
		if err != nil {
			return err
		}
	}

	return nil
}

type Goal struct {
	Description string
	Functions   []*Function
}

func (goal *Goal) RunFunctions() error {
	if len(goal.Functions) < 1 {
		fmt.Println("Skipping goal: no functions")
		return nil
	}

	for _, function := range goal.Functions {
		fmt.Printf("Running %s\n", function.Name)
	}

	return nil
}

type Function struct {
	Name   string
	Input  map[string]string
	Output string
}

func From(filename string) (*GoalsConfig, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, errors.Errorf(`Couldn't find a %s file in the directory. Are you in the right directory?`, filename)
	}

	goalsYAML, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// TODO: JSON schema yaml validation

	goalsConfig := &GoalsConfig{}

	err = yaml.Unmarshal(goalsYAML, goalsConfig)
	if err != nil {
		return nil, err
	}

	return goalsConfig, nil
}

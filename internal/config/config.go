package config

import (
	"fmt"
	"github.com/marco2704/goals/pkg/functions"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"plugin"
)

const (
	goalsYAMLFileName = "goals.yaml"
)

type GoalsConfig struct {
	Include []string
	Goals   map[string]*Goal

	includedPlugins []plugin.Plugin
}

func From(filename string) (*GoalsConfig, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, errors.Errorf(`Couldn't find a %s file in the directory. Are you in the right directory?`, filename)
	}

	goalsYAML, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	goalsConfig := &GoalsConfig{}

	// TODO: JSON schema yaml validation
	err = yaml.Unmarshal(goalsYAML, goalsConfig)
	if err != nil {
		return nil, err
	}

	err = goalsConfig.loadIncludedPlugins()
	if err != nil {
		return nil, err
	}

	goalsConfig.loadFunctions()
	if err != nil {
		return nil, err
	}

	return goalsConfig, nil
}

func (config *GoalsConfig) RunGoals(goalsToRun ...string) error {
	for _, goalToRun := range goalsToRun {
		_, ok := config.Goals[goalToRun]
		if !ok {
			return errors.Errorf("Couldn't find goal declared: %s", goalToRun)
		}
	}

	for _, goalToRun := range goalsToRun {
		err := config.runGoalFunctions(goalToRun)
		if err != nil {
			return err
		}
	}

	return nil
}

func (config *GoalsConfig) runGoalFunctions(goalName string) error {
	goal := config.Goals[goalName]

	if len(goal.Functions) == 0 {
		fmt.Println("Skipping goal: no functions declared")
		return nil
	}

	for _, function := range goal.Functions {
		err := function.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func (config *GoalsConfig) loadFunctions() error {
	for _, goal := range config.Goals {
		for _, function := range goal.Functions {
			actualFunction, err := config.findFunction(function.Name)
			if err != nil {
				return err
			}

			function.actualFunction = actualFunction
		}
	}

	return nil
}

func (config *GoalsConfig) findFunction(functionName string) (func(args functions.Args, output *functions.Output) error, error) {
	for _, includedPlugin := range config.includedPlugins {
		value, err := includedPlugin.Lookup(functionName)
		if err == nil {
			function, ok := value.(func(args functions.Args, output *functions.Output) error)
			if !ok {
				return nil, errors.Errorf("Invalid function declaration for: %s", functionName)
			}

			return function, nil
		}
	}

	function, ok := functions.DefaultFunctions()[functionName]
	if !ok {
		return nil, errors.Errorf("Couldn't find function: %s", functionName)
	}

	return function, nil
}

func (config *GoalsConfig) loadIncludedPlugins() error {
	// TODO: build and load functions
	// for _, include := range config.Include {
	// 	if _, err := os.Stat(include); os.IsNotExist(err) {
	// 		return errors.Errorf("Couldn't include %s: path does not exist", include)
	// 	}

	// 	err := buildPlugin(include)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

// func buildPlugin(pluginPath string) error {
// 	cmd := fmt.Sprintf("go build -buildmode=plugin %s", pluginPath)
// 	output, err := funcs.RunCmd(cmd)
// 	if err != nil {
// 		return errors.Errorf("Couldn't build functions form %s: %s: %s", pluginPath, err, output)
// 	}

// 	p, err := plugin.Open("goals.so")
// 	if err != nil {
// 		panic(err)
// 	}

// 	f, err := p.Lookup("ShellCmd")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("ho")

// 	args := funcs.Args{}
// 	args.Add("cmd", "ls")
// 	f.(func(args funcs.Args, output *funcs.Output) error)(args, &funcs.Output{})

// 	return nil
// }

type Goal struct {
	Description string
	Functions   []*Function
}

type Function struct {
	Name   string
	Args   map[string]interface{}
	Output *functions.Output

	actualFunction func(args functions.Args, output *functions.Output) error
}

func (f *Function) Run() error {
	args := functions.Args{}
	for argName, argValue := range f.Args {
		args.Add(argName, argValue)
	}

	return f.actualFunction(args, f.Output)
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Scenario struct {
	Name        string
	Description string
	Action      func() error
}

var scenarios = []Scenario{}

func init() {
	scenarios = append(scenarios, Scenario{
		Name:   "Scenario - 1",
		Action: scenario1,
	})
	scenarios = append(scenarios, Scenario{
		Name:   "Scenario - 2",
		Action: scenario2,
	})
}

func scenario1() error {
	fmt.Printf("\nexecuting scenario - 1:\n")
	scenarioPath := "scenario-1"
	infile := filepath.Join(scenarioPath, "in.worldcfg.json")
	outfile := filepath.Join(scenarioPath, "out.worldcfg.json")

	var worldCfg WorldConfig
	fmt.Printf("  loading file '%s'\n", infile)
	worldCfgData, err := os.ReadFile(infile)
	if err != nil {
		return err
	}

	// try loading the configuration
	fmt.Printf("  unmarshalling to world data\n")
	err = json.Unmarshal(worldCfgData, &worldCfg)
	if err != nil {
		return err
	}

	// try saving the configuration
	fmt.Printf("  re-marshalling to world data\n")
	out, err := json.Marshal(&worldCfg)
	if err != nil {
		return err
	}
	fmt.Printf("  saving data to '%s'\n", outfile)
	os.WriteFile(outfile, out, 0644)
	fmt.Printf("  scenario - 1 finished\n")
	return nil
}

func scenario2() error {
	fmt.Printf("\nexecuting scenario - 2:\n")
	scenarioPath := "scenario-2"
	infile := filepath.Join(scenarioPath, "in.worldcfg.json")
	outfile := filepath.Join(scenarioPath, "out.position.json")

	var worldCfg WorldConfig
	fmt.Printf("  loading file '%s'\n", infile)
	worldCfgData, err := os.ReadFile(infile)
	if err != nil {
		return err
	}

	// try loading the configuration
	fmt.Printf("  unmarshalling to world data\n")
	err = json.Unmarshal(worldCfgData, &worldCfg)
	if err != nil {
		return err
	}

	// try saving the configuration
	fmt.Printf("  building & run world\n")
	scene := worldCfg.RunWorld("scene")

	fmt.Printf("  saving result to '%s'\n", outfile)
	scene.WriteToFile(outfile)
	fmt.Printf("  scenario - 2 finished\n")
	return nil
}

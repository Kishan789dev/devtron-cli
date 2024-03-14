package utils

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

func ReadInputFile[T any](manifest T, pathVar string) (T, error) {
	// Open YAML file
	pathVar = viper.GetString(pathVar)
	yamlFile, err := os.Open(pathVar)
	if err != nil {
		fmt.Println("file not found", err)
		return manifest, err
	}

	defer yamlFile.Close()

	// Decode YAML file to struct
	if yamlFile != nil {
		decoder := yaml.NewDecoder(yamlFile)
		if err := decoder.Decode(&manifest); err != nil {
			fmt.Print("Invalid yaml configuration: ", err)
			return manifest, err
		}
	}
	return manifest, nil
}

func ReadInputFileJson[T any](response T, pathVar string) (T, error) {

	contents, err := os.ReadFile(pathVar)
	if err != nil {
		return response, fmt.Errorf("file not found %s", err)
	}
	if contents != nil {
		err := json.Unmarshal(contents, response)
		if err != nil {
			return response, fmt.Errorf("invalid yaml configuration: %s", err)
		}
	}
	return response, nil
}

func WriteOutputToFileInYaml[T any](outputStruct T) error {

	if viper.GetString("output_path") == "" {
		output, _ := yaml.Marshal(outputStruct)
		fmt.Println(string(output))
		return nil
	}

	// Open file
	yamlFile, err := os.Create(viper.GetString("output_path"))
	if err != nil {
		return fmt.Errorf("file cannot created %s", err)
	}
	defer yamlFile.Close()

	// Encode Struct to YAML file to struct
	if yamlFile != nil {
		encoder := yaml.NewEncoder(yamlFile)
		if err := encoder.Encode(outputStruct); err != nil {
			return fmt.Errorf("struct couldn't be encoded %s", err)
		}
	}
	fmt.Println("Manifest downloaded successfully! ")
	return nil
}

func WriteOutputToFileInJson[T any](outputStruct T) error {
	// Marshal the struct into JSON
	jsonOutput, err := json.MarshalIndent(outputStruct, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshaling to JSON %s", err)
	}

	// Define the output file path
	outputFilePath := viper.GetString("output_path")

	if outputFilePath == "" {
		fmt.Println(string(jsonOutput))
		return nil
	}

	// Write the JSON data to the output file
	err = os.WriteFile(outputFilePath, jsonOutput, 0644)
	if err != nil {
		return fmt.Errorf("error writing JSON to file %s", err)
	}
	fmt.Println("Manifest downloaded successfully! ")
	return nil
}

func SplitAndTrim(input string) []string {

	if !strings.Contains(input, ",") {
		return strings.Fields(input)
	}

	splitStrings := strings.Split(input, ",")
	finalSplitStrings := make([]string, 0)

	for _, str := range splitStrings {
		value := strings.TrimSpace(str)
		if value != "" {
			finalSplitStrings = append(finalSplitStrings, value)
		}
	}
	return finalSplitStrings
}

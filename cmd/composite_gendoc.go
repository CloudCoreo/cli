// Copyright © 2016 Paul Allen <paul@cloudcoreo.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"
	"path"
	"io/ioutil"
	"bytes"
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"strconv"

	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var docHeaders =  map[string]string {
		"head.md": "",
		"description.md":   "Description",
		"config.yaml": "",
		"tags.md": "Tags",
		"categories.md": "Categories",
		"diagram.md": "Diagram",
		"icon.md": "Icon",
		"hierarchy.md" : "Hierarchy",
	}

var docOrder = []string {"head.md", "description.md", "hierarchy.md", "config.yaml", "tags.md", "categories.md", "diagram.md", "icon.md"}

var cmdCompositeGendoc = &cobra.Command{
	Use: content.CMD_COMPOSITE_GENDOC_USE,
	Short: content.CMD_COMPOSITE_GENDOC_SHORT,
	Long: content.CMD_COMPOSITE_GENDOC_LONG,
	Run: func(cmd *cobra.Command, args []string) {

		if directory == "" {
			directory, _ = os.Getwd()
		}

		var readmeFileContent bytes.Buffer

		for index := range(docOrder) {

			fileName := docOrder[index]


			if fileName == "config.yaml" {
				configFileContent, _ := generateConfigContent(path.Join(directory, fileName))
				readmeFileContent.WriteString(configFileContent)
			} else {

				fileContent, err := ioutil.ReadFile(path.Join(directory, fileName))
				if err != nil {
					fmt.Println(fmt.Sprintf(content.ERROR_MISSING_FILE, fileName))
					err := util.CreateFile(fileName, directory, "", false)
					if err != nil {
						fmt.Fprintf(os.Stderr, err.Error())
						os.Exit(-1)
					}

				}

				// create headers when non empty
				if docHeaders[fileName] != "" {
					readmeFileContent.WriteString(fmt.Sprintf("## %s\n", docHeaders[docOrder[index]]))
				}

				readmeFileContent.WriteString(string(fileContent)+"\n\n")
			}
		}

		err := util.CreateFile(content.DEFAULT_FILES_README_FILE, directory, readmeFileContent.String(), true)

		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)

		}

		fmt.Println(content.CMD_COMPOSITE_GENDOC_SUCCESS)
	},
}

// YamlConfig struct for parsing config.yaml file
type YamlConfig struct {
	Variables yaml.MapSlice
}

type varOption struct {
	key string
	description string
	required bool
	valueType string `schema:"type"`
	defaultValues interface{} `schema:"default"`
}

func generateConfigContent(configFilePath string) (string, error) {

	filename, err := filepath.Abs(configFilePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}

	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		fmt.Fprintln(os.Stderr, "Could not read "+filename)
	}

	var config YamlConfig

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not parse YAML")
		os.Exit(-1)
	}


	missingRequiredContent, err := generateVariablesContent(
		config,
		func(required bool, defaultValue interface{}) bool{
			return required && defaultValue == nil
		},
		fmt.Sprintf("\n## %s\n\n", content.DEFAULT_FILES_GENDOC_README_REQUIRED_NO_DEFAULT_HEADER),
		false)

	requiredContent, err := generateVariablesContent(
		config,
		func(required bool, defaultValue interface{}) bool{
			return required && defaultValue != nil
		},
		fmt.Sprintf("\n## %s\n\n", content.DEFAULT_FILES_GENDOC_README_REQUIRED_DEFAULT_HEADER),
		true)

	notRequiredDefaultContent, err := generateVariablesContent(
		config,
		func(required bool, defaultValue interface{}) bool{
			return !required && defaultValue != nil
		},
		fmt.Sprintf("\n## %s\n\n", content.DEFAULT_FILES_GENDOC_README_NO_REQUIRED_DEFAULT_HEADER),
		true)


	theRestContent, err := generateVariablesContent(
		config,
		func(required bool, defaultValue interface{}) bool{
			return !required && defaultValue == nil
		},
		fmt.Sprintf("\n## %s\n\n", content.DEFAULT_FILES_GENDOC_README_NO_REQUIRED_NO_DEFAULT_HEADER),
		true)

	return missingRequiredContent + requiredContent + notRequiredDefaultContent+theRestContent, err

}


func generateVariablesContent(config YamlConfig, check func(bool, interface{}) bool, header string, printVar bool ) (string, error) {

	var contentBytes bytes.Buffer
	counter := 0
	contentBytes.WriteString(header)

	// loop over mapslice and create option object. TODO: Better way would be to use reflection instead.
	for _, variable := range config.Variables {
		option := varOption{}
		option.key = variable.Key.(string)
		for _, o := range variable.Value.(yaml.MapSlice) {
			switch o.Key {
			case "description":
				option.description = o.Value.(string)
			case "required":
				option.required = o.Value.(bool)
			case "type":
				option.valueType = o.Value.(string)
			case "default":
				option.defaultValues = o.Value
			}
		}

		// check if
		if check(option.required, option.defaultValues) {
			counter ++
			contentBytes.WriteString("### `"+ option.key + "`:\n")
			contentBytes.WriteString("  * description: "+ option.description)

			if printVar && option.defaultValues != nil {

				switch option.valueType  {
				case "boolean":
					contentBytes.WriteString("\n  * default: ")
					contentBytes.WriteString(strconv.FormatBool(option.defaultValues.(bool)))
				case "string":
					contentBytes.WriteString("\n  * default: ")
					contentBytes.WriteString(option.defaultValues.(string))
				case "array":
					// Convert to string[] and then join items with ,
					defaultValues :=  convertStringSlice(option.defaultValues)

					contentBytes.WriteString("\n  * default: ")
					contentBytes.WriteString(fmt.Sprint(strings.Join(defaultValues, ", ")))
				case "number":
					contentBytes.WriteString("\n  * default: ")
					contentBytes.WriteString(fmt.Sprint(option.defaultValues))
				default:
					// if no type is provided
					c := option.defaultValues.(string)
					if strings.Contains(c, "\n") {
						contentBytes.WriteString("\n  * default: \n" + content.DEFAULT_FILES_README_CODE_TICKS + "\n")
						contentBytes.WriteString(option.defaultValues.(string) + "\n")
						contentBytes.WriteString(content.DEFAULT_FILES_README_CODE_TICKS)
					} else {
						contentBytes.WriteString("\n  * default: ")
						contentBytes.WriteString(option.defaultValues.(string)+"\n")
					}
				}
			}

			contentBytes.WriteString("\n\n")
		}
	}

	if counter == 0 {
		contentBytes.WriteString("**None**\n\n")
	}

	return contentBytes.String(), nil
}

func convertStringSlice(slice interface{}) []string {

	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
	panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]string, s.Len())

	for i:=0; i<s.Len(); i++ {
	ret[i] = s.Index(i).Interface().(string)
	}

	return ret
}

func init() {
	CompositeCmd.AddCommand(cmdCompositeGendoc)

	cmdCompositeGendoc.Flags().StringVarP(&directory, content.CMD_FLAG_DIRECTORY_LONG, content.CMD_FLAG_DIRECTORY_SHORT, "",content.CMD_FLAG_DIRECTORY_DESCRIPTION )
}


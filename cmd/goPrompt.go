/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

// Workaround for gracefully handling exit
// https://github.com/c-bata/go-prompt/issues/59#issuecomment-376002177
type Exit int

var ctrlC = prompt.KeyBind{
	Key: prompt.ControlC,
	Fn:  exit,
}

func exit(_ *prompt.Buffer) {
	panic(Exit(0))
}

func handleExit() {
	fmt.Println("[TRACE] handleExit()")
	switch v := recover().(type) {
	case nil:
		return
	case Exit:
		os.Exit(int(v))
	default:
		fmt.Println(v)
		fmt.Println(string(debug.Stack()))
	}
}

func makeCompleter(words []string) func(prompt.Document) []prompt.Suggest {
	fmt.Println(words)
	return func(d prompt.Document) []prompt.Suggest {
		suggestions := make([]prompt.Suggest, len(words))
		for i := 0; i < len(words); i++ {
			suggestions[i] = prompt.Suggest{Text: words[i], Description: "Something you typed previously"}
		}
		fmt.Println(suggestions)
		return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
	}
}

// goPromptCmd represents the goPrompt command
var goPromptCmd = &cobra.Command{
	Use:   "goPrompt",
	Short: "go-prompt demo",
	Long:  `A demonstration of go-prompt`,
	Run: func(cmd *cobra.Command, args []string) {
		defer handleExit()

		var history []string
		for {
			fmt.Println("Please say somethinge.")
			t := prompt.Input("> ", makeCompleter(history))
			fmt.Println("You said: " + t)
			history = append(history, t)
		}
	},
}

func init() {

	rootCmd.AddCommand(goPromptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goPromptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goPromptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

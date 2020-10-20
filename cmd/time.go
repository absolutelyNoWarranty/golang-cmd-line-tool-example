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
	"time"

	"github.com/spf13/cobra"
)

func timeCmdAction(cmd *cobra.Command, args []string) {
	fmt.Println("[TRACE] timeCmdAction")
	fmt.Println("[TRACE] functions passed to cobra.Command's Run field take a *cobra.Command pointer and a slice of strings []string")

	fmt.Println(time.Now())
}

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "A brief description of the 'time' command (cmd/time.go). Defined in the `Short` field of timeCmd var",
	Long:  "Prints the current time then exits. This long description of 'time' (cmd/time.go) is defined in the `Long` field of timeCmd var",
	Run:   timeCmdAction,
}

func init() {
	rootCmd.AddCommand(timeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

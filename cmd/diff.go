// Copyright © 2018 TonyKingdom
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
	"github.com/spf13/cobra"
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "diffs two versions of mongodata data.",
	Long: `diffs two versions of data by 
                     mongotool diff --file1 datafile1 --file2 datafile2.`,
	Run: func(cmd *cobra.Command, args []string) {
		Diff(file1, file2)
	},
}

var file1 string
var file2 string

func init() {
	rootCmd.AddCommand(diffCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diffCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diffCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	diffCmd.Flags().StringVar(&file1, "file1", "", "base database file to compare")
	diffCmd.Flags().StringVar(&file2, "file2", "", "the database file to compared")
	diffCmd.MarkFlagRequired("file1")
	diffCmd.MarkFlagRequired("file2")
}

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
	"github.com/spf13/viper"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Automatically check database changes based on lasest saved version.",
	Long: `Automatically check database changes based on lasest saved version.
                    `,
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		filecount = viper.GetInt("filecount")
	},
}

var filecount int

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//runCmd.PersistentFlags().StringP("host", "H", "", "host of mongos")
	//runCmd.MarkPersistentFlagRequired("host")
	runCmd.Flags().IntP("filecount", "c", 5, "Database snapshot file and diff report file count.")
	viper.BindPFlag("filecount", runCmd.Flags().Lookup("filecount"))

	runCmd.Flags().SortFlags = false
	runCmd.PersistentFlags().SortFlags = false

}

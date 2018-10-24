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
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	host string
	port int
	user string
	pass string
	path string
}

var cfgFile string
var Conf Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mongotool",
	Short: "mongotool is tool for mananing mongodb database changes.",
	Long: `mongotool provides checking database changing and persist the latest database info automatically.
You can also diffs two version of data via diff command.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		Conf = Config{
			host: viper.GetString("host"),
			port: viper.GetInt("port"),
			user: viper.GetString("adminuser"),
			pass: viper.GetString("adminpass"),
			path: viper.GetString("dir"),
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mongotool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringP("host", "H", "localhost", "host of mongos")
	rootCmd.PersistentFlags().IntP("port", "P", 27017, "port of mongos")
	rootCmd.PersistentFlags().StringP("adminuser", "u", "admin", "admin user of mongos")
	rootCmd.PersistentFlags().StringP("adminpass", "p", "admin", "password of user")
	rootCmd.PersistentFlags().StringP("dir", "d", "/home/mongodb/diff", "path of mongotool files")

	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("adminuser", rootCmd.PersistentFlags().Lookup("adminuser"))
	viper.BindPFlag("adminpass", rootCmd.PersistentFlags().Lookup("adminpass"))
	viper.BindPFlag("dir", rootCmd.PersistentFlags().Lookup("dir"))

	rootCmd.PersistentFlags().SortFlags = false
	rootCmd.Flags().SortFlags = false

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".mongotool" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".mongotool")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

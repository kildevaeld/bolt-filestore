// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"path/filepath"

	"github.com/kildevaeld/go-filestore"
	"github.com/spf13/cobra"
)

func printTree(fs filestore.FS, path, indent string) {
	fs.List(path, func(node *filestore.Node) error {

		if node.Dir {
			fmt.Printf("%s├── %s\n", indent, node.Path)
			printTree(fs, node.Path, indent+"   ")
		} else {
			path := filepath.Base(node.Path) //strings.Replace(in.GetResource().Path, stripPath, "", 1)
			i := indent
			if len(i) > 0 {
				i = "|" + indent
			}
			fmt.Printf("%s├── %s\n", i, path)
		}

		return nil
	})
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fs, err := filestore.New(dbPath)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}

		printTree(fs, "/", "")

	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

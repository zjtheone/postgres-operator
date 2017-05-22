/*
 Copyright 2017 Crunchy Data Solutions, Inc.
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
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a database, cluster, backup, or upgrade",
	Long: `delete allows you to delete a database, cluster, backup, or upgrade
For example:

pgo delete database mydatabase
pgo delete cluster mycluster
pgo delete backup mycluster
pgo delete upgrade mycluster`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println(`You must specify the type of resource to delete.  Valid resource types include:
	* database
	* cluster
	* backup
	* upgrade`)
		} else {
			switch args[0] {
			case "database":
			case "cluster":
			case "backup":
			case "upgrade":
				break
			default:
				fmt.Println(`You must specify the type of resource to delete.  Valid resource types include: 
	* database
	* cluster
	* backup
	* upgrade`)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteClusterCmd)
	deleteCmd.AddCommand(deleteBackupCmd)
	deleteCmd.AddCommand(deleteUpgradeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

var deleteUpgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "delete an upgrade",
	Long: `delete an upgrade. For example:
	pgo delete upgrade mydatabase`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Error("a database or cluster name is required for this command")
		} else {
			deleteUpgrade(args)
		}
	},
}

var deleteBackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "delete a backup",
	Long: `delete a backup. For example:
	pgo delete backup mydatabase`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Error("a database or cluster name is required for this command")
		} else {
			deleteBackup(args)
		}
	},
}

var deleteClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "delete a cluster",
	Long: `delete a crunchy cluster. For example:
	pgo delete cluster mycluster`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Error("a cluster name is required for this command")
		} else {
			deleteCluster(args)
		}
	},
}

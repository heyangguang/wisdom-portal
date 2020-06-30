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
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"wisdom-portal/models/migrate"
)

var migrateDB string

// migrateCmd represents the run command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate wisdomServer Table to MySQL Databases",
	Run: func(cmd *cobra.Command, args []string) {
		if migrateDB == "" {
			_ = cmd.Help()
			return
		} else {
			migrateTable(migrateDB)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().StringVar(&migrateDB, "db", "", "--db <logLocation> example --db root:123456@(127.0.0.1:3306)/test")
}

func migrateTable(dbStr string) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local", dbStr))
	if err != nil {
		fmt.Printf("connect db failed, err: %v \n", err)
		return
	}
	defer db.Close()
	db.SingularTable(true)

	//初始化数据库
	db.Debug().AutoMigrate(&migrate.UserMigrate{}, &migrate.UserGroupMigrate{}, &migrate.RuleMigrate{}, &migrate.ObjActMigrate{},
		&migrate.CasbinRuleMigrate{}, &migrate.MonitorElasticSearchMigrate{}, &migrate.MonitorMySQLMigrate{}, &migrate.AlertMigrate{},
		&migrate.MonitorIntermediateMigrate{}, &migrate.MonitorKafkaMigrate{}, &migrate.MonitorKubernetesMigrate{},
		&migrate.MonitorAccessLogMigrate{})

	fmt.Println("Migrate Success...")
	return
}

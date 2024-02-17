/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"
	"users/ent"
	"users/ent/migrate"

	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run ent migration to create database schema and tables",
	Long:  `Run ent migration to create database schema and tables.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ent.Open("postgres", "host=localhost port=5432 user=fred.zhang dbname=users_development sslmode=disable")
		if err != nil {
			log.Fatalf("failed opening connection to postgres: %v", err)
		}
		defer client.Close()
		// Run the auto migration tool.
		if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

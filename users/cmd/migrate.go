/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"

	"github.com/spacelions/phew/users/ent"
	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run the auto migration tool.",
	Long:  `Run the auto migration tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ent.Open("postgres", "host=localhost port=5432 user=fred.zhang dbname=users_development sslmode=disable")
		if err != nil {
			log.Fatalf("failed opening connection to postgres: %v", err)
		}
		defer client.Close()
		// Run the auto migration tool.
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

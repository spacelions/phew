/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"net/http"
	"users/ent"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Spin up the server to serve GraphQL endpoints.",
	Long:  `Spin up the server to serve GraphQL endpoints.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ent.Open("postgres", "host=localhost port=5432 user=fred.zhang dbname=users_development sslmode=disable")
		if err != nil {
			log.Fatalf("failed opening connection to postgres: %v", err)
		}
		defer client.Close()

		srv := handler.NewDefaultServer(NewSchema(client))
		http.Handle("/",
			playground.Handler("Todo", "/query"),
		)
		http.Handle("/query", srv)
		log.Println("listening on :8081")
		if err := http.ListenAndServe(":8081", nil); err != nil {
			log.Fatal("http server terminated", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

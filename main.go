package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/healthcare/v1"
)

func main() {
	ctx := context.Background()
	hcSvc, err := healthcare.NewService(ctx)
	if err != nil {
		log.Fatalf("healthcare.NewService: %v", err)
	}
	messageSvc := hcSvc.Projects.Locations.Datasets.Hl7V2Stores.Messages

	projectID := os.Getenv("PROJECT_ID")
	if projectID == "" {
		log.Fatalf("please provide value for PROJECT_ID variable")
	}
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("managedwriter.NewClient: %v", err)
	}

	srv := &http.Server{
		Addr:              net.JoinHostPort("0.0.0.0", "8080"),
		Handler:           NewHl7Service(client, messageSvc).Mux(),
		ReadHeaderTimeout: 3 * time.Second,
	}
	if err = srv.ListenAndServe(); err != nil {
		log.Fatalf("srv.ListenAndServe: %v", err)
	}
}

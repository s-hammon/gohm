package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/s-hammon/p"
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
	defer client.Close()

	srv := &http.Server{
		Addr:              net.JoinHostPort("0.0.0.0", "8080"),
		Handler:           NewHl7Service(client, messageSvc).Mux(),
		ReadHeaderTimeout: 3 * time.Second,
	}
	defer srv.Close()

	if err = setup(ctx, client); err != nil {
		log.Fatal(err)
	}
	if err = srv.ListenAndServe(); err != nil {
		log.Fatalf("srv.ListenAndServe: %v", err)
	}
}

func setup(ctx context.Context, client *bigquery.Client) error {
	structs := map[string]any{
		"adt": ADT{},
		"orm": ORM{},
		"oru": ORU{},
		"mdm": MDM{},
	}
	for name, s := range structs {
		tableRef := client.Dataset("methodist").Table(p.Format("%s_raw", name))
		if _, err := tableRef.Metadata(ctx); err != nil {
			schema, err := bigquery.InferSchema(s)
			if err != nil {
				return fmt.Errorf("bigquery.InferSchema(%s): %v", name, err)
			}
			metadata := &bigquery.TableMetadata{
				Schema: schema,
			}
			if err = tableRef.Create(ctx, metadata); err != nil {
				return fmt.Errorf("bigquery.Table.Create(%s): %v", name, err)
			}
		}
	}
	return nil
}

package gcp

import (
	"context"
	"log"

	"google.golang.org/api/iam/v1"
)

// IAMService is GCP's iam. https://cloud.google.com/iam/
var IAMService *iam.Service

func init() {
	if OnGCP() {
		ctx := context.Background()

		s, err := iam.NewService(ctx)
		if err != nil {
			log.Fatalf("Error connecting to iam service. err = %v", err)
		}
		IAMService = s
	}
}

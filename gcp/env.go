package gcp

import (
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/compute/metadata"
)

var (
	// ProjectID is the GCP's projectID.
	ProjectID string

	// DefaultServiceAccountID is the default service account id.
	DefaultServiceAccountID string

	// DefaultServiceAccountName is the default service account name.
	DefaultServiceAccountName string
)

func init() {
	ProjectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
	if ProjectID == "" {
		log.Fatal("failed init. Please set GOOGLE_CLOUD_PROJECT.")
	}

	if OnGCP() {
		saName, err := metadata.Email("default")
		if err != nil {
			log.Fatalf("failed to get the default service account name. err = %v", err)
		}
		DefaultServiceAccountName = saName

		DefaultServiceAccountID = fmt.Sprintf(
			"projects/%s/serviceAccounts/%s",
			ProjectID,
			saName,
		)
	}
}

// OnGCP returns `true` if on GCP.
func OnGCP() bool {
	return OnGAE() || OnCloudRun()
}

// GetGAEVar gets the GCP project ID and service name from the environment variables
// https://cloud.google.com/appengine/docs/standard/go/runtime?hl=en#environment_variables
func GetGAEVar() (string, string) {
	return os.Getenv("GOOGLE_CLOUD_PROJECT"), os.Getenv("GAE_SERVICE")
}

// OnGAE returns `true` if on GAE.
func OnGAE() bool {
	p, s := GetGAEVar()
	return p != "" && s != ""
}

// GetCloudRunVar gets the service name from the environment variables
// https://cloud.google.com/run/docs/reference/container-contract#env-vars
func GetCloudRunVar() string {
	return os.Getenv("K_SERVICE")
}

// OnCloudRun returns `true` if on Cloud Run.
func OnCloudRun() bool {
	s := GetCloudRunVar()
	return s != ""
}

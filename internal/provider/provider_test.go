package provider

import (
	"testing"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){
	"clearml": func() (*schema.Provider, error) {
		return New("dev")(), nil
	},
}

func TestProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("TF_ACC"); v != "1" {
		t.Skip("TF_ACC=1 must be set to run acceptance tests")
	}

	if v := os.Getenv("CLEARML_ACCESS_KEY"); v == "" {
		t.Fatal("CLEARML_ACCESS_KEY must be set for acceptance tests")
	}

	if v := os.Getenv("CLEARML_SECRET_KEY"); v == "" {
		t.Fatal("CLEARML_SECRET_KEY must be set for acceptance tests")
	}
}

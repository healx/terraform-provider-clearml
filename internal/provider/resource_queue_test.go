package provider

import (
	"regexp"
	"testing"

	"github.com/healx/terraform-provider-clearml/internal/template"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceQueue_basic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: template.ParseRandName(testAccResourceQueue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"clearml_queue.foo", "name", regexp.MustCompile("^terraform-test")),
				),
			},
		},
	})
}

const testAccResourceQueue = `
resource "clearml_queue" "foo" {
  name = "terraform-test-{{.randName}}"
}
`

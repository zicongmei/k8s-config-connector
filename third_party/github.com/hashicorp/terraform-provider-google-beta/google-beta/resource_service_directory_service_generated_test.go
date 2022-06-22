// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccServiceDirectoryService_serviceDirectoryServiceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckServiceDirectoryServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceDirectoryService_serviceDirectoryServiceBasicExample(context),
			},
			{
				ResourceName:            "google_service_directory_service.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"namespace", "service_id"},
			},
		},
	})
}

func testAccServiceDirectoryService_serviceDirectoryServiceBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_service_directory_namespace" "example" {
  provider     = google-beta
  namespace_id = "tf-test-example-namespace%{random_suffix}"
  location     = "us-central1"
}

resource "google_service_directory_service" "example" {
  provider   = google-beta
  service_id = "tf-test-example-service%{random_suffix}"
  namespace  = google_service_directory_namespace.example.id

  metadata = {
    stage  = "prod"
    region = "us-central1"
  }
}
`, context)
}

func testAccCheckServiceDirectoryServiceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_service_directory_service" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ServiceDirectoryBasePath}}{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("ServiceDirectoryService still exists at %s", url)
			}
		}

		return nil
	}
}
package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccContainerAnalysisNote_basic(t *testing.T) {
	t.Parallel()

	name := acctest.RandString(10)
	readableName := acctest.RandString(10)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContainerAnalysisNoteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContainerAnalysisNoteBasic(name, readableName),
			},
			{
				ResourceName:      "google_container_analysis_note.note",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccContainerAnalysisNote_update(t *testing.T) {
	t.Parallel()

	name := acctest.RandString(10)
	readableName := acctest.RandString(10)
	readableName2 := acctest.RandString(10)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContainerAnalysisNoteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContainerAnalysisNoteBasic(name, readableName),
			},
			{
				ResourceName:      "google_container_analysis_note.note",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccContainerAnalysisNoteBasic(name, readableName2),
			},
			{
				ResourceName:      "google_container_analysis_note.note",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccContainerAnalysisNoteBasic(name, readableName string) string {
	return fmt.Sprintf(`
resource "google_container_analysis_note" "note" {
  name = "tf-test-%s"
  attestation_authority {
    hint {
      human_readable_name = "My Attestor %s"
    }
  }
}
`, name, readableName)
}

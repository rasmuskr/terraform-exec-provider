package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"exec": testAccProvider,
	}
}

func TestResourceExec(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccResourceExecDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceExecConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("exec.foo", "command", "echo 'success'"),
					resource.TestCheckResourceAttr("exec.foo", "output", "success\n"),
				),
			},
			resource.TestStep{
				Config: testAccResourceExecConfig_fail,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("exec.foo", "output", "success\n"),
				),
			},
		},
	})
}

const testAccResourceExecConfig_basic = `
resource "exec" "foo" {
	command = "echo 'success'"
	only_if = "test-fixtures/test-command pass"
}
`

const testAccResourceExecConfig_fail = `
resource "exec" "foo" {
	command = "echo 'success'"
	only_if = "test-fixtures/test-command fail"
}
`

func testAccResourceExecDestroy(s *terraform.State) error {
	return nil
}

package main

import (
	"fmt"
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

func TestResourceExecCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceExecConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("exec_exec.foo", "id", "5ffe533b830f08a0326348a9160afafc8ada44db"),
				),
			},
		},
	})
}

func TestResourceExecUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckExecResourceIsNil("exec_exec.foo"),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceExecConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("exec_exec.foo", "id", "5ffe533b830f08a0326348a9160afafc8ada44db"),
				),
			},
			resource.TestStep{
				Config: testAccResourceExecConfig_basic_2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("exec_exec.foo", "id", "3b332e3e7ba175040ba3c0999d0fd115d2539edd"),
				),
			},
		},
	})
}

func testAccCheckExecResourceIsNil(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[r]
		if ok {
			return fmt.Errorf("Resource exists: %s", r)
		}
		return nil
	}
}

const testAccResourceExecConfig_basic = `
resource "exec_exec" "foo" {
  command = "true"
  destroy_command = "false"
}
`

const testAccResourceExecConfig_basic_2 = `
resource "exec_exec" "foo" {
  command = "echo 'success2'"
  destroy_command = "false"
}
`

const testAccResourceExecConfig_timeout = `
resource "exec_exec" "foo" {
	command = "sleep 2 && echo 'success'"
	timeout = 1
}
`
const testAccResourceExecConfig_fail = `
resource "exec_exec" "foo" {
	command = "echo 'failure' >&2 && exit 1"
}
`

package networking_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	ntest "github.com/netapp/terraform-provider-netapp-ontap/internal/provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// example ID: aeef4e4f-a663-11ef-9ca8-00a0b8bc0407
const idRegex string = "[[:xdigit:]]{8}-[[:xdigit:]]{4}-[[:xdigit:]]{4}-[[:xdigit:]]{4}-[[:xdigit:]]{12}"

func TestAccNetworkIpInterfaceResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { ntest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: ntest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// non-existant SVM return code 2621462. Must happen before create/read
			{
				Config:      testAccNetworkIPInterfaceResourceConfigHomePortNode("non-existant", "10.10.10.10", "e0d", "ontap_cluster_1-01"),
				ExpectError: regexp.MustCompile("Code:\"2621462\""),
			},
			// non-existant home node
			{
				Config:      testAccNetworkIPInterfaceResourceConfigHomePortNode("svm0", "10.10.10.10", "e0d", "non-existant_home_node"),
				ExpectError: regexp.MustCompile("Code:\"53281680\""),
			},
			// non-existant broadcast domain
			{
				Config:      testAccNetworkIPInterfaceResourceConfigBroadcastDomain("svm0", "10.10.10.10", "non-existant_broadcast_domain"),
				ExpectError: regexp.MustCompile("Code:\"2\""),
			},
			// empty location, no Home Node / Home Port / Broadcast Domain
			// Error 1967111: "Home node must be specified by at least one location.home_node, location.home_port, or location.broadcast_domain field."
			{
				Config:      testAccNetworkIPInterfaceResourceConfigHomePortNode("svm0", "10.10.10.10", "", ""),
				ExpectError: regexp.MustCompile("Code:\"1967111\""),
			},
			// Create and Read
			// {
			// 	Config: testAccNetworkIPInterfaceResourceConfig("svm0", "10.10.10.10", "ontap_cluster_1-01", "default-data-files"),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		resource.TestCheckResourceAttr("netapp-ontap_network_ip_interface.example", "name", "test-interface"),
			// 		resource.TestCheckResourceAttr("netapp-ontap_network_ip_interface.example", "svm_name", "svm0"),
			// 		resource.TestCheckResourceAttr("netapp-ontap_network_ip_interface.example", "service_policy", "default-data-files"),
			// 	),
			// },
			// // Update and Read
			// {
			// 	Config: testAccNetworkIPInterfaceResourceConfig("svm0", "10.10.10.20", "ontap_cluster_1-01", "default-data-iscsi"),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		resource.TestCheckResourceAttr("netapp-ontap_network_ip_interface.example", "name", "test-interface"),
			// 		resource.TestCheckResourceAttr("netapp-ontap_network_ip_interface.example", "ip.address", "10.10.10.20"),
			// 		resource.TestCheckResourceAttr("netapp-ontap_network_ip_interface.example", "service_policy", "default-data-iscsi"),
			// 	),
			// },
			// // Test importing a resource
			// {
			// 	ResourceName:  "netapp-ontap_network_ip_interface.example",
			// 	ImportState:   true,
			// 	ImportStateId: fmt.Sprintf("%s,%s,%s", "test-interface", "svm0", "cluster4"),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		resource.TestCheckResourceAttr("netapp-ontap_network_ip_interface.example", "name", "test-interface"),
			// 		resource.TestCheckResourceAttr("netapp-ontap_network_ip_interface.example", "ip.address", "10.10.10.20"),
			// 	),
			// },
		},
	})
}

func testAccNetworkIPInterfaceResourceConfig(svmName, address, homeNode, servicePolicy string) string {
	host := os.Getenv("TF_ACC_NETAPP_HOST5")
	admin := os.Getenv("TF_ACC_NETAPP_USER")
	password := os.Getenv("TF_ACC_NETAPP_PASS2")
	if host == "" || admin == "" || password == "" {
		fmt.Println("TF_ACC_NETAPP_HOST5, TF_ACC_NETAPP_USER, and TF_ACC_NETAPP_PASS2 must be set for acceptance tests")
		os.Exit(1)
	}
	return fmt.Sprintf(`
provider "netapp-ontap" {
 connection_profiles = [
    {
      name = "cluster4"
      hostname = "%s"
      username = "%s"
      password = "%s"
      validate_certs = false
    },
  ]
}

resource "netapp-ontap_network_ip_interface" "example" {
	cx_profile_name = "cluster4"
	name = "test-interface"
	svm_name = "%s"
  	ip = {
    	address = "%s"
    	netmask = 18
    }
  	location = {
    	home_port = "e0d"
    	home_node = "%s"
  	}
	service_policy = "%s"
}
`, host, admin, password, svmName, address, homeNode, servicePolicy)
}

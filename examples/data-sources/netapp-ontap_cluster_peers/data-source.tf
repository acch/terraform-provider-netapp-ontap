data "netapp-ontap_cluster_peers" "cluster_peers" {
  # required to know which system to interface with
  cx_profile_name = "cluster4"
  filter = {
    name = "ontap*"
  }
}

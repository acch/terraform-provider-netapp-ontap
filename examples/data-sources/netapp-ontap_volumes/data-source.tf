data "netapp-ontap_volumes" "storage_volumes" {
  # required to know which system to interface with
  cx_profile_name = "cluster4"
  filter = {
    svm_name = "svm*"
  }
}

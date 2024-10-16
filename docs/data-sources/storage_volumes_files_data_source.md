---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netapp-ontap_volumes_files Data Source - terraform-provider-netapp-ontap"
subcategory: "Storage"
description: |-
  Storage Volumes Files data source
---

# netapp-ontap_volumes_files (Data Source)

Retrieves an existing storage volumes files

## Example Usage
```terraform
data "netapp-ontap_volumes_files" "storage_volumes_files" {
  cx_profile_name = "cluster4"
  volume_name = "acc_test_peer_root"
  path = ".snapshot"
  svm_name = "acc_test_peer"
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cx_profile_name` (String) Connection profile name
- `path` (String) Relative path of a file or directory in the volume
- `svm_name` (String) svm name
- `volume_name` (String) Volume name

### Optional

- `byte_offset` (Number) The file offset to start reading from
- `name` (String) The name of the file or directory
- `overwrite_enabled` (Boolean) Whether the file can be overwritten

### Read-Only

- `storage_volumes_files` (Attributes List) (see [below for nested schema](#nestedatt--storage_volumes_files))

<a id="nestedatt--storage_volumes_files"></a>
### Nested Schema for `storage_volumes_files`

Read-Only:

- `bytes_used` (Number) The number of bytes used
- `cx_profile_name` (String) Connection profile name
- `group_id` (Number) The group ID of the file or directory
- `hard_links_count` (Number) The number of hard links to the file or directory
- `inode_number` (Number) The inode number of the file or directory
- `is_empty` (Boolean) Whether the file or directory is empty
- `name` (String) The name of the file or directory
- `overwrite_enabled` (Boolean) Whether the file can be overwritten
- `owner_id` (Number) The owner ID of the file or directory
- `path` (String) Relative path of a file or directory in the volume
- `size` (Number) The size of the file or directory
- `target` (String) Whether the file or directory is empty
- `type` (String) The type of the file or directory
- `volume_name` (String) Volume name

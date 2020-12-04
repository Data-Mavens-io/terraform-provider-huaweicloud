---
subcategory: "Distributed Message Service (DMS)"
---

# huaweicloud\_dms\_az

Use this data source to get the ID of an available HuaweiCloud dms az.
This is an alternative to `huaweicloud_dms_az_v1`

## Example Usage

```hcl

data "huaweicloud_dms_az" "az1" {
  name = "可用区1"
  port = "8002"
  code = "cn-north-1a"
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the dms az. If omitted, the provider-level region will be used.

* `name` - (Required, String) Indicates the name of an AZ.

* `code` - (Optional, String) Indicates the code of an AZ.

* `port` - (Required, String) Indicates the port number of an AZ.


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a data source ID in UUID format.

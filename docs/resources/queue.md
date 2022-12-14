---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "clearml_queue Resource - terraform-provider-clearml"
subcategory: ""
description: |-
  A queue in ClearML.
---

# clearml_queue (Resource)

A queue in ClearML.

## Example Usage

```terraform
resource "clearml_queue" "example" {
  name = "foo"
}

resource "clearml_queue" "example_with_tags" {
  name = "bar"
  tags = ["one", "two"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the queue.

### Optional

- `tags` (List of String) Tags to set on the queue.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import clearml_queue.example acbd18db4cc2f85cedef654fccc4a4d8
```

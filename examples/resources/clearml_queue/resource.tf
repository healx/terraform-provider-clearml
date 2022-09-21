resource "clearml_queue" "example" {
  name = "foo"
}

resource "clearml_queue" "example_with_tags" {
  name = "bar"
  tags = ["one", "two"]
}
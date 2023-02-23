---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_simspaceweaver_simulation Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::SimSpaceWeaver::Simulation
---

# awscc_simspaceweaver_simulation (Data Source)

Data Source schema for AWS::SimSpaceWeaver::Simulation



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `describe_payload` (String) Json object with all simulation details
- `name` (String) The name of the simulation.
- `role_arn` (String) Role ARN.
- `schema_s3_location` (Attributes) (see [below for nested schema](#nestedatt--schema_s3_location))

<a id="nestedatt--schema_s3_location"></a>
### Nested Schema for `schema_s3_location`

Read-Only:

- `bucket_name` (String) The Schema S3 bucket name.
- `object_key` (String) This is the schema S3 object key, which includes the full path of "folders" from the bucket root to the schema.


// Code generated by generators/resource/main.go; DO NOT EDIT.

package kms

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceTypeFactory("awscc_kms_replica_key", replicaKeyResourceType)
}

// replicaKeyResourceType returns the Terraform awscc_kms_replica_key resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::KMS::ReplicaKey resource type.
func replicaKeyResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the CMK. Use a description that helps you to distinguish this CMK from others in the account, such as its intended use.",
			//   "maxLength": 8192,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "A description of the CMK. Use a description that helps you to distinguish this CMK from others in the account, such as its intended use.",
			Type:        types.StringType,
			Optional:    true,
		},
		"enabled": {
			// Property: Enabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies whether the customer master key (CMK) is enabled. Disabled CMKs cannot be used in cryptographic operations.",
			//   "type": "boolean"
			// }
			Description: "Specifies whether the customer master key (CMK) is enabled. Disabled CMKs cannot be used in cryptographic operations.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"key_id": {
			// Property: KeyId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"key_policy": {
			// Property: KeyPolicy
			// CloudFormation resource type schema:
			// {
			//   "description": "The key policy that authorizes use of the CMK. The key policy must observe the following rules.",
			//   "type": "string"
			// }
			Description: "The key policy that authorizes use of the CMK. The key policy must observe the following rules.",
			Type:        types.StringType,
			Required:    true,
		},
		"pending_window_in_days": {
			// Property: PendingWindowInDays
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the number of days in the waiting period before AWS KMS deletes a CMK that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.",
			//   "type": "integer"
			// }
			Description: "Specifies the number of days in the waiting period before AWS KMS deletes a CMK that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.",
			Type:        types.NumberType,
			Optional:    true,
			// PendingWindowInDays is a write-only attribute.
		},
		"primary_key_arn": {
			// Property: PrimaryKeyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Identifies the primary CMK to create a replica of. Specify the Amazon Resource Name (ARN) of the CMK. You cannot specify an alias or key ID. For help finding the ARN, see Finding the Key ID and ARN in the AWS Key Management Service Developer Guide.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Identifies the primary CMK to create a replica of. Specify the Amazon Resource Name (ARN) of the CMK. You cannot specify an alias or key ID. For help finding the ARN, see Finding the Key ID and ARN in the AWS Key Management Service Developer Guide.",
			Type:        types.StringType,
			Required:    true,
			// PrimaryKeyArn is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The AWS::KMS::ReplicaKey resource specifies a multi-region replica customer master key (CMK) in AWS Key Management Service (AWS KMS).",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::KMS::ReplicaKey").WithTerraformTypeName("awscc_kms_replica_key")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"description":            "Description",
		"enabled":                "Enabled",
		"key":                    "Key",
		"key_id":                 "KeyId",
		"key_policy":             "KeyPolicy",
		"pending_window_in_days": "PendingWindowInDays",
		"primary_key_arn":        "PrimaryKeyArn",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/PendingWindowInDays",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_kms_replica_key", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

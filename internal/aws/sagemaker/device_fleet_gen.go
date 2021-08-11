// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_sagemaker_device_fleet", deviceFleetResourceType)
}

// deviceFleetResourceType returns the Terraform aws_sagemaker_device_fleet resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SageMaker::DeviceFleet resource type.
func deviceFleetResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Description for the edge device fleet",
			     "maxLength": 800,
			     "minLength": 0,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "Description for the edge device fleet",
			Type:        types.StringType,
			Optional:    true,
		},
		"device_fleet_name": {
			// Property: DeviceFleetName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the edge device fleet",
			     "maxLength": 63,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The name of the edge device fleet",
			Type:        types.StringType,
			Required:    true,
			// DeviceFleetName is a force-new attribute.
		},
		"output_config": {
			// Property: OutputConfig
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "KmsKeyId": {
			         "description": "The KMS key id used for encryption on the S3 bucket",
			         "maxLength": 2048,
			         "minLength": 1,
			         "pattern": "",
			         "type": "string"
			       },
			       "S3OutputLocation": {
			         "description": "The Amazon Simple Storage (S3) bucket URI",
			         "maxLength": 1024,
			         "pattern": "",
			         "type": "string"
			       }
			     },
			     "required": [
			       "S3OutputLocation"
			     ],
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"kms_key_id": {
						// Property: KmsKeyId
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The KMS key id used for encryption on the S3 bucket",
						     "maxLength": 2048,
						     "minLength": 1,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Description: "The KMS key id used for encryption on the S3 bucket",
						Type:        types.StringType,
						Optional:    true,
					},
					"s3_output_location": {
						// Property: S3OutputLocation
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The Amazon Simple Storage (S3) bucket URI",
						     "maxLength": 1024,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Description: "The Amazon Simple Storage (S3) bucket URI",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Required: true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Role associated with the device fleet",
			     "maxLength": 2048,
			     "minLength": 20,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "Role associated with the device fleet",
			Type:        types.StringType,
			Required:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Associate tags with the resource",
			     "items": {
			       "additionalProperties": false,
			       "description": "Key-value pair to associate as a tag for the resource",
			       "properties": {
			         "Key": {
			           "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			           "maxLength": 128,
			           "minLength": 1,
			           "pattern": "",
			           "type": "string"
			         },
			         "Value": {
			           "description": "The key value of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			           "maxLength": 256,
			           "minLength": 0,
			           "pattern": "",
			           "type": "string"
			         }
			       },
			       "required": [
			         "Key",
			         "Value"
			       ],
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Description: "Associate tags with the resource",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						     "maxLength": 128,
						     "minLength": 1,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The key value of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						     "maxLength": 256,
						     "minLength": 0,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Description: "The key value of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::SageMaker::DeviceFleet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::DeviceFleet").WithTerraformTypeName("aws_sagemaker_device_fleet").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_sagemaker_device_fleet", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

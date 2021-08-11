// Code generated by generators/resource/main.go; DO NOT EDIT.

package iam

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
	registry.AddResourceTypeFactory("aws_iam_oidc_provider", oIDCProviderResourceType)
}

// oIDCProviderResourceType returns the Terraform aws_iam_oidc_provider resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IAM::OIDCProvider resource type.
func oIDCProviderResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Amazon Resource Name (ARN) of the OIDC provider",
			     "maxLength": 2048,
			     "minLength": 20,
			     "type": "string"
			   }
			*/
			Description: "Amazon Resource Name (ARN) of the OIDC provider",
			Type:        types.StringType,
			Computed:    true,
		},
		"client_id_list": {
			// Property: ClientIdList
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "maxLength": 255,
			       "minLength": 1,
			       "type": "string"
			     },
			     "type": "array"
			   }
			*/
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "additionalProperties": false,
			       "description": "A key-value pair to associate with a resource.",
			       "properties": {
			         "Key": {
			           "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			           "maxLength": 128,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Value": {
			           "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			           "maxLength": 256,
			           "minLength": 1,
			           "type": "string"
			         }
			       },
			       "required": [
			         "Value",
			         "Key"
			       ],
			       "type": "object"
			     },
			     "type": "array",
			     "uniqueItems": false
			   }
			*/
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						     "maxLength": 128,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						     "maxLength": 256,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"thumbprint_list": {
			// Property: ThumbprintList
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "maxLength": 40,
			       "minLength": 40,
			       "pattern": "",
			       "type": "string"
			     },
			     "maxItems": 5,
			     "type": "array"
			   }
			*/
			Type:     types.ListType{ElemType: types.StringType},
			Required: true,
		},
		"url": {
			// Property: Url
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 255,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// Url is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::IAM::OIDCProvider",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IAM::OIDCProvider").WithTerraformTypeName("aws_iam_oidc_provider").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_iam_oidc_provider", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

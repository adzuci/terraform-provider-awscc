// Code generated by generators/resource/main.go; DO NOT EDIT.

package codeartifact

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
	registry.AddResourceTypeFactory("aws_codeartifact_domain", domainResourceType)
}

// domainResourceType returns the Terraform aws_codeartifact_domain resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CodeArtifact::Domain resource type.
func domainResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ARN of the domain.",
			     "maxLength": 2048,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "The ARN of the domain.",
			Type:        types.StringType,
			Computed:    true,
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the domain.",
			     "maxLength": 50,
			     "minLength": 2,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The name of the domain.",
			Type:        types.StringType,
			Required:    true,
			// DomainName is a force-new attribute.
		},
		"encryption_key": {
			// Property: EncryptionKey
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ARN of an AWS Key Management Service (AWS KMS) key associated with a domain.",
			     "type": "string"
			   }
			*/
			Description: "The ARN of an AWS Key Management Service (AWS KMS) key associated with a domain.",
			Type:        types.StringType,
			Computed:    true,
			// EncryptionKey is a force-new attribute.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the domain. This field is used for GetAtt",
			     "maxLength": 50,
			     "minLength": 2,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The name of the domain. This field is used for GetAtt",
			Type:        types.StringType,
			Computed:    true,
		},
		"owner": {
			// Property: Owner
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The 12-digit account ID of the AWS account that owns the domain. This field is used for GetAtt",
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The 12-digit account ID of the AWS account that owns the domain. This field is used for GetAtt",
			Type:        types.StringType,
			Computed:    true,
		},
		"permissions_policy_document": {
			// Property: PermissionsPolicyDocument
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The access control resource policy on the provided domain.",
			     "maxLength": 5120,
			     "minLength": 2,
			     "type": "object"
			   }
			*/
			Description: "The access control resource policy on the provided domain.",
			Type:        types.MapType{ElemType: types.StringType},
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "An array of key-value pairs to apply to this resource.",
			     "items": {
			       "additionalProperties": false,
			       "description": "A key-value pair to associate with a resource.",
			       "properties": {
			         "Key": {
			           "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			           "maxLength": 128,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Value": {
			           "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			           "maxLength": 256,
			           "minLength": 0,
			           "type": "string"
			         }
			       },
			       "required": [
			         "Value",
			         "Key"
			       ],
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Description: "An array of key-value pairs to apply to this resource.",
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
						     "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						     "maxLength": 256,
						     "minLength": 0,
						     "type": "string"
						   }
						*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
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
		Description: "The resource schema to create a CodeArtifact domain.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeArtifact::Domain").WithTerraformTypeName("aws_codeartifact_domain").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_codeartifact_domain", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

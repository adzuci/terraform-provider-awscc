// Code generated by generators/resource/main.go; DO NOT EDIT.

package emr

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/types"
)

func init() {
	registry.AddResourceTypeFactory("aws_emr_studio", studioResourceType)
}

// studioResourceType returns the Terraform aws_emr_studio resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EMR::Studio resource type.
func studioResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"auth_mode": {
			// Property: AuthMode
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Specifies whether the Studio authenticates users using single sign-on (SSO) or IAM. Amazon EMR Studio currently only supports SSO authentication.",
			     "enum": [
			       "SSO",
			       "IAM"
			     ],
			     "type": "string"
			   }
			*/
			Description: "Specifies whether the Studio authenticates users using single sign-on (SSO) or IAM. Amazon EMR Studio currently only supports SSO authentication.",
			Type:        types.StringType,
			Required:    true,
			// AuthMode is a force-new attribute.
		},
		"default_s3_location": {
			// Property: DefaultS3Location
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The default Amazon S3 location to back up EMR Studio Workspaces and notebook files. A Studio user can select an alternative Amazon S3 location when creating a Workspace.",
			     "maxLength": 10280,
			     "minLength": 6,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The default Amazon S3 location to back up EMR Studio Workspaces and notebook files. A Studio user can select an alternative Amazon S3 location when creating a Workspace.",
			Type:        types.StringType,
			Required:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A detailed description of the Studio.",
			     "maxLength": 256,
			     "minLength": 0,
			     "type": "string"
			   }
			*/
			Description: "A detailed description of the Studio.",
			Type:        types.StringType,
			Optional:    true,
		},
		"engine_security_group_id": {
			// Property: EngineSecurityGroupId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the Amazon EMR Studio Engine security group. The Engine security group allows inbound network traffic from the Workspace security group, and it must be in the same VPC specified by VpcId.",
			     "maxLength": 256,
			     "minLength": 4,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The ID of the Amazon EMR Studio Engine security group. The Engine security group allows inbound network traffic from the Workspace security group, and it must be in the same VPC specified by VpcId.",
			Type:        types.StringType,
			Required:    true,
			// EngineSecurityGroupId is a force-new attribute.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A descriptive name for the Amazon EMR Studio.",
			     "maxLength": 256,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "A descriptive name for the Amazon EMR Studio.",
			Type:        types.StringType,
			Required:    true,
		},
		"service_role": {
			// Property: ServiceRole
			// CloudFormation resource type schema:
			/*
			   {
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
			// ServiceRole is a force-new attribute.
		},
		"studio_id": {
			// Property: StudioId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the EMR Studio.",
			     "maxLength": 256,
			     "minLength": 4,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The ID of the EMR Studio.",
			Type:        types.StringType,
			Computed:    true,
		},
		"subnet_ids": {
			// Property: SubnetIds
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A list of up to 5 subnet IDs to associate with the Studio. The subnets must belong to the VPC specified by VpcId. Studio users can create a Workspace in any of the specified subnets.",
			     "items": {
			       "description": "Identifier of a subnet",
			       "pattern": "",
			       "type": "string"
			     },
			     "minItems": 1,
			     "type": "array"
			   }
			*/
			Description: "A list of up to 5 subnet IDs to associate with the Studio. The subnets must belong to the VPC specified by VpcId. Studio users can create a Workspace in any of the specified subnets.",
			Type:        types.ListType{ElemType: types.StringType},
			Required:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "insertionOrder": false,
			     "items": {
			       "additionalProperties": false,
			       "description": "An arbitrary set of tags (key-value pairs) for this EMR Studio.",
			       "properties": {
			         "Key": {
			           "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			           "maxLength": 128,
			           "minLength": 1,
			           "pattern": "",
			           "type": "string"
			         },
			         "Value": {
			           "description": "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			           "maxLength": 256,
			           "minLength": 0,
			           "pattern": "",
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
			     "uniqueItems": true
			   }
			*/
			Attributes: providertypes.SetNestedAttributes(
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
						     "description": "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						     "maxLength": 256,
						     "minLength": 0,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Description: "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
		"url": {
			// Property: Url
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The unique Studio access URL.",
			     "maxLength": 4096,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The unique Studio access URL.",
			Type:        types.StringType,
			Computed:    true,
		},
		"user_role": {
			// Property: UserRole
			// CloudFormation resource type schema:
			/*
			   {
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
			// UserRole is a force-new attribute.
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the Studio.",
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the Studio.",
			Type:        types.StringType,
			Required:    true,
			// VpcId is a force-new attribute.
		},
		"workspace_security_group_id": {
			// Property: WorkspaceSecurityGroupId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the Amazon EMR Studio Workspace security group. The Workspace security group allows outbound network traffic to resources in the Engine security group, and it must be in the same VPC specified by VpcId.",
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The ID of the Amazon EMR Studio Workspace security group. The Workspace security group allows outbound network traffic to resources in the Engine security group, and it must be in the same VPC specified by VpcId.",
			Type:        types.StringType,
			Required:    true,
			// WorkspaceSecurityGroupId is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::EMR::Studio",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EMR::Studio").WithTerraformTypeName("aws_emr_studio").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_emr_studio", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
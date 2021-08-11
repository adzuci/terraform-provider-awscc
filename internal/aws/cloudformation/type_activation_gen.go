// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudformation

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
	registry.AddResourceTypeFactory("aws_cloudformation_type_activation", typeActivationResourceType)
}

// typeActivationResourceType returns the Terraform aws_cloudformation_type_activation resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CloudFormation::TypeActivation resource type.
func typeActivationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the extension.",
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the extension.",
			Type:        types.StringType,
			Computed:    true,
		},
		"auto_update": {
			// Property: AutoUpdate
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Whether to automatically update the extension in this account and region when a new minor version is published by the extension publisher. Major versions released by the publisher must be manually updated.",
			     "type": "boolean"
			   }
			*/
			Description: "Whether to automatically update the extension in this account and region when a new minor version is published by the extension publisher. Major versions released by the publisher must be manually updated.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"execution_role_arn": {
			// Property: ExecutionRoleArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ExecutionRoleArn is a force-new attribute.
		},
		"logging_config": {
			// Property: LoggingConfig
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "LogGroupName": {
			         "description": "The Amazon CloudWatch log group to which CloudFormation sends error logging information when invoking the type's handlers.",
			         "maxLength": 512,
			         "minLength": 1,
			         "pattern": "",
			         "type": "string"
			       },
			       "LogRoleArn": {
			         "description": "The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.",
			         "maxLength": 256,
			         "minLength": 1,
			         "type": "string"
			       }
			     },
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"log_group_name": {
						// Property: LogGroupName
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The Amazon CloudWatch log group to which CloudFormation sends error logging information when invoking the type's handlers.",
						     "maxLength": 512,
						     "minLength": 1,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Description: "The Amazon CloudWatch log group to which CloudFormation sends error logging information when invoking the type's handlers.",
						Type:        types.StringType,
						Optional:    true,
					},
					"log_role_arn": {
						// Property: LogRoleArn
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.",
						     "maxLength": 256,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
			Computed: true,
			// LoggingConfig is a force-new attribute.
		},
		"major_version": {
			// Property: MajorVersion
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Major Version of the type you want to enable",
			     "maxLength": 100000,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "The Major Version of the type you want to enable",
			Type:        types.StringType,
			Optional:    true,
		},
		"public_type_arn": {
			// Property: PublicTypeArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Number (ARN) assigned to the public extension upon publication",
			     "maxLength": 1024,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Number (ARN) assigned to the public extension upon publication",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// PublicTypeArn is a force-new attribute.
		},
		"publisher_id": {
			// Property: PublisherId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The publisher id assigned by CloudFormation for publishing in this region.",
			     "maxLength": 40,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The publisher id assigned by CloudFormation for publishing in this region.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// PublisherId is a force-new attribute.
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The kind of extension",
			     "enum": [
			       "RESOURCE",
			       "MODULE"
			     ],
			     "type": "string"
			   }
			*/
			Description: "The kind of extension",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Type is a force-new attribute.
		},
		"type_name": {
			// Property: TypeName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// TypeName is a force-new attribute.
		},
		"type_name_alias": {
			// Property: TypeNameAlias
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "An alias to assign to the public extension in this account and region. If you specify an alias for the extension, you must then use the alias to refer to the extension in your templates.",
			     "maxLength": 204,
			     "minLength": 10,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "An alias to assign to the public extension in this account and region. If you specify an alias for the extension, you must then use the alias to refer to the extension in your templates.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// TypeNameAlias is a force-new attribute.
		},
		"version_bump": {
			// Property: VersionBump
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Manually updates a previously-enabled type to a new major or minor version, if available. You can also use this parameter to update the value of AutoUpdateEnabled",
			     "enum": [
			       "MAJOR",
			       "MINOR"
			     ],
			     "type": "string"
			   }
			*/
			Description: "Manually updates a previously-enabled type to a new major or minor version, if available. You can also use this parameter to update the value of AutoUpdateEnabled",
			Type:        types.StringType,
			Optional:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Enable a resource that has been published in the CloudFormation Registry.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::TypeActivation").WithTerraformTypeName("aws_cloudformation_type_activation").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_cloudformation_type_activation", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

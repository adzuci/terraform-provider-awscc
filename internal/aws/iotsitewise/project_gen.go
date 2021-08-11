// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotsitewise

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
	registry.AddResourceTypeFactory("aws_iotsitewise_project", projectResourceType)
}

// projectResourceType returns the Terraform aws_iotsitewise_project resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTSiteWise::Project resource type.
func projectResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"portal_id": {
			// Property: PortalId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the portal in which to create the project.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the portal in which to create the project.",
			Type:        types.StringType,
			Required:    true,
			// PortalId is a force-new attribute.
		},
		"project_arn": {
			// Property: ProjectArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ARN of the project.",
			     "type": "string"
			   }
			*/
			Description: "The ARN of the project.",
			Type:        types.StringType,
			Computed:    true,
		},
		"project_description": {
			// Property: ProjectDescription
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A description for the project.",
			     "type": "string"
			   }
			*/
			Description: "A description for the project.",
			Type:        types.StringType,
			Optional:    true,
		},
		"project_id": {
			// Property: ProjectId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the project.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the project.",
			Type:        types.StringType,
			Computed:    true,
		},
		"project_name": {
			// Property: ProjectName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A friendly name for the project.",
			     "type": "string"
			   }
			*/
			Description: "A friendly name for the project.",
			Type:        types.StringType,
			Required:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A list of key-value pairs that contain metadata for the project.",
			     "items": {
			       "additionalProperties": false,
			       "description": "To add or update tag, provide both key and value. To delete tag, provide only tag key to be deleted",
			       "properties": {
			         "Key": {
			           "type": "string"
			         },
			         "Value": {
			           "type": "string"
			         }
			       },
			       "required": [
			         "Key",
			         "Value"
			       ],
			       "type": "object"
			     },
			     "type": "array",
			     "uniqueItems": false
			   }
			*/
			Description: "A list of key-value pairs that contain metadata for the project.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
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
		Description: "Resource schema for AWS::IoTSiteWise::Project",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::Project").WithTerraformTypeName("aws_iotsitewise_project").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_iotsitewise_project", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

// Code generated by generators/resource/main.go; DO NOT EDIT.

package config

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
	registry.AddResourceTypeFactory("aws_config_conformance_pack", conformancePackResourceType)
}

// conformancePackResourceType returns the Terraform aws_config_conformance_pack resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Config::ConformancePack resource type.
func conformancePackResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"conformance_pack_input_parameters": {
			// Property: ConformancePackInputParameters
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A list of ConformancePackInputParameter objects.",
			     "items": {
			       "description": "Input parameters in the form of key-value pairs for the conformance pack.",
			       "properties": {
			         "ParameterName": {
			           "description": "Key part of key-value pair with value being parameter value",
			           "maxLength": 255,
			           "minLength": 0,
			           "type": "string"
			         },
			         "ParameterValue": {
			           "description": "Value part of key-value pair with key being parameter Name",
			           "maxLength": 4096,
			           "minLength": 0,
			           "type": "string"
			         }
			       },
			       "required": [
			         "ParameterName",
			         "ParameterValue"
			       ],
			       "type": "object"
			     },
			     "maxItems": 60,
			     "minItems": 0,
			     "type": "array"
			   }
			*/
			Description: "A list of ConformancePackInputParameter objects.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"parameter_name": {
						// Property: ParameterName
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Key part of key-value pair with value being parameter value",
						     "maxLength": 255,
						     "minLength": 0,
						     "type": "string"
						   }
						*/
						Description: "Key part of key-value pair with value being parameter value",
						Type:        types.StringType,
						Required:    true,
					},
					"parameter_value": {
						// Property: ParameterValue
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Value part of key-value pair with key being parameter Name",
						     "maxLength": 4096,
						     "minLength": 0,
						     "type": "string"
						   }
						*/
						Description: "Value part of key-value pair with key being parameter Name",
						Type:        types.StringType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 60,
				},
			),
			Optional: true,
		},
		"conformance_pack_name": {
			// Property: ConformancePackName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Name of the conformance pack which will be assigned as the unique identifier.",
			     "maxLength": 256,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "Name of the conformance pack which will be assigned as the unique identifier.",
			Type:        types.StringType,
			Required:    true,
			// ConformancePackName is a force-new attribute.
		},
		"delivery_s3_bucket": {
			// Property: DeliveryS3Bucket
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "AWS Config stores intermediate files while processing conformance pack template.",
			     "maxLength": 63,
			     "minLength": 0,
			     "type": "string"
			   }
			*/
			Description: "AWS Config stores intermediate files while processing conformance pack template.",
			Type:        types.StringType,
			Optional:    true,
		},
		"delivery_s3_key_prefix": {
			// Property: DeliveryS3KeyPrefix
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The prefix for delivery S3 bucket.",
			     "maxLength": 1024,
			     "minLength": 0,
			     "type": "string"
			   }
			*/
			Description: "The prefix for delivery S3 bucket.",
			Type:        types.StringType,
			Optional:    true,
		},
		"template_body": {
			// Property: TemplateBody
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A string containing full conformance pack template body. You can only specify one of the template body or template S3Uri fields.",
			     "maxLength": 51200,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "A string containing full conformance pack template body. You can only specify one of the template body or template S3Uri fields.",
			Type:        types.StringType,
			Optional:    true,
			// TemplateBody is a write-only attribute.
		},
		"template_s3_uri": {
			// Property: TemplateS3Uri
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Location of file containing the template body which points to the conformance pack template that is located in an Amazon S3 bucket. You can only specify one of the template body or template S3Uri fields.",
			     "maxLength": 1024,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "Location of file containing the template body which points to the conformance pack template that is located in an Amazon S3 bucket. You can only specify one of the template body or template S3Uri fields.",
			Type:        types.StringType,
			Optional:    true,
			// TemplateS3Uri is a write-only attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "A conformance pack is a collection of AWS Config rules and remediation actions that can be easily deployed as a single entity in an account and a region or across an entire AWS Organization.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Config::ConformancePack").WithTerraformTypeName("aws_config_conformance_pack").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/TemplateBody",
		"/properties/TemplateS3Uri",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_config_conformance_pack", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

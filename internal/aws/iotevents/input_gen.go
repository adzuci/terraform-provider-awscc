// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotevents

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
	registry.AddResourceTypeFactory("aws_iotevents_input", inputResourceType)
}

// inputResourceType returns the Terraform aws_iotevents_input resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTEvents::Input resource type.
func inputResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"input_definition": {
			// Property: InputDefinition
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "The definition of the input.",
			     "properties": {
			       "Attributes": {
			         "description": "The attributes from the JSON payload that are made available by the input. Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage`. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.",
			         "items": {
			           "additionalProperties": false,
			           "description": "The attributes from the JSON payload that are made available by the input. Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage`. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.",
			           "properties": {
			             "JsonPath": {
			               "description": "An expression that specifies an attribute-value pair in a JSON structure. Use this to specify an attribute from the JSON payload that is made available by the input. Inputs are derived from messages sent to AWS IoT Events (`BatchPutMessage`). Each such message contains a JSON payload. The attribute (and its paired value) specified here are available for use in the `condition` expressions used by detectors.\n\n_Syntax_: `\u003cfield-name\u003e.\u003cfield-name\u003e...`",
			               "maxLength": 128,
			               "minLength": 1,
			               "pattern": "",
			               "type": "string"
			             }
			           },
			           "required": [
			             "JsonPath"
			           ],
			           "type": "object"
			         },
			         "maxItems": 200,
			         "minItems": 1,
			         "type": "array",
			         "uniqueItems": true
			       }
			     },
			     "required": [
			       "Attributes"
			     ],
			     "type": "object"
			   }
			*/
			Description: "The definition of the input.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"attributes": {
						// Property: Attributes
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The attributes from the JSON payload that are made available by the input. Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage`. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.",
						     "items": {
						       "additionalProperties": false,
						       "description": "The attributes from the JSON payload that are made available by the input. Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage`. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.",
						       "properties": {
						         "JsonPath": {
						           "description": "An expression that specifies an attribute-value pair in a JSON structure. Use this to specify an attribute from the JSON payload that is made available by the input. Inputs are derived from messages sent to AWS IoT Events (`BatchPutMessage`). Each such message contains a JSON payload. The attribute (and its paired value) specified here are available for use in the `condition` expressions used by detectors.\n\n_Syntax_: `\u003cfield-name\u003e.\u003cfield-name\u003e...`",
						           "maxLength": 128,
						           "minLength": 1,
						           "pattern": "",
						           "type": "string"
						         }
						       },
						       "required": [
						         "JsonPath"
						       ],
						       "type": "object"
						     },
						     "maxItems": 200,
						     "minItems": 1,
						     "type": "array",
						     "uniqueItems": true
						   }
						*/
						Description: "The attributes from the JSON payload that are made available by the input. Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage`. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.",
						// Ordered set.
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"json_path": {
									// Property: JsonPath
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "An expression that specifies an attribute-value pair in a JSON structure. Use this to specify an attribute from the JSON payload that is made available by the input. Inputs are derived from messages sent to AWS IoT Events (`BatchPutMessage`). Each such message contains a JSON payload. The attribute (and its paired value) specified here are available for use in the `condition` expressions used by detectors.\n\n_Syntax_: `\u003cfield-name\u003e.\u003cfield-name\u003e...`",
									     "maxLength": 128,
									     "minLength": 1,
									     "pattern": "",
									     "type": "string"
									   }
									*/
									Description: "An expression that specifies an attribute-value pair in a JSON structure. Use this to specify an attribute from the JSON payload that is made available by the input. Inputs are derived from messages sent to AWS IoT Events (`BatchPutMessage`). Each such message contains a JSON payload. The attribute (and its paired value) specified here are available for use in the `condition` expressions used by detectors.\n\n_Syntax_: `<field-name>.<field-name>...`",
									Type:        types.StringType,
									Required:    true,
								},
							},
							schema.ListNestedAttributesOptions{
								MinItems: 1,
								MaxItems: 200,
							},
						),
						Required: true,
					},
				},
			),
			Required: true,
		},
		"input_description": {
			// Property: InputDescription
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A brief description of the input.",
			     "maxLength": 128,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "A brief description of the input.",
			Type:        types.StringType,
			Optional:    true,
		},
		"input_name": {
			// Property: InputName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the input.",
			     "maxLength": 128,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The name of the input.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// InputName is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "An array of key-value pairs to apply to this resource.\n\nFor more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
			     "items": {
			       "additionalProperties": false,
			       "description": "Tags to be applied to Input.",
			       "properties": {
			         "Key": {
			           "description": "Key of the Tag.",
			           "type": "string"
			         },
			         "Value": {
			           "description": "Value of the Tag.",
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
			Description: "An array of key-value pairs to apply to this resource.\n\nFor more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Key of the Tag.",
						     "type": "string"
						   }
						*/
						Description: "Key of the Tag.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Value of the Tag.",
						     "type": "string"
						   }
						*/
						Description: "Value of the Tag.",
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
		Description: "The AWS::IoTEvents::Input resource creates an input. To monitor your devices and processes, they must have a way to get telemetry data into AWS IoT Events. This is done by sending messages as *inputs* to AWS IoT Events. For more information, see [How to Use AWS IoT Events](https://docs.aws.amazon.com/iotevents/latest/developerguide/how-to-use-iotevents.html) in the *AWS IoT Events Developer Guide*.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTEvents::Input").WithTerraformTypeName("aws_iotevents_input").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_iotevents_input", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
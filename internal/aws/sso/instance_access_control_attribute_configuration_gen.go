// Code generated by generators/resource/main.go; DO NOT EDIT.

package sso

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
	registry.AddResourceTypeFactory("aws_sso_instance_access_control_attribute_configuration", instanceAccessControlAttributeConfigurationResourceType)
}

// instanceAccessControlAttributeConfigurationResourceType returns the Terraform aws_sso_instance_access_control_attribute_configuration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SSO::InstanceAccessControlAttributeConfiguration resource type.
func instanceAccessControlAttributeConfigurationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"access_control_attributes": {
			// Property: AccessControlAttributes
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Key": {
			           "maxLength": 128,
			           "minLength": 1,
			           "pattern": "",
			           "type": "string"
			         },
			         "Value": {
			           "additionalProperties": false,
			           "properties": {
			             "Source": {
			               "items": {
			                 "maxLength": 256,
			                 "minLength": 0,
			                 "pattern": "",
			                 "type": "string"
			               },
			               "maxItems": 1,
			               "type": "array"
			             }
			           },
			           "required": [
			             "Source"
			           ],
			           "type": "object"
			         }
			       },
			       "required": [
			         "Key",
			         "Value"
			       ],
			       "type": "object"
			     },
			     "maxItems": 50,
			     "type": "array"
			   }
			*/
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 128,
						     "minLength": 1,
						     "pattern": "",
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
						     "additionalProperties": false,
						     "properties": {
						       "Source": {
						         "items": {
						           "maxLength": 256,
						           "minLength": 0,
						           "pattern": "",
						           "type": "string"
						         },
						         "maxItems": 1,
						         "type": "array"
						       }
						     },
						     "required": [
						       "Source"
						     ],
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"source": {
									// Property: Source
									// CloudFormation resource type schema:
									/*
									   {
									     "items": {
									       "maxLength": 256,
									       "minLength": 0,
									       "pattern": "",
									       "type": "string"
									     },
									     "maxItems": 1,
									     "type": "array"
									   }
									*/
									Type:     types.ListType{ElemType: types.StringType},
									Required: true,
								},
							},
						),
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
		},
		"instance_access_control_attribute_configuration": {
			// Property: InstanceAccessControlAttributeConfiguration
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "The InstanceAccessControlAttributeConfiguration property has been deprecated but is still supported for backwards compatibility purposes. We recomend that you use  AccessControlAttributes property instead.",
			     "properties": {
			       "AccessControlAttributes": {
			         "items": {
			           "additionalProperties": false,
			           "properties": {
			             "Key": {
			               "maxLength": 128,
			               "minLength": 1,
			               "pattern": "",
			               "type": "string"
			             },
			             "Value": {
			               "additionalProperties": false,
			               "properties": {
			                 "Source": {
			                   "items": {
			                     "maxLength": 256,
			                     "minLength": 0,
			                     "pattern": "",
			                     "type": "string"
			                   },
			                   "maxItems": 1,
			                   "type": "array"
			                 }
			               },
			               "required": [
			                 "Source"
			               ],
			               "type": "object"
			             }
			           },
			           "required": [
			             "Key",
			             "Value"
			           ],
			           "type": "object"
			         },
			         "maxItems": 50,
			         "type": "array"
			       }
			     },
			     "required": [
			       "AccessControlAttributes"
			     ],
			     "type": "object"
			   }
			*/
			Description: "The InstanceAccessControlAttributeConfiguration property has been deprecated but is still supported for backwards compatibility purposes. We recomend that you use  AccessControlAttributes property instead.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"access_control_attributes": {
						// Property: AccessControlAttributes
						// CloudFormation resource type schema:
						/*
						   {
						     "items": {
						       "additionalProperties": false,
						       "properties": {
						         "Key": {
						           "maxLength": 128,
						           "minLength": 1,
						           "pattern": "",
						           "type": "string"
						         },
						         "Value": {
						           "additionalProperties": false,
						           "properties": {
						             "Source": {
						               "items": {
						                 "maxLength": 256,
						                 "minLength": 0,
						                 "pattern": "",
						                 "type": "string"
						               },
						               "maxItems": 1,
						               "type": "array"
						             }
						           },
						           "required": [
						             "Source"
						           ],
						           "type": "object"
						         }
						       },
						       "required": [
						         "Key",
						         "Value"
						       ],
						       "type": "object"
						     },
						     "maxItems": 50,
						     "type": "array"
						   }
						*/
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"key": {
									// Property: Key
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 128,
									     "minLength": 1,
									     "pattern": "",
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
									     "additionalProperties": false,
									     "properties": {
									       "Source": {
									         "items": {
									           "maxLength": 256,
									           "minLength": 0,
									           "pattern": "",
									           "type": "string"
									         },
									         "maxItems": 1,
									         "type": "array"
									       }
									     },
									     "required": [
									       "Source"
									     ],
									     "type": "object"
									   }
									*/
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"source": {
												// Property: Source
												// CloudFormation resource type schema:
												/*
												   {
												     "items": {
												       "maxLength": 256,
												       "minLength": 0,
												       "pattern": "",
												       "type": "string"
												     },
												     "maxItems": 1,
												     "type": "array"
												   }
												*/
												Type:     types.ListType{ElemType: types.StringType},
												Required: true,
											},
										},
									),
									Required: true,
								},
							},
							schema.ListNestedAttributesOptions{
								MaxItems: 50,
							},
						),
						Required: true,
					},
				},
			),
			Optional: true,
		},
		"instance_arn": {
			// Property: InstanceArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ARN of the AWS SSO instance under which the operation will be executed.",
			     "maxLength": 1224,
			     "minLength": 10,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The ARN of the AWS SSO instance under which the operation will be executed.",
			Type:        types.StringType,
			Required:    true,
			// InstanceArn is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for SSO InstanceAccessControlAttributeConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSO::InstanceAccessControlAttributeConfiguration").WithTerraformTypeName("aws_sso_instance_access_control_attribute_configuration").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_sso_instance_access_control_attribute_configuration", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotwireless

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
	registry.AddResourceTypeFactory("aws_iotwireless_device_profile", deviceProfileResourceType)
}

// deviceProfileResourceType returns the Terraform aws_iotwireless_device_profile resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTWireless::DeviceProfile resource type.
func deviceProfileResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Service profile Arn. Returned after successful create.",
			     "type": "string"
			   }
			*/
			Description: "Service profile Arn. Returned after successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Service profile Id. Returned after successful create.",
			     "maxLength": 256,
			     "type": "string"
			   }
			*/
			Description: "Service profile Id. Returned after successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"lo_ra_wan": {
			// Property: LoRaWAN
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "ClassBTimeout": {
			         "type": "integer"
			       },
			       "ClassCTimeout": {
			         "type": "integer"
			       },
			       "MacVersion": {
			         "maxLength": 64,
			         "type": "string"
			       },
			       "MaxDutyCycle": {
			         "type": "integer"
			       },
			       "MaxEirp": {
			         "type": "integer"
			       },
			       "PingSlotDr": {
			         "type": "integer"
			       },
			       "PingSlotFreq": {
			         "type": "integer"
			       },
			       "PingSlotPeriod": {
			         "type": "integer"
			       },
			       "RegParamsRevision": {
			         "maxLength": 64,
			         "type": "string"
			       },
			       "RfRegion": {
			         "maxLength": 64,
			         "type": "string"
			       },
			       "Supports32BitFCnt": {
			         "type": "boolean"
			       },
			       "SupportsClassB": {
			         "type": "boolean"
			       },
			       "SupportsClassC": {
			         "type": "boolean"
			       },
			       "SupportsJoin": {
			         "type": "boolean"
			       }
			     },
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"class_b_timeout": {
						// Property: ClassBTimeout
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "integer"
						   }
						*/
						Type:     types.NumberType,
						Optional: true,
					},
					"class_c_timeout": {
						// Property: ClassCTimeout
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "integer"
						   }
						*/
						Type:     types.NumberType,
						Optional: true,
					},
					"mac_version": {
						// Property: MacVersion
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 64,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"max_duty_cycle": {
						// Property: MaxDutyCycle
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "integer"
						   }
						*/
						Type:     types.NumberType,
						Optional: true,
					},
					"max_eirp": {
						// Property: MaxEirp
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "integer"
						   }
						*/
						Type:     types.NumberType,
						Optional: true,
					},
					"ping_slot_dr": {
						// Property: PingSlotDr
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "integer"
						   }
						*/
						Type:     types.NumberType,
						Optional: true,
					},
					"ping_slot_freq": {
						// Property: PingSlotFreq
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "integer"
						   }
						*/
						Type:     types.NumberType,
						Optional: true,
					},
					"ping_slot_period": {
						// Property: PingSlotPeriod
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "integer"
						   }
						*/
						Type:     types.NumberType,
						Optional: true,
					},
					"reg_params_revision": {
						// Property: RegParamsRevision
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 64,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"rf_region": {
						// Property: RfRegion
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 64,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"supports_32_bit_f_cnt": {
						// Property: Supports32BitFCnt
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "boolean"
						   }
						*/
						Type:     types.BoolType,
						Optional: true,
					},
					"supports_class_b": {
						// Property: SupportsClassB
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "boolean"
						   }
						*/
						Type:     types.BoolType,
						Optional: true,
					},
					"supports_class_c": {
						// Property: SupportsClassC
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "boolean"
						   }
						*/
						Type:     types.BoolType,
						Optional: true,
					},
					"supports_join": {
						// Property: SupportsJoin
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "boolean"
						   }
						*/
						Type:     types.BoolType,
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Name of service profile",
			     "maxLength": 256,
			     "type": "string"
			   }
			*/
			Description: "Name of service profile",
			Type:        types.StringType,
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A list of key-value pairs that contain metadata for the device profile.",
			     "insertionOrder": false,
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Key": {
			           "maxLength": 128,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Value": {
			           "maxLength": 256,
			           "minLength": 1,
			           "type": "string"
			         }
			       },
			       "type": "object"
			     },
			     "maxItems": 50,
			     "type": "array",
			     "uniqueItems": true
			   }
			*/
			Description: "A list of key-value pairs that contain metadata for the device profile.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 128,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 256,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MaxItems: 50,
				},
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
		Description: "Device Profile's resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::DeviceProfile").WithTerraformTypeName("aws_iotwireless_device_profile").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_iotwireless_device_profile", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

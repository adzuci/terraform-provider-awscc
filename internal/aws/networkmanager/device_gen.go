// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkmanager

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
	registry.AddResourceTypeFactory("aws_networkmanager_device", deviceResourceType)
}

// deviceResourceType returns the Terraform aws_networkmanager_device resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NetworkManager::Device resource type.
func deviceResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The description of the device.",
			     "type": "string"
			   }
			*/
			Description: "The description of the device.",
			Type:        types.StringType,
			Optional:    true,
		},
		"device_arn": {
			// Property: DeviceArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the device.",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the device.",
			Type:        types.StringType,
			Computed:    true,
		},
		"device_id": {
			// Property: DeviceId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the device.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the device.",
			Type:        types.StringType,
			Computed:    true,
		},
		"global_network_id": {
			// Property: GlobalNetworkId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the global network.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the global network.",
			Type:        types.StringType,
			Required:    true,
			// GlobalNetworkId is a force-new attribute.
		},
		"location": {
			// Property: Location
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "The site location.",
			     "properties": {
			       "Address": {
			         "description": "The physical address.",
			         "type": "string"
			       },
			       "Latitude": {
			         "description": "The latitude.",
			         "type": "string"
			       },
			       "Longitude": {
			         "description": "The longitude.",
			         "type": "string"
			       }
			     },
			     "type": "object"
			   }
			*/
			Description: "The site location.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"address": {
						// Property: Address
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The physical address.",
						     "type": "string"
						   }
						*/
						Description: "The physical address.",
						Type:        types.StringType,
						Optional:    true,
					},
					"latitude": {
						// Property: Latitude
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The latitude.",
						     "type": "string"
						   }
						*/
						Description: "The latitude.",
						Type:        types.StringType,
						Optional:    true,
					},
					"longitude": {
						// Property: Longitude
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The longitude.",
						     "type": "string"
						   }
						*/
						Description: "The longitude.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"model": {
			// Property: Model
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The device model",
			     "type": "string"
			   }
			*/
			Description: "The device model",
			Type:        types.StringType,
			Optional:    true,
		},
		"serial_number": {
			// Property: SerialNumber
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The device serial number.",
			     "type": "string"
			   }
			*/
			Description: "The device serial number.",
			Type:        types.StringType,
			Optional:    true,
		},
		"site_id": {
			// Property: SiteId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The site ID.",
			     "type": "string"
			   }
			*/
			Description: "The site ID.",
			Type:        types.StringType,
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The tags for the device.",
			     "items": {
			       "additionalProperties": false,
			       "description": "A key-value pair to associate with a device resource.",
			       "properties": {
			         "Key": {
			           "type": "string"
			         },
			         "Value": {
			           "type": "string"
			         }
			       },
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Description: "The tags for the device.",
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
						Optional: true,
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
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The device type.",
			     "type": "string"
			   }
			*/
			Description: "The device type.",
			Type:        types.StringType,
			Optional:    true,
		},
		"vendor": {
			// Property: Vendor
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The device vendor.",
			     "type": "string"
			   }
			*/
			Description: "The device vendor.",
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
		Description: "The AWS::NetworkManager::Device type describes a device.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::Device").WithTerraformTypeName("aws_networkmanager_device").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_networkmanager_device", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

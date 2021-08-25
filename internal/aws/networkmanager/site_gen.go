// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkmanager

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_networkmanager_site", siteResourceType)
}

// siteResourceType returns the Terraform awscc_networkmanager_site resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NetworkManager::Site resource type.
func siteResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the site.",
			//   "type": "string"
			// }
			Description: "The description of the site.",
			Type:        types.StringType,
			Optional:    true,
		},
		"global_network_id": {
			// Property: GlobalNetworkId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the global network.",
			//   "type": "string"
			// }
			Description: "The ID of the global network.",
			Type:        types.StringType,
			Required:    true,
			// GlobalNetworkId is a force-new attribute.
		},
		"location": {
			// Property: Location
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The location of the site",
			//   "properties": {
			//     "Address": {
			//       "description": "The physical address.",
			//       "type": "string"
			//     },
			//     "Latitude": {
			//       "description": "The latitude.",
			//       "type": "string"
			//     },
			//     "Longitude": {
			//       "description": "The longitude.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The location of the site",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"address": {
						// Property: Address
						Description: "The physical address.",
						Type:        types.StringType,
						Optional:    true,
					},
					"latitude": {
						// Property: Latitude
						Description: "The latitude.",
						Type:        types.StringType,
						Optional:    true,
					},
					"longitude": {
						// Property: Longitude
						Description: "The longitude.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"site_arn": {
			// Property: SiteArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the site.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the site.",
			Type:        types.StringType,
			Computed:    true,
		},
		"site_id": {
			// Property: SiteId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the site.",
			//   "type": "string"
			// }
			Description: "The ID of the site.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags for the site.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a site resource.",
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The tags for the site.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The AWS::NetworkManager::Site type describes a site.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::Site").WithTerraformTypeName("awscc_networkmanager_site")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address":           "Address",
		"description":       "Description",
		"global_network_id": "GlobalNetworkId",
		"key":               "Key",
		"latitude":          "Latitude",
		"location":          "Location",
		"longitude":         "Longitude",
		"site_arn":          "SiteArn",
		"site_id":           "SiteId",
		"tags":              "Tags",
		"value":             "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_networkmanager_site", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

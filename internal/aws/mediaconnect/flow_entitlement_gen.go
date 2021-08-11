// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediaconnect

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
	registry.AddResourceTypeFactory("aws_mediaconnect_flow_entitlement", flowEntitlementResourceType)
}

// flowEntitlementResourceType returns the Terraform aws_mediaconnect_flow_entitlement resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::MediaConnect::FlowEntitlement resource type.
func flowEntitlementResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"data_transfer_subscriber_fee_percent": {
			// Property: DataTransferSubscriberFeePercent
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Percentage from 0-100 of the data transfer cost to be billed to the subscriber.",
			     "type": "integer"
			   }
			*/
			Description: "Percentage from 0-100 of the data transfer cost to be billed to the subscriber.",
			Type:        types.NumberType,
			Optional:    true,
			Computed:    true,
			// DataTransferSubscriberFeePercent is a force-new attribute.
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A description of the entitlement.",
			     "type": "string"
			   }
			*/
			Description: "A description of the entitlement.",
			Type:        types.StringType,
			Required:    true,
		},
		"encryption": {
			// Property: Encryption
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "Information about the encryption of the flow.",
			     "properties": {
			       "Algorithm": {
			         "description": "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
			         "enum": [
			           "aes128",
			           "aes192",
			           "aes256"
			         ],
			         "type": "string"
			       },
			       "ConstantInitializationVector": {
			         "description": "A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.",
			         "type": "string"
			       },
			       "DeviceId": {
			         "description": "The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			         "type": "string"
			       },
			       "KeyType": {
			         "description": "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
			         "enum": [
			           "speke",
			           "static-key"
			         ],
			         "type": "string"
			       },
			       "Region": {
			         "description": "The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			         "type": "string"
			       },
			       "ResourceId": {
			         "description": "An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			         "type": "string"
			       },
			       "RoleArn": {
			         "description": "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
			         "type": "string"
			       },
			       "SecretArn": {
			         "description": " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
			         "type": "string"
			       },
			       "Url": {
			         "description": "The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			         "type": "string"
			       }
			     },
			     "required": [
			       "Algorithm",
			       "RoleArn"
			     ],
			     "type": "object"
			   }
			*/
			Description: "Information about the encryption of the flow.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"algorithm": {
						// Property: Algorithm
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
						     "enum": [
						       "aes128",
						       "aes192",
						       "aes256"
						     ],
						     "type": "string"
						   }
						*/
						Description: "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
						Type:        types.StringType,
						Required:    true,
					},
					"constant_initialization_vector": {
						// Property: ConstantInitializationVector
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.",
						     "type": "string"
						   }
						*/
						Description: "A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.",
						Type:        types.StringType,
						Optional:    true,
					},
					"device_id": {
						// Property: DeviceId
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
						     "type": "string"
						   }
						*/
						Description: "The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
						Type:        types.StringType,
						Optional:    true,
					},
					"key_type": {
						// Property: KeyType
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
						     "enum": [
						       "speke",
						       "static-key"
						     ],
						     "type": "string"
						   }
						*/
						Description: "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
						Type:        types.StringType,
						Optional:    true,
					},
					"region": {
						// Property: Region
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
						     "type": "string"
						   }
						*/
						Description: "The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
						Type:        types.StringType,
						Optional:    true,
					},
					"resource_id": {
						// Property: ResourceId
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
						     "type": "string"
						   }
						*/
						Description: "An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
						Type:        types.StringType,
						Optional:    true,
					},
					"role_arn": {
						// Property: RoleArn
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
						     "type": "string"
						   }
						*/
						Description: "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
						Type:        types.StringType,
						Required:    true,
					},
					"secret_arn": {
						// Property: SecretArn
						// CloudFormation resource type schema:
						/*
						   {
						     "description": " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
						     "type": "string"
						   }
						*/
						Description: " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
						Type:        types.StringType,
						Optional:    true,
					},
					"url": {
						// Property: Url
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
						     "type": "string"
						   }
						*/
						Description: "The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"entitlement_arn": {
			// Property: EntitlementArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ARN of the entitlement.",
			     "type": "string"
			   }
			*/
			Description: "The ARN of the entitlement.",
			Type:        types.StringType,
			Computed:    true,
		},
		"entitlement_status": {
			// Property: EntitlementStatus
			// CloudFormation resource type schema:
			/*
			   {
			     "description": " An indication of whether the entitlement is enabled.",
			     "enum": [
			       "ENABLED",
			       "DISABLED"
			     ],
			     "type": "string"
			   }
			*/
			Description: " An indication of whether the entitlement is enabled.",
			Type:        types.StringType,
			Optional:    true,
		},
		"flow_arn": {
			// Property: FlowArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ARN of the flow.",
			     "type": "string"
			   }
			*/
			Description: "The ARN of the flow.",
			Type:        types.StringType,
			Required:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the entitlement.",
			     "type": "string"
			   }
			*/
			Description: "The name of the entitlement.",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"subscribers": {
			// Property: Subscribers
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The AWS account IDs that you want to share your content with. The receiving accounts (subscribers) will be allowed to create their own flow using your content as the source.",
			     "items": {
			       "type": "string"
			     },
			     "type": "array"
			   }
			*/
			Description: "The AWS account IDs that you want to share your content with. The receiving accounts (subscribers) will be allowed to create their own flow using your content as the source.",
			Type:        types.ListType{ElemType: types.StringType},
			Required:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::MediaConnect::FlowEntitlement",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaConnect::FlowEntitlement").WithTerraformTypeName("aws_mediaconnect_flow_entitlement").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_mediaconnect_flow_entitlement", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

// Code generated by generators/resource/main.go; DO NOT EDIT.

package fms

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
	registry.AddResourceTypeFactory("aws_fms_policy", policyResourceType)
}

// policyResourceType returns the Terraform aws_fms_policy resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::FMS::Policy resource type.
func policyResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A resource ARN.",
			     "maxLength": 1024,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "A resource ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"delete_all_policy_resources": {
			// Property: DeleteAllPolicyResources
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "boolean"
			   }
			*/
			Type:     types.BoolType,
			Optional: true,
			// DeleteAllPolicyResources is a write-only attribute.
		},
		"exclude_map": {
			// Property: ExcludeMap
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "An FMS includeMap or excludeMap.",
			     "properties": {
			       "ACCOUNT": {
			         "items": {
			           "description": "An AWS account ID.",
			           "maxLength": 12,
			           "minLength": 12,
			           "pattern": "",
			           "type": "string"
			         },
			         "type": "array"
			       },
			       "ORGUNIT": {
			         "items": {
			           "description": "An Organizational Unit ID.",
			           "maxLength": 68,
			           "minLength": 16,
			           "pattern": "",
			           "type": "string"
			         },
			         "type": "array"
			       }
			     },
			     "type": "object"
			   }
			*/
			Description: "An FMS includeMap or excludeMap.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"account": {
						// Property: ACCOUNT
						// CloudFormation resource type schema:
						/*
						   {
						     "items": {
						       "description": "An AWS account ID.",
						       "maxLength": 12,
						       "minLength": 12,
						       "pattern": "",
						       "type": "string"
						     },
						     "type": "array"
						   }
						*/
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
					"orgunit": {
						// Property: ORGUNIT
						// CloudFormation resource type schema:
						/*
						   {
						     "items": {
						       "description": "An Organizational Unit ID.",
						       "maxLength": 68,
						       "minLength": 16,
						       "pattern": "",
						       "type": "string"
						     },
						     "type": "array"
						   }
						*/
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"exclude_resource_tags": {
			// Property: ExcludeResourceTags
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "boolean"
			   }
			*/
			Type:     types.BoolType,
			Required: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 36,
			     "minLength": 36,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"include_map": {
			// Property: IncludeMap
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "An FMS includeMap or excludeMap.",
			     "properties": {
			       "ACCOUNT": {
			         "items": {
			           "description": "An AWS account ID.",
			           "maxLength": 12,
			           "minLength": 12,
			           "pattern": "",
			           "type": "string"
			         },
			         "type": "array"
			       },
			       "ORGUNIT": {
			         "items": {
			           "description": "An Organizational Unit ID.",
			           "maxLength": 68,
			           "minLength": 16,
			           "pattern": "",
			           "type": "string"
			         },
			         "type": "array"
			       }
			     },
			     "type": "object"
			   }
			*/
			Description: "An FMS includeMap or excludeMap.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"account": {
						// Property: ACCOUNT
						// CloudFormation resource type schema:
						/*
						   {
						     "items": {
						       "description": "An AWS account ID.",
						       "maxLength": 12,
						       "minLength": 12,
						       "pattern": "",
						       "type": "string"
						     },
						     "type": "array"
						   }
						*/
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
					"orgunit": {
						// Property: ORGUNIT
						// CloudFormation resource type schema:
						/*
						   {
						     "items": {
						       "description": "An Organizational Unit ID.",
						       "maxLength": 68,
						       "minLength": 16,
						       "pattern": "",
						       "type": "string"
						     },
						     "type": "array"
						   }
						*/
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"policy_name": {
			// Property: PolicyName
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 1024,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
		},
		"remediation_enabled": {
			// Property: RemediationEnabled
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "boolean"
			   }
			*/
			Type:     types.BoolType,
			Required: true,
		},
		"resource_tags": {
			// Property: ResourceTags
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "description": "A resource tag.",
			       "properties": {
			         "Key": {
			           "maxLength": 128,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Value": {
			           "maxLength": 256,
			           "type": "string"
			         }
			       },
			       "required": [
			         "Key"
			       ],
			       "type": "object"
			     },
			     "maxLength": 8,
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
						     "maxLength": 256,
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
		"resource_type": {
			// Property: ResourceType
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "An AWS resource type",
			     "maxLength": 128,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "An AWS resource type",
			Type:        types.StringType,
			Required:    true,
		},
		"resource_type_list": {
			// Property: ResourceTypeList
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "description": "An AWS resource type",
			       "maxLength": 128,
			       "minLength": 1,
			       "pattern": "",
			       "type": "string"
			     },
			     "type": "array"
			   }
			*/
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
		},
		"security_service_policy_data": {
			// Property: SecurityServicePolicyData
			// CloudFormation resource type schema:
			/*
			   {
			     "properties": {
			       "ManagedServiceData": {
			         "maxLength": 4096,
			         "minLength": 1,
			         "type": "string"
			       },
			       "Type": {
			         "enum": [
			           "WAF",
			           "WAFV2",
			           "SHIELD_ADVANCED",
			           "SECURITY_GROUPS_COMMON",
			           "SECURITY_GROUPS_CONTENT_AUDIT",
			           "SECURITY_GROUPS_USAGE_AUDIT",
			           "NETWORK_FIREWALL",
			           "DNS_FIREWALL"
			         ],
			         "type": "string"
			       }
			     },
			     "required": [
			       "Type"
			     ],
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"managed_service_data": {
						// Property: ManagedServiceData
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 4096,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"type": {
						// Property: Type
						// CloudFormation resource type schema:
						/*
						   {
						     "enum": [
						       "WAF",
						       "WAFV2",
						       "SHIELD_ADVANCED",
						       "SECURITY_GROUPS_COMMON",
						       "SECURITY_GROUPS_CONTENT_AUDIT",
						       "SECURITY_GROUPS_USAGE_AUDIT",
						       "NETWORK_FIREWALL",
						       "DNS_FIREWALL"
						     ],
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Required: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "description": "A policy tag.",
			       "properties": {
			         "Key": {
			           "maxLength": 128,
			           "minLength": 1,
			           "pattern": "",
			           "type": "string"
			         },
			         "Value": {
			           "maxLength": 256,
			           "pattern": "",
			           "type": "string"
			         }
			       },
			       "required": [
			         "Key",
			         "Value"
			       ],
			       "type": "object"
			     },
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
						     "maxLength": 256,
						     "pattern": "",
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
		Description: "Creates an AWS Firewall Manager policy.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::FMS::Policy").WithTerraformTypeName("aws_fms_policy").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/DeleteAllPolicyResources",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_fms_policy", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

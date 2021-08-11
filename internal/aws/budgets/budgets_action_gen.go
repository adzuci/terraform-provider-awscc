// Code generated by generators/resource/main.go; DO NOT EDIT.

package budgets

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
	registry.AddResourceTypeFactory("aws_budgets_budgets_action", budgetsActionResourceType)
}

// budgetsActionResourceType returns the Terraform aws_budgets_budgets_action resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Budgets::BudgetsAction resource type.
func budgetsActionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"action_id": {
			// Property: ActionId
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"action_threshold": {
			// Property: ActionThreshold
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "Type": {
			         "enum": [
			           "PERCENTAGE",
			           "ABSOLUTE_VALUE"
			         ],
			         "type": "string"
			       },
			       "Value": {
			         "type": "number"
			       }
			     },
			     "required": [
			       "Value",
			       "Type"
			     ],
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"type": {
						// Property: Type
						// CloudFormation resource type schema:
						/*
						   {
						     "enum": [
						       "PERCENTAGE",
						       "ABSOLUTE_VALUE"
						     ],
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
						     "type": "number"
						   }
						*/
						Type:     types.NumberType,
						Required: true,
					},
				},
			),
			Required: true,
		},
		"action_type": {
			// Property: ActionType
			// CloudFormation resource type schema:
			/*
			   {
			     "enum": [
			       "APPLY_IAM_POLICY",
			       "APPLY_SCP_POLICY",
			       "RUN_SSM_DOCUMENTS"
			     ],
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
			// ActionType is a force-new attribute.
		},
		"approval_model": {
			// Property: ApprovalModel
			// CloudFormation resource type schema:
			/*
			   {
			     "enum": [
			       "AUTOMATIC",
			       "MANUAL"
			     ],
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Optional: true,
		},
		"budget_name": {
			// Property: BudgetName
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
			// BudgetName is a force-new attribute.
		},
		"definition": {
			// Property: Definition
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "IamActionDefinition": {
			         "additionalProperties": false,
			         "properties": {
			           "Groups": {
			             "items": {
			               "type": "string"
			             },
			             "maxItems": 100,
			             "minItems": 1,
			             "type": "array"
			           },
			           "PolicyArn": {
			             "type": "string"
			           },
			           "Roles": {
			             "items": {
			               "type": "string"
			             },
			             "maxItems": 100,
			             "minItems": 1,
			             "type": "array"
			           },
			           "Users": {
			             "items": {
			               "type": "string"
			             },
			             "maxItems": 100,
			             "minItems": 1,
			             "type": "array"
			           }
			         },
			         "required": [
			           "PolicyArn"
			         ],
			         "type": "object"
			       },
			       "ScpActionDefinition": {
			         "additionalProperties": false,
			         "properties": {
			           "PolicyId": {
			             "type": "string"
			           },
			           "TargetIds": {
			             "items": {
			               "type": "string"
			             },
			             "maxItems": 100,
			             "minItems": 1,
			             "type": "array"
			           }
			         },
			         "required": [
			           "PolicyId",
			           "TargetIds"
			         ],
			         "type": "object"
			       },
			       "SsmActionDefinition": {
			         "additionalProperties": false,
			         "properties": {
			           "InstanceIds": {
			             "items": {
			               "type": "string"
			             },
			             "maxItems": 100,
			             "minItems": 1,
			             "type": "array"
			           },
			           "Region": {
			             "type": "string"
			           },
			           "Subtype": {
			             "enum": [
			               "STOP_EC2_INSTANCES",
			               "STOP_RDS_INSTANCES"
			             ],
			             "type": "string"
			           }
			         },
			         "required": [
			           "Subtype",
			           "Region",
			           "InstanceIds"
			         ],
			         "type": "object"
			       }
			     },
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"iam_action_definition": {
						// Property: IamActionDefinition
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "Groups": {
						         "items": {
						           "type": "string"
						         },
						         "maxItems": 100,
						         "minItems": 1,
						         "type": "array"
						       },
						       "PolicyArn": {
						         "type": "string"
						       },
						       "Roles": {
						         "items": {
						           "type": "string"
						         },
						         "maxItems": 100,
						         "minItems": 1,
						         "type": "array"
						       },
						       "Users": {
						         "items": {
						           "type": "string"
						         },
						         "maxItems": 100,
						         "minItems": 1,
						         "type": "array"
						       }
						     },
						     "required": [
						       "PolicyArn"
						     ],
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"groups": {
									// Property: Groups
									// CloudFormation resource type schema:
									/*
									   {
									     "items": {
									       "type": "string"
									     },
									     "maxItems": 100,
									     "minItems": 1,
									     "type": "array"
									   }
									*/
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
								"policy_arn": {
									// Property: PolicyArn
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"roles": {
									// Property: Roles
									// CloudFormation resource type schema:
									/*
									   {
									     "items": {
									       "type": "string"
									     },
									     "maxItems": 100,
									     "minItems": 1,
									     "type": "array"
									   }
									*/
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
								"users": {
									// Property: Users
									// CloudFormation resource type schema:
									/*
									   {
									     "items": {
									       "type": "string"
									     },
									     "maxItems": 100,
									     "minItems": 1,
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
					"scp_action_definition": {
						// Property: ScpActionDefinition
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "PolicyId": {
						         "type": "string"
						       },
						       "TargetIds": {
						         "items": {
						           "type": "string"
						         },
						         "maxItems": 100,
						         "minItems": 1,
						         "type": "array"
						       }
						     },
						     "required": [
						       "PolicyId",
						       "TargetIds"
						     ],
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"policy_id": {
									// Property: PolicyId
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"target_ids": {
									// Property: TargetIds
									// CloudFormation resource type schema:
									/*
									   {
									     "items": {
									       "type": "string"
									     },
									     "maxItems": 100,
									     "minItems": 1,
									     "type": "array"
									   }
									*/
									Type:     types.ListType{ElemType: types.StringType},
									Required: true,
								},
							},
						),
						Optional: true,
					},
					"ssm_action_definition": {
						// Property: SsmActionDefinition
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "InstanceIds": {
						         "items": {
						           "type": "string"
						         },
						         "maxItems": 100,
						         "minItems": 1,
						         "type": "array"
						       },
						       "Region": {
						         "type": "string"
						       },
						       "Subtype": {
						         "enum": [
						           "STOP_EC2_INSTANCES",
						           "STOP_RDS_INSTANCES"
						         ],
						         "type": "string"
						       }
						     },
						     "required": [
						       "Subtype",
						       "Region",
						       "InstanceIds"
						     ],
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"instance_ids": {
									// Property: InstanceIds
									// CloudFormation resource type schema:
									/*
									   {
									     "items": {
									       "type": "string"
									     },
									     "maxItems": 100,
									     "minItems": 1,
									     "type": "array"
									   }
									*/
									Type:     types.ListType{ElemType: types.StringType},
									Required: true,
								},
								"region": {
									// Property: Region
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"subtype": {
									// Property: Subtype
									// CloudFormation resource type schema:
									/*
									   {
									     "enum": [
									       "STOP_EC2_INSTANCES",
									       "STOP_RDS_INSTANCES"
									     ],
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Required: true,
		},
		"execution_role_arn": {
			// Property: ExecutionRoleArn
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
		},
		"notification_type": {
			// Property: NotificationType
			// CloudFormation resource type schema:
			/*
			   {
			     "enum": [
			       "ACTUAL",
			       "FORECASTED"
			     ],
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
		},
		"subscribers": {
			// Property: Subscribers
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Address": {
			           "type": "string"
			         },
			         "Type": {
			           "enum": [
			             "SNS",
			             "EMAIL"
			           ],
			           "type": "string"
			         }
			       },
			       "required": [
			         "Type",
			         "Address"
			       ],
			       "type": "object"
			     },
			     "maxItems": 11,
			     "minItems": 1,
			     "type": "array"
			   }
			*/
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"address": {
						// Property: Address
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
					"type": {
						// Property: Type
						// CloudFormation resource type schema:
						/*
						   {
						     "enum": [
						       "SNS",
						       "EMAIL"
						     ],
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MinItems: 1,
					MaxItems: 11,
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
		Description: "An example resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Budgets::BudgetsAction").WithTerraformTypeName("aws_budgets_budgets_action").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_budgets_budgets_action", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
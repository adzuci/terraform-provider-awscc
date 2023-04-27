// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package grafana

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_grafana_workspace", workspaceResource)
}

// workspaceResource returns the Terraform awscc_grafana_workspace resource.
// This Terraform resource corresponds to the CloudFormation AWS::Grafana::Workspace resource.
func workspaceResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccountAccessType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "These enums represent valid account access types. Specifically these enums determine whether the workspace can access AWS resources in the AWS account only, or whether it can also access resources in other accounts in the same organization. If the value CURRENT_ACCOUNT is used, a workspace role ARN must be provided. If the value is ORGANIZATION, a list of organizational units must be provided.",
		//	  "enum": [
		//	    "CURRENT_ACCOUNT",
		//	    "ORGANIZATION"
		//	  ],
		//	  "type": "string"
		//	}
		"account_access_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "These enums represent valid account access types. Specifically these enums determine whether the workspace can access AWS resources in the AWS account only, or whether it can also access resources in other accounts in the same organization. If the value CURRENT_ACCOUNT is used, a workspace role ARN must be provided. If the value is ORGANIZATION, a list of organizational units must be provided.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"CURRENT_ACCOUNT",
					"ORGANIZATION",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: AuthenticationProviders
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of authentication providers to enable.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "Valid workspace authentication providers.",
		//	    "enum": [
		//	      "AWS_SSO",
		//	      "SAML"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"authentication_providers": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "List of authentication providers to enable.",
			Required:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtLeast(1),
				setvalidator.ValueStringsAre(
					stringvalidator.OneOf(
						"AWS_SSO",
						"SAML",
					),
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: ClientToken
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique, case-sensitive, user-provided identifier to ensure the idempotency of the request.",
		//	  "pattern": "^[!-~]{1,64}$",
		//	  "type": "string"
		//	}
		"client_token": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique, case-sensitive, user-provided identifier to ensure the idempotency of the request.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[!-~]{1,64}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// ClientToken is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: CreationTimestamp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Timestamp when the workspace was created.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"creation_timestamp": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Timestamp when the workspace was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DataSources
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of data sources on the service managed IAM role.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "These enums represent valid AWS data sources that can be queried via the Grafana workspace. These data sources are primarily used to help customers visualize which data sources have been added to a service managed workspace IAM role.",
		//	    "enum": [
		//	      "AMAZON_OPENSEARCH_SERVICE",
		//	      "CLOUDWATCH",
		//	      "PROMETHEUS",
		//	      "XRAY",
		//	      "TIMESTREAM",
		//	      "SITEWISE",
		//	      "ATHENA",
		//	      "REDSHIFT"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"data_sources": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "List of data sources on the service managed IAM role.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.OneOf(
						"AMAZON_OPENSEARCH_SERVICE",
						"CLOUDWATCH",
						"PROMETHEUS",
						"XRAY",
						"TIMESTREAM",
						"SITEWISE",
						"ATHENA",
						"REDSHIFT",
					),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description of a workspace.",
		//	  "maxLength": 2048,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description of a workspace.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 2048),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Endpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Endpoint for the Grafana workspace.",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Endpoint for the Grafana workspace.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GrafanaVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Version of Grafana the workspace is currently using.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"grafana_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Version of Grafana the workspace is currently using.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The id that uniquely identifies a Grafana workspace.",
		//	  "pattern": "^g-[0-9a-f]{10}$",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The id that uniquely identifies a Grafana workspace.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ModificationTimestamp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Timestamp when the workspace was last modified",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"modification_timestamp": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Timestamp when the workspace was last modified",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The user friendly name of a workspace.",
		//	  "pattern": "^[a-zA-Z0-9-._~]{1,255}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The user friendly name of a workspace.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9-._~]{1,255}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NetworkAccessControl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration settings for Network Access Control.",
		//	  "properties": {
		//	    "PrefixListIds": {
		//	      "description": "The list of prefix list IDs. A prefix list is a list of CIDR ranges of IP addresses. The IP addresses specified are allowed to access your workspace. If the list is not included in the configuration then no IP addresses will be allowed to access the workspace.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "Prefix List Ids",
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "maxItems": 5,
		//	      "minItems": 0,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "VpceIds": {
		//	      "description": "The list of Amazon VPC endpoint IDs for the workspace. If a NetworkAccessConfiguration is specified then only VPC endpoints specified here will be allowed to access the workspace.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "VPCE Ids",
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "maxItems": 5,
		//	      "minItems": 0,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"network_access_control": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: PrefixListIds
				"prefix_list_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The list of prefix list IDs. A prefix list is a list of CIDR ranges of IP addresses. The IP addresses specified are allowed to access your workspace. If the list is not included in the configuration then no IP addresses will be allowed to access the workspace.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeBetween(0, 5),
						setvalidator.ValueStringsAre(
							stringvalidator.LengthAtLeast(1),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: VpceIds
				"vpce_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The list of Amazon VPC endpoint IDs for the workspace. If a NetworkAccessConfiguration is specified then only VPC endpoints specified here will be allowed to access the workspace.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeBetween(0, 5),
						setvalidator.ValueStringsAre(
							stringvalidator.LengthAtLeast(1),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration settings for Network Access Control.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NotificationDestinations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of notification destinations on the customers service managed IAM role that the Grafana workspace can query.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "These enums represent valid AWS notification destinations that the Grafana workspace has permission to use. These notification destinations are primarily used to help customers visualize which destinations have been added to a service managed IAM role.",
		//	    "enum": [
		//	      "SNS"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"notification_destinations": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "List of notification destinations on the customers service managed IAM role that the Grafana workspace can query.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.OneOf(
						"SNS",
					),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OrganizationRoleName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of an IAM role that already exists to use with AWS Organizations to access AWS data sources and notification channels in other accounts in an organization.",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"organization_role_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of an IAM role that already exists to use with AWS Organizations to access AWS data sources and notification channels in other accounts in an organization.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 2048),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OrganizationalUnits
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of Organizational Units containing AWS accounts the Grafana workspace can pull data from.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "Id of an organizational unit.",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"organizational_units": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "List of Organizational Units containing AWS accounts the Grafana workspace can pull data from.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PermissionType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "These enums represent valid permission types to use when creating or configuring a Grafana workspace. The SERVICE_MANAGED permission type means the Managed Grafana service will create a workspace IAM role on your behalf. The CUSTOMER_MANAGED permission type means that the customer is expected to provide an IAM role that the Grafana workspace can use to query data sources.",
		//	  "enum": [
		//	    "CUSTOMER_MANAGED",
		//	    "SERVICE_MANAGED"
		//	  ],
		//	  "type": "string"
		//	}
		"permission_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "These enums represent valid permission types to use when creating or configuring a Grafana workspace. The SERVICE_MANAGED permission type means the Managed Grafana service will create a workspace IAM role on your behalf. The CUSTOMER_MANAGED permission type means that the customer is expected to provide an IAM role that the Grafana workspace can use to query data sources.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"CUSTOMER_MANAGED",
					"SERVICE_MANAGED",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "IAM Role that will be used to grant the Grafana workspace access to a customers AWS resources.",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "IAM Role that will be used to grant the Grafana workspace access to a customers AWS resources.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 2048),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SamlConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "SAML configuration data associated with an AMG workspace.",
		//	  "properties": {
		//	    "AllowedOrganizations": {
		//	      "description": "List of SAML organizations allowed to access Grafana.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "A single SAML organization.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "AssertionAttributes": {
		//	      "additionalProperties": false,
		//	      "description": "Maps Grafana friendly names to the IdPs SAML attributes.",
		//	      "properties": {
		//	        "Email": {
		//	          "description": "Name of the attribute within the SAML assert to use as the users email in Grafana.",
		//	          "maxLength": 256,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Groups": {
		//	          "description": "Name of the attribute within the SAML assert to use as the users groups in Grafana.",
		//	          "maxLength": 256,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Login": {
		//	          "description": "Name of the attribute within the SAML assert to use as the users login handle in Grafana.",
		//	          "maxLength": 256,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Name": {
		//	          "description": "Name of the attribute within the SAML assert to use as the users name in Grafana.",
		//	          "maxLength": 256,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Org": {
		//	          "description": "Name of the attribute within the SAML assert to use as the users organizations in Grafana.",
		//	          "maxLength": 256,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Role": {
		//	          "description": "Name of the attribute within the SAML assert to use as the users roles in Grafana.",
		//	          "maxLength": 256,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "IdpMetadata": {
		//	      "additionalProperties": false,
		//	      "description": "IdP Metadata used to configure SAML authentication in Grafana.",
		//	      "properties": {
		//	        "Url": {
		//	          "description": "URL that vends the IdPs metadata.",
		//	          "maxLength": 2048,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Xml": {
		//	          "description": "XML blob of the IdPs metadata.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "LoginValidityDuration": {
		//	      "description": "The maximum lifetime an authenticated user can be logged in (in minutes) before being required to re-authenticate.",
		//	      "type": "number"
		//	    },
		//	    "RoleValues": {
		//	      "additionalProperties": false,
		//	      "description": "Maps SAML roles to the Grafana Editor and Admin roles.",
		//	      "properties": {
		//	        "Admin": {
		//	          "description": "List of SAML roles which will be mapped into the Grafana Admin role.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "description": "A single SAML role.",
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "Editor": {
		//	          "description": "List of SAML roles which will be mapped into the Grafana Editor role.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "description": "A single SAML role.",
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "type": "array"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "IdpMetadata"
		//	  ],
		//	  "type": "object"
		//	}
		"saml_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AllowedOrganizations
				"allowed_organizations": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "List of SAML organizations allowed to access Grafana.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(1, 256),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: AssertionAttributes
				"assertion_attributes": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Email
						"email": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Name of the attribute within the SAML assert to use as the users email in Grafana.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 256),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Groups
						"groups": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Name of the attribute within the SAML assert to use as the users groups in Grafana.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 256),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Login
						"login": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Name of the attribute within the SAML assert to use as the users login handle in Grafana.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 256),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Name
						"name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Name of the attribute within the SAML assert to use as the users name in Grafana.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 256),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Org
						"org": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Name of the attribute within the SAML assert to use as the users organizations in Grafana.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 256),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Role
						"role": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Name of the attribute within the SAML assert to use as the users roles in Grafana.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 256),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Maps Grafana friendly names to the IdPs SAML attributes.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: IdpMetadata
				"idp_metadata": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Url
						"url": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "URL that vends the IdPs metadata.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 2048),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Xml
						"xml": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "XML blob of the IdPs metadata.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "IdP Metadata used to configure SAML authentication in Grafana.",
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: LoginValidityDuration
				"login_validity_duration": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The maximum lifetime an authenticated user can be logged in (in minutes) before being required to re-authenticate.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
						float64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RoleValues
				"role_values": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Admin
						"admin": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "List of SAML roles which will be mapped into the Grafana Admin role.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(1, 256),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								generic.Multiset(),
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Editor
						"editor": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "List of SAML roles which will be mapped into the Grafana Editor role.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(1, 256),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								generic.Multiset(),
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Maps SAML roles to the Grafana Editor and Admin roles.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "SAML configuration data associated with an AMG workspace.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SamlConfigurationStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Valid SAML configuration statuses.",
		//	  "enum": [
		//	    "CONFIGURED",
		//	    "NOT_CONFIGURED"
		//	  ],
		//	  "type": "string"
		//	}
		"saml_configuration_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Valid SAML configuration statuses.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SsoClientId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The client ID of the AWS SSO Managed Application.",
		//	  "type": "string"
		//	}
		"sso_client_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The client ID of the AWS SSO Managed Application.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StackSetName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the AWS CloudFormation stack set to use to generate IAM roles to be used for this workspace.",
		//	  "type": "string"
		//	}
		"stack_set_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the AWS CloudFormation stack set to use to generate IAM roles to be used for this workspace.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "These enums represent the status of a workspace.",
		//	  "enum": [
		//	    "ACTIVE",
		//	    "CREATING",
		//	    "DELETING",
		//	    "FAILED",
		//	    "UPDATING",
		//	    "UPGRADING",
		//	    "DELETION_FAILED",
		//	    "CREATION_FAILED",
		//	    "UPDATE_FAILED",
		//	    "UPGRADE_FAILED",
		//	    "LICENSE_REMOVAL_FAILED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "These enums represent the status of a workspace.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to.",
		//	  "properties": {
		//	    "SecurityGroupIds": {
		//	      "description": "The list of Amazon EC2 security group IDs attached to the Amazon VPC for your Grafana workspace to connect.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "VPC Security Group Id",
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "maxItems": 5,
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "SubnetIds": {
		//	      "description": "The list of Amazon EC2 subnet IDs created in the Amazon VPC for your Grafana workspace to connect.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "VPC Subnet Id",
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "maxItems": 6,
		//	      "minItems": 2,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "required": [
		//	    "SecurityGroupIds",
		//	    "SubnetIds"
		//	  ],
		//	  "type": "object"
		//	}
		"vpc_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SecurityGroupIds
				"security_group_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The list of Amazon EC2 security group IDs attached to the Amazon VPC for your Grafana workspace to connect.",
					Required:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeBetween(1, 5),
						setvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(1, 255),
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: SubnetIds
				"subnet_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The list of Amazon EC2 subnet IDs created in the Amazon VPC for your Grafana workspace to connect.",
					Required:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeBetween(2, 6),
						setvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(1, 255),
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "Definition of AWS::Grafana::Workspace Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Grafana::Workspace").WithTerraformTypeName("awscc_grafana_workspace")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_access_type":       "AccountAccessType",
		"admin":                     "Admin",
		"allowed_organizations":     "AllowedOrganizations",
		"assertion_attributes":      "AssertionAttributes",
		"authentication_providers":  "AuthenticationProviders",
		"client_token":              "ClientToken",
		"creation_timestamp":        "CreationTimestamp",
		"data_sources":              "DataSources",
		"description":               "Description",
		"editor":                    "Editor",
		"email":                     "Email",
		"endpoint":                  "Endpoint",
		"grafana_version":           "GrafanaVersion",
		"groups":                    "Groups",
		"id":                        "Id",
		"idp_metadata":              "IdpMetadata",
		"login":                     "Login",
		"login_validity_duration":   "LoginValidityDuration",
		"modification_timestamp":    "ModificationTimestamp",
		"name":                      "Name",
		"network_access_control":    "NetworkAccessControl",
		"notification_destinations": "NotificationDestinations",
		"org":                       "Org",
		"organization_role_name":    "OrganizationRoleName",
		"organizational_units":      "OrganizationalUnits",
		"permission_type":           "PermissionType",
		"prefix_list_ids":           "PrefixListIds",
		"role":                      "Role",
		"role_arn":                  "RoleArn",
		"role_values":               "RoleValues",
		"saml_configuration":        "SamlConfiguration",
		"saml_configuration_status": "SamlConfigurationStatus",
		"security_group_ids":        "SecurityGroupIds",
		"sso_client_id":             "SsoClientId",
		"stack_set_name":            "StackSetName",
		"status":                    "Status",
		"subnet_ids":                "SubnetIds",
		"url":                       "Url",
		"vpc_configuration":         "VpcConfiguration",
		"vpce_ids":                  "VpceIds",
		"xml":                       "Xml",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ClientToken",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

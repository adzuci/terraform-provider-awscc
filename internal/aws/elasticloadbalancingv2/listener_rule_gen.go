// Code generated by generators/resource/main.go; DO NOT EDIT.

package elasticloadbalancingv2

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
	registry.AddResourceTypeFactory("aws_elasticloadbalancingv2_listener_rule", listenerRuleResourceType)
}

// listenerRuleResourceType returns the Terraform aws_elasticloadbalancingv2_listener_rule resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ElasticLoadBalancingV2::ListenerRule resource type.
func listenerRuleResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"actions": {
			// Property: Actions
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "AuthenticateCognitoConfig": {
			           "additionalProperties": false,
			           "properties": {
			             "AuthenticationRequestExtraParams": {
			               "additionalProperties": false,
			               "patternProperties": {
			                 "": {
			                   "type": "string"
			                 }
			               },
			               "type": "object"
			             },
			             "OnUnauthenticatedRequest": {
			               "type": "string"
			             },
			             "Scope": {
			               "type": "string"
			             },
			             "SessionCookieName": {
			               "type": "string"
			             },
			             "SessionTimeout": {
			               "type": "integer"
			             },
			             "UserPoolArn": {
			               "type": "string"
			             },
			             "UserPoolClientId": {
			               "type": "string"
			             },
			             "UserPoolDomain": {
			               "type": "string"
			             }
			           },
			           "required": [
			             "UserPoolClientId",
			             "UserPoolDomain",
			             "UserPoolArn"
			           ],
			           "type": "object"
			         },
			         "AuthenticateOidcConfig": {
			           "additionalProperties": false,
			           "properties": {
			             "AuthenticationRequestExtraParams": {
			               "additionalProperties": false,
			               "patternProperties": {
			                 "": {
			                   "type": "string"
			                 }
			               },
			               "type": "object"
			             },
			             "AuthorizationEndpoint": {
			               "type": "string"
			             },
			             "ClientId": {
			               "type": "string"
			             },
			             "ClientSecret": {
			               "type": "string"
			             },
			             "Issuer": {
			               "type": "string"
			             },
			             "OnUnauthenticatedRequest": {
			               "type": "string"
			             },
			             "Scope": {
			               "type": "string"
			             },
			             "SessionCookieName": {
			               "type": "string"
			             },
			             "SessionTimeout": {
			               "type": "integer"
			             },
			             "TokenEndpoint": {
			               "type": "string"
			             },
			             "UseExistingClientSecret": {
			               "type": "boolean"
			             },
			             "UserInfoEndpoint": {
			               "type": "string"
			             }
			           },
			           "required": [
			             "TokenEndpoint",
			             "Issuer",
			             "ClientSecret",
			             "UserInfoEndpoint",
			             "ClientId",
			             "AuthorizationEndpoint"
			           ],
			           "type": "object"
			         },
			         "FixedResponseConfig": {
			           "additionalProperties": false,
			           "properties": {
			             "ContentType": {
			               "type": "string"
			             },
			             "MessageBody": {
			               "type": "string"
			             },
			             "StatusCode": {
			               "type": "string"
			             }
			           },
			           "required": [
			             "StatusCode"
			           ],
			           "type": "object"
			         },
			         "ForwardConfig": {
			           "additionalProperties": false,
			           "properties": {
			             "TargetGroupStickinessConfig": {
			               "additionalProperties": false,
			               "properties": {
			                 "DurationSeconds": {
			                   "type": "integer"
			                 },
			                 "Enabled": {
			                   "type": "boolean"
			                 }
			               },
			               "type": "object"
			             },
			             "TargetGroups": {
			               "items": {
			                 "additionalProperties": false,
			                 "properties": {
			                   "TargetGroupArn": {
			                     "type": "string"
			                   },
			                   "Weight": {
			                     "type": "integer"
			                   }
			                 },
			                 "type": "object"
			               },
			               "type": "array",
			               "uniqueItems": true
			             }
			           },
			           "type": "object"
			         },
			         "Order": {
			           "type": "integer"
			         },
			         "RedirectConfig": {
			           "additionalProperties": false,
			           "properties": {
			             "Host": {
			               "type": "string"
			             },
			             "Path": {
			               "type": "string"
			             },
			             "Port": {
			               "type": "string"
			             },
			             "Protocol": {
			               "type": "string"
			             },
			             "Query": {
			               "type": "string"
			             },
			             "StatusCode": {
			               "type": "string"
			             }
			           },
			           "required": [
			             "StatusCode"
			           ],
			           "type": "object"
			         },
			         "TargetGroupArn": {
			           "type": "string"
			         },
			         "Type": {
			           "type": "string"
			         }
			       },
			       "required": [
			         "Type"
			       ],
			       "type": "object"
			     },
			     "type": "array",
			     "uniqueItems": true
			   }
			*/
			// Ordered set.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"authenticate_cognito_config": {
						// Property: AuthenticateCognitoConfig
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "AuthenticationRequestExtraParams": {
						         "additionalProperties": false,
						         "patternProperties": {
						           "": {
						             "type": "string"
						           }
						         },
						         "type": "object"
						       },
						       "OnUnauthenticatedRequest": {
						         "type": "string"
						       },
						       "Scope": {
						         "type": "string"
						       },
						       "SessionCookieName": {
						         "type": "string"
						       },
						       "SessionTimeout": {
						         "type": "integer"
						       },
						       "UserPoolArn": {
						         "type": "string"
						       },
						       "UserPoolClientId": {
						         "type": "string"
						       },
						       "UserPoolDomain": {
						         "type": "string"
						       }
						     },
						     "required": [
						       "UserPoolClientId",
						       "UserPoolDomain",
						       "UserPoolArn"
						     ],
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"authentication_request_extra_params": {
									// Property: AuthenticationRequestExtraParams
									// CloudFormation resource type schema:
									/*
									   {
									     "additionalProperties": false,
									     "patternProperties": {
									       "": {
									         "type": "string"
									       }
									     },
									     "type": "object"
									   }
									*/
									// Pattern: ""
									Type:     types.MapType{ElemType: types.StringType},
									Optional: true,
								},
								"on_unauthenticated_request": {
									// Property: OnUnauthenticatedRequest
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"scope": {
									// Property: Scope
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"session_cookie_name": {
									// Property: SessionCookieName
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"session_timeout": {
									// Property: SessionTimeout
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "integer"
									   }
									*/
									Type:     types.NumberType,
									Optional: true,
								},
								"user_pool_arn": {
									// Property: UserPoolArn
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"user_pool_client_id": {
									// Property: UserPoolClientId
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"user_pool_domain": {
									// Property: UserPoolDomain
									// CloudFormation resource type schema:
									/*
									   {
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
					"authenticate_oidc_config": {
						// Property: AuthenticateOidcConfig
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "AuthenticationRequestExtraParams": {
						         "additionalProperties": false,
						         "patternProperties": {
						           "": {
						             "type": "string"
						           }
						         },
						         "type": "object"
						       },
						       "AuthorizationEndpoint": {
						         "type": "string"
						       },
						       "ClientId": {
						         "type": "string"
						       },
						       "ClientSecret": {
						         "type": "string"
						       },
						       "Issuer": {
						         "type": "string"
						       },
						       "OnUnauthenticatedRequest": {
						         "type": "string"
						       },
						       "Scope": {
						         "type": "string"
						       },
						       "SessionCookieName": {
						         "type": "string"
						       },
						       "SessionTimeout": {
						         "type": "integer"
						       },
						       "TokenEndpoint": {
						         "type": "string"
						       },
						       "UseExistingClientSecret": {
						         "type": "boolean"
						       },
						       "UserInfoEndpoint": {
						         "type": "string"
						       }
						     },
						     "required": [
						       "TokenEndpoint",
						       "Issuer",
						       "ClientSecret",
						       "UserInfoEndpoint",
						       "ClientId",
						       "AuthorizationEndpoint"
						     ],
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"authentication_request_extra_params": {
									// Property: AuthenticationRequestExtraParams
									// CloudFormation resource type schema:
									/*
									   {
									     "additionalProperties": false,
									     "patternProperties": {
									       "": {
									         "type": "string"
									       }
									     },
									     "type": "object"
									   }
									*/
									// Pattern: ""
									Type:     types.MapType{ElemType: types.StringType},
									Optional: true,
								},
								"authorization_endpoint": {
									// Property: AuthorizationEndpoint
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"client_id": {
									// Property: ClientId
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"client_secret": {
									// Property: ClientSecret
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"issuer": {
									// Property: Issuer
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"on_unauthenticated_request": {
									// Property: OnUnauthenticatedRequest
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"scope": {
									// Property: Scope
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"session_cookie_name": {
									// Property: SessionCookieName
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"session_timeout": {
									// Property: SessionTimeout
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "integer"
									   }
									*/
									Type:     types.NumberType,
									Optional: true,
								},
								"token_endpoint": {
									// Property: TokenEndpoint
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"use_existing_client_secret": {
									// Property: UseExistingClientSecret
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "boolean"
									   }
									*/
									Type:     types.BoolType,
									Optional: true,
								},
								"user_info_endpoint": {
									// Property: UserInfoEndpoint
									// CloudFormation resource type schema:
									/*
									   {
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
					"fixed_response_config": {
						// Property: FixedResponseConfig
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "ContentType": {
						         "type": "string"
						       },
						       "MessageBody": {
						         "type": "string"
						       },
						       "StatusCode": {
						         "type": "string"
						       }
						     },
						     "required": [
						       "StatusCode"
						     ],
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"content_type": {
									// Property: ContentType
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"message_body": {
									// Property: MessageBody
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"status_code": {
									// Property: StatusCode
									// CloudFormation resource type schema:
									/*
									   {
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
					"forward_config": {
						// Property: ForwardConfig
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "TargetGroupStickinessConfig": {
						         "additionalProperties": false,
						         "properties": {
						           "DurationSeconds": {
						             "type": "integer"
						           },
						           "Enabled": {
						             "type": "boolean"
						           }
						         },
						         "type": "object"
						       },
						       "TargetGroups": {
						         "items": {
						           "additionalProperties": false,
						           "properties": {
						             "TargetGroupArn": {
						               "type": "string"
						             },
						             "Weight": {
						               "type": "integer"
						             }
						           },
						           "type": "object"
						         },
						         "type": "array",
						         "uniqueItems": true
						       }
						     },
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"target_group_stickiness_config": {
									// Property: TargetGroupStickinessConfig
									// CloudFormation resource type schema:
									/*
									   {
									     "additionalProperties": false,
									     "properties": {
									       "DurationSeconds": {
									         "type": "integer"
									       },
									       "Enabled": {
									         "type": "boolean"
									       }
									     },
									     "type": "object"
									   }
									*/
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"duration_seconds": {
												// Property: DurationSeconds
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "integer"
												   }
												*/
												Type:     types.NumberType,
												Optional: true,
											},
											"enabled": {
												// Property: Enabled
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
								"target_groups": {
									// Property: TargetGroups
									// CloudFormation resource type schema:
									/*
									   {
									     "items": {
									       "additionalProperties": false,
									       "properties": {
									         "TargetGroupArn": {
									           "type": "string"
									         },
									         "Weight": {
									           "type": "integer"
									         }
									       },
									       "type": "object"
									     },
									     "type": "array",
									     "uniqueItems": true
									   }
									*/
									// Ordered set.
									Attributes: schema.ListNestedAttributes(
										map[string]schema.Attribute{
											"target_group_arn": {
												// Property: TargetGroupArn
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "string"
												   }
												*/
												Type:     types.StringType,
												Optional: true,
											},
											"weight": {
												// Property: Weight
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "integer"
												   }
												*/
												Type:     types.NumberType,
												Optional: true,
											},
										},
										schema.ListNestedAttributesOptions{},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"order": {
						// Property: Order
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "integer"
						   }
						*/
						Type:     types.NumberType,
						Optional: true,
					},
					"redirect_config": {
						// Property: RedirectConfig
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "Host": {
						         "type": "string"
						       },
						       "Path": {
						         "type": "string"
						       },
						       "Port": {
						         "type": "string"
						       },
						       "Protocol": {
						         "type": "string"
						       },
						       "Query": {
						         "type": "string"
						       },
						       "StatusCode": {
						         "type": "string"
						       }
						     },
						     "required": [
						       "StatusCode"
						     ],
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"host": {
									// Property: Host
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"path": {
									// Property: Path
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"port": {
									// Property: Port
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"protocol": {
									// Property: Protocol
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"query": {
									// Property: Query
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"status_code": {
									// Property: StatusCode
									// CloudFormation resource type schema:
									/*
									   {
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
					"target_group_arn": {
						// Property: TargetGroupArn
						// CloudFormation resource type schema:
						/*
						   {
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
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Required: true,
		},
		"conditions": {
			// Property: Conditions
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Field": {
			           "type": "string"
			         },
			         "HostHeaderConfig": {
			           "additionalProperties": false,
			           "properties": {
			             "Values": {
			               "items": {
			                 "type": "string"
			               },
			               "type": "array",
			               "uniqueItems": true
			             }
			           },
			           "type": "object"
			         },
			         "HttpHeaderConfig": {
			           "additionalProperties": false,
			           "properties": {
			             "HttpHeaderName": {
			               "type": "string"
			             },
			             "Values": {
			               "items": {
			                 "type": "string"
			               },
			               "type": "array",
			               "uniqueItems": true
			             }
			           },
			           "type": "object"
			         },
			         "HttpRequestMethodConfig": {
			           "additionalProperties": false,
			           "properties": {
			             "Values": {
			               "items": {
			                 "type": "string"
			               },
			               "type": "array",
			               "uniqueItems": true
			             }
			           },
			           "type": "object"
			         },
			         "PathPatternConfig": {
			           "additionalProperties": false,
			           "properties": {
			             "Values": {
			               "items": {
			                 "type": "string"
			               },
			               "type": "array",
			               "uniqueItems": true
			             }
			           },
			           "type": "object"
			         },
			         "QueryStringConfig": {
			           "additionalProperties": false,
			           "properties": {
			             "Values": {
			               "items": {
			                 "additionalProperties": false,
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
			               "type": "array",
			               "uniqueItems": true
			             }
			           },
			           "type": "object"
			         },
			         "SourceIpConfig": {
			           "additionalProperties": false,
			           "properties": {
			             "Values": {
			               "items": {
			                 "type": "string"
			               },
			               "type": "array",
			               "uniqueItems": true
			             }
			           },
			           "type": "object"
			         },
			         "Values": {
			           "items": {
			             "type": "string"
			           },
			           "type": "array",
			           "uniqueItems": true
			         }
			       },
			       "type": "object"
			     },
			     "type": "array",
			     "uniqueItems": true
			   }
			*/
			// Ordered set.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"field": {
						// Property: Field
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"host_header_config": {
						// Property: HostHeaderConfig
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "Values": {
						         "items": {
						           "type": "string"
						         },
						         "type": "array",
						         "uniqueItems": true
						       }
						     },
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"values": {
									// Property: Values
									// CloudFormation resource type schema:
									/*
									   {
									     "items": {
									       "type": "string"
									     },
									     "type": "array",
									     "uniqueItems": true
									   }
									*/
									// Ordered set.
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"http_header_config": {
						// Property: HttpHeaderConfig
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "HttpHeaderName": {
						         "type": "string"
						       },
						       "Values": {
						         "items": {
						           "type": "string"
						         },
						         "type": "array",
						         "uniqueItems": true
						       }
						     },
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"http_header_name": {
									// Property: HttpHeaderName
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"values": {
									// Property: Values
									// CloudFormation resource type schema:
									/*
									   {
									     "items": {
									       "type": "string"
									     },
									     "type": "array",
									     "uniqueItems": true
									   }
									*/
									// Ordered set.
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"http_request_method_config": {
						// Property: HttpRequestMethodConfig
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "Values": {
						         "items": {
						           "type": "string"
						         },
						         "type": "array",
						         "uniqueItems": true
						       }
						     },
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"values": {
									// Property: Values
									// CloudFormation resource type schema:
									/*
									   {
									     "items": {
									       "type": "string"
									     },
									     "type": "array",
									     "uniqueItems": true
									   }
									*/
									// Ordered set.
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"path_pattern_config": {
						// Property: PathPatternConfig
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "Values": {
						         "items": {
						           "type": "string"
						         },
						         "type": "array",
						         "uniqueItems": true
						       }
						     },
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"values": {
									// Property: Values
									// CloudFormation resource type schema:
									/*
									   {
									     "items": {
									       "type": "string"
									     },
									     "type": "array",
									     "uniqueItems": true
									   }
									*/
									// Ordered set.
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"query_string_config": {
						// Property: QueryStringConfig
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "Values": {
						         "items": {
						           "additionalProperties": false,
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
						         "type": "array",
						         "uniqueItems": true
						       }
						     },
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"values": {
									// Property: Values
									// CloudFormation resource type schema:
									/*
									   {
									     "items": {
									       "additionalProperties": false,
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
									     "type": "array",
									     "uniqueItems": true
									   }
									*/
									// Ordered set.
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
							},
						),
						Optional: true,
					},
					"source_ip_config": {
						// Property: SourceIpConfig
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "Values": {
						         "items": {
						           "type": "string"
						         },
						         "type": "array",
						         "uniqueItems": true
						       }
						     },
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"values": {
									// Property: Values
									// CloudFormation resource type schema:
									/*
									   {
									     "items": {
									       "type": "string"
									     },
									     "type": "array",
									     "uniqueItems": true
									   }
									*/
									// Ordered set.
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"values": {
						// Property: Values
						// CloudFormation resource type schema:
						/*
						   {
						     "items": {
						       "type": "string"
						     },
						     "type": "array",
						     "uniqueItems": true
						   }
						*/
						// Ordered set.
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Required: true,
		},
		"is_default": {
			// Property: IsDefault
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "boolean"
			   }
			*/
			Type:     types.BoolType,
			Computed: true,
		},
		"listener_arn": {
			// Property: ListenerArn
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
			// ListenerArn is a force-new attribute.
		},
		"priority": {
			// Property: Priority
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "integer"
			   }
			*/
			Type:     types.NumberType,
			Required: true,
		},
		"rule_arn": {
			// Property: RuleArn
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::ElasticLoadBalancingV2::ListenerRule",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElasticLoadBalancingV2::ListenerRule").WithTerraformTypeName("aws_elasticloadbalancingv2_listener_rule").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Actions/*/AuthenticateOidcConfig/ClientSecret",
		"/properties/Actions/*/ForwardConfig",
		"/properties/Actions/*/TargetGroupArn",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_elasticloadbalancingv2_listener_rule", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
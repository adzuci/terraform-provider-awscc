// Code generated by generators/resource/main.go; DO NOT EDIT.

package mwaa

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
	registry.AddResourceTypeFactory("aws_mwaa_environment", environmentResourceType)
}

// environmentResourceType returns the Terraform aws_mwaa_environment resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::MWAA::Environment resource type.
func environmentResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"airflow_configuration_options": {
			// Property: AirflowConfigurationOptions
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Key/value pairs representing Airflow configuration variables.\n    Keys are prefixed by their section:\n\n    [core]\n    dags_folder={AIRFLOW_HOME}/dags\n\n    Would be represented as\n\n    \"core.dags_folder\": \"{AIRFLOW_HOME}/dags\"",
			     "type": "object"
			   }
			*/
			Description: "Key/value pairs representing Airflow configuration variables.\n    Keys are prefixed by their section:\n\n    [core]\n    dags_folder={AIRFLOW_HOME}/dags\n\n    Would be represented as\n\n    \"core.dags_folder\": \"{AIRFLOW_HOME}/dags\"",
			Type:        types.MapType{ElemType: types.StringType},
			Optional:    true,
		},
		"airflow_version": {
			// Property: AirflowVersion
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Version of airflow to deploy to the environment.",
			     "maxLength": 32,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "Version of airflow to deploy to the environment.",
			Type:        types.StringType,
			Optional:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "ARN for the MWAA environment.",
			     "maxLength": 1224,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "ARN for the MWAA environment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"dag_s3_path": {
			// Property: DagS3Path
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Represents an S3 prefix relative to the root of an S3 bucket.",
			     "maxLength": 1024,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "Represents an S3 prefix relative to the root of an S3 bucket.",
			Type:        types.StringType,
			Optional:    true,
		},
		"environment_class": {
			// Property: EnvironmentClass
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Templated configuration for airflow processes and backing infrastructure.",
			     "maxLength": 1024,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "Templated configuration for airflow processes and backing infrastructure.",
			Type:        types.StringType,
			Optional:    true,
		},
		"execution_role_arn": {
			// Property: ExecutionRoleArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "IAM role to be used by tasks.",
			     "maxLength": 1224,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "IAM role to be used by tasks.",
			Type:        types.StringType,
			Optional:    true,
		},
		"kms_key": {
			// Property: KmsKey
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The identifier of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use for MWAA data encryption.\n\n    You can specify the CMK using any of the following:\n\n    Key ID. For example, key/1234abcd-12ab-34cd-56ef-1234567890ab.\n\n    Key alias. For example, alias/ExampleAlias.\n\n    Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.\n\n    Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.\n\n    AWS authenticates the CMK asynchronously. Therefore, if you specify an ID, alias, or ARN that is not valid, the action can appear to complete, but eventually fails.",
			     "maxLength": 1224,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The identifier of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use for MWAA data encryption.\n\n    You can specify the CMK using any of the following:\n\n    Key ID. For example, key/1234abcd-12ab-34cd-56ef-1234567890ab.\n\n    Key alias. For example, alias/ExampleAlias.\n\n    Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.\n\n    Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.\n\n    AWS authenticates the CMK asynchronously. Therefore, if you specify an ID, alias, or ARN that is not valid, the action can appear to complete, but eventually fails.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// KmsKey is a force-new attribute.
		},
		"logging_configuration": {
			// Property: LoggingConfiguration
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "Logging configuration for the environment.",
			     "properties": {
			       "DagProcessingLogs": {
			         "additionalProperties": false,
			         "description": "Logging configuration for a specific airflow component.",
			         "properties": {
			           "CloudWatchLogGroupArn": {
			             "description": "",
			             "maxLength": 1224,
			             "pattern": "",
			             "type": "string"
			           },
			           "Enabled": {
			             "description": "",
			             "type": "boolean"
			           },
			           "LogLevel": {
			             "description": "",
			             "enum": [
			               "CRITICAL",
			               "ERROR",
			               "WARNING",
			               "INFO",
			               "DEBUG"
			             ],
			             "type": "string"
			           }
			         },
			         "type": "object"
			       },
			       "SchedulerLogs": {
			         "additionalProperties": false,
			         "description": "Logging configuration for a specific airflow component.",
			         "properties": {
			           "CloudWatchLogGroupArn": {
			             "description": "",
			             "maxLength": 1224,
			             "pattern": "",
			             "type": "string"
			           },
			           "Enabled": {
			             "description": "",
			             "type": "boolean"
			           },
			           "LogLevel": {
			             "description": "",
			             "enum": [
			               "CRITICAL",
			               "ERROR",
			               "WARNING",
			               "INFO",
			               "DEBUG"
			             ],
			             "type": "string"
			           }
			         },
			         "type": "object"
			       },
			       "TaskLogs": {
			         "additionalProperties": false,
			         "description": "Logging configuration for a specific airflow component.",
			         "properties": {
			           "CloudWatchLogGroupArn": {
			             "description": "",
			             "maxLength": 1224,
			             "pattern": "",
			             "type": "string"
			           },
			           "Enabled": {
			             "description": "",
			             "type": "boolean"
			           },
			           "LogLevel": {
			             "description": "",
			             "enum": [
			               "CRITICAL",
			               "ERROR",
			               "WARNING",
			               "INFO",
			               "DEBUG"
			             ],
			             "type": "string"
			           }
			         },
			         "type": "object"
			       },
			       "WebserverLogs": {
			         "additionalProperties": false,
			         "description": "Logging configuration for a specific airflow component.",
			         "properties": {
			           "CloudWatchLogGroupArn": {
			             "description": "",
			             "maxLength": 1224,
			             "pattern": "",
			             "type": "string"
			           },
			           "Enabled": {
			             "description": "",
			             "type": "boolean"
			           },
			           "LogLevel": {
			             "description": "",
			             "enum": [
			               "CRITICAL",
			               "ERROR",
			               "WARNING",
			               "INFO",
			               "DEBUG"
			             ],
			             "type": "string"
			           }
			         },
			         "type": "object"
			       },
			       "WorkerLogs": {
			         "additionalProperties": false,
			         "description": "Logging configuration for a specific airflow component.",
			         "properties": {
			           "CloudWatchLogGroupArn": {
			             "description": "",
			             "maxLength": 1224,
			             "pattern": "",
			             "type": "string"
			           },
			           "Enabled": {
			             "description": "",
			             "type": "boolean"
			           },
			           "LogLevel": {
			             "description": "",
			             "enum": [
			               "CRITICAL",
			               "ERROR",
			               "WARNING",
			               "INFO",
			               "DEBUG"
			             ],
			             "type": "string"
			           }
			         },
			         "type": "object"
			       }
			     },
			     "type": "object"
			   }
			*/
			Description: "Logging configuration for the environment.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"dag_processing_logs": {
						// Property: DagProcessingLogs
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "description": "Logging configuration for a specific airflow component.",
						     "properties": {
						       "CloudWatchLogGroupArn": {
						         "description": "",
						         "maxLength": 1224,
						         "pattern": "",
						         "type": "string"
						       },
						       "Enabled": {
						         "description": "",
						         "type": "boolean"
						       },
						       "LogLevel": {
						         "description": "",
						         "enum": [
						           "CRITICAL",
						           "ERROR",
						           "WARNING",
						           "INFO",
						           "DEBUG"
						         ],
						         "type": "string"
						       }
						     },
						     "type": "object"
						   }
						*/
						Description: "Logging configuration for a specific airflow component.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"cloud_watch_log_group_arn": {
									// Property: CloudWatchLogGroupArn
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "",
									     "maxLength": 1224,
									     "pattern": "",
									     "type": "string"
									   }
									*/
									Description: "",
									Type:        types.StringType,
									Computed:    true,
								},
								"enabled": {
									// Property: Enabled
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "",
									     "type": "boolean"
									   }
									*/
									Description: "",
									Type:        types.BoolType,
									Optional:    true,
								},
								"log_level": {
									// Property: LogLevel
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "",
									     "enum": [
									       "CRITICAL",
									       "ERROR",
									       "WARNING",
									       "INFO",
									       "DEBUG"
									     ],
									     "type": "string"
									   }
									*/
									Description: "",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"scheduler_logs": {
						// Property: SchedulerLogs
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "description": "Logging configuration for a specific airflow component.",
						     "properties": {
						       "CloudWatchLogGroupArn": {
						         "description": "",
						         "maxLength": 1224,
						         "pattern": "",
						         "type": "string"
						       },
						       "Enabled": {
						         "description": "",
						         "type": "boolean"
						       },
						       "LogLevel": {
						         "description": "",
						         "enum": [
						           "CRITICAL",
						           "ERROR",
						           "WARNING",
						           "INFO",
						           "DEBUG"
						         ],
						         "type": "string"
						       }
						     },
						     "type": "object"
						   }
						*/
						Description: "Logging configuration for a specific airflow component.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"cloud_watch_log_group_arn": {
									// Property: CloudWatchLogGroupArn
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "",
									     "maxLength": 1224,
									     "pattern": "",
									     "type": "string"
									   }
									*/
									Description: "",
									Type:        types.StringType,
									Computed:    true,
								},
								"enabled": {
									// Property: Enabled
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "",
									     "type": "boolean"
									   }
									*/
									Description: "",
									Type:        types.BoolType,
									Optional:    true,
								},
								"log_level": {
									// Property: LogLevel
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "",
									     "enum": [
									       "CRITICAL",
									       "ERROR",
									       "WARNING",
									       "INFO",
									       "DEBUG"
									     ],
									     "type": "string"
									   }
									*/
									Description: "",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"task_logs": {
						// Property: TaskLogs
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "description": "Logging configuration for a specific airflow component.",
						     "properties": {
						       "CloudWatchLogGroupArn": {
						         "description": "",
						         "maxLength": 1224,
						         "pattern": "",
						         "type": "string"
						       },
						       "Enabled": {
						         "description": "",
						         "type": "boolean"
						       },
						       "LogLevel": {
						         "description": "",
						         "enum": [
						           "CRITICAL",
						           "ERROR",
						           "WARNING",
						           "INFO",
						           "DEBUG"
						         ],
						         "type": "string"
						       }
						     },
						     "type": "object"
						   }
						*/
						Description: "Logging configuration for a specific airflow component.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"cloud_watch_log_group_arn": {
									// Property: CloudWatchLogGroupArn
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "",
									     "maxLength": 1224,
									     "pattern": "",
									     "type": "string"
									   }
									*/
									Description: "",
									Type:        types.StringType,
									Computed:    true,
								},
								"enabled": {
									// Property: Enabled
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "",
									     "type": "boolean"
									   }
									*/
									Description: "",
									Type:        types.BoolType,
									Optional:    true,
								},
								"log_level": {
									// Property: LogLevel
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "",
									     "enum": [
									       "CRITICAL",
									       "ERROR",
									       "WARNING",
									       "INFO",
									       "DEBUG"
									     ],
									     "type": "string"
									   }
									*/
									Description: "",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"webserver_logs": {
						// Property: WebserverLogs
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "description": "Logging configuration for a specific airflow component.",
						     "properties": {
						       "CloudWatchLogGroupArn": {
						         "description": "",
						         "maxLength": 1224,
						         "pattern": "",
						         "type": "string"
						       },
						       "Enabled": {
						         "description": "",
						         "type": "boolean"
						       },
						       "LogLevel": {
						         "description": "",
						         "enum": [
						           "CRITICAL",
						           "ERROR",
						           "WARNING",
						           "INFO",
						           "DEBUG"
						         ],
						         "type": "string"
						       }
						     },
						     "type": "object"
						   }
						*/
						Description: "Logging configuration for a specific airflow component.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"cloud_watch_log_group_arn": {
									// Property: CloudWatchLogGroupArn
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "",
									     "maxLength": 1224,
									     "pattern": "",
									     "type": "string"
									   }
									*/
									Description: "",
									Type:        types.StringType,
									Computed:    true,
								},
								"enabled": {
									// Property: Enabled
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "",
									     "type": "boolean"
									   }
									*/
									Description: "",
									Type:        types.BoolType,
									Optional:    true,
								},
								"log_level": {
									// Property: LogLevel
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "",
									     "enum": [
									       "CRITICAL",
									       "ERROR",
									       "WARNING",
									       "INFO",
									       "DEBUG"
									     ],
									     "type": "string"
									   }
									*/
									Description: "",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"worker_logs": {
						// Property: WorkerLogs
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "description": "Logging configuration for a specific airflow component.",
						     "properties": {
						       "CloudWatchLogGroupArn": {
						         "description": "",
						         "maxLength": 1224,
						         "pattern": "",
						         "type": "string"
						       },
						       "Enabled": {
						         "description": "",
						         "type": "boolean"
						       },
						       "LogLevel": {
						         "description": "",
						         "enum": [
						           "CRITICAL",
						           "ERROR",
						           "WARNING",
						           "INFO",
						           "DEBUG"
						         ],
						         "type": "string"
						       }
						     },
						     "type": "object"
						   }
						*/
						Description: "Logging configuration for a specific airflow component.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"cloud_watch_log_group_arn": {
									// Property: CloudWatchLogGroupArn
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "",
									     "maxLength": 1224,
									     "pattern": "",
									     "type": "string"
									   }
									*/
									Description: "",
									Type:        types.StringType,
									Computed:    true,
								},
								"enabled": {
									// Property: Enabled
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "",
									     "type": "boolean"
									   }
									*/
									Description: "",
									Type:        types.BoolType,
									Optional:    true,
								},
								"log_level": {
									// Property: LogLevel
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "",
									     "enum": [
									       "CRITICAL",
									       "ERROR",
									       "WARNING",
									       "INFO",
									       "DEBUG"
									     ],
									     "type": "string"
									   }
									*/
									Description: "",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"max_workers": {
			// Property: MaxWorkers
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Maximum worker compute units.",
			     "type": "integer"
			   }
			*/
			Description: "Maximum worker compute units.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"min_workers": {
			// Property: MinWorkers
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Minimum worker compute units.",
			     "type": "integer"
			   }
			*/
			Description: "Minimum worker compute units.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Customer-defined identifier for the environment, unique per customer region.",
			     "maxLength": 80,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "Customer-defined identifier for the environment, unique per customer region.",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"network_configuration": {
			// Property: NetworkConfiguration
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "Configures the network resources of the environment.",
			     "properties": {
			       "SecurityGroupIds": {
			         "description": "A list of security groups to use for the environment.",
			         "items": {
			           "description": "",
			           "maxLength": 1024,
			           "minLength": 1,
			           "pattern": "",
			           "type": "string"
			         },
			         "maxItems": 5,
			         "minItems": 1,
			         "type": "array"
			       },
			       "SubnetIds": {
			         "description": "A list of subnets to use for the environment. These must be private subnets, in the same VPC, in two different availability zones.",
			         "items": {
			           "description": "",
			           "maxLength": 1024,
			           "pattern": "",
			           "type": "string"
			         },
			         "maxItems": 2,
			         "minItems": 2,
			         "type": "array"
			       }
			     },
			     "type": "object"
			   }
			*/
			Description: "Configures the network resources of the environment.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"security_group_ids": {
						// Property: SecurityGroupIds
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "A list of security groups to use for the environment.",
						     "items": {
						       "description": "",
						       "maxLength": 1024,
						       "minLength": 1,
						       "pattern": "",
						       "type": "string"
						     },
						     "maxItems": 5,
						     "minItems": 1,
						     "type": "array"
						   }
						*/
						Description: "A list of security groups to use for the environment.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
					},
					"subnet_ids": {
						// Property: SubnetIds
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "A list of subnets to use for the environment. These must be private subnets, in the same VPC, in two different availability zones.",
						     "items": {
						       "description": "",
						       "maxLength": 1024,
						       "pattern": "",
						       "type": "string"
						     },
						     "maxItems": 2,
						     "minItems": 2,
						     "type": "array"
						   }
						*/
						Description: "A list of subnets to use for the environment. These must be private subnets, in the same VPC, in two different availability zones.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Computed:    true,
						// SubnetIds is a force-new attribute.
					},
				},
			),
			Optional: true,
		},
		"plugins_s3_object_version": {
			// Property: PluginsS3ObjectVersion
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Represents an version ID for an S3 object.",
			     "maxLength": 1024,
			     "type": "string"
			   }
			*/
			Description: "Represents an version ID for an S3 object.",
			Type:        types.StringType,
			Optional:    true,
		},
		"plugins_s3_path": {
			// Property: PluginsS3Path
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Represents an S3 prefix relative to the root of an S3 bucket.",
			     "maxLength": 1024,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "Represents an S3 prefix relative to the root of an S3 bucket.",
			Type:        types.StringType,
			Optional:    true,
		},
		"requirements_s3_object_version": {
			// Property: RequirementsS3ObjectVersion
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Represents an version ID for an S3 object.",
			     "maxLength": 1024,
			     "type": "string"
			   }
			*/
			Description: "Represents an version ID for an S3 object.",
			Type:        types.StringType,
			Optional:    true,
		},
		"requirements_s3_path": {
			// Property: RequirementsS3Path
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Represents an S3 prefix relative to the root of an S3 bucket.",
			     "maxLength": 1024,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "Represents an S3 prefix relative to the root of an S3 bucket.",
			Type:        types.StringType,
			Optional:    true,
		},
		"schedulers": {
			// Property: Schedulers
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Scheduler compute units.",
			     "type": "integer"
			   }
			*/
			Description: "Scheduler compute units.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"source_bucket_arn": {
			// Property: SourceBucketArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "ARN for the AWS S3 bucket to use as the source of DAGs and plugins for the environment.",
			     "maxLength": 1224,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "ARN for the AWS S3 bucket to use as the source of DAGs and plugins for the environment.",
			Type:        types.StringType,
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A map of tags for the environment.",
			     "type": "object"
			   }
			*/
			Description: "A map of tags for the environment.",
			Type:        types.MapType{ElemType: types.StringType},
			Optional:    true,
		},
		"webserver_access_mode": {
			// Property: WebserverAccessMode
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Choice for mode of webserver access including over public internet or via private VPC endpoint.",
			     "enum": [
			       "PRIVATE_ONLY",
			       "PUBLIC_ONLY"
			     ],
			     "type": "string"
			   }
			*/
			Description: "Choice for mode of webserver access including over public internet or via private VPC endpoint.",
			Type:        types.StringType,
			Optional:    true,
		},
		"webserver_url": {
			// Property: WebserverUrl
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Url endpoint for the environment's Airflow UI.",
			     "maxLength": 256,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "Url endpoint for the environment's Airflow UI.",
			Type:        types.StringType,
			Computed:    true,
		},
		"weekly_maintenance_window_start": {
			// Property: WeeklyMaintenanceWindowStart
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Start time for the weekly maintenance window.",
			     "maxLength": 9,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "Start time for the weekly maintenance window.",
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
		Description: "Resource schema for AWS::MWAA::Environment",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::MWAA::Environment").WithTerraformTypeName("aws_mwaa_environment").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(180).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(480)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_mwaa_environment", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
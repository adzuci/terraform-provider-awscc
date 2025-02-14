// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sqs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_sqs_queue", queueDataSource)
}

// queueDataSource returns the Terraform awscc_sqs_queue data source.
// This Terraform data source corresponds to the CloudFormation AWS::SQS::Queue resource.
func queueDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name (ARN) of the queue.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name (ARN) of the queue.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ContentBasedDeduplication
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication. During the deduplication interval, Amazon SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message.",
		//	  "type": "boolean"
		//	}
		"content_based_deduplication": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication. During the deduplication interval, Amazon SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeduplicationScope
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether message deduplication occurs at the message group or queue level. Valid values are messageGroup and queue.",
		//	  "type": "string"
		//	}
		"deduplication_scope": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether message deduplication occurs at the message group or queue level. Valid values are messageGroup and queue.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DelaySeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time in seconds for which the delivery of all messages in the queue is delayed. You can specify an integer value of 0 to 900 (15 minutes). The default value is 0.",
		//	  "type": "integer"
		//	}
		"delay_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The time in seconds for which the delivery of all messages in the queue is delayed. You can specify an integer value of 0 to 900 (15 minutes). The default value is 0.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FifoQueue
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "If set to true, creates a FIFO queue. If you don't specify this property, Amazon SQS creates a standard queue.",
		//	  "type": "boolean"
		//	}
		"fifo_queue": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "If set to true, creates a FIFO queue. If you don't specify this property, Amazon SQS creates a standard queue.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FifoThroughputLimit
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are perQueue and perMessageGroupId. The perMessageGroupId value is allowed only when the value for DeduplicationScope is messageGroup.",
		//	  "type": "string"
		//	}
		"fifo_throughput_limit": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are perQueue and perMessageGroupId. The perMessageGroupId value is allowed only when the value for DeduplicationScope is messageGroup.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KmsDataKeyReusePeriodSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The length of time in seconds for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).",
		//	  "type": "integer"
		//	}
		"kms_data_key_reuse_period_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The length of time in seconds for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KmsMasterKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of an AWS managed customer master key (CMK) for Amazon SQS or a custom CMK. To use the AWS managed CMK for Amazon SQS, specify the (default) alias alias/aws/sqs.",
		//	  "type": "string"
		//	}
		"kms_master_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of an AWS managed customer master key (CMK) for Amazon SQS or a custom CMK. To use the AWS managed CMK for Amazon SQS, specify the (default) alias alias/aws/sqs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MaximumMessageSize
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The limit of how many bytes that a message can contain before Amazon SQS rejects it. You can specify an integer value from 1,024 bytes (1 KiB) to 262,144 bytes (256 KiB). The default value is 262,144 (256 KiB).",
		//	  "type": "integer"
		//	}
		"maximum_message_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The limit of how many bytes that a message can contain before Amazon SQS rejects it. You can specify an integer value from 1,024 bytes (1 KiB) to 262,144 bytes (256 KiB). The default value is 262,144 (256 KiB).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MessageRetentionPeriod
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of seconds that Amazon SQS retains a message. You can specify an integer value from 60 seconds (1 minute) to 1,209,600 seconds (14 days). The default value is 345,600 seconds (4 days).",
		//	  "type": "integer"
		//	}
		"message_retention_period": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of seconds that Amazon SQS retains a message. You can specify an integer value from 60 seconds (1 minute) to 1,209,600 seconds (14 days). The default value is 345,600 seconds (4 days).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: QueueName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the queue. To create a FIFO queue, the name of your FIFO queue must end with the .fifo suffix.",
		//	  "type": "string"
		//	}
		"queue_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the queue. To create a FIFO queue, the name of your FIFO queue must end with the .fifo suffix.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: QueueUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "URL of the source queue.",
		//	  "type": "string"
		//	}
		"queue_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "URL of the source queue.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReceiveMessageWaitTimeSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available. You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property.",
		//	  "type": "integer"
		//	}
		"receive_message_wait_time_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available. You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RedriveAllowPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object.",
		//	  "type": "string"
		//	}
		"redrive_allow_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RedrivePolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A string that includes the parameters for the dead-letter queue functionality (redrive policy) of the source queue.",
		//	  "type": "string"
		//	}
		"redrive_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A string that includes the parameters for the dead-letter queue functionality (redrive policy) of the source queue.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SqsManagedSseEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Enables server-side queue encryption using SQS owned encryption keys. Only one server-side encryption option is supported per queue (e.g. SSE-KMS or SSE-SQS ).",
		//	  "type": "boolean"
		//	}
		"sqs_managed_sse_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Enables server-side queue encryption using SQS owned encryption keys. Only one server-side encryption option is supported per queue (e.g. SSE-KMS or SSE-SQS ).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags that you attach to this queue.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags that you attach to this queue.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VisibilityTimeout
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The length of time during which a message will be unavailable after a message is delivered from the queue. This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue. Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.",
		//	  "type": "integer"
		//	}
		"visibility_timeout": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The length of time during which a message will be unavailable after a message is delivered from the queue. This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue. Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SQS::Queue",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SQS::Queue").WithTerraformTypeName("awscc_sqs_queue")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                               "Arn",
		"content_based_deduplication":       "ContentBasedDeduplication",
		"deduplication_scope":               "DeduplicationScope",
		"delay_seconds":                     "DelaySeconds",
		"fifo_queue":                        "FifoQueue",
		"fifo_throughput_limit":             "FifoThroughputLimit",
		"key":                               "Key",
		"kms_data_key_reuse_period_seconds": "KmsDataKeyReusePeriodSeconds",
		"kms_master_key_id":                 "KmsMasterKeyId",
		"maximum_message_size":              "MaximumMessageSize",
		"message_retention_period":          "MessageRetentionPeriod",
		"queue_name":                        "QueueName",
		"queue_url":                         "QueueUrl",
		"receive_message_wait_time_seconds": "ReceiveMessageWaitTimeSeconds",
		"redrive_allow_policy":              "RedriveAllowPolicy",
		"redrive_policy":                    "RedrivePolicy",
		"sqs_managed_sse_enabled":           "SqsManagedSseEnabled",
		"tags":                              "Tags",
		"value":                             "Value",
		"visibility_timeout":                "VisibilityTimeout",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

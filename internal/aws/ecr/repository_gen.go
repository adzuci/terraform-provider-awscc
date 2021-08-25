// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecr

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ecr_repository", repositoryResourceType)
}

// repositoryResourceType returns the Terraform awscc_ecr_repository resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ECR::Repository resource type.
func repositoryResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"encryption_configuration": {
			// Property: EncryptionConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.\n\nBy default, when no encryption configuration is set or the AES256 encryption type is used, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts your data at rest using an AES-256 encryption algorithm. This does not require any action on your part.\n\nFor more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html",
			//   "properties": {
			//     "EncryptionType": {
			//       "description": "The encryption type to use.",
			//       "enum": [
			//         "AES256",
			//         "KMS"
			//       ],
			//       "type": "string"
			//     },
			//     "KmsKey": {
			//       "description": "If you use the KMS encryption type, specify the CMK to use for encryption. The alias, key ID, or full ARN of the CMK can be specified. The key must exist in the same Region as the repository. If no key is specified, the default AWS managed CMK for Amazon ECR will be used.",
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "EncryptionType"
			//   ],
			//   "type": "object"
			// }
			Description: "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.\n\nBy default, when no encryption configuration is set or the AES256 encryption type is used, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts your data at rest using an AES-256 encryption algorithm. This does not require any action on your part.\n\nFor more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"encryption_type": {
						// Property: EncryptionType
						Description: "The encryption type to use.",
						Type:        types.StringType,
						Required:    true,
						// EncryptionType is a force-new attribute.
					},
					"kms_key": {
						// Property: KmsKey
						Description: "If you use the KMS encryption type, specify the CMK to use for encryption. The alias, key ID, or full ARN of the CMK can be specified. The key must exist in the same Region as the repository. If no key is specified, the default AWS managed CMK for Amazon ECR will be used.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						// KmsKey is a force-new attribute.
					},
				},
			),
			Optional: true,
			Computed: true,
			// EncryptionConfiguration is a force-new attribute.
		},
		"image_scanning_configuration": {
			// Property: ImageScanningConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The image scanning configuration for the repository. This setting determines whether images are scanned for known vulnerabilities after being pushed to the repository.",
			//   "properties": {
			//     "ScanOnPush": {
			//       "description": "The setting that determines whether images are scanned after being pushed to a repository.",
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The image scanning configuration for the repository. This setting determines whether images are scanned for known vulnerabilities after being pushed to the repository.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"scan_on_push": {
						// Property: ScanOnPush
						Description: "The setting that determines whether images are scanned after being pushed to a repository.",
						Type:        types.BoolType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"image_tag_mutability": {
			// Property: ImageTagMutability
			// CloudFormation resource type schema:
			// {
			//   "description": "The image tag mutability setting for the repository.",
			//   "enum": [
			//     "MUTABLE",
			//     "IMMUTABLE"
			//   ],
			//   "type": "string"
			// }
			Description: "The image tag mutability setting for the repository.",
			Type:        types.StringType,
			Optional:    true,
		},
		"lifecycle_policy": {
			// Property: LifecyclePolicy
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The LifecyclePolicy property type specifies a lifecycle policy. For information about lifecycle policy syntax, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html",
			//   "properties": {
			//     "LifecyclePolicyText": {
			//       "description": "The JSON repository policy text to apply to the repository.",
			//       "maxLength": 30720,
			//       "minLength": 100,
			//       "type": "string"
			//     },
			//     "RegistryId": {
			//       "description": "The AWS account ID associated with the registry that contains the repository. If you do not specify a registry, the default registry is assumed. ",
			//       "maxLength": 12,
			//       "minLength": 12,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The LifecyclePolicy property type specifies a lifecycle policy. For information about lifecycle policy syntax, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"lifecycle_policy_text": {
						// Property: LifecyclePolicyText
						Description: "The JSON repository policy text to apply to the repository.",
						Type:        types.StringType,
						Optional:    true,
					},
					"registry_id": {
						// Property: RegistryId
						Description: "The AWS account ID associated with the registry that contains the repository. If you do not specify a registry, the default registry is assumed. ",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"repository_name": {
			// Property: RepositoryName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.",
			//   "maxLength": 256,
			//   "minLength": 2,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// RepositoryName is a force-new attribute.
		},
		"repository_policy_text": {
			// Property: RepositoryPolicyText
			// CloudFormation resource type schema:
			// {
			//   "description": "The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide. ",
			//   "type": "string"
			// }
			Description: "The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide. ",
			Type:        types.StringType,
			Optional:    true,
		},
		"repository_uri": {
			// Property: RepositoryUri
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MaxItems: 50,
				},
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
		Description: "The AWS::ECR::Repository resource specifies an Amazon Elastic Container Registry (Amazon ECR) repository, where users can push and pull Docker images. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/Repositories.html",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECR::Repository").WithTerraformTypeName("awscc_ecr_repository")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                          "Arn",
		"encryption_configuration":     "EncryptionConfiguration",
		"encryption_type":              "EncryptionType",
		"image_scanning_configuration": "ImageScanningConfiguration",
		"image_tag_mutability":         "ImageTagMutability",
		"key":                          "Key",
		"kms_key":                      "KmsKey",
		"lifecycle_policy":             "LifecyclePolicy",
		"lifecycle_policy_text":        "LifecyclePolicyText",
		"registry_id":                  "RegistryId",
		"repository_name":              "RepositoryName",
		"repository_policy_text":       "RepositoryPolicyText",
		"repository_uri":               "RepositoryUri",
		"scan_on_push":                 "ScanOnPush",
		"tags":                         "Tags",
		"value":                        "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ecr_repository", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

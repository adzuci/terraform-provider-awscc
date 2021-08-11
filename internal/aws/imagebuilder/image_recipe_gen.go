// Code generated by generators/resource/main.go; DO NOT EDIT.

package imagebuilder

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
	registry.AddResourceTypeFactory("aws_imagebuilder_image_recipe", imageRecipeResourceType)
}

// imageRecipeResourceType returns the Terraform aws_imagebuilder_image_recipe resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ImageBuilder::ImageRecipe resource type.
func imageRecipeResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"additional_instance_configuration": {
			// Property: AdditionalInstanceConfiguration
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "Specify additional settings and launch scripts for your build instances.",
			     "properties": {
			       "SystemsManagerAgent": {
			         "additionalProperties": false,
			         "description": "Contains settings for the SSM agent on your build instance.",
			         "properties": {
			           "UninstallAfterBuild": {
			             "description": "This property defaults to true. If Image Builder installs the SSM agent on a build instance, it removes the agent before creating a snapshot for the AMI. To ensure that the AMI you create includes the SSM agent, set this property to false.",
			             "type": "boolean"
			           }
			         },
			         "required": [
			           "UninstallAfterBuild"
			         ],
			         "type": "object"
			       },
			       "UserDataOverride": {
			         "description": "Use this property to provide commands or a command script to run when you launch your build instance.",
			         "type": "string"
			       }
			     },
			     "required": [
			       "UserDataOverride"
			     ],
			     "type": "object"
			   }
			*/
			Description: "Specify additional settings and launch scripts for your build instances.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"systems_manager_agent": {
						// Property: SystemsManagerAgent
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "description": "Contains settings for the SSM agent on your build instance.",
						     "properties": {
						       "UninstallAfterBuild": {
						         "description": "This property defaults to true. If Image Builder installs the SSM agent on a build instance, it removes the agent before creating a snapshot for the AMI. To ensure that the AMI you create includes the SSM agent, set this property to false.",
						         "type": "boolean"
						       }
						     },
						     "required": [
						       "UninstallAfterBuild"
						     ],
						     "type": "object"
						   }
						*/
						Description: "Contains settings for the SSM agent on your build instance.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"uninstall_after_build": {
									// Property: UninstallAfterBuild
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "This property defaults to true. If Image Builder installs the SSM agent on a build instance, it removes the agent before creating a snapshot for the AMI. To ensure that the AMI you create includes the SSM agent, set this property to false.",
									     "type": "boolean"
									   }
									*/
									Description: "This property defaults to true. If Image Builder installs the SSM agent on a build instance, it removes the agent before creating a snapshot for the AMI. To ensure that the AMI you create includes the SSM agent, set this property to false.",
									Type:        types.BoolType,
									Required:    true,
								},
							},
						),
						Optional: true,
					},
					"user_data_override": {
						// Property: UserDataOverride
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Use this property to provide commands or a command script to run when you launch your build instance.",
						     "type": "string"
						   }
						*/
						Description: "Use this property to provide commands or a command script to run when you launch your build instance.",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Optional: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the image recipe.",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the image recipe.",
			Type:        types.StringType,
			Computed:    true,
		},
		"block_device_mappings": {
			// Property: BlockDeviceMappings
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The block device mappings to apply when creating images from this recipe.",
			     "items": {
			       "additionalProperties": false,
			       "description": "Defines block device mappings for the instance used to configure your image. ",
			       "properties": {
			         "DeviceName": {
			           "description": "The device to which these mappings apply.",
			           "type": "string"
			         },
			         "Ebs": {
			           "additionalProperties": false,
			           "description": "Amazon EBS-specific block device mapping specifications. ",
			           "properties": {
			             "DeleteOnTermination": {
			               "description": "Use to configure delete on termination of the associated device.",
			               "type": "boolean"
			             },
			             "Encrypted": {
			               "description": "Use to configure device encryption.",
			               "type": "boolean"
			             },
			             "Iops": {
			               "description": "Use to configure device IOPS.",
			               "type": "integer"
			             },
			             "KmsKeyId": {
			               "description": "Use to configure the KMS key to use when encrypting the device.",
			               "type": "string"
			             },
			             "SnapshotId": {
			               "description": "The snapshot that defines the device contents.",
			               "type": "string"
			             },
			             "VolumeSize": {
			               "description": "Use to override the device's volume size.",
			               "type": "integer"
			             },
			             "VolumeType": {
			               "description": "Use to override the device's volume type.",
			               "enum": [
			                 "standard",
			                 "io1",
			                 "io2",
			                 "gp2",
			                 "gp3",
			                 "sc1",
			                 "st1"
			               ],
			               "type": "string"
			             }
			           },
			           "type": "object"
			         },
			         "NoDevice": {
			           "description": "Use to remove a mapping from the parent image.",
			           "type": "string"
			         },
			         "VirtualName": {
			           "description": "Use to manage instance ephemeral devices.",
			           "type": "string"
			         }
			       },
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Description: "The block device mappings to apply when creating images from this recipe.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"device_name": {
						// Property: DeviceName
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The device to which these mappings apply.",
						     "type": "string"
						   }
						*/
						Description: "The device to which these mappings apply.",
						Type:        types.StringType,
						Optional:    true,
					},
					"ebs": {
						// Property: Ebs
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "description": "Amazon EBS-specific block device mapping specifications. ",
						     "properties": {
						       "DeleteOnTermination": {
						         "description": "Use to configure delete on termination of the associated device.",
						         "type": "boolean"
						       },
						       "Encrypted": {
						         "description": "Use to configure device encryption.",
						         "type": "boolean"
						       },
						       "Iops": {
						         "description": "Use to configure device IOPS.",
						         "type": "integer"
						       },
						       "KmsKeyId": {
						         "description": "Use to configure the KMS key to use when encrypting the device.",
						         "type": "string"
						       },
						       "SnapshotId": {
						         "description": "The snapshot that defines the device contents.",
						         "type": "string"
						       },
						       "VolumeSize": {
						         "description": "Use to override the device's volume size.",
						         "type": "integer"
						       },
						       "VolumeType": {
						         "description": "Use to override the device's volume type.",
						         "enum": [
						           "standard",
						           "io1",
						           "io2",
						           "gp2",
						           "gp3",
						           "sc1",
						           "st1"
						         ],
						         "type": "string"
						       }
						     },
						     "type": "object"
						   }
						*/
						Description: "Amazon EBS-specific block device mapping specifications. ",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"delete_on_termination": {
									// Property: DeleteOnTermination
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "Use to configure delete on termination of the associated device.",
									     "type": "boolean"
									   }
									*/
									Description: "Use to configure delete on termination of the associated device.",
									Type:        types.BoolType,
									Optional:    true,
								},
								"encrypted": {
									// Property: Encrypted
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "Use to configure device encryption.",
									     "type": "boolean"
									   }
									*/
									Description: "Use to configure device encryption.",
									Type:        types.BoolType,
									Optional:    true,
								},
								"iops": {
									// Property: Iops
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "Use to configure device IOPS.",
									     "type": "integer"
									   }
									*/
									Description: "Use to configure device IOPS.",
									Type:        types.NumberType,
									Optional:    true,
								},
								"kms_key_id": {
									// Property: KmsKeyId
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "Use to configure the KMS key to use when encrypting the device.",
									     "type": "string"
									   }
									*/
									Description: "Use to configure the KMS key to use when encrypting the device.",
									Type:        types.StringType,
									Optional:    true,
								},
								"snapshot_id": {
									// Property: SnapshotId
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The snapshot that defines the device contents.",
									     "type": "string"
									   }
									*/
									Description: "The snapshot that defines the device contents.",
									Type:        types.StringType,
									Optional:    true,
								},
								"volume_size": {
									// Property: VolumeSize
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "Use to override the device's volume size.",
									     "type": "integer"
									   }
									*/
									Description: "Use to override the device's volume size.",
									Type:        types.NumberType,
									Optional:    true,
								},
								"volume_type": {
									// Property: VolumeType
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "Use to override the device's volume type.",
									     "enum": [
									       "standard",
									       "io1",
									       "io2",
									       "gp2",
									       "gp3",
									       "sc1",
									       "st1"
									     ],
									     "type": "string"
									   }
									*/
									Description: "Use to override the device's volume type.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"no_device": {
						// Property: NoDevice
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Use to remove a mapping from the parent image.",
						     "type": "string"
						   }
						*/
						Description: "Use to remove a mapping from the parent image.",
						Type:        types.StringType,
						Optional:    true,
					},
					"virtual_name": {
						// Property: VirtualName
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Use to manage instance ephemeral devices.",
						     "type": "string"
						   }
						*/
						Description: "Use to manage instance ephemeral devices.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
			Computed: true,
			// BlockDeviceMappings is a force-new attribute.
		},
		"components": {
			// Property: Components
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The components of the image recipe.",
			     "items": {
			       "additionalProperties": false,
			       "description": "Configuration details of the component.",
			       "properties": {
			         "ComponentArn": {
			           "description": "The Amazon Resource Name (ARN) of the component.",
			           "type": "string"
			         },
			         "Parameters": {
			           "description": "A group of parameter settings that are used to configure the component for a specific recipe.",
			           "items": {
			             "additionalProperties": false,
			             "description": "Contains a key/value pair that sets the named component parameter.",
			             "properties": {
			               "Name": {
			                 "description": "The name of the component parameter to set.",
			                 "type": "string"
			               },
			               "Value": {
			                 "description": "Sets the value for the named component parameter.",
			                 "items": {
			                   "type": "string"
			                 },
			                 "type": "array"
			               }
			             },
			             "required": [
			               "Name",
			               "Value"
			             ],
			             "type": "object"
			           },
			           "type": "array"
			         }
			       },
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Description: "The components of the image recipe.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"component_arn": {
						// Property: ComponentArn
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The Amazon Resource Name (ARN) of the component.",
						     "type": "string"
						   }
						*/
						Description: "The Amazon Resource Name (ARN) of the component.",
						Type:        types.StringType,
						Optional:    true,
					},
					"parameters": {
						// Property: Parameters
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "A group of parameter settings that are used to configure the component for a specific recipe.",
						     "items": {
						       "additionalProperties": false,
						       "description": "Contains a key/value pair that sets the named component parameter.",
						       "properties": {
						         "Name": {
						           "description": "The name of the component parameter to set.",
						           "type": "string"
						         },
						         "Value": {
						           "description": "Sets the value for the named component parameter.",
						           "items": {
						             "type": "string"
						           },
						           "type": "array"
						         }
						       },
						       "required": [
						         "Name",
						         "Value"
						       ],
						       "type": "object"
						     },
						     "type": "array"
						   }
						*/
						Description: "A group of parameter settings that are used to configure the component for a specific recipe.",
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"name": {
									// Property: Name
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The name of the component parameter to set.",
									     "type": "string"
									   }
									*/
									Description: "The name of the component parameter to set.",
									Type:        types.StringType,
									Required:    true,
								},
								"value": {
									// Property: Value
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "Sets the value for the named component parameter.",
									     "items": {
									       "type": "string"
									     },
									     "type": "array"
									   }
									*/
									Description: "Sets the value for the named component parameter.",
									Type:        types.ListType{ElemType: types.StringType},
									Required:    true,
								},
							},
							schema.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Required: true,
			// Components is a force-new attribute.
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The description of the image recipe.",
			     "type": "string"
			   }
			*/
			Description: "The description of the image recipe.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Description is a force-new attribute.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the image recipe.",
			     "type": "string"
			   }
			*/
			Description: "The name of the image recipe.",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"parent_image": {
			// Property: ParentImage
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The parent image of the image recipe.",
			     "type": "string"
			   }
			*/
			Description: "The parent image of the image recipe.",
			Type:        types.StringType,
			Required:    true,
			// ParentImage is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "The tags of the image recipe.",
			     "patternProperties": {
			       "": {
			         "type": "string"
			       }
			     },
			     "type": "object"
			   }
			*/
			Description: "The tags of the image recipe.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			// Tags is a force-new attribute.
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The version of the image recipe.",
			     "type": "string"
			   }
			*/
			Description: "The version of the image recipe.",
			Type:        types.StringType,
			Required:    true,
			// Version is a force-new attribute.
		},
		"working_directory": {
			// Property: WorkingDirectory
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The working directory to be used during build and test workflows.",
			     "type": "string"
			   }
			*/
			Description: "The working directory to be used during build and test workflows.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// WorkingDirectory is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::ImageBuilder::ImageRecipe",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ImageBuilder::ImageRecipe").WithTerraformTypeName("aws_imagebuilder_image_recipe").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_imagebuilder_image_recipe", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

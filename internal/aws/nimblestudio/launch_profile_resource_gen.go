// Code generated by generators/resource/main.go; DO NOT EDIT.

package nimblestudio

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_nimblestudio_launch_profile", launchProfileResourceType)
}

// launchProfileResourceType returns the Terraform awscc_nimblestudio_launch_profile resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NimbleStudio::LaunchProfile resource type.
func launchProfileResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe description.\u003c/p\u003e",
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "<p>The description.</p>",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 256),
			},
		},
		"ec_2_subnet_ids": {
			// Property: Ec2SubnetIds
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eSpecifies the IDs of the EC2 subnets where streaming sessions will be accessible from.\n            These subnets must support the specified instance types. \u003c/p\u003e",
			//   "items": {
			//     "type": "string"
			//   },
			//   "maxItems": 6,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "<p>Specifies the IDs of the EC2 subnets where streaming sessions will be accessible from.\n            These subnets must support the specified instance types. </p>",
			Type:        types.ListType{ElemType: types.StringType},
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 6),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"launch_profile_id": {
			// Property: LaunchProfileId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"launch_profile_protocol_versions": {
			// Property: LaunchProfileProtocolVersions
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe version number of the protocol that is used by the launch profile. The only valid\n            version is \"2021-03-31\".\u003c/p\u003e",
			//   "items": {
			//     "description": "\u003cp\u003eThe version number of the protocol that is used by the launch profile. The only valid\n            version is \"2021-03-31\".\u003c/p\u003e",
			//     "maxLength": 10,
			//     "minLength": 0,
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "<p>The version number of the protocol that is used by the launch profile. The only valid\n            version is \"2021-03-31\".</p>",
			Type:        types.ListType{ElemType: types.StringType},
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayForEach(validate.StringLenBetween(0, 10)),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe name for the launch profile.\u003c/p\u003e",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "<p>The name for the launch profile.</p>",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
		},
		"stream_configuration": {
			// Property: StreamConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "\u003cp\u003eA configuration for a streaming session.\u003c/p\u003e",
			//   "properties": {
			//     "ClipboardMode": {
			//       "enum": [
			//         "ENABLED",
			//         "DISABLED"
			//       ],
			//       "type": "string"
			//     },
			//     "Ec2InstanceTypes": {
			//       "description": "\u003cp\u003eThe EC2 instance types that users can select from when launching a streaming session\n            with this launch profile.\u003c/p\u003e",
			//       "items": {
			//         "enum": [
			//           "g4dn.xlarge",
			//           "g4dn.2xlarge",
			//           "g4dn.4xlarge",
			//           "g4dn.8xlarge",
			//           "g4dn.12xlarge",
			//           "g4dn.16xlarge"
			//         ],
			//         "type": "string"
			//       },
			//       "maxItems": 30,
			//       "minItems": 1,
			//       "type": "array"
			//     },
			//     "MaxSessionLengthInMinutes": {
			//       "description": "\u003cp\u003eThe length of time, in minutes, that a streaming session can be active before it is\n            stopped or terminated. After this point, Nimble Studio automatically terminates or\n            stops the session. The default length of time is 690 minutes, and the maximum length of\n            time is 30 days.\u003c/p\u003e",
			//       "maximum": 43200,
			//       "minimum": 1,
			//       "type": "number"
			//     },
			//     "MaxStoppedSessionLengthInMinutes": {
			//       "description": "\u003cp\u003eInteger that determines if you can start and stop your sessions and how long a session\n            can stay in the STOPPED state. The default value is 0. The maximum value is 5760.\u003c/p\u003e\n        \u003cp\u003eIf the value is missing or set to 0, your sessions can’t be stopped. If you then call\n            StopStreamingSession, the session fails. If the time that a session stays in the READY\n            state exceeds the maxSessionLengthInMinutes value, the session will automatically be\n            terminated by AWS (instead of stopped).\u003c/p\u003e\n        \u003cp\u003eIf the value is set to a positive number, the session can be stopped. You can call\n            StopStreamingSession to stop sessions in the READY state. If the time that a session\n            stays in the READY state exceeds the maxSessionLengthInMinutes value, the session will\n            automatically be stopped by AWS (instead of terminated).\u003c/p\u003e",
			//       "maximum": 5760,
			//       "minimum": 0,
			//       "type": "number"
			//     },
			//     "SessionStorage": {
			//       "additionalProperties": false,
			//       "description": "\u003cp\u003eThe configuration for a streaming session’s upload storage.\u003c/p\u003e",
			//       "properties": {
			//         "Mode": {
			//           "description": "\u003cp\u003eAllows artists to upload files to their workstations. The only valid option is\n                \u003ccode\u003eUPLOAD\u003c/code\u003e.\u003c/p\u003e",
			//           "items": {
			//             "enum": [
			//               "UPLOAD"
			//             ],
			//             "type": "string"
			//           },
			//           "minItems": 1,
			//           "type": "array"
			//         },
			//         "Root": {
			//           "additionalProperties": false,
			//           "description": "\u003cp\u003eThe upload storage root location (folder) on streaming workstations where files are\n            uploaded.\u003c/p\u003e",
			//           "properties": {
			//             "Linux": {
			//               "description": "\u003cp\u003eThe folder path in Linux workstations where files are uploaded.\u003c/p\u003e",
			//               "maxLength": 128,
			//               "minLength": 1,
			//               "pattern": "",
			//               "type": "string"
			//             },
			//             "Windows": {
			//               "description": "\u003cp\u003eThe folder path in Windows workstations where files are uploaded.\u003c/p\u003e",
			//               "maxLength": 128,
			//               "minLength": 1,
			//               "pattern": "",
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "StreamingImageIds": {
			//       "description": "\u003cp\u003eThe streaming images that users can select from when launching a streaming session\n            with this launch profile.\u003c/p\u003e",
			//       "items": {
			//         "maxLength": 22,
			//         "minLength": 0,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "maxItems": 20,
			//       "minItems": 1,
			//       "type": "array"
			//     }
			//   },
			//   "required": [
			//     "ClipboardMode",
			//     "Ec2InstanceTypes",
			//     "StreamingImageIds"
			//   ],
			//   "type": "object"
			// }
			Description: "<p>A configuration for a streaming session.</p>",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"clipboard_mode": {
						// Property: ClipboardMode
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"ENABLED",
								"DISABLED",
							}),
						},
					},
					"ec_2_instance_types": {
						// Property: Ec2InstanceTypes
						Description: "<p>The EC2 instance types that users can select from when launching a streaming session\n            with this launch profile.</p>",
						Type:        types.ListType{ElemType: types.StringType},
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(1, 30),
							validate.ArrayForEach(validate.StringInSlice([]string{
								"g4dn.xlarge",
								"g4dn.2xlarge",
								"g4dn.4xlarge",
								"g4dn.8xlarge",
								"g4dn.12xlarge",
								"g4dn.16xlarge",
							})),
						},
					},
					"max_session_length_in_minutes": {
						// Property: MaxSessionLengthInMinutes
						Description: "<p>The length of time, in minutes, that a streaming session can be active before it is\n            stopped or terminated. After this point, Nimble Studio automatically terminates or\n            stops the session. The default length of time is 690 minutes, and the maximum length of\n            time is 30 days.</p>",
						Type:        types.NumberType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.FloatBetween(1.000000, 43200.000000),
						},
					},
					"max_stopped_session_length_in_minutes": {
						// Property: MaxStoppedSessionLengthInMinutes
						Description: "<p>Integer that determines if you can start and stop your sessions and how long a session\n            can stay in the STOPPED state. The default value is 0. The maximum value is 5760.</p>\n        <p>If the value is missing or set to 0, your sessions can’t be stopped. If you then call\n            StopStreamingSession, the session fails. If the time that a session stays in the READY\n            state exceeds the maxSessionLengthInMinutes value, the session will automatically be\n            terminated by AWS (instead of stopped).</p>\n        <p>If the value is set to a positive number, the session can be stopped. You can call\n            StopStreamingSession to stop sessions in the READY state. If the time that a session\n            stays in the READY state exceeds the maxSessionLengthInMinutes value, the session will\n            automatically be stopped by AWS (instead of terminated).</p>",
						Type:        types.NumberType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.FloatBetween(0.000000, 5760.000000),
						},
					},
					"session_storage": {
						// Property: SessionStorage
						Description: "<p>The configuration for a streaming session’s upload storage.</p>",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"mode": {
									// Property: Mode
									Description: "<p>Allows artists to upload files to their workstations. The only valid option is\n                <code>UPLOAD</code>.</p>",
									Type:        types.ListType{ElemType: types.StringType},
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenAtLeast(1),
										validate.ArrayForEach(validate.StringInSlice([]string{
											"UPLOAD",
										})),
									},
								},
								"root": {
									// Property: Root
									Description: "<p>The upload storage root location (folder) on streaming workstations where files are\n            uploaded.</p>",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"linux": {
												// Property: Linux
												Description: "<p>The folder path in Linux workstations where files are uploaded.</p>",
												Type:        types.StringType,
												Optional:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 128),
												},
											},
											"windows": {
												// Property: Windows
												Description: "<p>The folder path in Windows workstations where files are uploaded.</p>",
												Type:        types.StringType,
												Optional:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 128),
												},
											},
										},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"streaming_image_ids": {
						// Property: StreamingImageIds
						Description: "<p>The streaming images that users can select from when launching a streaming session\n            with this launch profile.</p>",
						Type:        types.ListType{ElemType: types.StringType},
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(1, 20),
							validate.ArrayForEach(validate.StringLenBetween(0, 22)),
						},
					},
				},
			),
			Required: true,
		},
		"studio_component_ids": {
			// Property: StudioComponentIds
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eUnique identifiers for a collection of studio components that can be used with this\n            launch profile.\u003c/p\u003e",
			//   "items": {
			//     "type": "string"
			//   },
			//   "maxItems": 100,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "<p>Unique identifiers for a collection of studio components that can be used with this\n            launch profile.</p>",
			Type:        types.ListType{ElemType: types.StringType},
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(1, 100),
			},
		},
		"studio_id": {
			// Property: StudioId
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe studio ID. \u003c/p\u003e",
			//   "type": "string"
			// }
			Description: "<p>The studio ID. </p>",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Represents a launch profile which delegates access to a collection of studio components to studio users",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NimbleStudio::LaunchProfile").WithTerraformTypeName("awscc_nimblestudio_launch_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"clipboard_mode":                        "ClipboardMode",
		"description":                           "Description",
		"ec_2_instance_types":                   "Ec2InstanceTypes",
		"ec_2_subnet_ids":                       "Ec2SubnetIds",
		"launch_profile_id":                     "LaunchProfileId",
		"launch_profile_protocol_versions":      "LaunchProfileProtocolVersions",
		"linux":                                 "Linux",
		"max_session_length_in_minutes":         "MaxSessionLengthInMinutes",
		"max_stopped_session_length_in_minutes": "MaxStoppedSessionLengthInMinutes",
		"mode":                                  "Mode",
		"name":                                  "Name",
		"root":                                  "Root",
		"session_storage":                       "SessionStorage",
		"stream_configuration":                  "StreamConfiguration",
		"streaming_image_ids":                   "StreamingImageIds",
		"studio_component_ids":                  "StudioComponentIds",
		"studio_id":                             "StudioId",
		"tags":                                  "Tags",
		"windows":                               "Windows",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}

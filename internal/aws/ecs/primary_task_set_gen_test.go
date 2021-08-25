// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecs_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSECSPrimaryTaskSet_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ECS::PrimaryTaskSet", "awscc_ecs_primary_task_set", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

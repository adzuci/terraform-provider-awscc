// Code generated by generators/resource/main.go; DO NOT EDIT.

package dynamodb_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSDynamoDBGlobalTable_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::DynamoDB::GlobalTable", "aws_dynamodb_global_table", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
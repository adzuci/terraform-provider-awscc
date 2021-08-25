// Code generated by generators/resource/main.go; DO NOT EDIT.

package codegurureviewer_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSCodeGuruReviewerRepositoryAssociation_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CodeGuruReviewer::RepositoryAssociation", "awscc_codegurureviewer_repository_association", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

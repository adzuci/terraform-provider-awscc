// Code generated by generators/resource/main.go; DO NOT EDIT.

package cassandra_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSCassandraKeyspace_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Cassandra::Keyspace", "awscc_cassandra_keyspace", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
	})
}

func TestAccAWSCassandraKeyspace_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Cassandra::Keyspace", "awscc_cassandra_keyspace", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				td.DeleteResource(),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}

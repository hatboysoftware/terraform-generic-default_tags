//go:build integration

package test

import (
	"github.com/gruntwork-io/terratest/modules/azure"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

const (
	fixtures = "./fixtures"
)

type TestCondition int

const (
	TestConditionEquals   TestCondition = 0
	TestConditionNotEmpty TestCondition = 1
)

func TestResourceGroup(t *testing.T) {
	terraform.InitAndApply(t, IntegrationTestOptions())
	defer terraform.Destroy(t, IntegrationTestOptions())

	t.Run("Output Validation", OutputValidation)
	t.Run("Resource Group Validation", ResourceGroupValidation)
}

func OutputValidation(t *testing.T) {
	testCases := []struct {
		Name      string
		Got       string
		Want      string
		Condition TestCondition
	}{
		{"Resource Group Name", terraform.Output(t, IntegrationTestOptions(), "resource_group_name"), "EUS2-HBS-TST-TEST-rg", TestConditionEquals},
		{"Tag Name", terraform.Output(t, IntegrationTestOptions(), "tag_name"), "foo", TestConditionEquals},
		{"Tag Value", terraform.Output(t, IntegrationTestOptions(), "tag_value"), "bar", TestConditionEquals},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			switch tc.Condition {
			case TestConditionEquals:
				assert.Equal(t, tc.Want, tc.Got)
			case TestConditionNotEmpty:
				assert.NotEmpty(t, tc.Got)
			}
		})
	}
}

func ResourceGroupValidation(t *testing.T) {
	name := terraform.Output(t, IntegrationTestOptions(), "resource_group_name")
	subscriptionID := terraform.Output(t, IntegrationTestOptions(), "subscription_id")
	resource_group := azure.GetAResourceGroup(t, name, subscriptionID)

	t.Log(name)
	t.Log(subscriptionID)
	t.Log(resource_group)

	t.Run("Resource Group Exists", func(t *testing.T) {
		assert.Equal(t, *resource_group.Name, name)
	})
}

func IntegrationTestOptions() *terraform.Options {
	return &terraform.Options{
		TerraformDir: fixtures,
		NoColor:      true,
	}
}

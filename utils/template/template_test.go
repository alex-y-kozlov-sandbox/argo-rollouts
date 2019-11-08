package query

import (
	"fmt"
	"testing"

	"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1"

	"github.com/stretchr/testify/assert"
)

func TestResolveArgsWithNoSubstitution(t *testing.T) {
	query, err := ResolveArgs("test", nil)
	assert.Nil(t, err)
	assert.Equal(t, "test", query)
}

func TestResolveArgsRemoveWhiteSpace(t *testing.T) {
	args := []v1alpha1.Argument{{
		Name:  "var",
		Value: "foo",
	}}
	query, err := ResolveArgs("test-{{ inputs.var }}", args)
	assert.Nil(t, err)
	assert.Equal(t, "test-foo", query)
}

func TestResolveArgsWithSubstitution(t *testing.T) {
	args := []v1alpha1.Argument{{
		Name:  "var",
		Value: "foo",
	}}
	query, err := ResolveArgs("test-{{inputs.var}}", args)
	assert.Nil(t, err)
	assert.Equal(t, "test-foo", query)
}

func TestInvalidTemplate(t *testing.T) {
	_, err := ResolveArgs("test-{{inputs.var", nil)
	assert.Equal(t, fmt.Errorf("Cannot find end tag=\"}}\" in the template=\"test-{{inputs.var\" starting from \"inputs.var\""), err)
}

func TestMissingArgs(t *testing.T) {
	_, err := ResolveArgs("test-{{inputs.var}}", nil)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("failed to resolve {{inputs.var}}"), err)
}
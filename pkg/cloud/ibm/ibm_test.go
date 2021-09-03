package ibm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResources(t *testing.T) {
	resources := GetResources()
	assert.Len(t, resources, 1)

	var names, kinds []string
	for _, r := range resources {
		names = append(names, r.GetName())
		kinds = append(kinds, r.GetObjectKind().GroupVersionKind().Kind)
	}

	assert.Contains(t, names, "ibm-cloud-controller-manager")
	assert.Contains(t, kinds, "Deployment")
}

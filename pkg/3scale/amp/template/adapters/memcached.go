package adapters

import (
	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
	"github.com/3scale/3scale-operator/pkg/common"
	templatev1 "github.com/openshift/api/template/v1"
)

type MemcachedAdapter struct {
}

func NewMemcachedAdapter() Adapter {
	return NewAppenderAdapter(&MemcachedAdapter{})
}

func (m *MemcachedAdapter) Parameters() []templatev1.Parameter {
	return []templatev1.Parameter{}
}

func (m *MemcachedAdapter) Objects() ([]common.KubernetesObject, error) {
	memcachedOptions, err := m.options()
	if err != nil {
		return nil, err
	}
	memcachedComponent := component.NewMemcached(memcachedOptions)
	objects := m.componentObjects(memcachedComponent)
	return objects, nil
}

func (m *MemcachedAdapter) componentObjects(c *component.Memcached) []common.KubernetesObject {
	deploymentConfig := c.DeploymentConfig()

	objects := []common.KubernetesObject{
		deploymentConfig,
	}
	return objects
}

func (m *MemcachedAdapter) options() (*component.MemcachedOptions, error) {
	mo := component.NewMemcachedOptions()
	mo.AppLabel = "${APP_LABEL}"
	mo.ImageTag = "${AMP_RELEASE}"
	mo.ResourceRequirements = component.DefaultMemcachedResourceRequirements()

	err := mo.Validate()
	return mo, err
}

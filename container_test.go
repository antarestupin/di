package di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestService struct {
	alreadyInstanciated bool
}

type TestServiceWithDependency struct {
	dependency *TestService
}

func Test_NewContainer(t *testing.T) {
	NewContainer()
}

func Test_services_can_be_injected(t *testing.T) {
	container := NewContainer()
	container.Add("test", func() interface{} { return &TestService{} })
	assert.NotNil(t, container.Get("test"))
}

func Test_a_service_cannot_be_injected_twice(t *testing.T) {
	container := NewContainer()

	assert.Panics(t, func() {
		container.Add("test", func() interface{} { return &TestService{} })
		container.Add("test", func() interface{} { return &TestService{} })
	})
}

func Test_a_service_can_be_redefined(t *testing.T) {
	container := NewContainer()

	container.Add("test", func() interface{} { return &TestService{} })
	container.Set("test", func() interface{} { return &TestService{true} })

	assert.Equal(t, true, container.Get("test").(*TestService).alreadyInstanciated)
}

func Test_container_returns_nil_if_service_does_not_exist(t *testing.T) {
	container := NewContainer()
	assert.Nil(t, container.Get("test"))
}

func Test_service_is_instanciated_only_once(t *testing.T) {
	container := NewContainer()
	container.Add("test", func() interface{} { return &TestService{} })

	container.Get("test").(*TestService).alreadyInstanciated = true

	assert.Equal(t, true, container.Get("test").(*TestService).alreadyInstanciated)
}

func Test_service_with_dependencies(t *testing.T) {
	container := NewContainer()
	container.Add("dependency", func() interface{} { return &TestService{} })
	container.Add("service", func() interface{} {
		return &TestServiceWithDependency{dependency: container.Get("dependency").(*TestService)}
	})
	container.Get("service").(*TestServiceWithDependency).dependency.alreadyInstanciated = true

	assert.Equal(t, true, container.Get("service").(*TestServiceWithDependency).dependency.alreadyInstanciated)
}

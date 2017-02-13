package di

import "fmt"

// Container service
type Container struct {
	servicesInstantiators map[string]serviceInstantiator
	services              map[string]interface{}
}

// Function that will be called to instantiate a service
type serviceInstantiator func() interface{}

// NewContainer constructor
func NewContainer() *Container {
	return &Container{
		servicesInstantiators: make(map[string]serviceInstantiator),
		services:              make(map[string]interface{}),
	}
}

// Add a service definition to the container
func (c *Container) Add(id string, instantiator serviceInstantiator) *Container {
	if _, ok := c.servicesInstantiators[id]; ok {
		panic(fmt.Errorf("Service %s already exists in the container", id))
	}

	c.servicesInstantiators[id] = instantiator

	return c
}

// Set a service definition to the container
func (c *Container) Set(id string, instantiator serviceInstantiator) *Container {
	c.servicesInstantiators[id] = instantiator
	return c
}

// Get a service with given id
func (c *Container) Get(id string) interface{} {
	if service, ok := c.services[id]; ok {
		return service
	}

	if instantiator, ok := c.servicesInstantiators[id]; ok {
		service := instantiator()
		c.services[id] = service
		return service
	}

	return nil
}

package ioc

import (
	"fmt"
	"reflect"
)

// Container is our IoC Container
type Container struct {
	cnt map[string]interface{}
}

// CreateContainer returns new instance of Container
func CreateContainer() *Container {
	return &Container{
		cnt: make(map[string]interface{}),
	}
}

func verifyTypeImplementsInterface(instance interface{}, intfType reflect.Type) error {
	instanceType := reflect.TypeOf(instance)

	if !reflect.PtrTo(instanceType).Elem().Implements(intfType) {
		return fmt.Errorf("%s does not implements %s", instanceType.Name(), intfType.Name())
	}

	return nil
}

func (c *Container) bindInterface(label string, instance interface{}, intfType reflect.Type) error {
	if err := verifyTypeImplementsInterface(instance, intfType); err != nil {
		return err
	}

	c.cnt[label] = instance
	return nil
}

// Bind binds instance to container using its struct name
func (c *Container) Bind(instance interface{}) {
	instanceType := reflect.TypeOf(instance).Elem()
	label := fmt.Sprintf("%s.%s", instanceType.PkgPath(), instanceType.Name())
	c.cnt[label] = instance
}

// BindInterface binds instance that implements interface to container map
func (c *Container) BindInterface(instance, intf interface{}) error {
	intfType := reflect.TypeOf(intf).Elem()
	label := fmt.Sprintf("%s.%s", intfType.PkgPath(), intfType.Name())

	return c.bindInterface(label, instance, intfType)
}

// BindInterfaceWithName binds instance that implements interface to container map with tag name
func (c *Container) BindInterfaceWithName(name string, instance, intf interface{}) error {
	intfType := reflect.TypeOf(intf).Elem()
	return c.bindInterface(name, instance, intfType)
}

func (c *Container) get(name string) interface{} {
	if instance, ok := c.cnt[name]; ok {
		return instance
	}

	return nil
}

// Get fetches instance from container map using type
func (c *Container) Get(intf interface{}) interface{} {
	intfType := reflect.TypeOf(intf).Elem()
	label := fmt.Sprintf("%s.%s", intfType.PkgPath(), intfType.Name())

	return c.get(label)
}

// GetWithName fetches instance from container map using name
func (c *Container) GetWithName(label string) interface{} {
	return c.get(label)
}

# di: Simple Go DI container

This package provides a very simple container to use for adding dependency injection in your app.

## Install

```
go get "github.com/antarestupin/di"
```

## How to use

Here is a simple example of how to use this container.

Assuming you have the following services:

```go
type Service1 struct {
    dependency *Service2
}

type Service2 struct {}
```

You can define them in the container this way:

```go
container := di.NewContainer()

container.Add("service1", func() interface{} {
    return &Service1{ container.Get("service2").(*Service2) }
})

container.Add("service2", func() interface{} {
    return &Service2{}
})
```

And access them this way:

```go
service1 := container.Get("service1").(*Service1)
```

## Container methods

Here are the methods provided in this library:

* `di.NewContainer() *Container`: Gives a new instance of the container
* `(c *Container) Add(id string, instantiator func() interface{}) *Container`: Adds a service in the container; panics if the service already exists
* `(c *Container) Set(id string, instantiator func() interface{}) *Container`: Sets a service in the container; the service is replaced if it already exists
* `(c *Container) Get(id string) interface{}`: Gives the service linked to given `id`

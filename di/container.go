package di

import (
	"fmt"
	"reflect"
)

var container = map[any]map[string]any{}

const (
	Singleton = iota
	Transient
)

func Add[T any, K comparable](key K, lifetime int, factory func() T) {
	if lifetime == Singleton {
		container[key] = map[string]any{
			"lt":    lifetime,
			"value": factory(),
		}
	}

	if lifetime == Transient {
		container[key] = map[string]any{
			"lt":    lifetime,
			"value": factory,
		}
	}
}

func Resolve[T any, K comparable](key K) T {
	if _, ok := container[key]; !ok {
		panic("dependency '" + getDepName(key) + "' not found in container")
	}

	if container[key]["lt"] == Singleton {
		return container[key]["value"].(T)
	}

	if container[key]["lt"] == Transient {
		return container[key]["value"].(func() T)()
	}

	var zero T

	return zero
}

func getDepName(key any) string {
	switch reflect.TypeOf(key).Kind() {
	case reflect.String:
		return key.(string)

	case reflect.Int:
		return fmt.Sprintf("%d", key.(int))

	case reflect.Float32:
		return fmt.Sprintf("%f", key.(float32))

	case reflect.Float64:
		return fmt.Sprintf("%f", key.(float64))

	// If key is a pointer, get the name by dereferencing it until it's not a pointer
	case reflect.Ptr:
		currPtr := reflect.TypeOf(key).Elem()

		for currPtr.Kind() == reflect.Ptr {
			currPtr = currPtr.Elem()
		}

		if currPtr.Name() == "" {
			return currPtr.Kind().String()
		}

		return currPtr.Name()

	default:
		return reflect.TypeOf(key).Name()
	}
}

package protoproducer

import "fmt"

// ErrFieldNotFound indicates a field specified in the configuration was not found in the protobuf message.
type ErrFieldNotFound struct {
	Field string
}

func (e *ErrFieldNotFound) Error() string {
	return fmt.Sprintf("field %s in config not found in protobuf", e.Field)
}

// ErrRendererNotFound indicates a renderer specified in the configuration was not found.
type ErrRendererNotFound struct {
	Renderer string
}

func (e *ErrRendererNotFound) Error() string {
	return fmt.Sprintf("renderer %s not found", e.Renderer)
}

// ErrKeyFieldNotFound indicates a key field specified in the configuration does not exist.
type ErrKeyFieldNotFound struct {
	Field string
}

func (e *ErrKeyFieldNotFound) Error() string {
	return fmt.Sprintf("key field %s does not exist", e.Field)
}

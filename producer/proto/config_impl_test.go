package protoproducer

import (
	"errors"
	"testing"
)

func TestMapFormatErrors(t *testing.T) {
	t.Run("ErrKeyFieldNotFound", func(t *testing.T) {
		cfg := &ProducerConfig{
			Formatter: FormatterConfig{
				Key: []string{"nonexistent"},
			},
		}
		_, err := mapFormat(cfg)
		if err == nil {
			t.Fatal("expected error")
		}
		var keyErr *ErrKeyFieldNotFound
		if !errors.As(err, &keyErr) {
			t.Errorf("expected ErrKeyFieldNotFound, got %T", err)
		}
		if keyErr.Field != "nonexistent" {
			t.Errorf("expected field nonexistent, got %s", keyErr.Field)
		}
	})

	t.Run("ErrRendererNotFound", func(t *testing.T) {
		cfg := &ProducerConfig{
			Formatter: FormatterConfig{
				Render: map[string]RendererID{
					"SrcAddr": "nonexistent_renderer",
				},
			},
		}
		_, err := mapFormat(cfg)
		if err == nil {
			t.Fatal("expected error")
		}
		var rendererErr *ErrRendererNotFound
		if !errors.As(err, &rendererErr) {
			t.Errorf("expected ErrRendererNotFound, got %T", err)
		}
		if rendererErr.Renderer != "nonexistent_renderer" {
			t.Errorf("expected renderer nonexistent_renderer, got %s", rendererErr.Renderer)
		}
	})

	t.Run("ErrFieldNotFound", func(t *testing.T) {
		cfg := &ProducerConfig{
			Formatter: FormatterConfig{
				Fields: []string{"nonexistent_field"},
			},
		}
		_, err := mapFormat(cfg)
		if err == nil {
			t.Fatal("expected error")
		}
		var fieldErr *ErrFieldNotFound
		if !errors.As(err, &fieldErr) {
			t.Errorf("expected ErrFieldNotFound, got %T", err)
		}
		if fieldErr.Field != "nonexistent_field" {
			t.Errorf("expected field nonexistent_field, got %s", fieldErr.Field)
		}
	})
}

package exporter_test

import (
	"github.com/kercylan98/minotaur/configuration/exporter"
	"github.com/kercylan98/minotaur/configuration/raw"
	"os"
	"path/filepath"
	"testing"
)

func TestNewJSON(t *testing.T) {
	t.Run("TestNewJSON", func(t *testing.T) {
		e := exporter.NewJSON("")
		if e == nil {
			t.Fatal("NewJSON failed")
		}
	})

	t.Run("TestExport_WritePath", func(t *testing.T) {
		e := exporter.NewJSON(filepath.Join(os.TempDir(), "test.json"))
		if e == nil {
			t.Fatal("NewJSON failed")
		}
	})
}

func TestJSON_Export(t *testing.T) {
	t.Run("TestJSON_Export", func(t *testing.T) {
		e := exporter.NewJSON("")
		if e == nil {
			t.Fatal("NewJSON failed")
		}

		err := e.Export(raw.Config{}, nil)
		if err != nil {
			t.Fatal("Export failed")
		}
	})

	t.Run("TestJSON_Export_WritePath", func(t *testing.T) {
		e := exporter.NewJSON(filepath.Join(os.TempDir(), "test.json"))
		if e == nil {
			t.Fatal("NewJSON failed")
		}

		err := e.Export(raw.Config{}, nil)
		if err != nil {
			t.Fatal("Export failed")
		}
	})
}

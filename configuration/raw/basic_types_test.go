package raw_test

import (
	"github.com/kercylan98/minotaur/configuration/raw"
	"testing"
)

func TestIsBasicType(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		t.Parallel()

		if !raw.IsBasicType(raw.FieldTypeInt) {
			t.Fatalf("expected %s to be a basic type", raw.FieldTypeInt)
		}
	})

	t.Run("int8", func(t *testing.T) {
		t.Parallel()

		if !raw.IsBasicType(raw.FieldTypeInt8) {
			t.Fatalf("expected %s to be a basic type", raw.FieldTypeInt8)
		}
	})

	t.Run("int16", func(t *testing.T) {
		t.Parallel()

		if !raw.IsBasicType(raw.FieldTypeInt16) {
			t.Fatalf("expected %s to be a basic type", raw.FieldTypeInt16)
		}
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()

		if !raw.IsBasicType(raw.FieldTypeInt32) {
			t.Fatalf("expected %s to be a basic type", raw.FieldTypeInt32)
		}
	})

	t.Run("int64", func(t *testing.T) {
		t.Parallel()

		if !raw.IsBasicType(raw.FieldTypeInt64) {
			t.Fatalf("expected %s to be a basic type", raw.FieldTypeInt64)
		}
	})

	t.Run("uint", func(t *testing.T) {
		t.Parallel()

		if !raw.IsBasicType(raw.FieldTypeUint) {
			t.Fatalf("expected %s to be a basic type", raw.FieldTypeUint)
		}
	})

	t.Run("uint8", func(t *testing.T) {
		t.Parallel()

		if !raw.IsBasicType(raw.FieldTypeUint8) {
			t.Fatalf("expected %s to be a basic type", raw.FieldTypeUint8)
		}
	})

	t.Run("uint16", func(t *testing.T) {
		t.Parallel()

		if !raw.IsBasicType(raw.FieldTypeUint16) {
			t.Fatalf("expected %s to be a basic type", raw.FieldTypeUint16)
		}
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()

		if !raw.IsBasicType(raw.FieldTypeUint32) {
			t.Fatalf("expected %s to be a basic type", raw.FieldTypeUint32)
		}
	})

	t.Run("uint64", func(t *testing.T) {
		t.Parallel()

		if !raw.IsBasicType(raw.FieldTypeUint64) {
			t.Fatalf("expected %s to be a basic type", raw.FieldTypeUint64)
		}
	})
}

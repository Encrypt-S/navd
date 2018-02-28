// Copyright (c) 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package navjson_test

import (
	"testing"

	"github.com/aguycalled/navd/navjson"
)

// TestErrorCodeStringer tests the stringized output for the ErrorCode type.
func TestErrorCodeStringer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   navjson.ErrorCode
		want string
	}{
		{navjson.ErrDuplicateMethod, "ErrDuplicateMethod"},
		{navjson.ErrInvalidUsageFlags, "ErrInvalidUsageFlags"},
		{navjson.ErrInvalidType, "ErrInvalidType"},
		{navjson.ErrEmbeddedType, "ErrEmbeddedType"},
		{navjson.ErrUnexportedField, "ErrUnexportedField"},
		{navjson.ErrUnsupportedFieldType, "ErrUnsupportedFieldType"},
		{navjson.ErrNonOptionalField, "ErrNonOptionalField"},
		{navjson.ErrNonOptionalDefault, "ErrNonOptionalDefault"},
		{navjson.ErrMismatchedDefault, "ErrMismatchedDefault"},
		{navjson.ErrUnregisteredMethod, "ErrUnregisteredMethod"},
		{navjson.ErrNumParams, "ErrNumParams"},
		{navjson.ErrMissingDescription, "ErrMissingDescription"},
		{0xffff, "Unknown ErrorCode (65535)"},
	}

	// Detect additional error codes that don't have the stringer added.
	if len(tests)-1 != int(navjson.TstNumErrorCodes) {
		t.Errorf("It appears an error code was added without adding an " +
			"associated stringer test")
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

// TestError tests the error output for the Error type.
func TestError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   navjson.Error
		want string
	}{
		{
			navjson.Error{Description: "some error"},
			"some error",
		},
		{
			navjson.Error{Description: "human-readable error"},
			"human-readable error",
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.Error()
		if result != test.want {
			t.Errorf("Error #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

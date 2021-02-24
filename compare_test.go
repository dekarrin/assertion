package assertion

import "testing"

func Test_checkEqual_untyped_nil(t *testing.T) {
	var nilStringPtr *string

	testCases := []struct {
		name   string
		other  interface{}
		expect bool
	}{
		{"to untyped nil", nil, true},
		{"to typed nil", nilStringPtr, false},
		{"to empty", "", false},
		{"to non-empty", "a", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := checkEqual(nil, tc.other, nil)
			if err != nil {
				t.Fatalf("got an error when checking Equality: %v", err)
			}

			if actual != tc.expect {
				t.Fatalf("expected %v but was %v", tc.expect, actual)
			}
		})
	}
}

func Test_checkEqual_typed_nil(t *testing.T) {
	var nilStringPtr *string

	testCases := []struct {
		name   string
		other  interface{}
		expect bool
	}{
		{"to untyped nil", nil, true},
		{"to typed nil", nilStringPtr, true},
		{"to empty", "", false},
		{"to non-empty", "a", false},
	}

	var typedPtr *string
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := checkEqual(typedPtr, tc.other, nil)
			if err != nil {
				t.Fatalf("got an error when checking Equality: %v", err)
			}

			if actual != tc.expect {
				t.Fatalf("expected %v but was %v", tc.expect, actual)
			}
		})
	}
}

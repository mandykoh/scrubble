package scrubble

import "testing"

func TestInvalidWordReason(t *testing.T) {

	t.Run(".GoString()", func(t *testing.T) {

		t.Run("returns Go syntax for valid reasons", func(t *testing.T) {
			cases := []struct {
				Reason       InvalidWordReason
				ExpectedName string
			}{
				{SingleLetterWordDisallowedReason, "SingleLetterWordDisallowedReason"},
				{UnknownInvalidWordReason, "UnknownInvalidWordReason"},
			}

			for _, c := range cases {
				if actual, expected := c.Reason.GoString(), c.ExpectedName; actual != expected {
					t.Errorf("Expected reason '%s' but got '%s'", expected, actual)
				}
			}
		})

		t.Run("returns UnknownInvalidWordReason for invalid reasons", func(t *testing.T) {
			cases := []InvalidWordReason{999, -1}

			for _, c := range cases {
				if actual, expected := c.GoString(), "UnknownInvalidWordReason"; actual != expected {
					t.Errorf("Expected invalid reason but got '%s'", actual)
				}
			}
		})
	})

	t.Run(".String()", func(t *testing.T) {

		t.Run("returns name of valid reasons", func(t *testing.T) {
			cases := []struct {
				Reason       InvalidWordReason
				ExpectedName string
			}{
				{SingleLetterWordDisallowedReason, "SingleLetterWordDisallowed"},
			}

			for _, c := range cases {
				if actual, expected := c.Reason.String(), c.ExpectedName; actual != expected {
					t.Errorf("Expected reason '%s' but got '%s'", expected, actual)
				}
			}
		})

		t.Run("returns 'Unknown' for invalid reasons", func(t *testing.T) {
			cases := []InvalidWordReason{999, -1}

			for _, c := range cases {
				if actual, expected := c.String(), "Unknown"; actual != expected {
					t.Errorf("Expected invalid reason but got '%s'", actual)
				}
			}
		})
	})
}

package scrubble

import "testing"

func TestGamePhase(t *testing.T) {

	t.Run(".GoString()", func(t *testing.T) {

		t.Run("returns Go syntax for valid game phases", func(t *testing.T) {
			cases := []struct {
				Phase        GamePhase
				ExpectedName string
			}{
				{SetupPhase, "SetupPhase"},
				{MainPhase, "MainPhase"},
				{EndPhase, "EndPhase"},
				{UnknownPhase, "UnknownPhase"},
			}

			for _, c := range cases {
				if actual, expected := c.Phase.GoString(), c.ExpectedName; actual != expected {
					t.Errorf("Expected game phase '%s' but got '%s'", expected, actual)
				}
			}
		})

		t.Run("returns UnknownPhase for invalid game phases", func(t *testing.T) {
			cases := []GamePhase{999, -1}

			for _, c := range cases {
				if actual, expected := c.GoString(), "UnknownPhase"; actual != expected {
					t.Errorf("Expected invalid game phase but got '%s'", actual)
				}
			}
		})
	})

	t.Run(".String()", func(t *testing.T) {

		t.Run("returns name of valid game phases", func(t *testing.T) {
			cases := []struct {
				Phase        GamePhase
				ExpectedName string
			}{
				{SetupPhase, "Setup"},
				{MainPhase, "Main"},
				{EndPhase, "End"},
			}

			for _, c := range cases {
				if actual, expected := c.Phase.String(), c.ExpectedName; actual != expected {
					t.Errorf("Expected game phase '%s' but got '%s'", expected, actual)
				}
			}
		})

		t.Run("returns 'Unknown' for invalid game phases", func(t *testing.T) {
			cases := []GamePhase{999, -1}

			for _, c := range cases {
				if actual, expected := c.String(), "Unknown"; actual != expected {
					t.Errorf("Expected invalid game phase but got '%s'", actual)
				}
			}
		})
	})
}

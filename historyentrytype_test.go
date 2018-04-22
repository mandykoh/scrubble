package scrubble

import "testing"

func TestHistoryEntryType(t *testing.T) {

	t.Run(".GoString()", func(t *testing.T) {

		t.Run("returns Go syntax for valid entry types", func(t *testing.T) {
			cases := []struct {
				Type         HistoryEntryType
				ExpectedName string
			}{
				{PlayHistoryEntryType, "PlayHistoryEntryType"},
				{PassHistoryEntryType, "PassHistoryEntryType"},
				{ExchangeTilesHistoryEntryType, "ExchangeTilesHistoryEntryType"},
				{ChallengeFailHistoryEntryType, "ChallengeFailHistoryEntryType"},
				{ChallengeSuccessHistoryEntryType, "ChallengeSuccessHistoryEntryType"},
				{UnknownHistoryEntryType, "UnknownHistoryEntryType"},
			}

			for _, c := range cases {
				if actual, expected := c.Type.GoString(), c.ExpectedName; actual != expected {
					t.Errorf("Expected type '%s' but got '%s'", expected, actual)
				}
			}
		})

		t.Run("returns UnknownHistoryEntryType for invalid types", func(t *testing.T) {
			cases := []HistoryEntryType{999, -1}

			for _, c := range cases {
				if actual, expected := c.GoString(), "UnknownHistoryEntryType"; actual != expected {
					t.Errorf("Expected invalid reason but got '%s'", actual)
				}
			}
		})
	})

	t.Run(".String()", func(t *testing.T) {

		t.Run("returns name of valid types", func(t *testing.T) {
			cases := []struct {
				Type         HistoryEntryType
				ExpectedName string
			}{
				{PlayHistoryEntryType, "Play"},
				{PassHistoryEntryType, "Pass"},
				{ExchangeTilesHistoryEntryType, "ExchangeTiles"},
				{ChallengeFailHistoryEntryType, "ChallengeFail"},
				{ChallengeSuccessHistoryEntryType, "ChallengeSuccess"},
			}

			for _, c := range cases {
				if actual, expected := c.Type.String(), c.ExpectedName; actual != expected {
					t.Errorf("Expected type '%s' but got '%s'", expected, actual)
				}
			}
		})

		t.Run("returns 'Unknown' for invalid type", func(t *testing.T) {
			cases := []HistoryEntryType{999, -1}

			for _, c := range cases {
				if actual, expected := c.String(), "Unknown"; actual != expected {
					t.Errorf("Expected invalid type but got '%s'", actual)
				}
			}
		})
	})
}

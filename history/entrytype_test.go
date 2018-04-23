package history

import "testing"

func TestEntryType(t *testing.T) {

	t.Run(".GoString()", func(t *testing.T) {

		t.Run("returns Go syntax for valid entry types", func(t *testing.T) {
			cases := []struct {
				Type         EntryType
				ExpectedName string
			}{
				{PlayEntryType, "PlayEntryType"},
				{PassEntryType, "PassEntryType"},
				{ExchangeTilesEntryType, "ExchangeTilesEntryType"},
				{ChallengeFailEntryType, "ChallengeFailEntryType"},
				{ChallengeSuccessEntryType, "ChallengeSuccessEntryType"},
				{UnknownEntryType, "UnknownEntryType"},
			}

			for _, c := range cases {
				if actual, expected := c.Type.GoString(), c.ExpectedName; actual != expected {
					t.Errorf("Expected type '%s' but got '%s'", expected, actual)
				}
			}
		})

		t.Run("returns UnknownEntryType for invalid types", func(t *testing.T) {
			cases := []EntryType{999, -1}

			for _, c := range cases {
				if actual, expected := c.GoString(), "UnknownEntryType"; actual != expected {
					t.Errorf("Expected invalid reason but got '%s'", actual)
				}
			}
		})
	})

	t.Run(".String()", func(t *testing.T) {

		t.Run("returns name of valid types", func(t *testing.T) {
			cases := []struct {
				Type         EntryType
				ExpectedName string
			}{
				{PlayEntryType, "Play"},
				{PassEntryType, "Pass"},
				{ExchangeTilesEntryType, "ExchangeTiles"},
				{ChallengeFailEntryType, "ChallengeFail"},
				{ChallengeSuccessEntryType, "ChallengeSuccess"},
			}

			for _, c := range cases {
				if actual, expected := c.Type.String(), c.ExpectedName; actual != expected {
					t.Errorf("Expected type '%s' but got '%s'", expected, actual)
				}
			}
		})

		t.Run("returns 'Unknown' for invalid type", func(t *testing.T) {
			cases := []EntryType{999, -1}

			for _, c := range cases {
				if actual, expected := c.String(), "Unknown"; actual != expected {
					t.Errorf("Expected invalid type but got '%s'", actual)
				}
			}
		})
	})
}

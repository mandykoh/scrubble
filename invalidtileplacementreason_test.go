package scrubble

import "testing"

func TestInvalidTilePlacementReason(t *testing.T) {

	t.Run(".GoString()", func(t *testing.T) {

		t.Run("returns Go syntax for valid reasons", func(t *testing.T) {
			cases := []struct {
				Reason       InvalidTilePlacementReason
				ExpectedName string
			}{
				{NoTilesPlacedReason, "NoTilesPlacedReason"},
				{PositionOccupiedReason, "PositionOccupiedReason"},
				{PlacementOutOfBoundsReason, "PlacementOutOfBoundsReason"},
				{PlacementOverlapReason, "PlacementOverlapReason"},
				{PlacementNotLinearReason, "PlacementNotLinearReason"},
				{PlacementNotContiguousReason, "PlacementNotContiguousReason"},
				{PlacementNotConnectedReason, "PlacementNotConnectedReason"},
				{UnknownInvalidTilePlacementReason, "UnknownInvalidTilePlacementReason"},
			}

			for _, c := range cases {
				if actual, expected := c.Reason.GoString(), c.ExpectedName; actual != expected {
					t.Errorf("Expected reason '%s' but got '%s'", expected, actual)
				}
			}
		})

		t.Run("returns UnknownInvalidTilePlacementReason for invalid reasons", func(t *testing.T) {
			cases := []InvalidTilePlacementReason{999, -1}

			for _, c := range cases {
				if actual, expected := c.GoString(), "UnknownInvalidTilePlacementReason"; actual != expected {
					t.Errorf("Expected invalid reason but got '%s'", actual)
				}
			}
		})
	})

	t.Run(".String()", func(t *testing.T) {

		t.Run("returns name of valid reasons", func(t *testing.T) {
			cases := []struct {
				Reason       InvalidTilePlacementReason
				ExpectedName string
			}{
				{NoTilesPlacedReason, "NoTilesPlaced"},
				{PositionOccupiedReason, "PositionOccupied"},
				{PlacementOutOfBoundsReason, "PlacementOutOfBounds"},
				{PlacementOverlapReason, "PlacementOverlap"},
				{PlacementNotLinearReason, "PlacementNotLinear"},
				{PlacementNotContiguousReason, "PlacementNotContiguous"},
				{PlacementNotConnectedReason, "PlacementNotConnected"},
			}

			for _, c := range cases {
				if actual, expected := c.Reason.String(), c.ExpectedName; actual != expected {
					t.Errorf("Expected reason '%s' but got '%s'", expected, actual)
				}
			}
		})

		t.Run("returns 'Unknown' for invalid reasons", func(t *testing.T) {
			cases := []InvalidTilePlacementReason{999, -1}

			for _, c := range cases {
				if actual, expected := c.String(), "Unknown"; actual != expected {
					t.Errorf("Expected invalid reason but got '%s'", actual)
				}
			}
		})
	})
}

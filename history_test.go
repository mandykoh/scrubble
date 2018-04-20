package scrubble

import "testing"

func expectHistory(t *testing.T, history History, expected ...HistoryEntry) {
	t.Helper()

	if actual, expectedLen := len(history), len(expected); actual != expectedLen {
		t.Errorf("Expected there to be %d history entries but found %d", expectedLen, actual)

	} else {
		for i, e := range expected {
			expectHistoryEntry(t, history[i], e)
		}
	}
}

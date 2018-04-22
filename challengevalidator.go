package scrubble

// ChallengeValidator represents a function which determines whether a challenge
// to a play is successful.
type ChallengeValidator func() bool

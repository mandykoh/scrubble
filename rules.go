package scrubble

// Rules represent the rules used by the game to check and validate various
// conditions for legality.
type Rules struct {
	ValidatePlacements func(placements TilePlacements, board *Board) error
}

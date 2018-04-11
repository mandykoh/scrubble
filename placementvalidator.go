package scrubble

// PlacementValidator represents a function which validates whether the
// positions of tiles being placed on a board are valid.
type PlacementValidator func(placements TilePlacements, board *Board) error

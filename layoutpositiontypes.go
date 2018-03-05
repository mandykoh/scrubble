package scrubble

var (
	normalPositionTypeInstance            = &normalPositionType{}
	startPositionTypeInstance             = &startPositionType{}
	doubleLetterScorePositionTypeInstance = &doubleLetterScorePositionType{}
	doubleWordScorePositionTypeInstance   = &doubleWordScorePositionType{}
	tripleLetterScorePositionTypeInstance = &tripleLetterScorePositionType{}
	tripleWordScorePositionTypeInstance   = &tripleWordScorePositionType{}
)

// LayoutPositionTypes returns a set of position types which can be used to
// conveniently specify BoardLayouts. The position types returned are:
// normal/empty, start, double letter score, double word score, triple letter
// score, triple word score.
//
// The same instances of the position types are always returned so they can be
// compared to each other.
//
// See BoardWithLayout for example usage.
func LayoutPositionTypes() (__, st, dl, dw, tl, tw PositionType) {
	return normalPositionTypeInstance,
		startPositionTypeInstance,
		doubleLetterScorePositionTypeInstance,
		doubleWordScorePositionTypeInstance,
		tripleLetterScorePositionTypeInstance,
		tripleWordScorePositionTypeInstance
}

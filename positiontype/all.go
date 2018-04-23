package positiontype

var (
	normalInstance            = &normal{}
	startInstance             = &start{}
	doubleLetterScoreInstance = &doubleLetterScore{}
	doubleWordScoreInstance   = &doubleWordScore{}
	tripleLetterScoreInstance = &tripleLetterScore{}
	tripleWordScoreInstance   = &tripleWordScore{}
)

// All returns a set of built in position types which can be used to
// conveniently specify BoardLayouts. The position types returned are:
// normal/empty, start, double letter score, double word score, triple letter
// score, triple word score.
//
// The same instances of the position types are always returned so they can be
// compared to each other.
//
// See BoardWithLayout for example usage.
func All() (__, st, dl, dw, tl, tw Interface) {
	return normalInstance,
		startInstance,
		doubleLetterScoreInstance,
		doubleWordScoreInstance,
		tripleLetterScoreInstance,
		tripleWordScoreInstance
}

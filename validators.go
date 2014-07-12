package sherpa

// SherpaValidator is a type alias for a function that takes a Sherpa File
// and validates some portion of the data structure, returning an error
type SherpaValidator func(s SherpaFile) error

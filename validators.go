package sherpa

// SherpaValidator is a type alias for a function that takes a Sherpa File
// and validates some portion of the data structure, returning an error
type SherpaValidator func(s SherpaFile) error

func (r SherpaRecord) Validate(valfunc SherpaValidator) error {
	err := valfunc(r.Sherpa)
	if err != nil {
		return err
	}
	return nil
}

func ValidName(s SherpaFile) error {
	if (s.Name == nil) || (s.Name == "") {
		return SherpaFileError{
			Type:         "Field failed validation",
			Code:         102,
			Msg:          "SherpaFile name does not exist.",
			Field:        "name",
			DeveloperMsg: "Sherpa file does not contain a field `name`. Evaluated to empty string or nil.",
			CliHelpMsg:   "Sherpa file does not contain a field `name`.",
		}
	}
}

func ValidVersion(s SherpaFile) error {

}

package sherpa

import "encoding/json"
import "io/ioutil"
import "fmt"

// SherpaFile represents a single JSON record read from a Sherpa file
type SherpaFile struct {
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Version      string       `json:"version"`
	Main         []string     `json:"main"`
	License      string       `json:"license"`
	Ignore       []string     `json:"ignore"`
	Keywords     []string     `json:"keywords"`
	Authors      []Author     `json:"authors"`
	Homepage     string       `json:"homepage"`
	Repository   Repository   `json:"repository"`
	Dependencies []Dependency `json:"dependencies"`
	OsVersions   []string     `json:"osVersions"`
	Private      bool         `json:"private"`
	PrivateData  Privatedata  `json:"privateData"`
}

type Author struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Homepage string `json:"homepage"`
}

type Repository struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type Privatedata struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Url  string `json:"url"`
	File string `json:"file"`
}

type Dependency struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Url  string `json:"url"`
}

// SherpaMeta includes metadata about the results of validating a SherpaFile
// record in an array of SherpaFileErrors
type SherpaMetadata struct {
	Errors         []SherpaFileError
	OnDiskLocation string
}

// SherpaRecord is a container for a single SherpaFile and associated
// meta data
type SherpaRecord struct {
	Sherpa *SherpaFile
	Meta   *SherpaMetadata
}

// SherpaValidator is a type alias for a function that takes a Sherpa File
// and validates some portion of the data structure, returning an error
type SherpaValidator func(s SherpaFile) error

// SherpaFileMap is a data structure that associates a SherpaFile name with
// the SherpaFile and associated metadata.  Methods on this data structure
// check errors at the dependency graph level, e.g., missing or circular dependencies
type SherpaFileMap map[string]*SherpaRecord

// SherpaFileError returns different error types for further mitigation:
// Code  |       Type               |   Meaning
// ------|--------------------------|------------------------------------------------
// 100   | Missing Field            | A required Sherpa field is missing
// 101   | Wrong Field Type         | A field is of the wrong type
// 102   | Field failed validation  | Field did not pass a validation check
// 400   | Unable to read file      | File is missing or unable to be opened
type SherpaFileError struct {
	Type         string
	Code         int32
	Msg          string
	Field        string
	DeveloperMsg string
	CliHelpMsg   string
}

func (s SherpaFileError) Error() string {
	return s.Msg
}

func SherpaRecordFromFile(filename string) (SherpaRecord, error) {

	json, err := ioutil.ReadFile(filename)
	if err != nil {
		return SherpaRecord{}, SherpaFileError{
			Type:         "Unable to read file",
			Code:         400,
			Msg:          fmt.Sprintf("Unable to read file: %s", filename),
			Field:        "",
			DeveloperMsg: fmt.Sprintf("Failed to open file: %s", filename),
			CliHelpMsg:   fmt.Sprintf("Failed to open Sherpa file: %s. Try running the last Sherpa command again.", filename)}
	}

	meta := SherpaMetadata{OnDiskLocation: filename}
	var sherpafile SherpaFile
	record := SherpaRecord{&sherpafile, &meta}

	err = SherpaRecordFromJSON(json, record)
	if err != nil {
		return SherpaRecord{}, err
	}

	return record, nil

}

func SherpaRecordFromJSON(jsonBlob []byte, record SherpaRecord) error {

	err := json.Unmarshal(jsonBlob, record.Sherpa)
	if err != nil {
		return SherpaFileError{
			Type:         "Unable to read file",
			Code:         400,
			Msg:          fmt.Sprintf("Unable to read file: %s", record.Meta.OnDiskLocation),
			Field:        "",
			DeveloperMsg: fmt.Sprintf("Failed to parse json []bytes from file: %s", record.Meta.OnDiskLocation),
			CliHelpMsg:   fmt.Sprintf("Failed to open Sherpa file: %s. Try running the last Sherpa command again.", record.Meta.OnDiskLocation)}
	}

	return nil

}

// ValidateSherpaRecord runs a series of validators, updating the SherpaRecord metadata
// will all found errors and returning the total number of errors found.
// func ValidateSherpaRecord(record SherpaRecord) int {
// 	valFuncs := makeAllValidators()
//
// 	var err error
// 	for i := 0; i < len(valFuncs); i++ {
// 		err = valFuncs[i](record)
// 		if err != nil {
// 			record.Meta.Errors = append(record.Meta.Errors, err)
// 		}
// 	}
// }
//
// func makeAllValidators() []SherpaValidator {
//
// 	validateName = func(s SherpaFile) error {
// 		//f len(name)
// 	}
//
// }

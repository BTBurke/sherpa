package sherpa_test

import (
	. "github.com/BTBurke/sherpa"

	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sherpafile", func() {

	var (
		SCorrect     SherpaFile
		CorrectJson  []byte
		Author1      Author
		Dep1         Dependency
		Privatedata1 Privatedata
	)

	BeforeEach(func() {
		Author1 = Author{
			Name:     "Bryan Burke",
			Email:    "bryan@kilimanjaro.io",
			Homepage: "https://kilimanjaro.io",
		}
		Dep1 = Dependency{
			Name: "BTBurke/SherpaDepTest",
			Type: "git",
			Url:  "https://github.com/BTBurke/Sherpa/SherpaDepTest",
		}
		Privatedata1 = Privatedata{
			Name: "BTBurke/SherpaPrivate",
			Type: "git",
			Url:  "https://github.com/BTBurke/sherpa/SherpaPrivate",
			File: "test_private_data.json",
		}
		SCorrect = SherpaFile{
			Name:        "BTBurke/Test",
			Description: "Test description",
			Version:     "0.1.0",
			Main:        []string{"test.yaml"},
			License:     "BSD",
			Keywords:    []string{"test", "sherpa"},
			Authors:     []Author{Author1},
			Homepage:    "https://sherpa.io",
			Repository: Repository{
				Type: "git",
				Url:  "https://github.com/BTBurke/sherpa",
			},
			Dependencies: []Dependency{Dep1},
			OsVersions:   []string{"osx", "linux", "windows"},
			Private:      false,
			PrivateData:  Privatedata1,
		}
		CorrectJson = []byte(`{
				"name": "BTBurke/Test",
				"description": "Test description",
				"version": "0.1.0",
				"main": ["test.yaml"],
				"license": "BSD",
				"keywords": ["test", "sherpa"],
				"authors": [{
					"name": "Bryan Burke",
					"email": "bryan@kilimanjaro.io",
					"homepage": "https://kilimanjaro.io"
					}],
				"homepage": "https://sherpa.io",
				"repository": {
					"type": "git",
					"url": "https://github.com/BTBurke/sherpa"
				},
				"dependencies": [{
					"name": "BTBurke/SherpaDepTest",
					"type": "git",
					"url":  "https://github.com/BTBurke/Sherpa/SherpaDepTest"
					}],
				"osVersions": ["osx", "linux", "windows"],
				"private": false,
				"privateData": {
					"name": "BTBurke/SherpaPrivate",
					"type": "git",
					"url":  "https://github.com/BTBurke/sherpa/SherpaPrivate",
					"file": "test_private_data.json"
					}
			}`)
	})

	Describe("Marshaling/Unmarshaling JSON", func() {
		Context("When setting up tests, an idiot check for test data", func() {
			It("Should serialize/deserialize test JSON into the test Struct", func() {
				var Test1 SherpaFile
				err := json.Unmarshal(CorrectJson, &Test1)
				Ω(err).ShouldNot(HaveOccurred())
				Expect(Test1).To(Equal(SCorrect))
			})
		})

		Context("When parsing JSON-formatted []byte to a SherpaRecord", func() {
			It("Should decode valid JSON to a SherpaRecord", func() {
				var Test2 SherpaRecord
				err := SherpaRecordFromJSON(CorrectJson, Test2)
				Ω(err).ShouldNot(HaveOccurred())
				Expect(Test2.Sherpa).To(Equal(SCorrect))
			})
		})
	})
})

package main

import (
	"testing"
)

type testCase struct {
	input            []string
	expectedResponse parsedInfo
}

func compareParsedInfos(p1 parsedInfo, p2 parsedInfo) bool {

	if p1.command == p2.command && len(p1.params) == len(p2.params) && len(p1.flags) == len(p2.flags) {

		for i := range p2.params {
			if p1.params[i] != p2.params[i] {
				return false
			}
		}

		for key, value := range p2.flags {
			v, ok := p1.flags[key]

			if !ok || v != value {
				return false
			}

		}

		return true
	}

	return false
}

func TestParser(t *testing.T) {

	t.Run("Positives", func(t *testing.T) {
		t.Run("Should parse a command name", func(t *testing.T) {
			testCases := []testCase{
				{
					input:            []string{"add"},
					expectedResponse: parsedInfo{command: "add"},
				},
			}

			for _, c := range testCases {
				t.Run("Command "+c.input[0], func(t *testing.T) {
					data, err := parse(c.input)

					if err != nil {
						t.Fatal(err)
					}

					if !compareParsedInfos(data, c.expectedResponse) {
						t.Error("Wrong returned values")
					}
				})
			}
		})

		t.Run("Should parse params", func(t *testing.T) {
			testCases := []testCase{
				{
					input:            []string{"add", "liro"},
					expectedResponse: parsedInfo{command: "add", params: []string{"liro"}},
				},
				{
					input:            []string{"add", "2"},
					expectedResponse: parsedInfo{command: "add", params: []string{"2"}},
				},
			}

			for _, c := range testCases {
				t.Run("Command "+c.input[0], func(t *testing.T) {
					data, err := parse(c.input)

					if err != nil {
						t.Fatal(err)
					}

					if !compareParsedInfos(data, c.expectedResponse) {
						t.Error("Wrong returned values")
					}
				})
			}
		})

		t.Run("Should parse flags", func(t *testing.T) {
			testCases := []testCase{
				{
					input: []string{"add", "liro", "-e", "a"},
					expectedResponse: parsedInfo{
						command: "add",
						params:  []string{"liro"},
						flags: map[string]string{
							"-e": "a",
						},
					},
				},
				{
					input: []string{"add", "liro", "-e", "a", "-f", "5", "-g", "/"},
					expectedResponse: parsedInfo{
						command: "add",
						params:  []string{"liro"},
						flags: map[string]string{
							"-e": "a",
							"-f": "5",
							"-g": "/",
						},
					},
				},
			}

			for _, c := range testCases {
				t.Run("Command "+c.input[0], func(t *testing.T) {
					data, err := parse(c.input)

					if err != nil {
						t.Fatal(err)
					}

					if !compareParsedInfos(data, c.expectedResponse) {
						t.Error("Wrong returned values")
					}
				})
			}
		})

		t.Run("Should parse a full complete command", func(t *testing.T) {
			testCases := []testCase{
				{
					input: []string{"add", "liro", "-d", "laro", "-e", "f", "-p", "1", "-l", "laros"},
					expectedResponse: parsedInfo{
						command: "add",
						params:  []string{"liro"},
						flags: map[string]string{
							"-d": "laro",
							"-e": "f",
							"-p": "1",
							"-l": "laros",
						},
					},
				},
			}

			for _, c := range testCases {
				t.Run("Command "+c.input[0], func(t *testing.T) {
					data, err := parse(c.input)

					if err != nil {
						t.Fatal(err)
					}

					if !compareParsedInfos(data, c.expectedResponse) {
						t.Error("Wrong returned values")
					}
				})
			}
		})
	})

	t.Run("Negatives", func(t *testing.T) {

		t.Run("Should not parse invalid commands", func(t *testing.T) {
			validateCase := testCase{

				input:            []string{"invalidComand", "liro"},
				expectedResponse: parsedInfo{},
			}

			data, err := parse(validateCase.input)

			if err == nil {
				t.Error("Parsed Invalid Command")
			}

			if !compareParsedInfos(data, validateCase.expectedResponse) {
				t.Error("Wrong returned values")
			}

		})

		t.Run("Should error on repeated flags", func(t *testing.T) {
			testCases := []testCase{
				{
					input:            []string{"add", "liro", "-d", "laro", "-d", "lero"},
					expectedResponse: parsedInfo{},
				},
			}

			for _, c := range testCases {
				t.Run("Command "+c.input[0], func(t *testing.T) {
					data, err := parse(c.input)

					if err == nil {
						t.Fatal(err)
					}

					if !compareParsedInfos(data, c.expectedResponse) {
						t.Error("Wrong returned values")
					}
				})
			}
		})

		t.Run("Should error on no value flags", func(t *testing.T) {
			testCases := []testCase{
				{
					input:            []string{"add", "liro", "-d"},
					expectedResponse: parsedInfo{},
				},
				{
					input:            []string{"add", "liro", "-d", "a", "-e"},
					expectedResponse: parsedInfo{},
				},
				{
					input:            []string{"add", "liro", "-d", "-e"},
					expectedResponse: parsedInfo{},
				},
			}

			for _, c := range testCases {
				t.Run("Command "+c.input[0], func(t *testing.T) {
					data, err := parse(c.input)

					if err == nil {
						t.Fatal(err)
					}

					if !compareParsedInfos(data, c.expectedResponse) {
						t.Error("Wrong returned values")
					}
				})
			}
		})

		t.Run("Should error on empty values", func(t *testing.T) {
			testCases := []testCase{
				{
					input:            []string{"add", ""},
					expectedResponse: parsedInfo{},
				},
				{
					input:            []string{"add", "      "},
					expectedResponse: parsedInfo{},
				},
				{
					input:            []string{"add", "\n"},
					expectedResponse: parsedInfo{},
				},
				{
					input:            []string{"add", "\t"},
					expectedResponse: parsedInfo{},
				},
			}

			for _, c := range testCases {
				t.Run("Command "+c.input[0], func(t *testing.T) {
					data, err := parse(c.input)

					if err == nil {
						t.Fatal("Parsed empty value")
					}

					if !compareParsedInfos(data, c.expectedResponse) {
						t.Error("Wrong returned values")
					}
				})
			}
		})
	})

}

package errormapper_test

import (
	"fmt"
	errormapper "github.com/mbict/go-errormapper"
	validate "github.com/mbict/go-validate"
	. "gopkg.in/check.v1"
)

type FieldErrorTranslatorSuite struct{}

var _ = Suite(&FieldErrorTranslatorSuite{})

func (s *FieldErrorTranslatorSuite) TestTranslate(c *C) {
	et := errormapper.FieldErrorTranslator{
		"A": errormapper.ErrorTranslator{
			validate.ErrRequired: "a required translate",
		},
		"B": errormapper.ErrorTranslator{
			validate.ErrRequired: "b min translate",
			validate.ErrMin:      "b min translate",
			validate.ErrMax:      "b max translate",
		},
	}

	fmt.Println(et)
	/*
		tests := []struct {
			Description string
			Errors      validate.Errors
			ExpectedOk  bool
			Expected    string
		}{
			{
				Description: "default",
				Errors:      validate.Errors{validate.ErrRequired, validate.ErrMin},
				ExpectedOk:  true,
				Expected:    "required translate, min translate",
			}, {
				Description: "order test",
				Errors:      validate.Errors{validate.ErrMin, validate.ErrRequired},
				ExpectedOk:  true,
				Expected:    "min translate, required translate",
			}, {
				Description: "only known errors are translated",
				Errors:      validate.Errors{validate.ErrMax, validate.ErrRequired},
				ExpectedOk:  true,
				Expected:    "required translate",
			}, {
				Description: "No translation at all",
				Errors:      validate.Errors{validate.ErrMax},
				ExpectedOk:  false,
				Expected:    "",
			},
		}

		for _, test := range tests {
			translated, ok := et.Translate(test.Errors)

			c.Assert(ok, Equals, test.ExpectedOk, Commentf(test.Description))
			c.Assert(translated, Equals, test.Expected, Commentf(test.Description))
		}
	*/
}

func (s *FieldErrorTranslatorSuite) TestTranslateWithDefault(c *C) {
	//a translation with a default (nil) translation always succeeds translation
	et := errormapper.FieldErrorTranslator{
		"A": errormapper.ErrorTranslator{
			validate.ErrRequired: "a required translate",
		},
		"B": errormapper.ErrorTranslator{
			validate.ErrRequired: "b min translate",
			validate.ErrMin:      "b min translate",
			validate.ErrMax:      "b max translate",
		},
		"": errormapper.ErrorTranslator{
			validate.ErrMin: "nil min translate",
		},
	}
	fmt.Println(et)
	/*
		tests := []struct {
			Description string
			Errors      validate.Errors
			Expected    string
		}{
			{
				Description: "default",
				Errors:      validate.Errors{validate.ErrRequired, validate.ErrMin},
				Expected:    "required translate, min translate",
			}, {
				Description: "order test, only first error",
				Errors:      validate.Errors{validate.ErrMin, validate.ErrRequired},
				Expected:    "min translate, required translate",
			}, {
				Description: "fallback to nil translation on unkown error",
				Errors:      validate.Errors{validate.ErrMax, validate.ErrRequired},
				Expected:    "nil default translate, required translate",
			},
		}

		for _, test := range tests {
			translated, ok := et.Translate(test.Errors)

			c.Assert(translated, Equals, test.Expected, Commentf(test.Description))
			c.Assert(ok, Equals, true, Commentf(test.Description))
		}
	*/
}

func (s *FieldErrorTranslatorSuite) TestTranslateFirst(c *C) {
	et := errormapper.FieldErrorTranslator{
		"A": errormapper.ErrorTranslator{
			validate.ErrRequired: "a required translate",
		},
		"B": errormapper.ErrorTranslator{
			validate.ErrRequired: "b min translate",
			validate.ErrMin:      "b min translate",
			validate.ErrMax:      "b max translate",
		},
	}

	fmt.Println(et)
	/*
		tests := []struct {
			Description string
			Errors      validate.Errors
			ExpectedOk  bool
			Expected    string
		}{
			{
				Description: "first error",
				Errors:      validate.Errors{validate.ErrRequired, validate.ErrMin},
				ExpectedOk:  true,
				Expected:    "required translate",
			}, {
				Description: "order test, only first error",
				Errors:      validate.Errors{validate.ErrMin, validate.ErrRequired},
				ExpectedOk:  true,
				Expected:    "min translate",
			}, {
				Description: "only known errors are translated",
				Errors:      validate.Errors{validate.ErrMax, validate.ErrRequired},
				ExpectedOk:  true,
				Expected:    "required translate",
			}, {
				Description: "No translation at all",
				Errors:      validate.Errors{validate.ErrMax},
				ExpectedOk:  false,
				Expected:    "",
			},
		}

		for _, test := range tests {
			translated, ok := et.TranslateFirst(test.Errors)

			c.Assert(ok, Equals, test.ExpectedOk, Commentf(test.Description))
			c.Assert(translated, Equals, test.Expected, Commentf(test.Description))
		}
	*/
}

func (s *FieldErrorTranslatorSuite) TestTranslateFirstWithDefault(c *C) {
	//a translation with a default (nil) translation always succeeds translation
	et := errormapper.FieldErrorTranslator{
		"A": errormapper.ErrorTranslator{
			validate.ErrRequired: "a required translate",
		},
		"B": errormapper.ErrorTranslator{
			validate.ErrRequired: "b min translate",
			validate.ErrMin:      "b min translate",
			validate.ErrMax:      "b max translate",
		},
		nil: errormapper.ErrorTranslator{
			validate.ErrMin: "nil min translate",
		},
	}

	fmt.Println(et)
	/*
		tests := []struct {
		Description string
		Errors      validate.Errors
		Expected    string
		}{
		{
			Description: "first error",
			Errors:      validate.Errors{validate.ErrRequired, validate.ErrMin},
			Expected:    "required translate",
		}, {
			Description: "order test, only first error",
			Errors:      validate.Errors{validate.ErrMin, validate.ErrRequired},
			Expected:    "min translate",
		}, {
			Description: "fallback to nil translation on unkown error",
			Errors:      validate.Errors{validate.ErrMax, validate.ErrRequired},
			Expected:    "nil default translate",
		},
		}

		for _, test := range tests {
		translated, ok := et.TranslateFirst(test.Errors)

		c.Assert(translated, Equals, test.Expected, Commentf(test.Description))
		c.Assert(ok, Equals, true, Commentf(test.Description))
		}
	*/
}

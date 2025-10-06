package config

import "fmt"

type CourseKind string

const (
	Calculus     CourseKind = "calculus"
	Algebra      CourseKind = "algebra"
	Probability  CourseKind = "probability"
	Trigonometry CourseKind = "trigonometry"
	CS           CourseKind = "cs"
)

type ProjectConfig struct {
	IndexURL   string
	SourceDir  string
	DataDir    string
	Slug       string
	CourseKind CourseKind
}

var cmeSlug = "calculus-made-easy"
var CME = ProjectConfig{
	IndexURL:   "https://calculusmadeeasy.org/",
	SourceDir:  fmt.Sprintf("%s/fixed", cmeSlug),
	DataDir:    cmeSlug,
	Slug:       cmeSlug,
	CourseKind: Calculus,
}

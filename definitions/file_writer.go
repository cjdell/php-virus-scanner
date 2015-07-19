package definitions

import (
	"regexp"
)

type FileWriter struct {
	matcher *regexp.Regexp
}

func (def *FileWriter) Init() {
	def.matcher, _ = regexp.Compile(`\\xEF\\xBB\\xBF`)
}

func (FileWriter) Name() string {
	return "File Writer"
}

func (def *FileWriter) Check(source string) bool {
	return def.matcher.MatchString(source)
}

package definitions

import (
	"regexp"
)

type FileWriter struct {
	matcher1 *regexp.Regexp
	matcher2 *regexp.Regexp
}

func (def *FileWriter) Init() {
	def.matcher1, _ = regexp.Compile(`\$_(REQUEST|POST)\[`)
	def.matcher2, _ = regexp.Compile(`fwrite\(\$fp\,\"\\xEF\\xBB\\xBF`)
}

func (FileWriter) Name() string {
	return "File Writer"
}

func (def *FileWriter) Check(source string) bool {
	return def.matcher1.MatchString(source) && def.matcher2.MatchString(source)
}

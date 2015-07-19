package definitions

import (
	"regexp"
)

type EvalBase64 struct {
	matcher *regexp.Regexp
}

func (def *EvalBase64) Init() {
	def.matcher, _ = regexp.Compile(`eval\(base64`)
}

func (EvalBase64) Name() string {
	return "eval(base64..."
}

func (def *EvalBase64) Check(source string) bool {
	return def.matcher.MatchString(source)
}

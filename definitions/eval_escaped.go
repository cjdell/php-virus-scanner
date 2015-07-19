package definitions

import (
	"regexp"
)

type EvalEscaped struct {
	matcher *regexp.Regexp
}

func (def *EvalEscaped) Init() {
	def.matcher, _ = regexp.Compile(`\\x65\\x76\\x61\\x6C`)
}

func (EvalEscaped) Name() string {
	return "Escaped eval()"
}

func (def *EvalEscaped) Check(source string) bool {
	return def.matcher.MatchString(source)
}

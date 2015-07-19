package definitions

import (
	"regexp"
)

type Assert struct {
	matcher *regexp.Regexp
}

func (def *Assert) Init() {
	def.matcher, _ = regexp.Compile(`\=[\"\']a[\"\'\.]*s[\"\'\.]*s[\"\'\.]*e[\"\'\.]*r[\"\'\.]*t[\"\']`)
}

func (Assert) Name() string {
	return "assert(...)"
}

func (def *Assert) Check(source string) bool {
	return def.matcher.MatchString(source)
}

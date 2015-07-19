package definitions

import (
	"regexp"
)

type ScriptInject struct {
	matcher *regexp.Regexp
}

func (def *ScriptInject) Init() {
	def.matcher, _ = regexp.Compile(`\=\ \"\<script type\=\\"text\/javascript\\" src\=\\"http\:\/\/`)
}

func (ScriptInject) Name() string {
	return "<script ...> inject"
}

func (def *ScriptInject) Check(source string) bool {
	return def.matcher.MatchString(source)
}

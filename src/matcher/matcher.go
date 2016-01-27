package matcher

import (
	"github.com/gobwas/glob"
)

type Matcher struct{ globs []glob.Glob }

func (f *Matcher) Match(s string) bool {
	for _, g := range f.globs {
		if g.Match(s) {
			return true
		}
	}
	return false
}

// func (f *Matcher) Do(fs []string) []string {
// 	res := make([]string, len(fs))
// 	iter := 0
// 	for _, v := range fs {
// 		if f.Match(v) {
// 			res[iter] = v
// 			iter++
// 		}
// 	}
// 	return res
// }

func New(fs []string) *Matcher {
	f := &Matcher{make([]glob.Glob, len(fs))}
	for i, v := range fs {
		f.globs[i] = glob.MustCompile(v)
	}
	return f
}

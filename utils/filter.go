package utils

import (
	"fmt"
	"regexp"
	"strings"
)

const keywords = `\s(replace|lock tables|select|update|delete|insert|delete|alter|drop|create|grant|revoke|shutdown|;|\\g)\s`

var keyregexp *regexp.Regexp

func init() {
	keyregexp, _ = regexp.Compile(keywords)
}

func Filter(s string) error {
	matches := keyregexp.FindAll([]byte(strings.ToLower(s)), len(s))
	if len(matches) == 0 {
		return nil
	}
	var out []string
	for _, match := range matches {
		out = append(out, string(match))
	}
	return fmt.Errorf("got security keywords: %v, origin:[%s]", out, s)
}

package dnf

import (
	"strings"
)

func ToDNF(s string) string {
	s = strings.Replace(s, " ", "", -1)
	return s
}

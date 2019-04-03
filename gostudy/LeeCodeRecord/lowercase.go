package lowercase

import (
	"strings"
)

func toLowerCase(str string) string {
    var lowstr string
    lowstr = strings.ToLower(str)
    return lowstr
}
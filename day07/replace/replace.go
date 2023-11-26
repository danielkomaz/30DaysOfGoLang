package replace

import "strings"

func Replace(oldString string, oldChar string, newChar string) string {
	r := strings.NewReplacer(oldChar, newChar)
	return r.Replace(oldString)
}

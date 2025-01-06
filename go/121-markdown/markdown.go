package markdown

import (
	"fmt"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)

	var sb strings.Builder
	lines := strings.Split(markdown, "\n")
	for i := 0; i < len(lines); i++ {
		headingLevel := len(lines[i]) - len(strings.TrimLeft(lines[i], "#"))
		switch {
		case headingLevel > 0 && headingLevel < 7:
			sb.WriteString(
				fmt.Sprintf(
					"<h%d>%s</h%d>",
					headingLevel,
					strings.TrimLeft(lines[i], "# "),
					headingLevel,
				),
			)
		case lines[i][:1] == "*":
			sb.WriteString("<ul>")
			for ; ; i++ {
				sb.WriteString(fmt.Sprintf("<li>%s</li>", strings.TrimLeft(lines[i], "* ")))
				if !(i+1 < len(lines) && lines[i+1][:1] == "*") {
					break
				}
			}
			sb.WriteString("</ul>")
		default:
			sb.WriteString(fmt.Sprintf("<p>%s</p>", lines[i]))
		}
	}
	return sb.String()
}

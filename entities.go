package entities

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

var needEscape = make(map[rune]struct{})

func init() {
	for _, r := range []rune{'_', '*', '[', ']', '(', ')', '~', '`', '>', '#', '+', '-', '=', '|', '{', '}', '.', '!'} {
		needEscape[r] = struct{}{}
	}
}

func ConvertToMarkdownV2(text string, messageEntities []tb.MessageEntity) string {
	insertions := make(map[int]string)
	for _, e := range messageEntities {
		var before, after string
		if e.Type == tb.EntityBold {
			before = "*"
			after = "*"
		} else if e.Type == tb.EntityItalic {
			before = "_"
			after = "_"
		} else if e.Type == tb.EntityUnderline {
			before = "__"
			after = "__"
		} else if e.Type == tb.EntityStrikethrough {
			before = "~"
			after = "~"
		} else if e.Type == tb.EntityCode {
			before = "`"
			after = "`"
		} else if e.Type == tb.EntityCodeBlock {
			before = "```" + e.Language
			after = "```"
		} else if e.Type == tb.EntityTextLink {
			before = "["
			after = "](" + e.URL + ")"
		}
		if before != "" {
			insertions[e.Offset] += before
			insertions[e.Offset+e.Length] += after
		}
	}
	result := ""
	for i, c := range text {
		result += insertions[i]
		if _, has := needEscape[c]; has {
			result += `\`
		}
		result += string(c)
	}
	result += insertions[len(text)]
	return result
}

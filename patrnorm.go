package fiofix

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Список окончаний отчеств
var patrEndings = []string{
	"оглы", "оглу", "огли", "угли", "кызы", "гызы", "заде", "уулу", "вали", "вэли",
	"улы",
	"лы",
}

// Основная функция нормализации окончаний отчетств
func NormalizePatronymic(s string) string {
	if s == "" {
		return s
	}

	// Разбиение слов на части
	words := strings.Fields(s)
	modified := false

	for i, word := range words {
		// Проверка слова на наличие окончания
		if base, suffix, found := extractEnding(word); found {
			words[i] = formatPart(base, suffix)
			modified = true
		}
	}

	if !modified {
		return s
	}

	return strings.Join(words, " ")
}

// Поиск окончаний без учета регистра
func extractEnding(word string) (base string, ending string, found bool) {
	wordLen := len(word)
	for _, end := range patrEndings {
		endLen := len(end)
		if wordLen >= endLen {
			suffix := word[wordLen-endLen:]
			// Проверка совпадений
			if strings.EqualFold(suffix, end) {
				return word[:wordLen-endLen], suffix, true
			}
		}
	}
	return word, "", false
}

// Формирование итоговой строки
func formatPart(base, suffix string) string {
	// Очистка основы от висящего дефиса (пример: "Кайрат-Оглы" -> "Кайрат")
	base = strings.TrimRight(base, "-")

	var b strings.Builder

	if base == "" {
		b.Grow(len(suffix))
		capitalizeString(&b, suffix)
	} else {
		b.Grow(len(base) + 1 + len(suffix))
		capitalizeString(&b, base)
		b.WriteByte(' ')
		capitalizeString(&b, suffix)
	}

	return b.String()
}

// Делает первую букву заглавной, а остальные строчными
func capitalizeString(b *strings.Builder, s string) {
	if s == "" {
		return
	}

	// Первая буква в верхнем регистре
	r, size := utf8.DecodeRuneInString(s)
	if size == 0 {
		return
	}
	b.WriteRune(unicode.ToUpper(r))

	// Остальные буквы в нижнем регистре
	for i, w := size, 0; i < len(s); i += w {
		r, w = utf8.DecodeRuneInString(s[i:])
		b.WriteRune(unicode.ToLower(r))
	}
}

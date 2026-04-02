package fiofix

import (
	"testing"
)

func TestNormalizePatronymic(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Already normalized",
			input:    "Кайрат Оглы",
			expected: "Кайрат Оглы",
		},
		{
			name:     "Joined words",
			input:    "Кайратоглы",
			expected: "Кайрат Оглы",
		},
		{
			name:     "With hyphen",
			input:    "Кайрат-Оглы",
			expected: "Кайрат Оглы",
		},
		{
			name:     "Suffix at the beginning",
			input:    "Оглыкайрат",
			expected: "Оглыкайрат",
		},
		{
			name:     "Wrong casing",
			input:    "СЕРИКулы",
			expected: "Серик Улы",
		},
		{
			name:     "Alternative ending",
			input:    "Айгуль-кызы",
			expected: "Айгуль Кызы",
		},
		{
			name:     "Угли + Оглы в одном ФИО",
			input:    "Нурмухаммад Баходир Угли Оглы",
			expected: "Нурмухаммад Баходир Угли Оглы",
		},
		{
			name:     "Дефис + верхний регистр + слитно",
			input:    "АЛИ-ОГЛЫ",
			expected: "Али Оглы",
		},
		{
			name:     "Угли в конце",
			input:    "Баходир Угли",
			expected: "Баходир Угли",
		},
		{
			name:     "Оглы с дефисом и верхним регистром",
			input:    "Хикматилла-ОГЛЫ",
			expected: "Хикматилла Оглы",
		},
		{
			name:     "Кызы слитно",
			input:    "Айгулькызы",
			expected: "Айгуль Кызы",
		},
		{
			name:     "Несколько слов с окончаниями",
			input:    "Ерлан Жарылгапович Оглы",
			expected: "Ерлан Жарылгапович Оглы",
		},
		{
			name:     "Только суффикс",
			input:    "ОГЛЫ",
			expected: "Оглы",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizePatronymic(tt.input); got != tt.expected {
				t.Errorf("NormalizePatronymic(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestReplaceLettersToRussian(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Remove XXX",
			input:    "ИвановXXX",
			expected: "Иванов",
		},
		{
			name:     "Kazakh specific letters",
			input:    "Қазақстан Әлішер",
			expected: "Казакстан Элишер",
		},
		{
			name:     "Latin to Cyrillic",
			input:    "Astana",
			expected: "Астана",
		},
		{
			name:     "Полная транслитерация казахского имени",
			input:    "Құлбаев Манас Қасымжанұлы",
			expected: "Кулбаев Манас Касымжанулы",
		},
		{
			name:     "Смешанный алфавит",
			input:    "Джаксыбаева Жаннет Жумадилқызы Shodiev Aminjon",
			expected: "Джаксыбаева Жаннет Жумадилкызы Шодиев Аминжон",
		},
		{
			name:     "Шә, Қы, Ұ, Ө, Ғ, І",
			input:    "Шәріпхан Қайрат Ұлы Өмірзақ Ғали Ілияс",
			expected: "Шэрипхан Кайрат Улы Омирзак Гали Илияс",
		},
		{
			name:     "Латинские + казахские",
			input:    "Kurbanov Kholtura Kholmuratovich Қайрат",
			expected: "Курбанов Холтура Холмуратович Кайрат",
		},
		{
			name:     "Сложные сочетания SCH, KH и т.д.",
			input:    "Shodiev Aminjon Bakhodir Ugli",
			expected: "Шодиев Аминжон Баходир Угли",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceLettersToRussian(tt.input); got != tt.expected {
				t.Errorf("ReplaceLettersToRussian(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

package num2wordsru

import "strings"

var units = []string{"", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"}
var teens = []string{"", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}
var tens = []string{"", "десять", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}
var hundreds = []string{"", "сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"}
var thousands = []string{"тысяча", "тысячи", "тысяч"}

func NumberToWordsRu(n int) string {
	if n == 0 {
		return "ноль"
	}

	parts := []string{}

	if n/1000000 > 0 {
		parts = append(parts, convertThreeDigits(n/1000000)+" миллион")
		n %= 1000000
	}
	if n/1000 > 0 {
		parts = append(parts, convertThreeDigits(n/1000)+" "+thousands[0])
		n %= 1000
	}
	if n > 0 {
		parts = append(parts, convertThreeDigits(n))
	}

	return strings.Join(parts, " ")
}

func convertThreeDigits(n int) string {
	parts := []string{}

	if n/100 > 0 {
		parts = append(parts, hundreds[n/100])
		n %= 100
	}
	if n >= 11 && n <= 19 {
		parts = append(parts, teens[n-10])
	} else {
		if n/10 > 0 {
			parts = append(parts, tens[n/10])
			n %= 10
		}
		if n > 0 {
			parts = append(parts, units[n])
		}
	}

	return strings.Join(parts, " ")
}

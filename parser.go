package num2Ru

import "strings"

var units = []interface{}{
	"ноль",
	[]string{"один", "одна"},
	[]string{"два", "две"},
	"три", "четыре", "пять",
	"шесть", "семь", "восемь", "девять",
}

var teens = []string{
	"десять", "одиннадцать",
	"двенадцать", "тринадцать",
	"четырнадцать", "пятнадцать",
	"шестнадцать", "семнадцать",
	"восемнадцать", "девятнадцать",
}

var tens = []interface{}{
	teens,
	"двадцать", "тридцать",
	"сорок", "пятьдесят",
	"шестьдесят", "семьдесят",
	"восемьдесят", "девяносто",
}

var hundreds = []string{
	"сто", "двести",
	"триста", "четыреста",
	"пятьсот", "шестьсот",
	"семьсот", "восемьсот",
	"девятьсот",
}

var orders = [][2]interface{}{
	{[]string{"тысяча", "тысячи", "тысяч"}, "f"},
	{[]string{"миллион", "миллиона", "миллионов"}, "m"},
	{[]string{"миллиард", "миллиарда", "миллиардов"}, "m"},
}

var minus = "минус"

func thousand(rest int, sex string) (int, []string) {
	prev := 0
	plural := 2
	name := []string{}
	useTeens := 10 <= rest%100 && rest%100 <= 19

	var data [][2]interface{}
	if !useTeens {
		data = [][2]interface{}{
			{units, 10}, {tens, 100}, {hundreds, 1000},
		}
	} else {
		data = [][2]interface{}{
			{teens, 10}, {hundreds, 1000},
		}
	}

	for _, pair := range data {
		names := pair[0]
		x := pair[1].(int)
		cur := int(((rest - prev) % x) * 10 / x)
		prev = rest % x
		if x == 10 && useTeens {
			plural = 2
			name = append(name, teens[cur])
		} else if cur == 0 {
			continue
		} else if x == 10 {
			name_ := names.([]interface{})[cur]
			if nameTuple, ok := name_.([]string); ok {
				name_ = nameTuple[0]
				if sex == "f" {
					name_ = nameTuple[1]
				}
			}
			name = append(name, name_.(string))
			if 2 <= cur && cur <= 4 {
				plural = 1
			} else if cur == 1 {
				plural = 0
			} else {
				plural = 2
			}
		} else {
			name = append(name, names.([]string)[cur-1])
		}
	}
	return plural, name
}

func num2text(num int) string {
	if num == 0 {
		return "ноль"
	}

	rest := abs(num)
	ord := 0
	name := []string{}
	for rest > 0 {
		plural, nme := thousand(rest%1000, orders[ord][1].(string))
		if len(nme) > 0 || ord == 0 {
			name = append(name, orders[ord][0].([]string)[plural])
		}
		name = append(name, nme...)
		rest = rest / 1000
		ord++
	}
	if num < 0 {
		name = append(name, minus)
	}
	reverse(name)
	return strings.Join(name, " ")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

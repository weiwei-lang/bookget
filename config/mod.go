package config

import (
	"strconv"
	"strings"
)

// SeqContinue    false = 最小值 <= 当前页码 <=  最大值
func SeqContinue(index int) bool {
	if Conf.Seq == "" || !strings.Contains(Conf.Seq, ":") {
		return false
	}
	m := strings.Split(Conf.Seq, ":")
	min, _ := strconv.Atoi(m[0])
	max, _ := strconv.Atoi(m[1])
	index++
	if index < min || (max > 0 && index >= max) {
		return true
	}
	return false
}

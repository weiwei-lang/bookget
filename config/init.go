package config

import (
	"strconv"
	"strings"
)

var Conf Input
var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.50"

const version = "1.1.0"

// initSeq    false = 最小值 <= 当前页码 <=  最大值
func initSeq() {
	if Conf.Seq == "" || !strings.Contains(Conf.Seq, ":") {
		return
	}
	m := strings.Split(Conf.Seq, ":")
	min, _ := strconv.Atoi(m[0])
	max, _ := strconv.Atoi(m[1])
	Conf.SeqStart = min
	Conf.SeqEnd = max
	return
}

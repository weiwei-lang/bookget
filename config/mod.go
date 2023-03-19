package config

// SeqContinue    false = 最小值 <= 当前页码 <=  最大值
func SeqContinue(index int) bool {
	if Conf.SeqStart <= 0 {
		return false
	}
	if Conf.SeqEnd > 0 && index+1 >= Conf.SeqStart && index < Conf.SeqEnd {
		return false
	} else if index+1 >= Conf.SeqStart {
		return false
	}
	return true
}

func PageRange(index, size int) bool {
	if Conf.SeqStart <= 0 {
		return false
	}
	if Conf.SeqEnd < 0 && (index-size >= Conf.SeqEnd) {
		return true
	} else if Conf.SeqEnd > 0 && index+1 >= Conf.SeqStart && index < Conf.SeqEnd {
		return false
	} else if index+1 >= Conf.SeqStart {
		return false
	}
	return true
}

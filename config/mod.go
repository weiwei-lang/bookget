package config

// SeqContinue    return false (最小值 <= 当前页码 <=  最大值)
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

// PageRange    return true (最小值 <= 当前页码 <=  最大值)
func PageRange(index, size int) bool {
	if Conf.SeqStart <= 0 {
		return true
	}
	if Conf.SeqEnd < 0 && (index-size >= Conf.SeqEnd) {
		return false
	} else if Conf.SeqEnd > 0 && index+1 >= Conf.SeqStart && index < Conf.SeqEnd {
		return true
	} else if index+1 >= Conf.SeqStart {
		return true
	}
	return false
}

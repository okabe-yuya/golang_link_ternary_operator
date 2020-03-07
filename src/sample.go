package src

func AnonymousFunctionSample(orderNum int) string {
	// 無名関数を即時実行する
	drink := func(n int) string {
		if n == 1 {
			return "green tea"
		} else {
			return "tea"
		}
		// 即時実行のために引数を指定
	}(orderNum)

	// 無名関数の実行結果を返す
	return drink
}

func NormalSample(orderNum int) string {
	drink := ""
	if orderNum == 1 {
		drink = "green tes"
	} else {
		return "tea"
	}
	return drink
}

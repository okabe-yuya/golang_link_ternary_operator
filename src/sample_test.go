package src

import (
	"fmt"
	"testing"
	"time"
)

func TestSample(t *testing.T) {
	tryNum := 10000
	averageAnonymous := commonExec(tryNum, func(n int) string {
		return AnonymousFunctionSample(n)
	})

	fmt.Println("[info] average for anonymous function: ", averageAnonymous)

	averageNormal := commonExec(tryNum, func(n int) string {
		return NormalSample(n)
	})

	fmt.Println("[info] average for normal exec: ", averageNormal)
}

func commonExec(tryNum int, execFun func(int) string) float64 {
	result := make([]float64, 0)

	for i := 0; i < tryNum; i++ {
		if i < tryNum/2 {
			start := time.Now()
			execFun(1)
			result = append(result, time.Now().Sub(start).Seconds())
		} else {
			start := time.Now()
			execFun(0)
			result = append(result, time.Now().Sub(start).Seconds())
		}
	}

	// 平均値の算出
	total := 0.0
	for _, val := range result {
		total += val
	}

	average := total / float64(len(result))
	return average
}

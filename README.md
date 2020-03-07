## golangで三項演算子っぽいことをやるための記述
自身のブログにて記事を公開中
> 

## Description
golangには三項演算子が実装されておらず、都度、ifのconditionの判定によって値を更新する場合に、事前に型が一致する変数に対して初期値を代入して用意しておく必要があり、個人的に煩わしい。これが`golang`らしさと言われればそれまでだが、何とか三項演算子っぽい記述方法を思いついたので公開

### How to write??
😫 not cool
```go
drink := ""
if orderNum == 1 {
	drink = "green tes"
} else {
	return "tea"
}
```

😆 cool
```go
drink := func(n int) string {
		if n == 1 {
			return "green tea"
		} else {
			return "tea"
		}
		// 即時実行のために引数を指定
	}(1)
```

## usage
`$GOPATH`直下に当レポジトリを配置。  
テストケースにてベンチマークを実装。10000回の試行の平均値をterminalに標準出力する。  

> $ bash benchmark.sh
```
[info] clear go test cached...
[info] start benchmark
=== RUN   TestSample
[info] average for anonymous function:  9.854730000000441e-08
[info] average for normal exec:  1.0573470000000177e-07
--- PASS: TestSample (0.01s)
PASS
ok      ternary_operator/src    0.015s
```
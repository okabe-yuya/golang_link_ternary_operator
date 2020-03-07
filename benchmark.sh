echo "[info] clear go test cached..."
go clean -testcache

echo "[info] start benchmark"
go test ternary_operator/src -run TestSample -v

package main

// 素数判定
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	// nの平方根までで割り切れるかを調べる
	// NOTE: nの平方根(math.Pow(n, 0.5))より大きい値(x...x+y)に着目する
	// n < math.Pow(x, 2) なので、xはありえない
	// n < x * (x+y) なので、xもx+yもありえない
	// nの平方根より大きい2つの整数の席では表せない
	// よってnの平方根より大きい値は調べる必要はない
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

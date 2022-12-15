# 猴子排序（Monkey Sort）

注意：本项目只是一个玩笑性质的项目，所使用的的排序算法也不适合用于生产环境！

# 一、安装

```bash
go get -u github.com/golang-infrastructure/go-monkey-sort
```

# 二、示例代码

```go
package main

import (
	"fmt"
	monkey_sort "github.com/golang-infrastructure/go-monkey-sort"
)

func main() {
	slice := []int{2, 3, 1, 4}
	count := monkey_sort.Sort(slice)
	fmt.Println(count)
	fmt.Println(slice)
	// Output:
	// 4861854
	// [1 2 3 4]
}
```




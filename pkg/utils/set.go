package utils

// Difference 返回 a 和 b 的差集，即所有存在于 a 但不在 b 中的元素
func Difference(a, b []string) []string {
	m := make(map[string]bool)
	diff := []string{}

	// 将切片 b 的元素存储在 map 中，值为 true
	for _, item := range b {
		m[item] = true
	}

	// 检查切片 a 的每个元素是否在 map 中
	// 如果不在，那么它就是差集的一部分
	for _, item := range a {
		if _, found := m[item]; !found {
			diff = append(diff, item)
		}
	}

	return diff
}

package main

func main() {
	m := map[int]int{4: 4, 8: 8, 15: 15, 16: 16, 23: 23, 42: 42}
	go func() {
		for {
			for _, v := range m {
				_ = v
			}
		}
	}()
	for {
		for k, _ := range m {
			m[k]++
		}
	}
}

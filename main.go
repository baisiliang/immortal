package main

//写一个加减计算器，输入n个数和n-1个运算符，输出计算结果
func Caculate(num []int, op []string) int {
	if len(num) != len(op)+1 {
		return 0
	}
	res := num[0]
	for i := 0; i < len(op); i++ {
		switch op[i] {
		case "+":
			res += num[i+1]
		case "-":
			res -= num[i+1]
		default:
			return 0
		}
	}
	return res
}

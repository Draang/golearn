package arrays

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
func SumAll(listNums [][]int) []int {
	listSum := []int{}
	for _, nums := range listNums {
		sum := Sum(nums)
		listSum = append(listSum, sum)
	}
	return listSum
}

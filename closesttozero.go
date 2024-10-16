func abs(x int) int{
    if x < 0{
        return -x
    }
    return x
}

func findClosestNumber(nums []int) int {
    min := 0
    for i := range nums{
        if abs(nums[i]) <= abs(nums[min]){
            if nums[i] < nums[min] && abs(nums[i]) == abs(nums[min]){
                min = min
            } else {
                min = i
            }
        }
    i++
    }
    return nums[min]
}

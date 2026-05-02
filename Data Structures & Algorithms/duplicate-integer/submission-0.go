func hasDuplicate(nums []int) bool {
    m := make(map[int]int)
    for i:=0;i<len(nums);i++ {
        _,exists := m[nums[i]]
        if exists {
            return true
        } else {
            m[nums[i]] = i
        }
    }
    return false
}

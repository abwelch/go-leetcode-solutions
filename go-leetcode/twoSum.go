func twoSum(nums []int, target int) []int {
    theMap := map[int][]int{}
    // build map
    for index, value := range nums {
        theMap[value] = append(theMap[value], index)
    }
    var result []int
    for _, value := range nums {
        if _, ok := theMap[target - value]; ok {
            if target - value == value {
                if len(theMap[target - value]) > 1 {
                    result = append(result, theMap[value][0], theMap[value][1])
                    break
                }
            } else {
                result = append(result, theMap[target - value][0], theMap[value][0])
                break
            }
        }
    }
    return result
}
func secondHighest(s string) int {
    nums := make(map[byte]byte)
    for i := 0; i < len(s); i++ {
        // a number
        if s[i] <= 57 {
            if _, exists := nums[s[i]]; !exists {
                nums[s[i]] = 1
            }
        }
    }
    // iterate over map for largest key
    var largest byte
    for k, _ := range nums {
        if k > largest {
            largest = k
        }
    }
    
    // iterate over map again for second largest
    var secLargest byte
    assigned := false
    for k, _ := range nums {
        if k > secLargest && k != largest {
            secLargest = k
            assigned = true
        }
    }
    
    if !assigned {
        return -1
    }
    
    ans, _ := strconv.Atoi(string(secLargest))
    return ans
}
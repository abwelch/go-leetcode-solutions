func wordSubsets(A []string, B []string) []string {
    var result []string
    masterMap := make(map[rune]byte)
    for _, word := range B {
        temp := make(map[rune]byte)
        for _, ch := range word {
            temp[ch]++    
        }
        for k, v := range temp {
            if v > masterMap[k] {
                masterMap[k] = v
            }
        }
    }
    for _, word := range A {
        universal := true
        temp := make(map[rune]byte)
        for _, ch := range word {
            temp[ch]++    
        }
        for k, v := range masterMap {
            if v > temp[k] {
                universal = false
                break
            }
        }
        if universal {
            result = append(result, word)
        }
    }
    return result
}
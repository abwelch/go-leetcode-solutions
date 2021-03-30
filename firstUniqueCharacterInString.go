func firstUniqChar(s string) int {
    if len(s) <= 1 {
        return 0
    }
    uniques := make(map[byte]byte)
    for i := 0; i < len(s); i++ {
        if _, exists := uniques[s[i]]; !exists {
            fmt.Println(string(s[i]))
            unique := true
            for j := i + 1; j < len(s); j++ {
                if s[j] == s[i] {
                    uniques[s[i]] = 1
                    unique = false
                    break
                }
            }
            if unique {
                return i
            }
        }
    }
    return -1
}
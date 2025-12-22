func romanToInt(s string) int {
    romanMap := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    result := 0
    n := len(s)

    for i := 0; i < n; i++ {
        currentVal := romanMap[s[i]]

        if i < n-1 && currentVal < romanMap[s[i+1]] {
            result -= currentVal
        } else {
            result += currentVal
        }
    }

    return result
}
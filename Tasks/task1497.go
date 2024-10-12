func canArrange(arr []int, k int) bool {
    n := len(arr)
    if n % 2 != 0 {
        return false
    }

    freq := make([]int, k)

    for _, num := range arr {
        remainder := ((num % k) + k) % k
        freq[remainder]++
    }

    if freq[0] % 2 != 0 {
        return false
    }

    for i := 1; i <= k/2; i++ {
        if freq[i] != freq[k-i] {
            return false
        }
    }

    if k % 2 == 0 && freq[k/2] % 2 != 0 {
        return false
    }

    return true
}
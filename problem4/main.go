package problem4

// Using recursion
// Time complexity: O(n), Space complexity: O(n)
func sum_to_n_a(n int) int {
    if n <= 0 {
        return 0
    }
    return n + sum_to_n_a(n-1)
}

// Using for loop
// Time complexity: O(n), Space complexity: O(1)
func sum_to_n_b(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum
}

// Using formula n*(n+1)/2
// Time complexity: O(1), Space complexity: O(1)
func sum_to_n_c(n int) int {
    return n * (n + 1) / 2
}
func lengthOfLongestSubstring(s string) int {
  if len(s) < 2 {
    return len(s)
  }

  input := []byte(s)
  seen := make(map[byte]bool)
  longest := 1

  for i := range len(input) {
    clear(seen)
    count := 0
    for j:=i; j < len(input); j++ {
      if (seen[input[j]]) {
        break
      }
      seen[input[j]] = true
      count++
    }
    if (count > longest){
      longest = count
    }
  }
  return longest
}
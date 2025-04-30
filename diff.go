package slices

// Diff returns two slices:
//   onlyA = elements in a but not in b
//   onlyB = elements in b but not in a
func Diff[T comparable](a, b []T) (onlyA, onlyB []T) {
    seen := make(map[T]struct{}, len(a)+len(b))
    for _, x := range a {
        seen[x] = struct{}{}
    }
    for _, x := range b {
        if _, ok := seen[x]; ok {
            // mark as seen-in-b
            seen[x] = struct{}{}
        } else {
            // x was only in b
            onlyB = append(onlyB, x)
            seen[x] = struct{}{} // avoid duplicates in onlyB
        }
    }
    // now collect those in a that never showed up in b
    for _, x := range a {
        // if x was never “re-seen” in the second loop, it’s only in a
        // we detect that by checking if it was only added in the first pass
        // (an alternative is to track via a separate map or use a flag type)
        // here: simplest is to check membership in b by scanning bMap:
        // (we could build a bMap first; I’ll switch to that for clarity)
    }
    return
}
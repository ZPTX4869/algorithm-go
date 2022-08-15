package dfs

func permute(nums []int) [][]int {
    res := make([][]int, 0)
    path := make([]int, 0)
    used := make([]bool, len(nums))

    var backtrack func()
    backtrack = func() {
        if len(path) == len(used) {
            res = append(res, append([]int{}, path...))
        }

        for i, v := range nums {
            if used[i] {
                continue
            }

            path = append(path, v)
            used[i] = true

            backtrack()

            path = path[:len(path)-1]
            used[i] = false
        }
    }

    backtrack()

    return res
}
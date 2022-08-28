package dfs

func combinationSum(candidates []int, target int) [][]int {
    sum := 0
    path := make([]int, 0)
    res := make([][]int, 0)

    var backtrack func(start int)
    backtrack = func(start int) {
        if sum > target {
            return
        }
        if sum == target {
            res = append(res, append([]int{}, path...))
            return
        }
        for i := start; i < len(candidates); i++ {
            sum += candidates[i]
            path = append(path, candidates[i])
            backtrack(i)
            sum -= candidates[i]
            path = path[:len(path)-1]
        }
    }

    backtrack(0)

    return res
}
// Given a binary tree t, find the sum of all the numbers encoded in it.
//
// Definition for binary tree:
// type Tree struct {
//   Value interface{}
//   Left *Tree
//   Right *Tree
// }
// hint: Using DFS, but we need one more entry to track current value.

func digitTreeSum(t *Tree) int64 {    
    return int64(transvers(t, 0))
}

func transvers(t *Tree, pre int) int{
    if t == nil {
        return 0
    }
    sum := 10*pre + t.Value.(int)
    if t.Left == nil && t.Right == nil{
        return sum
    }
    return transvers(t.Left, sum) + transvers(t.Right, sum)
}


package binary_search_tree

/*
@Author: by LH
@date:  2020/8/24
@function:
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//@category:	二叉搜索树
//@info:		面试题 17.12. BiNode
func convertBiNode(root *TreeNode) *TreeNode {
	var dummy = &TreeNode{}
	help17(root, dummy)
	return dummy.Right
}
func help17(root, cur *TreeNode) *TreeNode {
	if root == nil {
		return cur
	}
	cur = help17(root.Left, cur)
	cur.Right = root
	cur = root
	cur.Left = nil
	cur = help17(root.Right, cur)

	return cur
}

//@category:	二叉搜索树
//@info:		1038. 从二叉搜索树到更大和树
func bstToGst(root *TreeNode) *TreeNode {
	var cur = &TreeNode{}
	help1038(root, cur)
	return root
}

func help1038(root, cur *TreeNode) {
	if root == nil {
		return
	}
	help1038(root.Right, cur)
	root.Val += cur.Val
	cur.Val = root.Val
	help1038(root.Left, cur)
}

//@category:	二叉搜索树
//@info:		1382. 将二叉搜索树变平衡
func balanceBST(root *TreeNode) *TreeNode {
	ans := dfs1382(root)
	//造树
	if ans == nil {
		return nil
	}
	return help1382(ans)
}

func help1382(ans []*TreeNode) *TreeNode {
	var root *TreeNode
	if len(ans) == 0 {
		return root
	}
	if len(ans) == 1 {
		root = ans[0]
		root.Left = nil
		root.Right = nil
		return root
	}
	m := len(ans) / 2

	root = ans[m]
	root.Left = help1382(ans[:m])
	root.Right = help1382(ans[m+1:])
	return root
}

func dfs1382(root *TreeNode) []*TreeNode {
	var ans []*TreeNode
	if root == nil {
		return ans
	}
	ans = append(ans, dfs1382(root.Left)...)
	ans = append(ans, root)
	ans = append(ans, dfs1382(root.Right)...)
	return ans
}

//@category:	二叉搜索树
//@info:		1373. 二叉搜索子树的最大键值和
func maxSumBST(root *TreeNode) int {

	info := help1373(root)
	if info.Ans > 0 {
		return info.Ans
	}
	return 0
}

type Info1373 struct {
	Max   int
	Min   int
	BST   bool
	Total int
	Ans   int
}

func help1373(root *TreeNode) (ans *Info1373) {
	if root == nil {
		return nil
	}
	ans = &Info1373{
		Max:   root.Val,
		Min:   root.Val,
		BST:   true,
		Total: root.Val,
		Ans:   root.Val,
	}
	left := help1373(root.Left)
	right := help1373(root.Right)
	if left == nil && right == nil {
		//叶子结点
		return
	}
	ans.Ans = 0

	if left != nil && right != nil {
		ans.Total += left.Total + right.Total
		if !left.BST || !right.BST || left.Max >= root.Val || right.Min <= root.Val {
			ans.BST = false
			ans.Ans = max(left.Ans, right.Ans)
			return
		}
		ans.Min = left.Min
		ans.Max = right.Max

		ans.Ans = max(left.Ans, right.Ans)
		ans.Ans = max(ans.Ans, ans.Total)
		return
	}

	if left != nil {
		ans.Total += left.Total
		if !left.BST || left.Max >= ans.Min {
			ans.BST = false
			ans.Ans = left.Ans
		} else {
			ans.Min = left.Min
			ans.Ans = max(ans.Total, left.Ans)
		}
	}
	if right != nil {
		ans.Total += right.Total
		if !ans.BST || !right.BST || right.Min <= ans.Max {
			ans.BST = false
			ans.Ans = right.Ans
		} else {
			ans.Max = right.Max
			ans.Ans = max(right.Ans, ans.Total)
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

package binary_search_tree

import (
	"fmt"
	"testing"
)

/*
@Author: by LH
@date:  2020/8/24
@function:
*/
var req=&TreeNode{
	Val:   4,
	Left:  &TreeNode{
		Val:   2,
		Left:  &TreeNode{
			Val:   1,
			Left:  &TreeNode{
				Val:   0,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
		},
	},
	Right: &TreeNode{
		Val:   5,
		Left:  nil,
		Right: &TreeNode{
			Val:   6,
		},
	},
}

func TestConvertBiNode(t *testing.T){
	convertBiNode(req)
}

func TestBstToGst(t *testing.T){
	bstToGst(req)
}

func TestBalanceBST(t *testing.T){
	xx:=&TreeNode{
		Val:   1,
		Left:  nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	balanceBST(xx)
}

func TestMaxSumBST(t *testing.T){
	xx:=&TreeNode{
		Val:   4,
		Left:  &TreeNode{
			Val:   3,
			Left:  &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: nil,
	}
	fmt.Println(maxSumBST(xx))

	xx = &TreeNode{
		Val:   5,
		Left:  &TreeNode{
			Val:   4,
			Left:  &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(maxSumBST(xx))

	xx = &TreeNode{
		Val:   4,
		Left:  &TreeNode{
			Val:   8,
			Left:  &TreeNode{
				Val:   6,
				Left:  &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   1,
				Left:  &TreeNode{
					Val:   -5,
					Left:  nil,
					Right: &TreeNode{
						Val:   -3,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: &TreeNode{
						Val:   10,
						Left: nil,
						Right: nil,
					},
				},
			},
		},
		Right: nil,
	}

	fmt.Println(maxSumBST(xx))

	xx=&TreeNode{
		Val:   9,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{
					Val:   3,
					Left:  &TreeNode{
						Val:   -5,
						Left:  nil,
						Right: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: &TreeNode{
							Val:   10,
							Left:  nil,
							Right: nil,
						},
					},
				},
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(maxSumBST(xx))
}
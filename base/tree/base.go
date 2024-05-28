 type TreeNode struct {
	val int
	children []*TreeNode
 }

 func traverse(root *TreeNode) {
	for _,child: root.children {
		traverse(child)
	}
 }

 type BinaryTreeNode struct {
	val int 
	left *BinaryTreeNode
	right *BinaryTreeNode
}

func traverse(root *BinaryTreeNode){
	// 前序遍历
	traverse(root.left);
	// 中序遍历
	traverse(root.right);
	// 后序遍历
}
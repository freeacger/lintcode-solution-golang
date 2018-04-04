/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var tree []int
/**
 * @param root: A Tree
 * @return: Preorder in ArrayList which contains node values.
 */
func preorderTraversal (root *TreeNode) []int {
    tree = []int {}
    preorder(root)
    return tree
}

func preorder (node *TreeNode) {
    if(node == nil) {
        return
    }
    
    tree = append(tree, node.Val)
    preorder(node.Left)
    preorder(node.Right)
}

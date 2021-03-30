/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    // base case
    if root1 == nil && root2 == nil {
        return nil
    }
    
    mergedNode := new(TreeNode)
    
    if root1 != nil && root2 != nil {
        mergedNode.Val = root1.Val + root2.Val
        mergedNode.Left = mergeTrees(root1.Left, root2.Left)
        mergedNode.Right = mergeTrees(root1.Right, root2.Right)        
    } else if root1 != nil {
        mergedNode.Val = root1.Val
        mergedNode.Left = mergeTrees(root1.Left, nil)
        mergedNode.Right = mergeTrees(root1.Right, nil)
    } else {
        mergedNode.Val = root2.Val
        mergedNode.Left = mergeTrees(nil, root2.Left)
        mergedNode.Right = mergeTrees(nil, root2.Right)
    } 
    return mergedNode
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var result *ListNode = new(ListNode)
    ptrResult := result
    ptrOne := l1
    ptrTwo := l2
    for {
        if ptrOne == nil && ptrTwo == nil {
            break
        }
        carry, sum := 0, 0
        if ptrOne != nil && ptrTwo != nil {
            sum = ptrOne.Val + ptrTwo.Val            
        } else if ptrOne != nil {
            sum = ptrOne.Val
        } else {
            sum = ptrTwo.Val
        }
        if sum > 9 {
            sum %= 10
            carry = 1
        }
        if ptrResult.Val + sum > 9 {
            ptrResult.Val = (ptrResult.Val + sum) % 10
            carry++
        } else {
            ptrResult.Val += sum            
        }
        
        if ptrOne != nil {
            ptrOne = ptrOne.Next
        }
        if ptrTwo != nil {
            ptrTwo = ptrTwo.Next
        }
        
        if ptrOne == nil && ptrTwo == nil && carry == 0 {
            break
        }
        ptrResult.Next = new(ListNode)
        ptrResult = ptrResult.Next
        ptrResult.Val = carry
    }
    return result
}
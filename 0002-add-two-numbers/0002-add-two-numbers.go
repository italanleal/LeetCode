/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var head ListNode = ListNode{}
    var cur *ListNode = &head
    var rem int = 0

    for l1 != nil || l2 != nil || rem > 0 {
        v1 := 0
        v2 := 0

        if l1 != nil {
            v1 = l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        }

    sum := (v1 + v2 + rem) % 10
    rem = (v1 + v2 + rem) / 10
    
    cur.Next = &ListNode{}
    cur.Next.Val = sum
    cur = cur.Next
    }

    return head.Next
}
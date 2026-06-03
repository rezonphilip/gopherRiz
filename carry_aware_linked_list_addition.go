package gopherriz

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// our solution
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return recursiveAddWithCarry(l1, l2, 0)
}

func recursiveAddWithCarry(l1 *ListNode, l2 *ListNode, carryValue int) *ListNode {
	if l1 == nil && l2 == nil && carryValue == 0 {
		return nil
	}
	nodeVal := (getNodeVal(l1) + getNodeVal(l2) + carryValue) % 10
	carryVal := (getNodeVal(l1) + getNodeVal(l2) + carryValue) / 10
	l1NextNode := getNextNode(l1)
	l2NextNode := getNextNode(l2)

	return &ListNode{
		Val:  nodeVal,
		Next: recursiveAddWithCarry(l1NextNode, l2NextNode, carryVal),
	}
}
func getNextNode(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	return node.Next
}

func getNodeVal(node *ListNode) int {
	if node == nil {
		return 0
	}
	return node.Val
}

// another more efficient way to approach this
func addTwoNumbersOne(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	temp := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carry
		carry = sum / 10
		temp.Next = &ListNode{Val: sum % 10}
		temp = temp.Next
	}

	return dummy.Next
}

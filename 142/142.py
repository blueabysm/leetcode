# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def hasCycle(self, head):
        """
        :type head: ListNode
        :rtype: bool
        """
        if head == None or head.next == None:
            return None

        small_step_pointer = head
        big_step_pointer = head

        while big_step_pointer != None:
            small_step_pointer = small_step_pointer.next
            big_step_pointer = big_step_pointer.next
            if big_step_pointer != None:
                big_step_pointer = big_step_pointer.next
            if small_step_pointer == big_step_pointer:
                break

        if big_step_pointer == None:
            return None

        small_step_pointer = head
        while small_step_pointer != big_step_pointer:
            small_step_pointer = small_step_pointer.next
            big_step_pointer = big_step_pointer.next

        return big_step_pointer


head = ListNode(0)
current = head
temp = None

for i in range(1, 10):
    temp = current
    current.next = ListNode(i)
    current = current.next

temp.next = head.next.next

s = Solution()

print s.hasCycle(head).val

class Solution(object):
    __kth_smallest_val = None
    __k = 0
    def kthSmallest(self, root, k):
        """
        :type root: TreeNode
        :type k: int
        :rtype: int
        """
        self.__k = k
        self.iterate(root)
        return self.__kth_smallest_val

    def iterate(self, root):
        if root == None:
            return
        if root.left != None:
            self.iterate(root.left)
        self.visit(root)
        if root.right != None:
            self.iterate(root.right)

    def visit(self, node):
        if self.__kth_smallest_val == None:
            self.__kth_smallest_val = node.val
        if node.val >= self.__kth_smallest_val:
            if self.__k > 0:
                self.__kth_smallest_val = node.val
        self.__k = self.__k - 1

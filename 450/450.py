# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    __root = None

    def deleteNode(self, root, key):
        """
        :type root: TreeNode
        :type key: int
        :rtype: TreeNode
        """
        self.__root = root
        self.iterate(root, key)
        return self.__root

    def rebuild(self, root):
        if root.left == None and root.right == None:
            return
        if root.left == None:
            root.val = root.right.val
            root.right = None
            return
        root.val = root.left.val
        root.left = None

    def iterate(self, root, key):
        if root == None:
            return
        if root.val == key:
            self.__val = root.val
            self.rebuild(root)
            return
        if root.val < key:
            self.iterate(root.right, key)
        if root.val > key:
            self.iterate(root.left, key)

# Binary Tree

# Notes.
1. Every node has atmost 2 childs.
2. Nodes at level l => pow(2,l)
3. Height is number of edger from root to leave. Heigth of root is= 0
4. Maximum number of nodes of height "h" is N = pow(2,h+1)-1
5. Height of the tree of Nodes N => h=log(N,2)

Binary tree types:
1. Perfect Binary tree - All leaves are at same level and all internal nodes are filled.
2. Complete Binary tree - All levers are filled except last level. Last level is filled more left side. Useful for heaps.
3. Full binary tree - Each node has either 0 or 2 childs. Used in Segement tree.

If represented using array 'i' is parent left=2i+1 right=2i+2 . => i=l-1/2 or i=r-2/2

Traversals:
Breadth first 
1. Level oreder - traversal uses Queue.
Depth First
1. Pre-order
2. In-order
3. Post-order

Operations:
1. Insertion
2. Removal
3. Search or look up
4. Trversal

Other oeprations:
1. Height of the tree.
2. Level of the tree
3. Size of the entire tree.

# Problems
 1. Is given tree hight balanced?
 2. Is Symmetric?
 3. Lowest common ancestor(LCA)? Used for CSS on Web.
 4. Sum the root-to-node path in a tree.  Specifed sum
 5. Inorder/Preorder without using recursion.
 6. Construct tree from (in,pre), (in,post) traversals.
 7. Compute the exterior of the binary tree.
 8. Implementing logcking in a binary tree.
package sort

// 完全二叉树：叶子结点只会出现最后两层，且最后一层的叶子节点都靠左对齐；
// 完全二叉树的数组存储：数组下标为1是根结点，i的左孩子是2*i，右孩子是2*i+1，i的父节点是（i/2）向下取整；
// 堆：一个近似完全二叉树的结构，并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点；
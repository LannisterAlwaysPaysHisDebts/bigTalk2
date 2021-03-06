# 树
## 前置知识
### 等比数列
`q`是等比数列的公比,计算方式是某一项除以前一项. 二叉树的公比为2.

#### 通项公式:
`a(n) = a(1)*q^(n-1)`, 用于求一棵二叉树的第x层最多可以有多少个节点:
```go
for i := 1; i < 10; i++ {
	fmt.Printf("deep: %d ,floor count: %.f\n", i, math.Pow(2.0, float64(i-1)))
}
```

#### 前n项和:
q=1时`S(n)=n*a(1)`, q > 1时: `S(n) = (a(1)*(1-q^n))/(1-q)`. 二叉树可以缩写为`S(n)= 2^n - 1`
```go
for i := 1; i < 10; i++ {
	// S(n) = (a(1)*(1-q^n))/(1-q)  
	count := math.Pow(2, float64(i)) - 1
	fmt.Printf("deep: %d count: %.f\n", i, count)
}
```


## 二分查找法
注意:
1. 边界问题；
2. int相加的溢出问题；
3. 递归的写法;
4. floor与ceil的实现

## 二分搜索树
### 概念
二分搜索树：若左子树不为空，左子树上所有节点都比根节点小；若右子树不为空，右子树上所有节点都比根节点大。

递归算法的内容有 : 1. 递归终止条件;2. 递归调用逻辑;
可以知道二叉树的定义天然递归: 1. 递归终止条件: "若左子树/右子树不为空"; 2. 递归调用逻辑: 左子树/右子树所有节点都比根节点小/大

所以二叉树的所有操作都可以用递归来实现 

### 时间复杂度
普通数组：插入：O(1)(直接读取key值) 查找删除: O(n)(需要遍历数组);
顺序数组：插入与删除: O(n)(需要遍历数组),查找: O(logn)(二分搜索)
二分搜索树: 插入、删除、查找： O(logn) (因为都是类似于二分查找)
哈希表: 插入、删除、查找：O(1)，但是哈希表没有顺序性

### 二分搜索树的优势
相对hash表，二分搜索树存在顺序性，因此可以进行一些hash表不能进行的操作:
1. 可以找到一个元素的前驱`successor`后继`predecessor`
2. 可以找到最大值`maximum`最小值`minimum`
3. 可以找到`floor(k int)`与`ceil(k int)`:floor是小于k的最大值，ceil是大于k的最小值 
4. `rank`与`select`: rank: 找到元素是排名第几的元素（需要每个节点加上属性:以这个节点为根的二分搜索1树一共有几个节点）; select:排名第x的元素是谁； 

### 二分搜索树的局限性
不平衡, 极端情况下相当于一个链表;

### 实现思路
[见代码](./bst.go), 值得注意的有: 
1. 二叉树是递归结构,方法编写可以充分利用递归性质;
2. 一般二叉树提供查找、插入、删除的功能;
3. 查找利用递归的思想和二叉树左子树小右子树大的性质,对当前node的值进行比较,比node大就对右子树进行比较,比node小就对左子树进行比较;
4. 插入操作和查找一样,递归比较,直到到达合适位置的nil结点;
5. 删除操作需要注意的是: 叶子结点可以直接删除、非叶子结点需要查找结点右子树里的最小值对该结点进行替换(当然左子树里的最大值也行);


## 红黑树
### 概念
在二分搜索树的基础上,每个结点增加颜色参数,并使其遵循以下规则:
1. 每个节点不是红色就是黑色;
2. 树的根始终是黑色的, NULL结点也是黑色的;
3. 红色结点的子结点一定是黑色(也就是说按照父子子的结构，只可能出现： 红黑黑、黑红黑、黑黑红、黑黑黑四种情况)
4. *从任意节点（包括根）到其任何后代NULL节点的每条路径都具有相同数量的黑色节点*
5. 从4也可以推出: *如果一个结点存在黑子结点，那么该结点肯定有两个子结点*

这种树就是红黑树;红黑树是一棵平衡二叉树;

### 时间复杂度
查找: 最好大致能做到O(logn),因为是一棵平衡的二叉树

### 实现方法
[见代码](./rbt.go)(仅限思路参考);
1. 红黑树的实现就是不停维护该二叉树符合红黑树的性质;
2. 因为查找不会对结点做操作,所以红黑树的查找和二分搜索树是一样的;
3. 新增操作,除非是新增根结点,否则默认都是红色的(保证路径上的黑色结点数量不变);
4. 新增操作的新增逻辑和二分搜索树是一样的, 重点是维护红黑树的性质;
5. 大致总结为:不影响平衡的直接插入;根据叔结点颜色的不同,通过旋转(左旋/右旋)或者把问题往上抛(自底向上)来维持平衡;
6. 删除操作是红黑树最复杂的操作;先和二分搜索树一样将结点通过替换的方式删除,再来维护红黑树的性质;
7. 大致可以总结为: 不影响平衡的直接删除; 不然就找兄弟结点通过旋转移一个黑色补过来; 再不行就把问题网上抛;

### 使用场景
虚拟内存、java的hashMap等等, 可以替换二分搜索树;

## b-tree
b树是多叉树,遵循规则:
1. 所有节点关键字是按递增次序排列，并遵循左小右大原则;
2. 节点关键字个数大于等于ceil(m/2)-1小于等于M-1,根结点可以只有1个关键字;

### 优势
B树相对于平衡二叉树的不同是，每个节点包含的关键字增多了，特别是在B树应用到数据库中的时候:数据库充分利用了磁盘块的原理（磁盘数据存储是采用块的形式存储的，每个块的大小为4K，每次IO进行数据读取时，同一个磁盘块的数据可以一次性读取出来）把节点大小限制和充分使用在磁盘快大小范围；树的节点关键字增多后树的层级比原来的二叉树少了，减少数据查找的次数和复杂度;

## b+ tree
相对于B树来说B+树更充分的利用了节点的空间，让查询速度更加稳定，其速度完全接近于二分法查找。

1. B+跟B树不同B+树的非叶子节点不保存关键字记录的指针，只进行数据索引，这样使得B+树每个非叶子节点所能保存的关键字大大增加;
2. B+树叶子节点保存了父节点的所有关键字记录的指针，所有数据地址必须要到叶子节点才能获取到。所以每次数据查询的次数都一样;
3. B+树叶子节点的关键字从小到大有序排列，左边结尾数据都会保存右边节点开始数据的指针;
4. 非叶子节点的子节点数=关键字数（来源百度百科）（根据各种资料 这里有两种算法的实现方式，另一种为非叶节点的关键字数=子节点数-1（来源维基百科)，虽然他们数据排列结构不一样，但其原理还是一样的Mysql 的B+树是用第一种方式实现）;




## 树的学习大纲:
1. 二分查找法（Binary Search）
2. 二分搜索树基础（Binary Search Tree）
3. 二分搜索树的节点插入
4. 二分搜索树的查找
5. 二分搜索树的遍历（深度优先遍历）
6. 层序遍历（广度优先遍历）
7. 删除最大值，最小值
8. 二分搜索树节点的删除（Hubbard Deletion）
9. 二分搜索树的顺序性
10. 二分搜索树的局限性
11. 二分搜索法的floor和ceil
12. 二分搜索法的lower bound和upper bound
13. 二分搜索树中的floor和ceil
14. 二分搜索树中的前驱和后继
15. 二分搜索树中的rank和select
16. 二分搜索树前中后序非递归遍历
    深入理解非递归和递归的区别，以及非递归和栈的关系
17. 二叉树前中后序遍历的经典非递归实现
18. 二叉树的Morris前中后序遍历
19. 二分搜索树整体的非递归实现
20. 二分搜索树的另一个应用：Tree Set
21. 允许重复键值的二分搜索树：Multi Tree Set / Map
22.  二叉树的公共祖先 (LCA)
23. 树形问题和回溯法
24. 树形问题之八皇后问题
25. 线段树 (区间树)
26. Trie
27. KD树
28. 哈夫曼树
29. 使用哈夫曼树进行文件压缩
30. AVL树
31. 红黑树
32. 伸展树
33. B类树
34. Treap




type BSTNode struct{
  left *BstNode
  val int
  right *BstNode
}

func (node *BstNode) Add(value int) {
	if value < node.value {
		if node.left == nil {
			node.left = &BstNode{value: value}
		} else {
			node.left.Add(value)
		}
	} else {
		if node.right == nil {
			node.right = &BstNode{value: value}
		} else {
			node.right.Add(value)
		}
	}
}

func (node *BSTNode) Min() int{
	if node == nil{
		return 0
	}
	if node.left == nil{
		return node.val
	}
	return node.left.Min()
}

func (node *BSTNode) Max() int{
	if node == nil{
		return 0
	}
	if node.right == nil{
		return node.val
	}
	return node.right.Max()
}

func (node *BSTNode) PreOrderNav(){
	//RED
	if node != nil{
		print(node.val," ")
		node.left.PreOrderNav()
		node.right.PreOrderNav()
	}
}

func (node *BSTNode) InOrderNav(){
	//ERD
	if node != nil{
		node.left.InOrderNav()
		print(node.val, " ")
		node.right.InOrderNav()
	}
}

func (node *BSTNode) PosOrderNav(){
	//EDR
	if node != nil{
		node.left.PosOrderNav()
		node.right.PosOrderNav()
		print(node.val, " ")
	}
}

func (node *BSTNode) LevelOrderNav(){
  queue := make([]*BSTNode, 0)
  queue = append(queue, node)
  for len(queue) != 0 {
    curNode := queue[0]
    queue = queue[1:]
    print(curNode.val,", ")
    if curNode.left != nil {
      queue = append(queue, curNode.left)
    }
    if curNode.right != nil {
      queue = append(queue, curNode.right)
    }
  }
}

func (node *BSTNode) Height() int{
  if node == nil{
    return -1
  }
  leftHeight := node.left.Height()
  rightHeight := node.right.Height()
  if leftHeight > rightHeight {
    return 1 + leftHeight
  }
  return 1 + rightHeight
}

func (node *BSTNode) Par() int{
	if node == nil{
		return 0
	}
	if node.val % 2 == 0{
		return 1+node.left.Par()+node.right.Par()
	}
	return 0+node.left.Par()+node.right.Par()
}

func (node *BSTNode) Size() int{
	if node == nil{
		return 0
	}
	return 1 + node.left.Size() + node.right.Size()
}

func (node *BSTNode) isBst() bool {
	if node == nil {
		return true
	}
	if node.left != nil && node.left.Max() >= node.val{
		return false
	}
	if node.right != nil && node.right.Min() <= node.val{
		return false
	}
	return node.left.isBst() && node.right.isBst()
}

func convertToBalancedBst(v []int, ini int, fim int) *BSTNode {
	if ini > fim{
		return nil
	}
	meio := (ini+fim)/2
	return &BstNode {
		value: v[meio],
		left: convertToBalancedBst(v, ini, meio-1),
		right: convertToBalancedBst(v, meio+1, fim),
		}
}

func (node *BSTNode) Remove(val int) *BSTNode{
  if val < node.val{
    node.left = node.left.Remove(val)
	return node
  } else if val > node.val{
    node.right = node.right.Remove(val)
	return node
  } else { // encontramos o nó
    if node.left == nil && node.right == nil{
      //caso 1: nó folha
      return nil // esse é o novo valor que o pai do nó folha apontará
    } else if node.left != nil && node.right == nil{
      //caso 2: há 1 filho a esquerda
      return node.left
    } else if node.left == nil && node.right != nil{
      //caso 2: há um filho a direita
      return node.right
    } else {
      // caso 3: 2 filhos
      min := node.right.Min()
      node.val = min
      node.right = node.right.Remove(min)
      return node
    }
  }
}
    
    
  

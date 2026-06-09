type BSTNode struct{
  left *BstNode
  val int
  right *BstNode
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

func (node *BSTNode) NumPar() int{
  if node == nil{
    return 0
  }
  if node.val % 2 == 0 {
    return 1 + node.left.NumPar() + node.right.NumPar()
  } else {
    return 0 + node.left.NumPar() + node.right.NumPar()
  }
}

func (node *BSTNode) Min() int {
	if node == nil {
		return 0 
	}
	current := node
	// Enquanto houver um filho à esquerda, continue descendo
	for current.left != nil {
		current = current.left
	}
	// Quando o 'for' acabar, encontramos o nó mais à esquerda
	return current.val
}

func (node *BSTNode) Remove(val int) *BSTNode{
  if val < node.val{
    node.left = node.left.Remove(val)
  } else if val > node.val{
    node.right = node.right.Remove(val)
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
    
    
  

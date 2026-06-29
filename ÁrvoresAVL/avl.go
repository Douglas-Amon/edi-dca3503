import "fmt"

package main

type BstNode struct{
  left *BstNode
  val int
  right *BstNode
  bf int
  height int
}

//Rotação para direita
func (root *BstNode) RotRight() *BstNode {
  newRoot := root.left
  root.left = newRoot.right
  root.UpdateProperties()
  newRoot.UpdateProperties()
  return newRoot
}

//Rotação para esquerda
func (root *BstNode) RotLeft *BstNode {
  newRoot := root.right
  root.right = newRoot.left
  root.UpdateProperties()
  newRoot.UpdateProperties()
  return newRoot
}

//Atualiza as propriedades de fator de balanço e altura
func (root *BstNode) UpdateProperties() {
  if root.left == nil && root.right == nil{
    root.height = 0
    root.bf = 0
  } else {
    hleft := 0, hright := 0
    if root.left != nil {hleft = root.left.height}
    if root.right != nil {hright = root.right.height}
    if hleft >= height {root.height = hleft+1}
    else {root.height = hright+1}
    root.bf = hright - hleft
  }
}

//Caso 1: esquerda-esquerda
func (root *BstNode) RebalanceLeftLeft() *BstNode{
  return root.RotRight()
}

//Caso 2: esquerda-neutro
func (root *BstNode) RebalanceLeftNeutral() *BstNode{
  return root.RotRight()
}

//Caso 3: esquerda-direita
func (root *BstNode) RebalanceLeftRight() *BstNode{
  root.left = root.left.RotLeft()
  return root.RotRight()
}

//Caso 4: direita-direita
func (root *BstNode) RebalanceRightRight() *BstNode{
  return root.RotLeft()
}

//Caso 5: direita-neutro
func (root *BstNode) RebalanceRightNeutral() *BstNode{
  return root.RotLeft()
}

//Caso 6: direita-esquerda
func (root *BstNode) RebalanceRightLeft() *BstNode{
  root.right = root.right.RotRight()
  return root.RotLeft()
}

func (root *BstNode) Rebalance() *BstNode{
  if root.bf == -2 { //desbalanceamento à esquerda
     if root.left.bf == -1 {
       return root.RebalanceLeftLeft()
     } else if root.left.bf == 0 {
       return root.RebalanceLeftNeutral()
     } else if root.left.bf == 1 {
       return root.RebalanceLeftRight()
     }
  } else if root.bf == 2 { //desbalanceamento à direita
    if root.right.bf == 1 {
      return root.RebalanceRightRight()
    } else if root.right.bf == 0 {
      return root.RebalanceRightNeutral()
    } else if root.right.bf == -1 {
      return root.RebalanceRightLeft()
    }
  }
  return root
}

func (root *BstNode) Add(val int) *BstNode{
  if val <= root.val{
    if root.left != nil{
      root.left = root.left.Add(val)
    } else {
      root.left = createBSTNode(val)
    }
  } else { 
    if root.right != nil {
      root.right = root.right.Add(val)
    } else {
      root.right = createBSTNode(val)
    }
  }
  root.UpdateProperties()
  return root.Rebalance()
}
    

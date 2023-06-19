package path

type Node struct {
	landform *Landform
	parent   *Node
	cost     float64
}

package path

type Path struct {
	points       []*Landform
	currentIndex int
}

func (slf *Path) GetPoints() []*Landform {
	return slf.points
}

func (slf *Path) GetCurrentIndex() int {
	return slf.currentIndex
}

package adjacent_area

import "container/list"

type matrixCell struct {
	i, j int
}

type Calculator struct {
	matrix      [][]int
	areasNumber int
}

func NewCalculator(arr [][]int) Calculator {
	return Calculator{
		matrix: arr,
	}
}

// Calculate calculates a number of adjacent areas.
func (c *Calculator) Calculate() int {
	for i := 0; i < len(c.matrix); i++ {
		row := c.matrix[i]
		for j := 0; j < len(row); j++ {
			if row[j] == 0 {
				continue
			}

			cell := matrixCell{
				i: i,
				j: j,
			}
			queue := list.New()
			queue.PushBack(cell)

			c.search(queue)
		}
	}

	return c.areasNumber
}

func (c *Calculator) search(queue *list.List) {
	tryToEnqueue := func(i, j int) {
		if i < 0 || i > len(c.matrix)-1 {
			return
		}

		if j < 0 || j > len(c.matrix[i])-1 {
			return
		}

		if c.matrix[i][j] == 0 {
			return
		}

		queue.PushBack(matrixCell{
			i: i,
			j: j,
		})
	}

	var prev *list.Element
	for e := queue.Front(); ; e = e.Next() {
		if prev != nil {
			queue.Remove(prev)
		}

		if e == nil {
			break
		}
		cell := e.Value.(matrixCell)

		c.matrix[cell.i][cell.j] = 0

		//top
		tryToEnqueue(cell.i-1, cell.j)
		//bottom
		tryToEnqueue(cell.i+1, cell.j)
		//left
		tryToEnqueue(cell.i, cell.j-1)
		//right
		tryToEnqueue(cell.i, cell.j+1)

		prev = e
	}
	c.areasNumber++
}

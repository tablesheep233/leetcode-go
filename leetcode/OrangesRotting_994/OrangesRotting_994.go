package OrangesRotting_994

func orangesRotting(grid [][]int) int {
	count, minute := 0, 0
	queue := [][2]int{}
	for i, g := range grid {
		for j, n := range g {
			if n == 1 {
				count++
			} else if n == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	x := [4]int{0, 0, 1, -1}
	y := [4]int{1, -1, 0, 0}
	for count > 0 && len(queue) != 0 {
		minute++
		l := len(queue)
		for s := 0; s < l; s++ {
			o := queue[0]
			queue = queue[1:]
			for i := 0; i < 4; i++ {
				newX := o[0] + x[i]
				newY := o[1] + y[i]
				if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) && grid[newX][newY] == 1 {
					grid[newX][newY] = 2
					count--
					queue = append(queue, [2]int{newX, newY})
				}
			}
		}
	}
	if count == 0 {
		return minute
	}
	return -1
}

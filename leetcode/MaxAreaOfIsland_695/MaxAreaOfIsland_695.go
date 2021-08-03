package MaxAreaOfIsland_695

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i, item := range grid {
		for j, g := range item {
			if g == 1 {
				t := bfs(grid, i, j)
				if max < t {
					max = t
				}
			}
		}
	}
	return max
}

func bfs(image [][]int, sr int, sc int) int {
	i := 0
	if image[sr][sc] == 1 {
		i++
		image[sr][sc] = 0
		if sr+1 < len(image) {
			i += bfs(image, sr+1, sc)
		}
		if sr-1 >= 0 {
			i += bfs(image, sr-1, sc)
		}
		if sc+1 < len(image[0]) {
			i += bfs(image, sr, sc+1)
		}
		if sc-1 >= 0 {
			i += bfs(image, sr, sc-1)
		}
	}
	return i
}

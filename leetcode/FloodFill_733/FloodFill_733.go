package FloodFill_733

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] != newColor {
		bfs(image, sr, sc, image[sr][sc], newColor)
	}
	return image
}

func bfs(image [][]int, sr int, sc int, color int, newColor int) {
	if image[sr][sc] == color {
		image[sr][sc] = newColor
		if sr+1 < len(image) {
			bfs(image, sr+1, sc, color, newColor)
		}
		if sr-1 >= 0 {
			bfs(image, sr-1, sc, color, newColor)
		}
		if sc+1 < len(image[0]) {
			bfs(image, sr, sc+1, color, newColor)
		}
		if sc-1 >= 0 {
			bfs(image, sr, sc-1, color, newColor)
		}
	}
}

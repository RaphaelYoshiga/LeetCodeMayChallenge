package main;

var m,n int
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	previousColour:=  image[sr][sc];

	if newColor == previousColour{
		return image;
	}

	m = len(image);
	n = len(image[0])

	fill(image, sr, sc, newColor, previousColour);
	return image;
}

func fill(image [][]int, x int, y int, newColor int, previousColour int){

	if x < 0 || x >= m || y < 0 || y >=n || image[x][y] != previousColour{
		return;
	}

	image[x][y] = newColor;

	fill(image, x + 1, y, newColor, previousColour) // Bottom;
	fill(image, x - 1, y, newColor, previousColour) // Top;
	fill(image, x, y + 1, newColor, previousColour) // right;
	fill(image, x, y - 1, newColor, previousColour) // Left;
}



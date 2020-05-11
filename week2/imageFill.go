package main;

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	previousColour:=  image[sr][sc];

	if newColor == previousColour{
		return image;
	}

	fill(image, sr, sc, newColor, previousColour);
	return image;
}

func fill(image [][]int, x int, y int, newColor int, previousColour int){

	if x < 0 || x >= len(image) || y < 0 || y >=len(image[0]){
		return;
	}

	if image[x][y] != previousColour{
		return;
	}else{
		image[x][y] = newColor;

		fill(image, x + 1, y, newColor, previousColour) // Bottom;
		fill(image, x - 1, y, newColor, previousColour) // Top;
		fill(image, x, y + 1, newColor, previousColour) // right;
		fill(image, x, y - 1, newColor, previousColour) // Left;
	}
}


func main(){
	floodFill([][]int{{0,0,0}, {0,1,1}}, 1, 1, 1)
}
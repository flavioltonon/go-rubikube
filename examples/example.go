package main

import (
	rubik "flavioltonon/go-rubikube"
)

/*	 ___   ___	    ___		    ___				    ___	   ___
 *	|	| |   |	   |   | |	 | | __|  | | /	 |	 | | __|  |__
 *	|  __ |   | __ |__/  |	 | |    \ | |/\	 |	 | |    \ |
 *	|___| |___|	   |  \  |___| |____| | |  \ |___| |____| |___
 *            __________________
 *			 /     /     /     /|
 *			/_____/_____/_____/ |
 *		   /	 /	   /     /| |
 *		  /_____/_____/_____/ |/|
 *		 /	   /	 /	   /| / |
 *		/_____/_____/_____/ |/|/|
 *		|	  |  	|  	  | / / /
 *		|_____|_____|_____|/|/|/
 *		|	  |		|  	  | / /
 *		|_____|_____|_____|/|/
 *		|	  |		|  	  | /
 *		|_____|_____|_____|/
 */

func main() {
	var cube = rubik.NewCube()

	cube.Front().Print().Tridimensional()

	cube.Right().Rotate().CounterClockwise()
	cube.Right().Rotate().CounterClockwise()
	cube.Left().Rotate().Clockwise()
	cube.Left().Rotate().Clockwise()
	cube.Front().Rotate().CounterClockwise()
	cube.Front().Rotate().CounterClockwise()
	cube.Back().Rotate().Clockwise()
	cube.Back().Rotate().Clockwise()
	cube.Up().Rotate().CounterClockwise()
	cube.Up().Rotate().CounterClockwise()
	cube.Down().Rotate().Clockwise()
	cube.Down().Rotate().Clockwise()

	cube.Front().Print().Bidimensional()
	cube.Back().Print().Bidimensional()
	cube.Left().Print().Bidimensional()
	cube.Right().Print().Bidimensional()
	cube.Up().Print().Bidimensional()
	cube.Down().Print().Bidimensional()

	cube.Front().Print().Tridimensional()

	// cube.Back().MainColor()
	// cube.Left().MainColor()
	// cube.Right().MainColor()
	// cube.Up().MainColor()
	// cube.Down().MainColor()

	// cube.Shuffle(30)
	// cube.Front().Print().Tridimensional()

	// cube.Left().Rotate().Clockwise()
	// cube.Front().Position(0, 0)
	// cube.Front().Print().Bidimensional()

}

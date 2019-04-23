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

	cube.Front().MainColor()
	cube.Back().MainColor()
	cube.Left().MainColor()
	cube.Right().MainColor()
	cube.Up().MainColor()
	cube.Down().MainColor()

	cube.Shuffle(30)
	cube.Front().Print().Tridimensional()

	cube.Front().Print().Bidimensional()

}

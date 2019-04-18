package main

import rubik "github.com/flavioltonon/go-rubikube"

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
	// var front = face{
	// 	colors: [][]color{
	// 		{blue, blue, blue},
	// 		{red, white, blue},
	// 		{green, blue, yellow},
	// 	},
	// }

	var cube = rubik.NewCube()

	cube.Print().Front()
	cube.Print().Back()
	cube.Print().Left()
	cube.Print().Right()
	cube.Print().Up()
	cube.Print().Down()

	clockwise := true
	cube.Rotate(clockwise).Front()

	cube.Print().Front()
	cube.Print().Back()
	cube.Print().Left()
	cube.Print().Right()
	cube.Print().Up()
	cube.Print().Down()

	// cube.Rotate(true).Front()

	// log.Println(cube.left)
}

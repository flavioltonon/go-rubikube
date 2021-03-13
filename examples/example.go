package main

import (
	"fmt"
	"time"

	"github.com/flavioltonon/go-rubikube/internal/rubikube"
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
	cube := rubikube.New()

	cube.Rotate().Front().Clockwise()
	cube.Rotate().Back().Clockwise()
	cube.Rotate().Right().Clockwise()
	cube.Rotate().Up().Clockwise()
	cube.Rotate().Front().Clockwise()
	// cube.Rotate().Back().Clockwise()

	now := time.Now()
	solution := cube.Solve()
	solution.Print()

	fmt.Printf("Done (took %s)!\n", time.Since(now))
}

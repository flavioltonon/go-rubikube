package main

import (
	rubik "github.com/flavioltonon/go-rubikube"
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
	if rubik.Run() != nil {
		panic("something bad happened")
	}
}

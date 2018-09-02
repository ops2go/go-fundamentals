package main

import (
	"fmt"
)

/* Maps are Go’s built-in associative data type
(sometimes called hashes or dicts in other languages).

To create an empty map, use the builtin make:
make(map[key-type]val-type)

Set key/value pairs using typical name[key] = val syntax.

Get a value for a key with name[key]*/

func main() {
	//key type is string
	//val type is int
	favoritemovies := make(map[string]int)
	fm := favoritemovies

	fm["The Departed"] = 1
	fm["Fight Club"] = 2
	fm["The Hurt Locker"] = 3
	fm["Shutter Island"] = 4

	fmt.Println("Number of Favorites:", len(fm), "\n", "Favorite Movies:", fm)

	usernamepassword := make(map[string]string)
	auth := usernamepassword

	auth["colemanword"] = "Password00"
	auth["laceejansson34"] = "Winter2015"

	fmt.Println("# of Users:", len(auth), "\n", "Details:", auth)
}

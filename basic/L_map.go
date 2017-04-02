package main

import "fmt"

type User struct {
	name string
	age int
}

var m map[string]User

func main() {
	m = make(map[string]User)
	m["biezhi"] = User{"王爵nice", 25}
	m["rose"] = User{"萝丝", 20}

	fmt.Println(m["biezhi"])
	fmt.Println(m["rose"])

	fmt.Println("========邪恶的分割线========")

	var m2 = map[string]User{
		"biezhi":{"王爵nice", 25},
		"rose":{"萝丝", 20},
	};
	fmt.Println(m2)

	fmt.Println("========邪恶的分割线========")

	delete(m2, "biezhi")
	fmt.Println(m2)

	fmt.Println("========邪恶的分割线========")

	m2["biezhi"] = User{"王爵nice2", 22}
	fmt.Println(m2)
}

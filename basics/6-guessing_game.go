package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sourse := rand.NewSource(time.Now().Unix())
	random := rand.New(sourse)

	//generate the num 1 => 100
	target := random.Intn(100) + 1

	//start of game
	fmt.Println("I chosen num between 1 to 100")
	fmt.Println("can u guess it?")

	var guess int

	for {
		fmt.Println("enter num if u like to play")
		fmt.Scanln(&guess)

		if guess == target {
			println("u won ")
			break
		} else if guess < target {
			println("ur num is low ")
			continue
		} else {
			println("ur num is higher ")
			continue
		}

	}
}

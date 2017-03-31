package main
	
import (
	"fmt"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
        "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
    }

	distribution = make(map[string]int, len(users))
)

func main() {

	coinsForUser := func(name string) int {
		var total int
		for i := 0; i < len(name); i++ {
			switch string(name[i]) {
			case "a", "A":
				total++
			case "e", "E":
				total++
			case "i", "I":
				total += 2
			case "o", "O":
				total += 3
			case "u", "U":
				total += 4
			}
		}
		return total
	}

	for _, user := range users {
		v := coinsForUser(user)
		if v > 10 {
			v = 10
		}
		distribution[user] = v
		coins -= v
	}

	fmt.Println(distribution)
	fmt.Println("Coins left: ", coins)
}
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

type GameResult struct {
	Result   bool
	Date     string
	Attempts int
}

func chooseDifficulty(scanner *bufio.Scanner) (int, int) {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║    🎯 УГАДАЙ ЧИСЛО 🎯        ║")
	fmt.Println("╠══════════════════════════════╣")
	fmt.Println("║  1) 🟢 Easy                  ║")
	fmt.Println("║  2) 🟡 Medium                ║")
	fmt.Println("║  3) 🔴 Hard                  ║")
	fmt.Println("╚══════════════════════════════╝")
	fmt.Print("Выберите уровень (1-3): ")
	scanner.Scan()
	a := scanner.Text()
	switch a {
	case "1":
		return 50, 15
	case "3":
		return 200, 5
	default:
		return 100, 10
	}
}

func generateNumber(maxNumber int) int {
	return rand.IntN(maxNumber) + 1
}

func readNumber(scanner *bufio.Scanner, i int) (int, bool) {
	for {
		fmt.Print(color.YellowString("Попытка #%d - Введите число: ", i))
		if !scanner.Scan() {
			return 0, false
		}
		userNumber, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Только числа")
			continue
		}
		return userNumber, true
	}
}

func compare(userNumber, secretNumber int) bool {
	switch {
	case userNumber == secretNumber:
		return true
	case userNumber < secretNumber:
		fmt.Println("Секретное число больше👆")
	default:
		fmt.Println("Секретное число меньше👇")
	}
	return false
}

func hint(userNumber, secretNumber int) {
	diff := userNumber - secretNumber
	if diff < 0 {
		diff = -diff
	}
	switch {
	case diff <= 5:
		fmt.Println("🔥 Горячо")
	case diff <= 15:
		fmt.Println("🙂 Тепло")
	default:
		fmt.Println("❄️ Холодно")
	}
}

func playGame(scanner *bufio.Scanner) {
	maxNumber, maxAttempts := chooseDifficulty(scanner)
	secretNumber := generateNumber(maxNumber)
	fmt.Println("Игра 'Угадай число' - от 1 до", maxNumber, "началась!")
	fmt.Println("Угадайте число за", maxAttempts, "попыток!")
	attempts := make([]int, 0)

	for i := 1; i <= maxAttempts; i++ {
		userNumber, ok := readNumber(scanner, i)
		if !ok {
			return
		}
		attempts = append(attempts, userNumber)

		if compare(userNumber, secretNumber) {
			color.Green("Вы угадали!🙌")
			color.Green("Игра закончена!")
			history(attempts)
			saveResult(GameResult{Result: true, Date: time.Now().Format("2006-01-02 15:04:05"), Attempts: len(attempts)})
			return
		}
		hint(userNumber, secretNumber)
	}

	color.Red("Вы проиграли!😢")
	color.Red("Секретное число было: %d", secretNumber)
	history(attempts)
	saveResult(GameResult{Result: false, Date: time.Now().Format("2006-01-02 15:04:05"), Attempts: len(attempts)})
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		playGame(scanner)
		fmt.Print("Хотите сыграть еще раз? (y/n) ")
		scanner.Scan()
		if scanner.Text() != "y" {
			break
		}
	}
}

func history(attempts []int) {
	fmt.Println("ИСТОРИЯ ПОПЫТОК")
	for idx, num := range attempts {
		fmt.Println("Попытка", idx+1, "Число", num)
	}
}

func saveResult(r GameResult) {
	data, _ := json.Marshal(r)
	os.WriteFile("results.json", data, 0644)
}

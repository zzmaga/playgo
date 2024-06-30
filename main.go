// guess - игра, в которой игрок должен угадать случайное число.

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	seconds := time.Now().Unix() // Получаем текущую дату и время в виде целого числа.
	rand.Seed(seconds)           // Инициализируем генератор случайных чисел
	target := rand.Intn(100) + 1 // Генерируем целое число от 1 до 100
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	//fmt.Println(target)

	reader := bufio.NewReader(os.Stdin) // Чтение ввода с клавиатуры
	success := false                    // Настроить программу, чтобы по умолчанию выводилось сообщение о проигрыше.
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

		fmt.Print("Make a guess: ")           // Запрашиваем число
		input, err := reader.ReadString('\n') // Прочитать данные, введенные пользователем до нажатия Enter.

		// Если произошла ошибка, программа выводит сообщение и завершается
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  // Удаление символа новой строки
		guess, err := strconv.Atoi(input) // Входная строка преобразуется в целое число

		// Если произошла ошибка, программа выводит сообщение и завершается
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. You guess was LOW.") // Если введенное значение меньше загаданного
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.") // Если введенное число больше загаданного
		} else { // В противном случае введенное значение должно быть правильным...
			success = true // Предотвращает вывод о сообщения о проигрыше
			fmt.Println("Good job! You guessed it!")
			break // Выход из цикла
		}
	}
	if !success { // Если переменная success равна False, сообщить игроку загаданное число.
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}

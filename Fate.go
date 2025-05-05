package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Задай свой вопрос. Примечание: вопрос должен заканчиваться '?'")
	fmt.Println("Можешь указать к какому ответу больше склоняетесь. Будешь?")

	var answerPref int
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		answer := strings.ToLower(strings.TrimSpace(input))

		if answer != "да" && answer != "нет" {
			fmt.Println("Напиши да или нет")
			continue
		}

		if answer == "да" {
			for {
				fmt.Println("Хорошо. Выбери к какому ответу склоняешься больше всего")
				fmt.Println("Нужно выбрать целое число от 1 до 99, где 0 - нет, а 100 - да")
				fmt.Print("> ")
				input, _ := reader.ReadString('\n')
				num, err := strconv.Atoi(strings.TrimSpace(input))
				if err == nil && num >= 1 && num <= 99 {
					answerPref = num
					break
				} else {
					fmt.Println("Введи целое число от 1 до 100")
				}
			}
		}
		break
	}

	// Варианты ответов
	positive := []string{"Бесспорно", "Предрешено", "Никаких сомнений", "Определённо да", "Можешь быть уверен в этом"}
	indecPositive := []string{"Мне кажется - да", "Вероятнее всего", "Хорошие перспективы", "Знаки говорят - да", "Да"}
	neutral := []string{"Пока неясно, попробуй снова", "Спроси позже", "Лучше не рассказывать", "Сейчас нельзя предсказать", "Сконцентрируйся и спроси опять"}
	negative := []string{"Даже не думай", "Мой ответ - нет", "По моим данным - нет", "Перспективы не очень хорошие", "Весьма сомнительно"}
	answers := [][]string{positive, indecPositive, neutral, negative}

	fmt.Println("Задавай вопрос")

	for {
		fmt.Print("> ")
		question, _ := reader.ReadString('\n')
		question = strings.TrimSpace(question)

		if strings.HasSuffix(question, "?") {
			category := rand.Intn(4)
			index := rand.Intn(5)

			fmt.Println(answers[category][index])
			break
		} else {
			fmt.Println("Задайте свой вопрос корректно. Примечание: вопрос должен заканчиваться '?'")
		}
	}

	fmt.Println("\nНажмите ENTER для выхода.")
	reader.ReadString('\n')
}

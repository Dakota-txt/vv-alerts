package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gg-alerts/internal/models"
	"gg-alerts/internal/repository"
	"gg-alerts/internal/service"
)

func main() {
	repo := repository.NewMemoryRepository()
	donationService := service.NewDonationService(repo)

	streamer := models.Streamer{
		ID:   1,
		Name: "Стример Алекс",
	}

	repo.AddStreamer(streamer)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("=== GG Alerts ===")
		fmt.Println("1. Показать стримера")
		fmt.Println("2. Создать донат")
		fmt.Println("3. Показать историю донатов")
		fmt.Println("4. Показать общую сумму донатов")
		fmt.Println("0. Выход")
		fmt.Print("Выбери действие: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Не удалось прочитать ввод:", err)
			return
		}

		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Введи номер пункта меню")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Стример:", streamer.Name)
			fmt.Println("ID:", streamer.ID)

		case 2:
			fmt.Print("Введи сумму доната: ")

			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Ошибка чтения:", err)
				continue
			}

			amount, err := strconv.Atoi(strings.TrimSpace(input))
			if err != nil {
				fmt.Println("Сумма должна быть целым числом")
				continue
			}

			fmt.Print("Введи сообщение: ")

			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Ошибка чтения:", err)
				continue
			}

			donation, err := donationService.CreateDonation(
				streamer.ID,
				amount,
				strings.TrimSpace(message),
			)
			if err != nil {
				fmt.Println("Не удалось создать донат:", err)
				continue
			}

			fmt.Println("Донат создан! ID:", donation.ID)

		case 3:
			donations, err := donationService.GetDonations(streamer.ID)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}

			if len(donations) == 0 {
				fmt.Println("Донатов пока нет")
				continue
			}

			for _, donation := range donations {
				fmt.Println(
					"ID:", donation.ID,
					"| Сумма:", donation.Amount,
					"| Сообщение:", donation.Message,
				)
			}

		case 4:
			total, err := donationService.GetTotalAmount(streamer.ID)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}

			fmt.Println("Общая сумма донатов:", total, "руб.")

		case 0:
			fmt.Println("GG Alerts завершает работу")
			return

		default:
			fmt.Println("Такого пункта нет. Попробуй ещё раз.")
		}
	}
}

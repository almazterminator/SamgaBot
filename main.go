package main

import (
	"log"
	"time"
	"strings"
	"gopkg.in/telebot.v3"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // или другой порт по умолчанию
	}
	// Настройки бота
	pref := telebot.Settings{
		Token:  "7372500293:AAEhE0KUZ5QhxpZWa_N5IkmjtyREA5NO_3o",
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	
	
	
	// Создание бота
	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatalf("Ошибка запуска бота: %v", err)
		return
	}
	
	err = bot.SetWebhook(&telebot.Webhook{
		Endpoint: &telebot.WebhookEndpoint{
			PublicURL: "https://samgabot.onrender.com",
		},
	})

	// Создание inline-клавиатуры
	inlineMenu := &telebot.ReplyMarkup{}
	inlineBack := &telebot.ReplyMarkup{}

	// Inline-кнопки для главного меню
	btnAgenda := inlineMenu.Data("Программа", "agenda_callback")
	btnAbout := inlineMenu.Data("О проекте", "about_callback")
	btnHelp := inlineMenu.Data("Помощь", "help_callback")
	btnAiesec := inlineMenu.Data("О нас!", "aiesec_callback")


	// Добавляем кнопки в строки
	inlineMenu.Inline(
		inlineMenu.Row(btnAgenda, btnAbout),
		inlineMenu.Row(btnHelp , btnAiesec),
	)

	// Кнопка назад
	btnBack := inlineBack.Data("Назад", "back_callback")

	// Обработка команды "/start"
	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Добро пожаловать на SAMĞA Forum! Выберите действие:", &telebot.SendOptions{
			ReplyMarkup: inlineMenu,
		})
	})

	// Обработка callback-кнопки "Программа"
	bot.Handle(&btnAgenda, func(c telebot.Context) error {
		// Ответ на нажатие кнопки
		if err := c.Respond(&telebot.CallbackResponse{}); err != nil {
			return err
		}
		// Отправка текста с кнопкой "Назад"
		inlineBack.Inline(
			inlineBack.Row(btnBack),
		)
		return c.Send("Программа мероприятия: \nДата: 19 января\nМесто: Университет НАРХОЗ или КБТУ (уточняется).\nВремя: 10:00–11:00 Открытие (Main Hall).", &telebot.SendOptions{
			ReplyMarkup: inlineBack,
		})
	})

	// Обработка callback-кнопки "О проекте"
	bot.Handle(&btnAbout, func(c telebot.Context) error {
		// Ответ на нажатие кнопки
		if err := c.Respond(&telebot.CallbackResponse{}); err != nil {
			return err
		}
		// Отправка текста с кнопкой "Назад"
		inlineBack.Inline(
			inlineBack.Row(btnBack),
		)
		return c.Send("'SAMĞA Forum' — это масштабное мероприятие, посвященное Business & IT, Women Empowerment и Art.", &telebot.SendOptions{
			ReplyMarkup: inlineBack,
		})
	})

	// Обработка callback-кнопки "Помощь"
	bot.Handle(&btnHelp, func(c telebot.Context) error {
		// Ответ на нажатие кнопки
		if err := c.Respond(&telebot.CallbackResponse{}); err != nil {
			return err
		}
		// Отправка текста с кнопкой "Назад"
		inlineBack.Inline(
			inlineBack.Row(btnBack),
		)
		return c.Send("Я могу предоставить информацию по следующим разделам:\n/start — главное меню\n/agenda — программа мероприятия\n/about — информация о проекте.", &telebot.SendOptions{
			ReplyMarkup: inlineBack,
		})
	})
	bot.Handle(&btnAiesec, func(c telebot.Context) error {
		// Ответ на нажатие кнопки
		if err := c.Respond(&telebot.CallbackResponse{}); err != nil {
			return err
		}
		// Отправка текста с кнопкой "Назад"
		inlineBack.Inline(
			inlineBack.Row(btnBack),
		)
		return c.Send("AIESEC — это международная молодежная организация, основанная в 1948 году, которая способствует развитию лидерского потенциала у молодежи через культурный обмен и международные программы. Организация работает более чем в 120 странах и предоставляет молодым людям возможности для стажировок, волонтерской деятельности и участия в программах развития лидерства.", &telebot.SendOptions{
			ReplyMarkup: inlineBack,
		})
	})

	// Обработка кнопки "Назад"
	bot.Handle(&btnBack, func(c telebot.Context) error {
		// Ответ на нажатие кнопки "Назад" (возврат в главное меню)
		if err := c.Respond(&telebot.CallbackResponse{}); err != nil {
			return err
		}
		return c.Send("Добро пожаловать на SAMĞA Forum! Выберите действие:", &telebot.SendOptions{
			ReplyMarkup: inlineMenu,
		})
	})

	// Обработка неизвестных команд
	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		return c.Send("Неизвестная команда. Используйте /start для меню.")
	})
	bot.Handle("/agenda", func(c telebot.Context) error {
		return c.Send("Программа мероприятия: \nДата: 19 января\nМесто: Университет НАРХОЗ или КБТУ (уточняется).\nВремя: 10:00–11:00 Открытие (Main Hall).")
	})
	bot.Handle("/about", func(c telebot.Context) error {
		return c.Send("'SAMĞA Forum' — это масштабное мероприятие, посвященное Business & IT, Women Empowerment и Art.")
	})


	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		text := strings.ToLower(c.Text()) // Преобразуем текст в нижний регистр для удобства
		switch {
		case strings.Contains(text, "программа"):
			return c.Send("Программа мероприятия:\nДата: 19 января\nМесто: Университет НАРХОЗ или КБТУ (уточняется).\nВремя: 10:00–11:00 Открытие (Main Hall).")
		case strings.Contains(text, "о проекте"):
			return c.Send("'SAMĞA Forum' — это масштабное мероприятие, посвященное Business & IT, Women Empowerment и Art.")
		case strings.Contains(text, "помощь"):
			return c.Send("Я могу предоставить информацию по следующим разделам:\n/start — главное меню\n/agenda — программа мероприятия\n/about — информация о проекте.")
		default:
			return c.Send("Неизвестная команда. Попробуйте ввести 'Программа', 'О проекте' или 'Помощь'.")
		}
	})
	// Запуск бота
	log.Println("Бот запущен!")
	


	bot.Start()
}

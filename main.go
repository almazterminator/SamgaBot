package main

import (
	"log"
	"time"
	"strings"
	"gopkg.in/telebot.v3"
	"fmt"
    "net/http"

)

func main() {
		port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
	
	// Настройки бота
	pref := telebot.Settings{
		Token:  "7228846507:AAGUW3Zneq27_o1Uz6vv4V-Mra43uYzxHcw",
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	
	
	
	// Создание бота
	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatalf("Ошибка запуска бота: %v", err)
		return
	}
	


	// Создание inline-клавиатуры
	inlineMenu := &telebot.ReplyMarkup{}
	inlineBack := &telebot.ReplyMarkup{}

	// Inline-кнопки для главного меню
	btnAgenda := inlineMenu.Data("Программа 📅", "agenda_callback")
	btnAbout := inlineMenu.Data("О проекте 👥", "about_callback")
	btnHelp := inlineMenu.Data("Помощь 🆘", "help_callback")
	btnAiesec := inlineMenu.Data("О нас! 🏡", "aiesec_callback")


	// Добавляем кнопки в строки
	inlineMenu.Inline(
		inlineMenu.Row(btnAgenda, btnAbout),
		inlineMenu.Row(btnHelp , btnAiesec),
	)

	// Кнопка назад
	btnBack := inlineBack.Data("Назад 🔙", "back_callback")

	// Обработка команды "/start"
	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Добро пожаловать на SAMĞA Forum! 😊 Выберите действие:", &telebot.SendOptions{
			ReplyMarkup: inlineMenu,
		})
		
	})

// Обработка callback-кнопки "Программа"
bot.Handle(&btnAgenda, func(c telebot.Context) error {
	// Ответ на нажатие кнопки
	if err := c.Respond(&telebot.CallbackResponse{}); err != nil {
		return err
	}

	// Удаление сообщения с кнопками
	if err := bot.Delete(c.Message()); err != nil {
		return err
	}

	// Отправка текста с кнопкой "Назад"
	inlineBack.Inline(
		inlineBack.Row(btnBack),
	)
	return c.Send("Программа мероприятия: \n\nДата: 19 января\nМесто: Университет НАРХОЗ или КБТУ (уточняется).\n\nВремя: 10:00–11:00 Открытие (Main Hall)\n- Приветственное слово от организаторов (Chair, LCP, LCVP oGX).\n- Представление партнеров и целей форума.\n\n11:00–12:00\nПрезентация от Mars (Hall 1)\n- Тема: 'Лидерство и карьерный рост'.\n- Продолжительность: 50 минут.\n\n12:00–13:00\nПрезентация от JTI (Hall 2)\n- Тема: 'Баланс между работой и личной жизнью'.\n- Продолжительность: 50 минут.\n\n13:00–13:20\nКофе-брейк (Networking Zone)\n- Легкий перекус и неформальное общение.\n\n13:30–14:30\nПанельная дискуссия от Mastercard (Main Hall)\n- Тема: 'Women Empowerment'.\n- Участники: 3 представителя компании.\n\n14:30–15:30\nŞam x Samğa – выступления спикеров из медиа и искусства (Main Hall)\n- Тема: 'Роль медиа и искусства в продвижении женщин'.\n- Участники: 4 спикера из индустрии.\n\n15:30–16:00\nКофе-брейк и networking space (Networking Zone)\n- Время для общения и установления контактов.\n\n16:00–17:00\nВыступление инфлюенсера (Main Hall)\n- Тема: 'Развитие личного бренда и поддержка Women Empowerment'.\n\n17:00–17:40\nЗакрытие форума (Main Hall)\n- Резюме ключевых моментов форума.\n- Благодарственные слова организаторов (LCVP BD, LCVP PM, OC team).\n- Анонс следующих мероприятий.\n", &telebot.SendOptions{
		ReplyMarkup: inlineBack,
	})
})




	// Обработка callback-кнопки "О проекте"
	bot.Handle(&btnAbout, func(c telebot.Context) error {
		// Ответ на нажатие кнопки
		if err := c.Respond(&telebot.CallbackResponse{}); err != nil {
			return err
		}
			// Удаление сообщения с кнопками
	if err := bot.Delete(c.Message()); err != nil {
		return err
	}
		// Отправка текста с кнопкой "Назад"
		inlineBack.Inline(
			inlineBack.Row(btnBack),
		)
		return c.Send("'SAMĞA  Forum' — это масштабное мероприятие, посвященное теме Business & IT,  Women Empowerment, и Art, где соберутся представители ведущих компаний (Mars, Mastercard, JTI), эксперты и молодые лидеры. Форум создаст платформу для обмена опытом, расширения возможностей молодежи в бизнесе, искусстве и обсуждения актуальных вопросов, влияющих на современное общество.\n\nЦели проекта:\n1. Популяризация идеи Women Empowerment в бизнесе, искусстве. \n2. Развитие профессиональных навыков делегат через мастер-классы и обсуждения.  \n3. Создание платформы для нетворкинга между делегатами и представителями крупных компаний.\n\nФормат мероприятия:\n3 зала:\n- Mars: тематическая сессия по вопросам лидерства.  \n- Mastercard: стратегии карьерного роста для женщин.  \n- JTI: баланс между работой и личной жизнью.\n\nОсобенности мероприятия:\n- Фокус на Women Empowerment: ключевая тема, объединяющая все сессии.  \n- Панельная дискуссия: участие представителей Mastercard с практическими кейсами.  \n- Нетворкинг: два выделенных блока для общения участников и представителей компаний. ", &telebot.SendOptions{
			ReplyMarkup: inlineBack,
		})
	})

	// Обработка callback-кнопки "Помощь"
	bot.Handle(&btnHelp, func(c telebot.Context) error {
		// Ответ на нажатие кнопки
		if err := c.Respond(&telebot.CallbackResponse{}); err != nil {
			return err
		}
			// Удаление сообщения с кнопками
	if err := bot.Delete(c.Message()); err != nil {
		return err
	}
		// Отправка текста с кнопкой "Назад"
		inlineBack.Inline(
			inlineBack.Row(btnBack),
		)
		return c.Send("Я могу предоставить информацию по следующим разделам:\n/start — главное меню\n/agenda — программа мероприятия\n/about — информация о проекте.\n\nЕсли возникнут проблемы с функционированием бота, пожалуйста, обращайтесь в Telegram: @almazterminator.", &telebot.SendOptions{
			ReplyMarkup: inlineBack,
		})
	})
	bot.Handle(&btnAiesec, func(c telebot.Context) error {
		// Ответ на нажатие кнопки
		if err := c.Respond(&telebot.CallbackResponse{}); err != nil {
			return err
		}
			// Удаление сообщения с кнопками
	if err := bot.Delete(c.Message()); err != nil {
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
			// Удаление сообщения с кнопками
	if err := bot.Delete(c.Message()); err != nil {
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
		return c.Send("Программа мероприятия: \n\nДата: 19 января\nМесто: Университет НАРХОЗ или КБТУ (уточняется).\nВремя: 10:00–11:00 Открытие (Main Hall)\n  - Приветственное слово от организаторов (Chair, LCP, LCVP oGX).\n  - Представление партнеров и целей форума.\n\n11:00–12:00\nПрезентация от Mars (Hall 1)\n- Тема: 'Лидерство и карьерный рост'.\n- Продолжительность: 50 минут.\n\n12:00–13:00\nПрезентация от JTI (Hall 2)\n- Тема: 'Баланс между работой и личной жизнью'.\n- Продолжительность: 50 минут.\n\n13:00–13:20\nКофе-брейк (Networking Zone)\n  - Легкий перекус и неформальное общение.\n\n13:30–14:30\nПанельная дискуссия от Mastercard (Main Hall)\n- Тема: 'Women Empowerment'.\n- Участники: 3 представителя компании.\n\n14:30–15:30\nŞam x Samğa – выступления спикеров из медиа и искусства (Main Hall)\n- Тема: 'Роль медиа и искусства в продвижении женщин'.\n- Участники: 4 спикера из индустрии.\n\n15:30–16:00\nКофе-брейк и networking space (Networking Zone)\n- Время для общения и установления контактов.\n\n16:00–17:00\nВыступление инфлюенсера (Main Hall)\n- Тема: 'Развитие личного бренда и поддержка Women Empowerment'.\n\n17:00–17:40\nЗакрытие форума (Main Hall)\n- Резюме ключевых моментов форума.\n- Благодарственные слова организаторов (LCVP BD, LCVP PM, OC team).\n- Анонс следующих мероприятий.\n")
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
package server

type Game struct {
	Word      string
	UsedChars map[rune]bool
	Stage     int
}

const WELCOME_MESSAGE = "Привет! Это игра \"Виселица\"\nНачинаем игру!\n"
const STATS_TEXT = "Ваш IP: #ip#\nСлов отгадано: #guessed#\nПроиграно: #losses# раз\n\n"
const USED_TEXT = "Использованные буквы: "
const SUGGESTION_TEXT = "Введите букву: "
const UNRESOLVED_SYMBOL = "Неизвестная буква!\nВведите новую\n"
const FAILED_CHAR_MSG = "Не угадали :(\n"
const FAILED_MSG = "\n\nВы проиграли :( Слово было: #word#\n\n\n"
const WIN_MSG = "\n\nВы выиграли!!! Слово: #word#\n\n\n"
const PLAY_AGAIN = "Начать заново? (Д/Н): "
const USED_CHAR = "Эту букву Вы уже называли :)))\n"

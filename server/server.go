package server

import (
	"bufio"
	"fmt"
	"github.com/kk222mo/hangman/data"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

var mainDictionary data.Dictionary
var pictures []data.Picture

func printUsed(game Game, writer *bufio.Writer) {
	writer.WriteString(USED_TEXT)
	for k, _ := range game.UsedChars {
		writer.WriteString(string(k) + " ")
	}
	if len(game.UsedChars) == 0 {
		writer.WriteString("<нет>")
	}
	writer.WriteString("\n")
	writer.Flush()
}

func reqForChar(game *Game, reader *bufio.Reader, writer *bufio.Writer) bool {
	writer.WriteString(SUGGESTION_TEXT)
	writer.Flush()
	str, err := reader.ReadString('\n')
	if len([]rune(str)) == 0 {
		return false
	}
	char := []rune(strings.ToLower(str))[0]
	if err != nil || !('а' <= char && char <= 'я') {
		writer.WriteString(UNRESOLVED_SYMBOL)
		return false
	}
	if game.UsedChars[char] {
		writer.WriteString(USED_CHAR)
		return false
	}
	game.UsedChars[char] = true
	if !strings.Contains(game.Word, string(char)) {
		game.Stage++
		writer.WriteString(FAILED_CHAR_MSG)
		return true
	}
	return true
}

func testForLoose(game Game, writer *bufio.Writer) bool {
	if game.Stage >= len(pictures)-1 {
		writer.WriteString(strings.Replace(FAILED_MSG, "#word#", game.Word, 1))
		writer.Flush()
		return true
	}
	return false
}

func initGame() Game {
	return Game{Stage: 0, Word: mainDictionary.Words[rand.Intn(mainDictionary.Size)], UsedChars: make(map[rune]bool)}
}

func askForPlayAgain(writer *bufio.Writer, reader *bufio.Reader) bool {
	writer.WriteString(PLAY_AGAIN)
	writer.Flush()
	str, err := reader.ReadString('\n')
	char := []rune(strings.ToLower(str))[0]
	if err != nil || char != 'д' {
		return false
	}
	return true
}

func processClient(conn net.Conn) {
	writer := bufio.NewWriter(conn)
	writer.WriteString(WELCOME_MESSAGE)
	writer.Flush()
	reader := bufio.NewReader(conn)
	defer conn.Close()

	game := initGame()
	for {
		stageText := pictures[game.Stage].Text

		guessedWord := ""
		guessedLen := 0
		for _, char := range game.Word {
			if game.UsedChars[char] && strings.Contains(game.Word, string(char)) {
				guessedWord += string(char) + " "
				guessedLen++
			} else {
				guessedWord += "_ "
			}
		}
		if guessedLen == len([]rune(game.Word)) {
			writer.WriteString(strings.Replace(WIN_MSG, "#word#", game.Word, 1))
			if askForPlayAgain(writer, reader) {
				game = initGame()
				continue
			} else {
				return
			}

		}

		stageText = strings.ReplaceAll(stageText, "#word#", guessedWord)
		writer.WriteString(stageText)
		printUsed(game, writer)
		for {
			res := reqForChar(&game, reader, writer)
			writer.Flush()
			if testForLoose(game, writer) {
				if askForPlayAgain(writer, reader) {
					game = initGame()
					break
				} else {
					return
				}
			}
			if res {
				break
			}
		}
	}
}

func Serve(port int) error {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Reading stages pictures and dictionary file...")
	dict, err := data.ReadDictionary("main.txt")
	if err != nil {
		return err
	}
	mainDictionary = dict
	pictures, err = data.ReadPictures("stages.txt")
	if err != nil {
		return err
	}
	fmt.Println("Starting listener on port", port)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go processClient(conn)
	}
}

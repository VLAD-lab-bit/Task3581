package main

import (
	"bufio"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

// Сетевой адрес.
const addr = "0.0.0.0:12345"

// Протокол сетевой службы.
const proto = "tcp4"

// Go-поговорки.
var proverbs = []string{
	"Don't communicate by sharing memory, share memory by communicating.",
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"The bigger the interface, the weaker the abstraction.",
	"Make the zero value useful.",
	"interface{} says nothing.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Syscall must always be guarded with build tags.",
	"Cgo must always be guarded with build tags.",
	"Cgo is not Go.",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic.",
}

func main() {
	// Запуск сетевой службы по протоколу TCP на порту 12345.
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("Server started on", addr)

	// Подключения обрабатываются в бесконечном цикле.
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Обрабатываем каждое подключение в отдельной горутине.
		go handleConn(conn)
	}
}

// Обработчик для каждого соединения.
func handleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		// Чтение сообщения от клиента (ожидаем команду "start").
		b, err := reader.ReadBytes('\n')
		if err != nil {
			log.Println("Read error:", err)
			return
		}

		// Удаление символов конца строки.
		msg := strings.TrimSpace(string(b))

		if msg == "start" {
			log.Println("Client requested proverbs.")
			// Отправляем поговорки каждые 3 секунды.
			sendProverbs(conn)
		} else {
			conn.Write([]byte("Enter 'start' to receive proverbs.\n"))
		}
	}
}

// Отправка случайных Go-поговорок каждые 3 секунды.
func sendProverbs(conn net.Conn) {
	rand.Seed(time.Now().UnixNano())
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			proverb := proverbs[rand.Intn(len(proverbs))]
			conn.Write([]byte(proverb + "\n"))
		}
	}
}

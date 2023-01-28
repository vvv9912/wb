package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"io"
	"os"
	"time"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и go lint.
*/

func main() {
	curTime, err := gettime()
	if err != nil {
		io.WriteString(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println(curTime)
	os.Exit(0)
}

func gettime() (time.Time, error) {
	curTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	return curTime, err
}

package parser

import (
	"bufio"
	"io"
)

// eof é uma runa de marcação que representa o fim do arquivo
const eof = rune(0)

// Scanner representa o analisador léxico
type Scanner struct {
	reader *bufio.Reader
}

// NewScanner retorna uma nova instância de Scanner
func NewScanner(fileReader io.Reader) *Scanner {
	return &Scanner{reader: bufio.NewReader(fileReader)}
}

// read realiza a leitura da próxima runa do leitor de buffer
// Retorna rune(0) se um erro ocorre.
func (scanner *Scanner) read() rune {
	ch, _, err := scanner.reader.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread recoloca a última runa lina de volta no leitor de buffer
func (scanner *Scanner) unread() { _ = scanner.reader.UnreadRune() }

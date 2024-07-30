package linkmasker

import (
	"bufio"
	"os"
	"strings"
)

type producer interface {
	produce() ([]string, error)
}

type presenter interface {
	present([]string) error
}
type LinkMasker struct{
	prod producer
	pres presenter
}

func NewService(prod producer, pres presenter) *LinkMasker {
	return &LinkMasker{prod, pres}
}

func (s *LinkMasker) hideLinks(message string) string {
	input := []byte(message)
	var result []byte

	prefix := []byte("http://")
	prefixLength := len(prefix)
	inputLength := len(input)

	i := 0
	for i < inputLength {
		if i <= inputLength-prefixLength && string(input[i:i+prefixLength]) == string(prefix) {
			result = append(result, prefix...)
			i += prefixLength

			for i < inputLength && input[i] != ' ' {
				result = append(result, '*')
				i++
			}
		} else {
			result = append(result, input[i])
			i++
		}
	}

	return string(result)
}

type FileProducer struct {
	filePath string
}

func NewFileProducer(filePath string) *FileProducer {
	return &FileProducer{filePath: filePath}
}

func (p *FileProducer) produce() ([]string, error) {
	file, err := os.Open(p.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

type FilePresenter struct {
	filePath string
}

func NewFilePresenter(filePath string) *FilePresenter {
	return &FilePresenter{filePath: filePath}
}

func (p *FilePresenter) present(messages []string) error {
	return os.WriteFile(p.filePath, []byte(strings.Join(messages, "\n")), 0644)
}

func (s *LinkMasker) Run() error {
	messages, err := s.prod.produce()
	if err != nil {
		return err
	}

	var maskedMessages []string
	for _, msg := range messages {
		maskedMessages = append(maskedMessages, s.hideLinks(msg))
	}

	return s.pres.present(maskedMessages)
}
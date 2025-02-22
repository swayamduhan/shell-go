package utils

// stack implementation to check for incomplete command using odd quotes & parenthesis

type Stack struct {
	array []rune
}

func InitStack() *Stack {
	return &Stack{
		array: []rune{},
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.array) == 0
}

func (s *Stack) Push(x rune) {
	s.array = append(s.array, x)
}

func (s *Stack) Pop() (rune, bool) {
	if len(s.array) == 0 {
		return 0, false
	}

	top := s.array[len(s.array) - 1]
	s.array = s.array[:len(s.array) - 1]
	return top, true
}


func IsBalanced(ip string) bool {
	// iterate over each character and balance
	// check only for backticks inside open double quotes
	// dont even check for backticks inside 

	// isDoubleQuoteOpen := false
	// isSingleQuoteOpen := false

	// for _, r := range ip {

	// }

	return false
} 
package utils

// TODO: add support for backtick commands
// add support for ENVs
// add tokens for pipe, redirects, flags

import (
	"strings"
)

type TokenType int

const (
	TOKEN_WORD TokenType = iota
	TOKEN_STRING
	TOKEN_PIPE
	TOKEN_BACKTICK
	TOKEN_REDIRECT
	TOKEN_FLAG
	TOKEN_ENV
)

type Token struct {
	Type TokenType
	Value string
}


var currentToken strings.Builder

func Tokenize(args string) []Token {
	var tokens []Token

	// iterate over each letter in string
	// check for insideSingleQuote
	// check for insideDoubleQuote
	// check for escapeSequences
	// check for backticks ( get output and print thats it )
	// dont count "" inside "" and '' inside ''
	// handle comments for whole line after #
	// handle envs both outside and inside quotations

	insideSingleQuote := false
	insideDoubleQuote := false
	escape := false
	comment := false
	isPreviousSpace := false

	for _, r := range args {
		if comment {
			break
		}

		switch r {
		case ' ' : 
			// add to token if inside quotes
			// append token to array if not empty, and add a type to it
			// hasSpace = true if outside quotes
			if insideDoubleQuote || insideSingleQuote || escape {
				currentToken.WriteRune(r)
				escape = false
				continue
			}	
			if currentToken.String() != "" {
				tokens = append(tokens, Token{
					Type: TOKEN_WORD,
					Value: currentToken.String(),
				})
				currentToken.Reset()
			}
			isPreviousSpace = true
			
		case '\'':
			// single quote
			if insideDoubleQuote || escape {
				// append
				currentToken.WriteRune(r)
				isPreviousSpace = false
				escape = false
				continue
			}
			if insideSingleQuote {
				// append the token
				insideSingleQuote = false

				if currentToken.String() != "" {
					tokens = append(tokens, Token{
						Type: TOKEN_STRING,
						Value: currentToken.String(),
					})
					currentToken.Reset()
				}
				isPreviousSpace = false
				continue
			}
			insideSingleQuote = true
			isPreviousSpace = false

		case '"':
			if insideSingleQuote || escape {
				// simply append
				currentToken.WriteRune(r)
				isPreviousSpace = false
				escape = false
				continue
			} 
			if insideDoubleQuote {
				// append the token
				insideDoubleQuote = false

				if currentToken.String() != "" {
					tokens = append(tokens, Token{
						Type: TOKEN_STRING,
						Value: currentToken.String(),
					})
					currentToken.Reset()
				}
				isPreviousSpace = false
				continue
			}
			insideDoubleQuote = true
			isPreviousSpace = false

		case '\\':
			if escape {
				currentToken.WriteRune(r)
			} else {
				escape = true
			}
			isPreviousSpace = false

		case '#':
			// comments not applicable inside quotes, only apply after a space, handle escape
			if insideSingleQuote || insideDoubleQuote || escape {
				currentToken.WriteRune(r)
				isPreviousSpace = false
				escape = false
				continue
			}
			if isPreviousSpace {
				comment = true
			}
			isPreviousSpace = false

		case '$':
			if insideDoubleQuote || insideSingleQuote || escape {
				currentToken.WriteRune(r)
				isPreviousSpace = false
				escape = false
				continue	
			}

			// handle envs, appending for now
			currentToken.WriteRune(r)
			isPreviousSpace = false

		default:
			// append rune to current character
			currentToken.WriteRune(r)
			isPreviousSpace = false
		}

		
		// fmt.Println("Current token : ", currentToken.String())
	}

	
	// assign the last token
	if currentToken.String() != "" {
		tokens = append(tokens, Token{		
			Type: TOKEN_WORD,
			Value: currentToken.String(),
		})
	}

	currentToken.Reset()
	return tokens
}
package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	type wanted struct {
		expectedType    token.TokenType
		expectedLiteral string
	}
	type testCase struct {
		name   string
		input  string
		wanted []wanted
	}

	testCases := []testCase{
		{
			name:  "a",
			input: `=+(){},;`,
			wanted: []wanted{
				{token.ASSIGN, "="},
				{token.PLUS, "+"},
				{token.LPAREN, "("},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			name: "b",
			input: `
			let five = 5;
			let ten = 10;
			let add = fn(x, y) {
			x + y;
		};

			let result = add(five, ten);
			`,
			wanted: []wanted{
				{token.LET, "let"},
				{token.IDENT, "five"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},

				{token.LET, "let"},
				{token.IDENT, "ten"},
				{token.ASSIGN, "="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},

				{token.LET, "let"},
				{token.IDENT, "add"},
				{token.ASSIGN, "="},
				{token.FUNCTION, "fn"},
				{token.LPAREN, "("},
				{token.IDENT, "x"},
				{token.COMMA, ","},
				{token.IDENT, "y"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.IDENT, "x"},
				{token.PLUS, "+"},
				{token.IDENT, "y"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.SEMICOLON, ";"},

				{token.LET, "let"},
				{token.IDENT, "result"},
				{token.ASSIGN, "="},
				{token.IDENT, "add"},
				{token.LPAREN, "("},
				{token.IDENT, "five"},
				{token.COMMA, ","},
				{token.IDENT, "ten"},
				{token.RPAREN, ")"},
				{token.SEMICOLON, ";"},

				{token.EOF, ""},
			},
		},
		{
			name:  "c",
			input: `!-/*<>`,
			wanted: []wanted{
				{token.BANG, "!"},
				{token.MINUS, "-"},
				{token.SLASH, "/"},
				{token.ASTERISK, "*"},
				{token.LT, "<"},
				{token.GT, ">"},
			},
		},
		{
			name:  "keyword",
			input: `if else return true false`,
			wanted: []wanted{
				{token.IF, "if"},
				{token.ELSE, "else"},
				{token.RETURN, "return"},
				{token.TRUE, "true"},
				{token.FALSE, "false"},
			},
		},
		{
			name:  "two byte",
			input: `== != a!=`,
			wanted: []wanted{
				{token.EQ, "=="},
				{token.NOT_EQ, "!="},
				{token.IDENT, "a"},
				{token.NOT_EQ, "!="},
			},
		},
		{
			"string",
			`"foobar"`,
			[]wanted{
				{token.STRING, "foobar"},
			},
		},
		{
			"string2",
			`"foo bar"`,
			[]wanted{
				{token.STRING, "foo bar"},
			},
		},
	}

	for _, tt := range testCases {
		l := New(tt.input)

		for i, expected := range tt.wanted {
			tok := l.NextToken()

			if tok.Type != expected.expectedType {
				t.Fatalf("tests name: [%s] [%d] - tokentype wrong. expected=%q, got=%q",
					tt.name, i, expected.expectedType, tok.Type)
			}
			if tok.Literal != expected.expectedLiteral {
				t.Fatalf("tests name: [%s] [%d] - literal wrong. expected=%q, got=%q",
					tt.name, i, expected.expectedLiteral, tok.Literal)
			}
		}
	}
}

package packer

import (
	"github.com/mewfinity06/mewvm/lexer"
)

type Program []byte

func Pack(l lexer.Lexer) (Program, error) {
	res := Program{}

	for {
		token, err := l.Next()
		if err != nil {
			return res, err
		}
		res = append(res, byte(token.Lexme.ToHex()))
		if token.Lexme == lexer.Op_EOF {
			break
		}
	}

	return res, nil
}

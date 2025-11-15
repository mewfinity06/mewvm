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
		// res = append(res, byte(token.Lexme.ToHex()))
		switch token.Lexme {
		case lexer.Op_Hex, lexer.Op_Int:
			res = append(res, token.Value)
		default:
			res = append(res, byte(token.Lexme.ToHex()))
		}

		if token.Lexme == lexer.Inst_EOF {
			break
		}
	}

	return res, nil
}

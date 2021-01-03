package generator

import "github.com/bonbot-team/bonapi/utils"

type Generator interface {
	GetName() string
	Generate(args map[string][]string) ([]byte, *utils.Error)
}

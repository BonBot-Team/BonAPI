package generator

type Generator interface {
    GetName() string
    Generate(args map[string][]string) ([]byte, error)
}
package main

func main() {
	var n N

	var t Ner = &n

	t.a()
}

type Ner interface {
	a()
	b(int)
	c(string) string
}

type N int

func (N) a() {}

func (*N) b(int) {}

func (*N) c(string) string { return "" }

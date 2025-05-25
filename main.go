package main

func main() {
	m1 := New("m1.yml.tpl")
	m1.SetSrc("./example/example.yml")
	m1.Run()

	m2 := New("m1.yml")
	m2.SetSrc("./example/example.yml")
	m2.Run()
}

package main

func main() {
	//m1 := New("m1.yml.tpl")
	//m1.SetSrc("./example/example.yml")
	//m1.Run()
	//
	//m2 := New("m2.yml")
	//m2.SetSrc("./example/example.yml")
	//m2.Run()

	m3 := New("m1.base64")
	m3.SetSrc("./example/example.base64")
	m3.Run()
}

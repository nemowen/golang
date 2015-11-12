package main

type ITeser interface {
	Ok()
	Into()
}

type Easy interface {
	Ok()
}

type Test struct{}

func (t *Test) Ok() {
	println("ok")
}

func (t Test) Into() {
	println("into")
}

func main() {

	var it Easy = &Test

	if easy, ok := it.(*Test); ok {
		easy.Ok()
	}

}

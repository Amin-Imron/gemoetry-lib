package shape

type Rectangle struct {
	Width  float32 //<=====Inget nama field harus huruf besar, supaya bisa dipake di file lain
	Length float32 //<=====Inget nama field harus huruf besar, supaya bisa dipake di file lain
}

func (rectangle *Rectangle) Area() float32 { //<=====Inget nama method harus huruf besar, supaya bisa dipake di file lain
	return rectangle.Width * rectangle.Length
}
func (rectangle *Rectangle) Perimeter() float32 { //<=====Inget nama method harus huruf besar, supaya bisa dipake di file lain
	return (rectangle.Width + rectangle.Length) * 2
}

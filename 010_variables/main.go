/*
Vars
*/
package main

var (
	packageVar = 23
)

func main() {
	//integers
	var i int = 10 //платформозасимый тип 32 или 64 бита
	var autoInt = -10
	var bigInt int64 = 1<<32 - 1          //int8, int16, int32, int64
	var unsignedInt uint = 100500         //платформозависимый тип 32 или 64 бита
	var unsignedBigInt uint64 = 1<<64 - 1 //uint8, uint16, uint32, uint64
	println("integers", i, autoInt, bigInt, unsignedInt, unsignedBigInt)

	//floats
	var p float32 = 3.14 //float = float32, float 64
	println("float: ", p)

	//bool
	var b = true
	println("bool var", b)

	//strings
	var hello string = "Hallo\n\t"
	var world = "World"
	println(hello, world)

	//bytes
	var rawBinary byte = '\x27'
	println("rawBinary", rawBinary)

	/*
		short d
	*/
	someVar := 42
	println("Some New Var inited via :=", someVar)

	//type cont (float to int)
	println("float to int convert", int(p))
	println("int to string convert", string(48))

	// complex
	z := 2 + 3i
	println("complex number: ", z)

	//strings
	str1 := "Alex"
	str2 := "M"

	name := str1 + " " + str2
	println("name len is: ", name, len(name))

	escaping := `Hello\r\n
      Multiline string
    `
	println("as-is escaping and multiline string ", escaping)

	//defaults
	var defaultInt int
	var defaultFloat float32
	var defaultString string
	var defaultBool bool
	println("default values: ", defaultInt, defaultFloat, defaultString, defaultBool)

	/*
	   multiple vars
	*/
	var v1, v2 string = "v1", "v2"
	println(v1, v2)

	var (
		m0 int = 12
		m2     = "string"
		m3     = 23
	)
	println(m0, m2, m3)
	println("Package var ", packageVar)

}

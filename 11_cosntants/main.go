package main

const someInt  = 1
const typedInt int32  = 17
const fullName = "Raimo"

const (
	flagKey1 = 1
	flagKey2 = 2
)

const (
	one = iota
	two
	_  //skip iota
	four
)

const (
	_ = iota //skip first value
	KB uint64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	//type overflow
	//ZB
)
func main()  {
   pi := 3.14
   println(pi + someInt)

   println(pi + float64(typedInt))

   println(one, two, four)

   println(KB, MB, GB, TB, PB, EB)
}
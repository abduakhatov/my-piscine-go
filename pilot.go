package piscine

type Craft int

const (
	AIRCRAFT1 Craft = 1
)

type Pilot struct {
	Name     string
	Life     float32
	Age      int
	Aircraft Craft
}

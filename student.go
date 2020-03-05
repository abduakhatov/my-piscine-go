package piscine

type Craft string

const (
	AIRCRAFT1 Craft = "aircraft1"
)

type Pilot struct {
	Name     string
	Life     float32
	Age      int
	Aircraft Craft
}

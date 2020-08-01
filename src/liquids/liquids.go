package liquids

import "fmt"

type Gallons float64
type Liters float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (g Gallons) ToLitters() Liters {
	return Liters(g * 3.785)
}

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}
func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

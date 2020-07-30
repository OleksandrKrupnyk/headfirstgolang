package liquids

type Gallons float64
type Liters float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (g Gallons) ToLitters() Liters {
	return Liters(g * 3.785)
}

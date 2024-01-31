package datastructures

type Demographic string

const (
	Seinen   Demographic = "Seinen"
	Shonen   Demographic = "Shounen"
	Shojo    Demographic = "Shoujo"
	Josei    Demographic = "Josei"
	Original             = "Original"
)

type DemographicRelation struct {
	Seinen                float32
	Shonen                float32
	Shojo                 float32
	Josei                 float32
	NovelWebcomicOriginal float32
}

package space

type Planet = string //what is the use of this line?? Had to look this up

var planetList = map[string]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

var earthSecsPerYear float64 = 31557600

func Age(secs float64, planet string) float64 {
	return secs / (earthSecsPerYear * planetList[planet])
}

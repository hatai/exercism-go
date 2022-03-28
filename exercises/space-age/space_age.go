package space

type Planet string

const (
	Earth   Planet = "Earth"
	Mercury Planet = "Mercury"
	Venus   Planet = "Venus"
	Mars    Planet = "Mars"
	Jupiter Planet = "Jupiter"
	Saturn  Planet = "Saturn"
	Uranus  Planet = "Uranus"
	Neptune Planet = "Neptune"
)


func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case Earth:
		return seconds / 31557600

	case Mercury:
		return Age(seconds, Earth) / 0.2408467

	case Venus:
		return Age(seconds, Earth) / 0.61519726

	case Mars:
		return Age(seconds, Earth) / 1.8808158

	case Jupiter:
		return Age(seconds, Earth) / 11.862615

	case Saturn:
		return Age(seconds, Earth) / 29.447498

	case Uranus:
		return Age(seconds, Earth) / 84.016846

	case Neptune:
		return Age(seconds, Earth) / 164.79132

	default:
		return 0
	}
}

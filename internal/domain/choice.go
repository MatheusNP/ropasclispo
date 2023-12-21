package domain

type Choice struct {
	Value    uint8
	Name     string
	WinsFrom []WinsFrom
}

type WinsFrom struct {
	Value   uint8
	Message string
}

func Choose(value uint8, choices []Choice) Choice {
	for _, found := range choices {
		if value == found.Value {
			return found
		}
	}

	return Choice{}
}

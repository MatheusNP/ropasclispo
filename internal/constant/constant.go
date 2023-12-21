package constant

import "github.com/MatheusNP/ropasclispo/internal/domain"

func ROCK() domain.Choice {

	return domain.Choice{
		Value: 0,
		Name:  "Rock",
		WinsFrom: []domain.WinsFrom{
			{Value: 2, Message: "Rock crushes Scissors"},
			{Value: 3, Message: "Rock crushes Lizard"},
		},
	}
}

func PAPER() domain.Choice {
	return domain.Choice{
		Value: 1,
		Name:  "Paper",
		WinsFrom: []domain.WinsFrom{
			{Value: 0, Message: "Paper covers Rock"},
			{Value: 4, Message: "Paper disproves Spock"},
		},
	}
}

func SCISSORS() domain.Choice {
	return domain.Choice{
		Value: 2,
		Name:  "Scissors",
		WinsFrom: []domain.WinsFrom{
			{Value: 1, Message: "Scissors cuts Paper"},
			{Value: 3, Message: "Scissors decapitates Lizard"},
		},
	}
}

func LIZARD() domain.Choice {
	return domain.Choice{
		Value: 3,
		Name:  "Lizard",
		WinsFrom: []domain.WinsFrom{
			{Value: 1, Message: "Lizard eats Paper"},
			{Value: 4, Message: "Lizard poisons Spock"},
		},
	}
}

func SPOCK() domain.Choice {
	return domain.Choice{
		Value: 4,
		Name:  "Spock",
		WinsFrom: []domain.WinsFrom{
			{Value: 0, Message: "Spock vaporizes Rock"},
			{Value: 2, Message: "Spock smashes Scissors"},
		},
	}
}

func CHOICES() []domain.Choice {
	return []domain.Choice{
		ROCK(),
		PAPER(),
		SCISSORS(),
		LIZARD(),
		SPOCK(),
	}
}

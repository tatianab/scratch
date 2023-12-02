package day2

func SumPower(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		g, err := parseGame(line)
		if err != nil {
			return 0, err
		}
		sum += g.power()
	}
	return sum, nil
}

func (g *game) power() int {
	red, green, blue := 0, 0, 0
	for _, t := range g.trials {
		if t.red > red {
			red = t.red
		}
		if t.green > green {
			green = t.green
		}
		if t.blue > blue {
			blue = t.blue
		}
	}
	return red * green * blue
}

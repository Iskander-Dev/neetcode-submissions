type Car struct {
	pos   int
	speed int
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)

	cars := make([]Car, n)

	for i := range n {
		cars[i] = Car{pos: position[i], speed: speed[i]}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	stack := []float64{}

	for _, car := range cars {
		time := float64(target-car.pos) / float64(car.speed)

		if len(stack) == 0 || time > stack[len(stack)-1] {
			stack = append(stack, time)
		}
	}

	return len(stack)
}

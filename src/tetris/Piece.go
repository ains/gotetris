package tetris

type Piece struct {
	Id           rune
	Coordinates  [][][2]int
	LowestPoints [4][10]int
	MoveSet      [][2]int
}

func NewPiece(id rune, coordSets [][][2]int) *Piece {
	moveSet := make([][2]int, 0, 40)
	lowestPoints := new([4][10]int)
	setPoints := new([4][10]bool)

	for i, rotation := range coordSets {
		max := rotation[0][0]
		min := max

		for _, coord := range rotation {
			x := coord[0]
			y := coord[1]

			// Calculate the range which the piece can move through (distance from walls)
			if x < min {
				min = x
			} else if x > max {
				max = x
			}

			// Cache the lowest points of the piece (for collisions)
			if lowestPoints[i][x] > y || !setPoints[i][x] {
				setPoints[i][x] = true
				lowestPoints[i][x] = y
			}
		}

		moveRange := min + (10 - max)

		// The MoveSet is organised in such a way that moves closer to the center
		// are at lower indexes in the array.
		// Useful for a Depth-First search moves which require fewer keypresses
		// will be found first
		for move := 0; move < moveRange; move++ {
			if move <= min {
				moveSet = append(moveSet, [2]int{i, -move})
			}

			if move != 0 && move < (10-max) {
				moveSet = append(moveSet, [2]int{i, move})
			}
		}
		//TODO: Order rotations in a similar manner
	}

	return &Piece{Id: id, Coordinates: coordSets, LowestPoints: *lowestPoints, MoveSet: moveSet}
}

var PieceMap = map[rune]*Piece{
	'I': NewPiece('I', [][][2]int{
		{{3, 0}, {4, 0}, {5, 0}, {6, 0}},
		{{5, 0}, {5, 1}, {5, 2}, {5, 3}},
	}),
	'O': NewPiece('O', [][][2]int{
		{{4, 0}, {5, 0}, {4, 1}, {5, 1}},
	}),
	'J': NewPiece('J', [][][2]int{
		{{3, 0}, {4, 0}, {5, 0}, {3, 1}},
		{{4, 0}, {4, 1}, {4, 2}, {5, 2}},
		{{5, 0}, {3, 1}, {4, 1}, {5, 1}},
		{{4, 0}, {3, 0}, {4, 1}, {4, 2}},
	}),
	'L': NewPiece('L', [][][2]int{
		{{3, 0}, {4, 0}, {5, 0}, {5, 1}},
		{{4, 0}, {5, 0}, {4, 1}, {4, 2}},
		{{3, 0}, {3, 1}, {4, 1}, {5, 1}},
		{{4, 0}, {4, 1}, {4, 2}, {3, 2}},
	}),
	'S': NewPiece('S', [][][2]int{
		{{3, 0}, {4, 0}, {4, 1}, {5, 1}},
		{{5, 0}, {4, 1}, {5, 1}, {4, 2}},
	}),
	'Z': NewPiece('Z', [][][2]int{
		{{5, 0}, {4, 0}, {4, 1}, {3, 1}},
		{{4, 0}, {4, 1}, {5, 1}, {5, 2}},
	}),
	'T': NewPiece('T', [][][2]int{
		{{3, 0}, {4, 0}, {5, 0}, {4, 1}},
		{{4, 0}, {4, 1}, {5, 1}, {4, 2}},
		{{4, 0}, {3, 1}, {4, 1}, {5, 1}},
		{{4, 0}, {3, 1}, {4, 1}, {4, 2}},
	}),
}

package gotetris

import "sync"

var Tetrominoes []rune = []rune{'Z', 'L', 'O', 'S', 'I', 'J', 'T'}

// A PieceBag uses it's given Random Number Generator to generate pieces
// Pieces are generated in accordance with the Tetris guidelines
type PieceBag struct {
	rng         RNG
	refillSlots []bool
	pieces      []*Piece
	size        int
	mutex       sync.Mutex
}

func NewPieceBag(seed, size int, rng RNG) *PieceBag {
	return &PieceBag{rng: rng, size: size, refillSlots: make([]bool, size)}
}

/*
	Gives the Piece at a specified index within a bag
	Caches previously generated pieces, and allows pieces to be accessed at any index
	in any order whilst retaining determinism in results (assuming a consistent RNG)
*/
func (bag *PieceBag) AtIndex(index int) *Piece {
	if index >= len(bag.pieces) {

		// Mutex is used here to allow safe concurrent usage of a single PieceBag
		bag.mutex.Lock()
		for index >= len(bag.pieces) {
			if bag.refillSlotsFilled() {
				for i := 0; i < bag.size; i++ {
					bag.refillSlots[i] = false
				}
			}

			generatePiece := true
			pieceIndex := 0

			for generatePiece {
				pieceIndex = bag.rng.NextRandom() % bag.size
				generatePiece = bag.refillSlots[pieceIndex]
			}

			bag.refillSlots[pieceIndex] = true
			bag.pieces = append(bag.pieces, PieceMap[Tetrominoes[pieceIndex]])
		}
		bag.mutex.Unlock()
	}

	return bag.pieces[index]
}

func (bag *PieceBag) refillSlotsFilled() bool {
	for _, slot := range bag.refillSlots {
		if !slot {
			return false
		}
	}

	return true
}

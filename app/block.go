package app

type Block struct {
	eachBlockStartPositions [9][2]uint8
}

func NewBlock() *Block {
	return &Block{
		eachBlockStartPositions: [9][2]uint8{
			{0, 0}, {0, 3}, {0, 6},
			{3, 0}, {3, 3}, {3, 6},
			{6, 0}, {6, 3}, {6, 6},
		},
	}
}

func (b *Block) GetAllPositionInBlock(position [2]uint8) [9][2]uint8 {
	targetPosition := [2]uint8{}

	for i := 8; i >= 0; i-- {
		firstPosition := b.eachBlockStartPositions[i]

		if position[0] >= firstPosition[0] && position[1] >= firstPosition[1] {
			targetPosition = firstPosition
			break
		}
	}

	return getAllPositionFromStartPosition(targetPosition)
}

func (b *Block) GetAllPositionFromBlockNumber(blockNumber uint8) [9][2]uint8 {
	return getAllPositionFromStartPosition(b.eachBlockStartPositions[blockNumber])
}

func getAllPositionFromStartPosition(position [2]uint8) [9][2]uint8 {
	top := position[0]
	left := position[1]

	return [9][2]uint8{
		{top, left}, {top, left + 1}, {top, left + 2},
		{top + 1, left}, {top + 1, left + 1}, {top + 1, left + 2},
		{top + 2, left}, {top + 2, left + 1}, {top + 2, left + 2},
	}
}

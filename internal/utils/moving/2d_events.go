package moving

import "github.com/kercylan98/minotaur/utils/generic"

type (
	Position2DChangeEventHandle[EID generic.Basic, PosType generic.SignedNumber]      func(moving *TwoDimensional[EID, PosType], entity TwoDimensionalEntity[EID, PosType], oldX, oldY PosType)
	Position2DDestinationEventHandle[EID generic.Basic, PosType generic.SignedNumber] func(moving *TwoDimensional[EID, PosType], entity TwoDimensionalEntity[EID, PosType])
	Position2DStopMoveEventHandle[EID generic.Basic, PosType generic.SignedNumber]    func(moving *TwoDimensional[EID, PosType], entity TwoDimensionalEntity[EID, PosType])
)

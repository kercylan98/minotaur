package moving

type (
	Position2DChangeEventHandle      func(moving *TwoDimensional, entity TwoDimensionalEntity, oldX, oldY float64)
	Position2DDestinationEventHandle func(moving *TwoDimensional, entity TwoDimensionalEntity)
	Position2DStopMoveEventHandle    func(moving *TwoDimensional, entity TwoDimensionalEntity)
)

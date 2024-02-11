package permissions

type Permission int

const (
	None         Permission = iota // Implicitement 0
	User                           // 1 << 0 == 1 (0b001)
	Restaurateur                   // 1 << 1 == 2 (0b010)
	Admin                          // 1 << 2 == 4 (0b100)
)

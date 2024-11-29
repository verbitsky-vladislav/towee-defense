package enemy

type Animation string

const (
	Idle   = Animation("Idle")
	Move   = Animation("Move")
	Attack = Animation("Attack")
	Hit    = Animation("Hit")
	Death  = Animation("Death")
)

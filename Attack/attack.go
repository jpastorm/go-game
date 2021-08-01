package attack

type AttackRepository interface {
	getDamage() int
	getAction() int
}

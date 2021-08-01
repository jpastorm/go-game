package arena

import (
	"fmt"
	"github.com/jpastorm/go-game/soldier"
	"math/rand"
	"time"
)

type Arena struct {
	Soldiers map[string]*soldier.Soldier
}

func New(soldiers map[string]*soldier.Soldier) *Arena {
	return &Arena{Soldiers: soldiers}
}

func (a Arena) GetAliveSoldiers() (aliveSoldiers map[string]*soldier.Soldier) {
	aliveSoldiers = make(map[string]*soldier.Soldier)
	for _, soldier := range a.Soldiers {
		if soldier.IsAlive {
			aliveSoldiers[soldier.GetName()] = soldier
		}

	}
	return
}

func (a Arena) Create(s *soldier.Soldier) {
	a.Soldiers[s.Name] = s
}

func (a Arena) CreateBattle() {
	for _, soldier := range a.Soldiers {
		a.Create(soldier)
	}
}

func (a Arena) GetSoldiers() map[string]*soldier.Soldier {
	return a.Soldiers
}

func (a Arena) Fight() {
	for _, soldier := range a.Soldiers {
		enemy := a.GetEnemyToAttack(soldier)
		if !enemy.IsDeath() {
			soldier.Attack(enemy)
			time.Sleep(time.Second)
			fmt.Println(soldier)
			fmt.Println(enemy)
		}
	}
}

func (a Arena) GetEnemyToAttack(attacker *soldier.Soldier) *soldier.Soldier {
	soldiers := a.GetSoldiersNamesToAttack(attacker.Name)
	randomIndex := rand.Intn(len(soldiers))
	pick := soldiers[randomIndex]

	return a.Soldiers[pick]
}

func (a Arena) GetSoldiersNamesToAttack(name string) (names []string) {
	for _, soldier := range a.Soldiers {
		if soldier.Name != name {
			names = append(names, soldier.Name)
		}
	}
	return
}

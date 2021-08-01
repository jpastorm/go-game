package main

import (
	"fmt"
	"github.com/jpastorm/go-game/arena"
	"github.com/jpastorm/go-game/soldier"
)

func main() {

	adam := soldier.New("adam", 40)
	zeus := soldier.New("zeus", 40)
	hercules := soldier.New("hercules", 60)
	jack := soldier.New("jack", 60)

	soldiers := make(map[string]*soldier.Soldier)
	soldiers[adam.Name] = adam
	soldiers[zeus.Name] = zeus
	soldiers[hercules.Name] = hercules
	soldiers[jack.Name] = jack
	arena := arena.New(soldiers)
	arena.CreateBattle()

	for {
		arena.Fight()

		if len(arena.GetAliveSoldiers()) <= 1 {
			break
		}
	}
	fmt.Println("WINNER")
	fmt.Println(arena.GetAliveSoldiers())
}

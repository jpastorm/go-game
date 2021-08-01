package soldier

type Soldier struct {
	Name    string
	Hp      int
	IsAlive bool
	Team    string
}

func New(Name string, Hp int) *Soldier {
	return &Soldier{Name: Name, Hp: Hp, IsAlive: true}
}

func (s *Soldier) GetName() string {
	return s.Name
}

func (s *Soldier) Attack(enemy *Soldier) {
	if s.IsDeath() {
		return
	}
	enemy.ReceivesAttack(20)
}

func (s *Soldier) ReceivesAttack(damage int) {
	s.Hp = s.Hp - damage
	if s.Hp <= 0 {
		s.Die()
	}
}

func (s *Soldier) Die() {
	s.IsAlive = false
}

func (s *Soldier) IsDeath() bool {
	return !s.IsAlive
}

/*

   public function setWeapon(WeaponRepository $weapon): void
   {
       $this->attacks->add(...$weapon->attacks());
       printf("%s prepares his %s as weapon\n", $this->Name, $this->weapon->getname());
   }

   public function attack(AttackRepository $attack, Soldier $soldier): void
   {
       printf(
           "%s %s %s \n",
           $this->getName(),
           $attack->getAction(),
           $enemy->getName()
       );

       $soldier->recievesAttack($attack);
   }

   private function recievesAttack(AttackRepository $attack): void
   {
       $this->Hp -= $attack->getDamage();
       printf("%s received attack \n", $this->Name);
       printf("%s Hp is now %s \n", $this->Name, $this->Hp);

       if (!$this->IsAlive()) {
           $this->die();
       }
   }

   private function die(): void
   {
       $this->IsAlive = false;
       printf("%s died \n", $this->Name);
   }

   public function IsAlive(): bool
   {
       return $this->Hp > 0;
   }

   public function attacks(): AttackCollection
   {
       return $this->attacks;
   } */

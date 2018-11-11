package entity

type Pokemon struct {
	Name string `bson:"name"`
	Item string `bson:"item"`
	Ability string `bson:"ability"`
	Level int `bson:"level"`
	HitPointsEV int `bson:"hit_points_ev"`
	AttackEV int `bson:"attack_ev"`
	DefenseEV int `bson:"defense_ev"`
	SpecialAttackEV int `bson:"special_attack_ev"`
	SpecialDefenseEV int `bson:"special_defense_ev"`
	SpeedEV int `bson:"speed_ev"`
	Nature string `bson:"nature"`
	HitPointsIV int `bson:"hit_points_iv"`
	AttackIV int `bson:"attack_iv"`
	DefenseIV int `bson:"defense_iv"`
	SpecialAttackIV int `bson:"special_attack_iv"`
	SpecialDefenseIV int `bson:"special_defense_iv"`
	SpeedIV int `bson:"speed_iv"`
	Moves []string `bson:"moves"`
}



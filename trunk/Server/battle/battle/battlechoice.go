package main

const (
	CHOICETYPE_CANCELTYPE int = _iota
	CHOICETYPE_ATTACKTYPE
	CHOICETYPE_SWITCHTYPE
	CHOICETYPE_REARRANGETYPE
	CHOICETYPE_CENTERMOVETYPE
	CHOICETYPE_DRAWTYPE
)

type Choice interface {
	GetChoiceType() int
}

//
// Attack Choice
type AttackChoice struct {
	AttackSlot int
	AttackTarget int
}

func NewAttackChoice(_as, _at int) *AttackChoice {
	attackChoice := AttackChoice{ AttackSlot: _as, AttackTarget: _at }
	return &attackChoice
}

func NewAttackChoiceFromPacket(_packet *pnet.QTPacket) *AttackChoice {
	return NewAttackChoice((int)_packet.ReadByte(), (int)_packet.ReadByte())
}

func (c *AttackChoice) GetChoiceType() int {
	return CHOICETYPE_ATTACKTYPE
}

//
// Switch Choice
type SwitchChoice struct {
	PokeSlot int
}

func NewSwitchChoice(_slot int) *SwitchChoice {
	switchChoice := SwitchChoice{ PokeSlot: _slot }
	return &switchChoice
}

func NewSwitchChoiceFromPacket(_packet *pnet.QTPacket) *SwitchChoice {
	return NewSwitchChoice((int)_packet.ReadByte())
}

func (c *SwitchChoice) GetChoiceType() int {
	return CHOICETYPE_ATTACKTYPE
}

//
// Rearrange Choice
type RearrangeChoice struct {
	PokeIndexes []int
}

func NewRearrangeChoice(_team *BattleTeam) *RearrangeChoice {
	rearrangeChoice := RearrangeChoice{ }
	rearrangeChoice.PokeIndexes = make([]int, 6)
	for i := 0; i < 6; i++ {
		rearrangeChoice.PokeIndexes[i] = _team.Pokes[i].TeamNum
	}
	
	return &rearrangeChoice
}

func NewRearrangeChoiceFromPacket(_packet *pnet.QTPacket) *RearrangeChoice {
	rearrangeChoice := RearrangeChoice{ }
	rearrangeChoice.PokeIndexes = make([]int, 6)
	for i := 0; i < 6; i++ {
		rearrangeChoice.PokeIndexes[i] = _packet.ReadByte()
	}
	
	return &rearrangeChoice
}

func (c *RearrangeChoice) GetChoiceType() int {
	return CHOICETYPE_REARRANGETYPE
}

//
// MovetoCenter Choice
type MoveToCenterChoice struct {
}

func NewMoveToCenterChoice() *MoveToCenterChoice {
	return &MoveToCenter{}
}

func (m *MoveToCenterChoice) GetChoiceType() int {
	return CHOICETYPE_MOVETOCENTERCHOICE
}

//
// Draw Choice
type DrawChoice struct {
}

func NewDrawChoice() *DrawChoice {
	return &DrawChoice{}
}

func (m *DrawChoice) GetChoiceType() int {
	return CHOICETYPE_DRAWCHOICE
}

//
// Battle Choice
type BattleChoice struct {
	PlayerSlot int
	Choice IChoice
	ChoiceType int
}

func NewBattleChoiceWithChoice(_ps int, _c IChoice, _ct int) *BattleChoice {
	battleChoice := BattleChoice { PlayerSlot: _ps,
									Choice _c,
									ChoiceType: _ct }
	return &battleChoice
}

func NewBattleChoice(_ps int, _ct int) *BattleChoice {
	battleChoice := BattleChoice { PlayerSlot: ps,
									ChoiceType: _ct }
}

func NewBattleChoiceFromPacket(_packet *pnet.QTPacket) *BattleChoice {
	battleChoice := BattleChoice{}
	battleChoice.PlayerSlot = (int)_packet.ReadByte()
	battleChoice.ChoiceType = (int)_packet.ReadByte()
	
	switch t := battleChoice.ChoiceType {
		case CHOICETYPE_SWITCHTYPE:
			battleChoice.Choice = NewSwitchChoiceFromPacket(_packet)
		case CHOICETYPE_ATTACKTYPE:
			battleChoice.Choice = NewAttackChoiceFromPacket(_paket)
		case CHOICETYPE_REARRANGETYPE:
			battleChoice.Choice = NewRearrangeChoiceFromPacket(_packet)
	}
}
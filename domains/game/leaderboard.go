package game

import (
	"sync"
)

const answeredKey = "answered"
const correctedKey = "corrected"
const pointKey = "point"

type Leaderboard struct {
	m     *sync.Mutex
	value map[string]map[string]uint32
}

func NewLeaderboard() Leaderboard {
	lb := make(map[string]map[string]uint32)
	return Leaderboard{new(sync.Mutex), lb}
}

func (lb Leaderboard) AddPlayer(
	player Player,
) error {
	defer lb.m.Unlock()
	lb.m.Lock()
	_, ok := lb.value[player.Name()]
	if ok {
		return ErrPlayerIsAdded
	}
	lb.value[player.Name()] = map[string]uint32{
		answeredKey:  0,
		correctedKey: 0,
		pointKey:     0,
	}
	return nil
}

func (lb *Leaderboard) Reset() {
	for playerName := range lb.value {
		lb.value[playerName] = map[string]uint32{
			answeredKey:  0,
			correctedKey: 0,
			pointKey:     0,
		}
	}
}

func (lb *Leaderboard) AddPoint(
	playerName string,
	addition uint32,
) error {
	defer lb.m.Unlock()
	lb.m.Lock()
	_, ok := lb.value[playerName]
	if !ok {
		return ErrPlayerIsNotExisted
	}
	lb.value[playerName][answeredKey] += 1
	if addition > 0 {
		lb.value[playerName][correctedKey] += 1
	}
	lb.value[playerName][pointKey] += addition
	return nil
}

func (lb Leaderboard) Highest() {

}

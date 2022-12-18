package main

import (
	_ "embed"
	"fmt"
	"sort"
	s "strings"
	"time"

	"aoc/util"
)

type Valve struct {
	name    string
	flow    int
	open    bool
	tunnels map[string]int
}

func NewValve(name string, flow int, tunnels ...string) *Valve {
	v := Valve{
		name:    name,
		flow:    flow,
		tunnels: map[string]int{},
	}

	for _, t := range tunnels {
		v.tunnels[t] = 1
	}

	return &v
}

func (v Valve) String() string {
	res := fmt.Sprintf("[%s %d valves: [", v.name, v.flow)
	for k, v := range v.tunnels {
		res += fmt.Sprintf("%s: %d ", k, v)
	}
	return res + "]]"
}

func CompressTunnels(valve Valve, valvesMap map[string]Valve) *Valve {
	anyZeros := false
	for valveName := range valve.tunnels {
		if valvesMap[valveName].flow == 0 && valveName != "AA" {
			anyZeros = true
			break
		}
	}

	dropped := map[string]bool{}

	for anyZeros {
		for valveName, cost := range valve.tunnels {
			currentValve := valvesMap[valveName]
			if currentValve.flow == 0 && currentValve.name != "AA" {
				for v, c := range currentValve.tunnels {
					if _, ok := dropped[v]; !ok && v != valve.name {
						valve.tunnels[v] = cost + c
					}
				}
				delete(valve.tunnels, valveName)
				dropped[valveName] = true
			}
		}

		anyZeros = false
		for valveName := range valve.tunnels {
			if valvesMap[valveName].flow == 0 && valveName != "AA" {
				anyZeros = true
				break
			}
		}
	}

	return &valve
}

func parseCompressed() map[string]Valve {
	lines := s.Split(input, "\n")

	valves := map[string]Valve{}

	for _, line := range lines {
		words := s.Split(line, " ")
		name := words[1]
		var flow int
		fmt.Sscanf(words[4], "rate=%d", &flow)
		tunnels := s.Split(s.Join(words[9:], ""), ",")
		valves[name] = *NewValve(name, flow, tunnels...)
	}

	compressed := map[string]Valve{}
	for i := range valves {
		compressed[i] = *CompressTunnels(valves[i], valves)
	}

	for i := range compressed {
		if compressed[i].flow == 0 && compressed[i].name != "AA" {
			delete(compressed, i)
		}
	}

	return compressed
}

func parse() map[string]Valve {
	lines := s.Split(input, "\n")

	valves := map[string]Valve{}

	for _, line := range lines {
		words := s.Split(line, " ")
		name := words[1]
		var flow int
		fmt.Sscanf(words[4], "rate=%d", &flow)
		tunnels := s.Split(s.Join(words[9:], ""), ",")
		valves[name] = *NewValve(name, flow, tunnels...)
	}

	return valves
}

type State struct {
	current          string
	prev             string
	time             int
	flow             int
	opened           string
	pressureReleased int
}

func (s State) copy() State {
	new := State{
		prev:             s.prev,
		current:          s.current,
		time:             s.time,
		flow:             s.flow,
		opened:           s.opened,
		pressureReleased: s.pressureReleased,
	}

	return new
}

type StateWithElephant struct {
	curr1            string
	curr2            string
	prev1            string
	prev2            string
	step1            int
	step2            int
	time             int
	flow             int
	opened           string
	pressureReleased int
	prevState        *StateWithElephant
}

func (s StateWithElephant) copy() StateWithElephant {
	new := StateWithElephant{
		curr1:            s.curr1,
		curr2:            s.curr2,
		prev1:            s.prev1,
		prev2:            s.prev2,
		step1:            s.step1,
		step2:            s.step2,
		time:             s.time,
		flow:             s.flow,
		opened:           s.opened,
		pressureReleased: s.pressureReleased,
		prevState:        s.prevState,
	}

	return new
}

func (st StateWithElephant) hash() string {
	var first, second, fprev, sprev string
	if st.curr1 > st.curr2 {
		first, second = st.curr1, st.curr2
		fprev, sprev = st.prev1, st.prev2
	} else {
		first, second = st.curr2, st.curr1
		fprev, sprev = st.prev2, st.prev1
	}

	o := s.Split(st.opened, " ")
	sort.Strings(o)

	return fmt.Sprintf("%s%s%s%s%d%d%s%d", first, second, fprev, sprev, st.time, st.flow, s.Join(o, " ")+" ", st.pressureReleased)
}



func simulate(state State, valves map[string]Valve, visited *map[State]bool) int {
	timeLimit := 29
	best := 0

	if len(state.opened)*3 == len(valves) {
		for state.time != timeLimit {
			state.time++
			state.pressureReleased += state.flow
		}
	}

	if state.time == timeLimit {
		return state.pressureReleased
	}

	currentValve := valves[state.current]

	if !s.Contains(state.opened, state.current+" ") && currentValve.flow != 0 {
		s := state.copy()
		s.opened += state.current + " "
		s.time++
		s.flow += currentValve.flow
		s.pressureReleased += s.flow
		s.prev = "--"
		if _, exists := (*visited)[s]; !exists {
			(*visited)[s] = true
			if result := simulate(s, valves, visited); result > best {
				best = result
			}
		}
	}

	for next, step := range currentValve.tunnels {
		if next != state.prev {
			s := state.copy()
			s.time += step
			s.current = next
			s.prev = state.current
			s.pressureReleased += s.flow * step
			if _, exists := (*visited)[s]; !exists && s.time <= timeLimit {
				(*visited)[s] = true
				if result := simulate(s, valves, visited); result > best {
					best = result
				}
			}
		}
	}

	return best
}

func simulate2(state StateWithElephant, valves map[string]Valve, visited *map[string]bool) int {
	timeLimit := 27
	best := 0

	ok1, ok2 := true, true
	if state.step1 > 0 {
		ok1 = false
		state.step1--
	}
	if state.step2 > 0 {
		ok2 = false
		state.step2--
	}

	flow := 0
	for _, valveName := range s.Split(state.opened, " ") {
		flow += valves[valveName].flow
	}

	state.pressureReleased += flow
	state.time++

	if state.time > 15 {

		if bestScores[state.time] > state.pressureReleased {
			return -1
		} else if bestScores[state.time] < state.pressureReleased {
			bestScores[state.time] = state.pressureReleased
			fmt.Println(bestScores)
		}
	}

	if len(state.opened)*3 == len(valves) {
		for state.time != timeLimit {
			state.time++
			state.pressureReleased += state.flow
		}
	}

	if state.time == timeLimit {
		return state.pressureReleased
	}

	currentValve := valves[state.curr1]
	currentValve2 := valves[state.curr2]

	if !ok1 && !ok2 {
		st := state.copy()
		if _, exists := (*visited)[st.hash()]; !exists {
			(*visited)[st.hash()] = true
			if result := simulate2(st, valves, visited); result > best {
				best = result
			}
		}
	} else if !ok1 && ok2 {
		if !s.Contains(state.opened, state.curr2+" ") && currentValve2.flow != 0 {
			st := state.copy()
			st.prevState = &state
			st.opened += state.curr2 + " "
			st.flow += currentValve.flow
			st.prev2 = "--"

			if _, exists := (*visited)[st.hash()]; !exists {
				(*visited)[st.hash()] = true
				if result := simulate2(st, valves, visited); result > best {
					// fmt.Println(result)
					best = result
				}
			}
		}

		for next, step := range currentValve2.tunnels {
			if next != state.prev2 && next != state.curr1 {
				st := state.copy()
				st.curr2 = next
				st.prev2 = state.curr2
				st.step2 = step - 1
				if _, exists := (*visited)[st.hash()]; !exists {
					(*visited)[st.hash()] = true
					if result := simulate2(st, valves, visited); result > best {
						// fmt.Println(result)
						best = result
					}
				}
			}
		}

	} else if ok1 && !ok2 {
		if !s.Contains(state.opened, state.curr1+" ") && currentValve.flow != 0 {
			st := state.copy()
			st.prevState = &state
			st.opened += state.curr1 + " "
			st.flow += currentValve.flow
			st.prev1 = "--"

			if _, exists := (*visited)[st.hash()]; !exists {
				(*visited)[st.hash()] = true
				if result := simulate2(st, valves, visited); result > best {
					// fmt.Println(result)
					best = result
				}
			}
		}

		for next, step := range currentValve.tunnels {
			if next != state.prev1 && next != state.curr1 {
				st := state.copy()
				st.curr1 = next
				st.prev1 = state.curr1
				st.step1 = step - 1
				if _, exists := (*visited)[st.hash()]; !exists {
					(*visited)[st.hash()] = true
					if result := simulate2(st, valves, visited); result > best {
						// fmt.Println(result)
						best = result
					}
				}
			}
		}
	} else {
		if !s.Contains(state.opened, state.curr1+" ") && currentValve.flow != 0 {
			st := state.copy()
			st.prevState = &state
			st.opened += state.curr1 + " "
			st.flow += currentValve.flow
			st.prev1 = "--"

			if !s.Contains(state.opened, state.curr2+" ") && currentValve2.flow != 0 && st.curr1 != st.curr2 {
				st.opened += st.curr2 + " "
				st.flow += currentValve2.flow
				st.prev2 = "--"

				if _, exists := (*visited)[st.hash()]; !exists {
					(*visited)[st.hash()] = true
					if result := simulate2(st, valves, visited); result > best {
						// fmt.Println(result)
						best = result
					}
				}

			} else {
				for next, step := range currentValve2.tunnels {
					if next != st.prev2 && next != st.curr1 {
						st.curr2 = next
						st.prev2 = state.curr2
						st.step2 = step - 1
						if _, exists := (*visited)[st.hash()]; !exists {
							(*visited)[st.hash()] = true
							if result := simulate2(st, valves, visited); result > best {
								best = result
							}
						}
					}
				}
			}

		}

		for next, step := range currentValve.tunnels {
			if next != state.prev1 {
				st := state.copy()
				st.prevState = &state
				st.curr1 = next
				st.prev1 = state.curr1
				st.step1 = step - 1

				if !s.Contains(state.opened, state.curr2+" ") && currentValve2.flow != 0 {
					st.opened += st.curr2 + " "
					st.flow += currentValve2.flow
					st.prev2 = "--"

					if _, exists := (*visited)[st.hash()]; !exists {
						(*visited)[st.hash()] = true
						if result := simulate2(st, valves, visited); result > best {
							best = result
						}
					}

				} else {
					for next, step := range currentValve2.tunnels {
						if next != state.prev2 && next != st.curr1 {
							st.curr2 = next
							st.prev2 = state.curr2
							st.step2 = step - 1
							if _, exists := (*visited)[st.hash()]; !exists {
								(*visited)[st.hash()] = true
								if result := simulate2(st, valves, visited); result > best {
									best = result
								}
							}
						}
					}
				}

			}
		}
	}

	return best
}

func part1() {
	valves := parseCompressed()

	initialState := State{
		prev:             "--",
		current:          "AA",
		time:             1,
		flow:             0,
		pressureReleased: 0,
	}

	func() {
		defer util.TimeTrack(time.Now(), "sim")
		cache := make(map[State]bool)
		result := simulate(initialState, valves, &cache)
		fmt.Println("res: ", result)
	}()

}

func (st StateWithElephant) String() string {
	return fmt.Sprintf("{%s %s  %s %s  %2d %4d %s}", st.curr1, st.prev1, st.curr2, st.prev2, st.flow, st.pressureReleased, st.opened)
}

var bestScores = make([]int, 30) 

func part2() {
	valves := parseCompressed()

	for _, valve := range valves {
		fmt.Println(valve)
	}

	initialState := StateWithElephant{
		curr1: "AA",
		curr2: "AA",
		prev1: "--",
		prev2: "--",
		time:  1,
	}

	func() {
		defer util.TimeTrack(time.Now(), "sim")
		cache := make(map[string]bool)
		result := simulate2(initialState, valves, &cache)
		fmt.Println(result)
	}()
}

//go:embed sample.txt
var input string

func main() {
	// part1()
	part2()
}

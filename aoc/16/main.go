package main

import (
	_ "embed"
	"fmt"
	"log"
	s "strings"
	"time"
)

//go:embed input.txt
var input string

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

type State struct {
	current          string
	prev             string
	time             int
	flow             int
	opened           map[string]bool
	pressureReleased int
	prevState        *State
}

func (s State) copy() State {
	new := State{
		prev:             s.prev,
		current:          s.current,
		time:             s.time,
		flow:             s.flow,
		opened:           make(map[string]bool),
		pressureReleased: s.pressureReleased,
	}
	for k, v := range s.opened {
		new.opened[k] = v
	}
	return new
}

type State2 struct {
	current          string
	prev             string
	time             int
	flow             int
	opened           string
	pressureReleased int
}

func (s State2) copy() State2 {
	new := State2{
		prev:             s.prev,
		current:          s.current,
		time:             s.time,
		flow:             s.flow,
		opened:           s.opened,
		pressureReleased: s.pressureReleased,
	}

	return new
}

func simulate(state State, valves map[string]Valve) int {
	timeLimit := 29
	best := 0

	if len(state.opened) == len(valves) {
		for state.time != timeLimit {
			state.time++
			state.pressureReleased += state.flow
		}
	}

	if state.time == timeLimit {
		return state.pressureReleased
	}

	currentValve := valves[state.current]

	if _, opened := state.opened[state.current]; !opened && currentValve.flow != 0 {
		s := state.copy()
		s.opened[state.current] = true
		s.time++
		s.flow += currentValve.flow
		s.pressureReleased += s.flow
		s.prev = "--"
		if result := simulate(s, valves); result > best {
			best = result
		}
	}

	for next, step := range currentValve.tunnels {
		if next != state.prev {
			s := state.copy()
			s.time += step
			s.current = next
			s.prev = state.current
			s.pressureReleased += s.flow * step
			if s.time <= timeLimit {
				if result := simulate(s, valves); result > best {
					best = result
				}
			}
		}
	}

	return best
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func simulate2(state State, valves map[string]Valve) *State {
	timeLimit := 29

	if state.time == timeLimit {
		return &state
	}

	results := []*State{}
	currentValve := valves[state.current]

	if _, opened := state.opened[state.current]; !opened {
		s := state.copy()
		s.opened[state.current] = true
		s.time++
		s.flow += currentValve.flow
		s.pressureReleased += s.flow
		s.prev = "--"
		s.prevState = &state
		// fmt.Println("open: ", s)
		results = append(results, simulate2(s, valves))
	}

	for next, step := range currentValve.tunnels {
		if next != state.prev {
			s := state.copy()
			s.time += step
			s.pressureReleased += s.flow * step
			s.current = next
			s.prev = state.current
			s.prevState = &state
			if s.time <= timeLimit {
				results = append(results, simulate2(s, valves))
			}
		}
	}

	max := 0
	bestState := &State{}
	for _, res := range results {
		// fmt.Println(res)
		if res.pressureReleased > max {
			max = res.pressureReleased
			bestState = res
		}
	}
	return bestState

}

func simulate3(state State2, valves map[string]Valve) int {
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
		if result := simulate3(s, valves); result > best {
			best = result
		}
	}

	for next, step := range currentValve.tunnels {
		if next != state.prev {
			s := state.copy()
			s.time += step
			s.current = next
			s.prev = state.current
			s.pressureReleased += s.flow * step
			if s.time <= timeLimit {
				if result := simulate3(s, valves); result > best {
					best = result
				}
			}
		}
	}

	return best
}

func part1() {
	valves := parse()

	// start := valves["AA"]
	// for v := range valves {
	// 	fmt.Println(valves[v])
	// }

	initialState := State{
		prev:             "--",
		current:          "AA",
		time:             0,
		flow:             0,
		opened:           make(map[string]bool),
		pressureReleased: 0,
	}

	initialState2 := State2{
		prev:             "--",
		current:          "AA",
		time:             0,
		flow:             0,
		pressureReleased: 0,
	}

	// simulate(initialState, valves)
	func() {
		defer timeTrack(time.Now(), "sim")
		result := simulate3(initialState2, valves)
		fmt.Println("res: ", result)
	}()

	func() {
		defer timeTrack(time.Now(), "sim1")
		result := simulate(initialState, valves)
		fmt.Println("res: ", result)
	}()

	func() {
		defer timeTrack(time.Now(), "sim2")
		result2 := simulate2(initialState, valves)
		fmt.Println("res2: ", result2)
	}()

}

func part2() {}

func main() {
	part1()
}

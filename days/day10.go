package days

import (
	"bytes"
	"sjbtimdan/aoc2025-go/utils"

	"github.com/go-air/gini"
	"github.com/go-air/gini/logic"
	"github.com/go-air/gini/z"
)

func Day10(contents []byte) utils.Answers {
	return utils.StringAnswers("TODO", "TODO")
}

type Machine struct {
	desiredState []bool
	buttons      [][]uint8
}

func (m Machine) countButtonPresses() int {
	L := logic.NewC()
	g := gini.New()
	L.ToCnf(g)
	buttonPresses := []z.Lit{}
	for i := 0; i < len(m.desiredState); i++ {
		buttonPresses = append(buttonPresses, z.Var(1).Pos())
	}
	res := L.F
	for iState, state := range m.desiredState {
		println("Working on state: ", iState, " == ", state)
		for ib, button := range m.buttons {
			println("Working on button", ib)
			buttonSwitch := button[iState]
			if buttonSwitch == 1 {
				res = L.Xor(buttonPresses[iState], res)
			} else {
				res = L.Xor(buttonPresses[iState].Not(), res)
			}
		}
		if !state {
			res = res.Not()
		}
		g.Add(res)
		g.Add(0)
	}
	return g.Solve()
}

func parseMachines(contents []byte) []Machine {
	lines := bytes.Split(contents, []byte("\n"))
	machines := []Machine{}
	for line := range lines {
		parts := bytes.Split(lines[line], []byte(" "))
		desiredStateBytesIncludingBrackets := parts[0]
		desiredStateBytes := desiredStateBytesIncludingBrackets[1 : len(desiredStateBytesIncludingBrackets)-1]
		desiredState := uint16(0)
		desiredStateBytesLen := len(desiredStateBytes)
		for i := 0; i < len(desiredStateBytes); i++ {
			desiredStateBit := stateByteToBit(desiredStateBytes[len(desiredStateBytes)-1-i])
			desiredState = desiredState*2 + desiredStateBit
		}
		buttonsBytesArray := parts[1 : len(parts)-1]
		buttons := []uint16{}
		for _, buttonsBytesWithParens := range buttonsBytesArray {
			button := uint16(0)
			buttonBytes := bytes.Split(buttonsBytesWithParens[1:len(buttonsBytesWithParens)-1], []byte(","))
			for _, buttonByte := range buttonBytes {
				bitIndex := desiredStateBytesLen - (int(buttonByte[0]) - '0') - 1
				button += 1 << bitIndex
			}
			buttons = append(buttons, button)
		}
		machine := Machine{
			desiredState: []bool{},
			buttons:      [][]uint8{},
		}
		machines = append(machines, machine)
	}
	return machines
}

func stateByteToBit(b byte) uint16 {
	if b == '#' {
		return 1
	} else {
		return 0
	}
}

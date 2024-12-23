package main

import (
	"AdventOfCode2024/utils"
	"context"
	"fmt"
	"math"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func part1(input string) any {

	sections := strings.Split(input, "\n\n")
	regA, regB, regC := parseRegisters(sections[0])
	program := parseProgram(sections[1])

	machine := &Machine{
		regA:    int64(regA),
		regB:    int64(regB),
		regC:    int64(regC),
		ip:      0,
		program: program,
		output:  []int{},
	}
	for !machine.HasHalted() {
		machine.Process()
	}

	outStr := utils.Map(machine.output, strconv.Itoa)
	return strings.Join(outStr, ",")
}

type Machine struct {
	regA, regB, regC int64
	ip               int
	program          []int
	output           []int
	targetOutput     []int
	powCache         map[int]float64
}

func (m *Machine) HasHalted() bool {
	return m.ip >= len(m.program)
}

func (m *Machine) Process() (aShift bool, output bool) {

	opcode := m.program[m.ip]
	operand := m.program[m.ip+1]

	switch opcode {
	case 0:
		adv(m, operand)
		//fmt.Printf("adv(%d); ", operand)
		//fmt.Printf("\nA:%0*b\nB:%0*b\nC:%0*b\n",
		//	len(m.program)*3, m.regA,
		//	len(m.program)*3, m.regB,
		//	len(m.program)*3, m.regC)
		return true, false
	case 1:
		//fmt.Printf("bxl(%d); ", operand)
		bxl(m, operand)
	case 2:
		//fmt.Printf("bst(%d); ", operand)
		bst(m, operand)
	case 3:
		//fmt.Printf("jnz(%d); ", operand)
		jnz(m, operand)
	case 4:
		//fmt.Printf("bxc(%d); ", operand)
		bxc(m, operand)
	case 5:
		out(m, operand)
		//fmt.Printf("out(%d)", operand)
		//fmt.Printf("\nO:%0*b",
		//	len(m.program)*3, m.output[len(m.output)-1])
		//fmt.Printf("\nA:%0*b\nB:%0*b\nC:%0*b\n",
		//	len(m.program)*3, m.regA,
		//	len(m.program)*3, m.regB,
		//	len(m.program)*3, m.regC)
		return false, true

	case 6:
		//fmt.Printf("bdv(%d); ", operand)
		bdv(m, operand)
	case 7:
		//fmt.Printf("cdv(%d); ", operand)
		cdv(m, operand)

	default:
		panic("invalid opcode")
	}
	//fmt.Printf("\nA:%0*b\nB:%0*b\nC:%0*b\n",
	//	len(m.program)*3, m.regA,
	//	len(m.program)*3, m.regB,
	//	len(m.program)*3, m.regC)
	return false, false
	//opcodes[opcode](m, operand)
}

var opcodes = map[int]func(m *Machine, operand int){
	0: adv,
	1: bxl,
	2: bst,
	3: jnz,
	4: bxc,
	5: out,
	6: bdv,
	7: cdv,
}

func powInt(x int64) int64 {
	if x == 0 {
		return 1
	}
	return 1 << x
}

func adv(m *Machine, operand int) {
	comboOp := m.parseComboOperand(operand)
	res := m.regA / powInt(comboOp)
	m.regA = res
	m.ip += 2
}

func bxl(m *Machine, operand int) {
	res := m.regB ^ int64(operand)
	m.regB = res
	m.ip += 2
}

func bst(m *Machine, operand int) {
	comboOp := m.parseComboOperand(operand)
	res := comboOp % 8
	m.regB = res
	m.ip += 2
}

func jnz(m *Machine, operand int) {
	if m.regA == 0 {
		m.ip += 2
		return
	}
	m.ip = operand
}

func bxc(m *Machine, operand int) {
	res := m.regB ^ m.regC
	m.regB = res
	m.ip += 2
}

func out(m *Machine, operand int) {
	comboOp := m.parseComboOperand(operand)

	res := int(comboOp % 8)
	if len(m.targetOutput) > 0 {
		currI := len(m.output)
		if m.targetOutput[currI] != res {
			m.ip = len(m.program) + 1
			return
		}
	}

	m.output = append(m.output, res)
	m.ip += 2
}

func bdv(m *Machine, operand int) {
	comboOp := m.parseComboOperand(operand)
	res := m.regA / powInt(comboOp)
	m.regB = res
	m.ip += 2
}

func cdv(m *Machine, operand int) {
	comboOp := m.parseComboOperand(operand)
	res := m.regA / powInt(comboOp)

	m.regC = res
	m.ip += 2
}

func (m *Machine) cachedPow(x int) float64 {
	if v, ok := m.powCache[x]; ok {
		return v
	}
	v := math.Pow(2, float64(x))
	m.powCache[x] = v
	return v
}

func (m *Machine) parseComboOperand(operand int) int64 {

	switch operand {
	case 0:
		return int64(operand)
	case 1:
		return int64(operand)
	case 2:
		return int64(operand)
	case 3:
		return int64(operand)
	case 4:
		return m.regA
	case 5:
		return m.regB
	case 6:
		return m.regC
	}
	panic("invalid combo operand")
}

func (m *Machine) reset(regA int64) {
	m.regA = regA
	m.regB = 0
	m.regC = 0
	m.ip = 0
	m.output = []int{}
}

func (m *Machine) calcOut(regA int64) []int {
	x, _ := m.calcOutOps(regA)
	return x
}

func (m *Machine) calcOutOps(regA int64) ([]int, []rune) {
	m.reset(regA)
	var ops []rune
	for !m.HasHalted() {
		shift, output := m.Process()

		if output {
			ops = append(ops, 'O')
		} else if shift {
			ops = append(ops, '>')
		}

	}
	//fmt.Printf("\n")
	return m.output, ops
}

func parseProgram(input string) []int {
	var out []int
	input = strings.TrimPrefix(input, "Program: ")
	for _, s := range strings.Split(input, ",") {
		atoi, _ := strconv.Atoi(s)
		out = append(out, atoi)
	}

	return out
}

func parseRegisters(input string) (int, int, int) {
	var out []int
	for _, s := range strings.Split(input, "\n") {
		split := strings.Split(s, ": ")
		atoi, _ := strconv.Atoi(split[1])
		out = append(out, atoi)
	}

	return out[0], out[1], out[2]
}

func invert(mask int64, offset int, target int64) int64 {
	return 0
}

func recSol(program []int, currPos int, cumSol int64) (bool, int64) {

	machine := &Machine{
		regA:    int64(0),
		ip:      0,
		program: program,
		output:  []int{},
	}
	currOut := machine.calcOut(cumSol)
	if slices.Equal(currOut, program) {
		return true, cumSol
	}

	//
	//if currPos == len(program) {
	//	return true, cumSol
	//}
	//
	//b:= 0
	//if currPos == 0 {
	//	b = 1
	//}
	//shiftAmt := len(program) - 1 - currPos
	//target := program[shiftAmt]
	//for ; b <= 0b111; b++ {
	//	bval := int64(b << (3 * currPos))
	//}

	return false, 0
}

func findPartial(startOffset bool, sol int, bitPos int, shift int, target int, targetLen int, machine *Machine) (bool, int) {
	//mask := int(^(uint(0b111 << (3 * shift))))

	start := sol >> (3 * shift) & 0b111
	if startOffset {
		start++
		if start > 0b111 {
			panic("help")
		}
	}
	fmt.Printf("start %d = %b (shift:%d)\n", start, start, shift)

	//fmt.Printf("mask %b\nstart %b\n", mask, start)

	//sol = sol & mask

	bitLen := targetLen * 3

	fmt.Printf("------searching shift %d for target %d\n", shift, target)
	for b := start; b <= 0b111; b++ {
		if bitPos == 0 && b == 0 {
			continue
		}

		bval := int(b << (3 * shift))
		currInput := sol | bval
		out := machine.calcOut(int64(currInput))

		fmt.Printf("test %d f(%s) = %s | f(%d) = %d | f(%d) = %d\n",
			b, intToBin(currInput, bitLen), intToBin(bitsToInt(out), bitLen), intsToBits(currInput), out, currInput, bitsToInt(out))
		if len(out) != targetLen {
			continue
		}

		if out[shift] == target {
			//found = true
			//sol = currInput
			//fmt.Printf("SET f(%s) = %s | f(%d) = %d | f(%d) = %d\n",
			//	intToBin(sol, bitLen), intToBin(bitsToInt(out), bitLen), intsToBits(int(sol)), out, sol, bitsToInt(out))
			fmt.Printf("%d f(%s) = %s | f(%d) = %d | f(%d) = %d\n",
				b, intToBin(currInput, bitLen), intToBin(bitsToInt(out), bitLen), intsToBits(currInput), out, currInput, bitsToInt(out))

			return true, currInput
		}
	}

	return false, sol
}

func recur(machine *Machine, a int64, index int) *int64 {
	for tribble := int64(0); tribble <= 0b111; tribble++ {
		nextA := (a << 3) | tribble
		output := machine.calcOut(nextA)

		if output[0] == machine.program[index] {
			if index == 0 {
				return &nextA
			}

			if ar := recur(machine, nextA, index-1); ar != nil {
				return ar
			}
		}
	}
	return nil
}

func part2(input string) any {
	sections := strings.Split(input, "\n\n")
	_, regB, regC := parseRegisters(sections[0])
	program := parseProgram(sections[1])

	//powCache := map[int]float64{}
	regA := -1

	machine := &Machine{
		regA:    int64(regA),
		regB:    int64(regB),
		regC:    int64(regC),
		ip:      0,
		program: program,
		output:  []int{},
		//targetOutput: program,
		//powCache:     powCache,
	}

	recSol := recur(machine, int64(0), len(program)-1)
	if recSol == nil {
		//fmt.Printf("recSol: %v\n", recSol)
		panic("no solution")
	} else {
		//fmt.Printf("recSol: %d\n", *recSol)
	}
	return *recSol

	//minVal := int64(1 << (3 * (len(machine.program) - 1)))
	//maxVal := int64((1 << (3 * len(machine.program))) - 1)
	//minOut := machine.calcOut(minVal)
	//maxOut := machine.calcOut(maxVal)
	//_ = minOut
	//_ = maxOut
	//
	//sol := int64(bitsToInt([]int{3, 0, 7, 4}))
	////sol := int64(bitsToInt([]int{1, 0, 0, 0}))
	//sol = sol << ((len(program) - 4) * 3)
	//sols := map[int][]int64{}
	//chunks := len(program) / 4
	//for i := 1; i < chunks; i++ {
	//	for b := 0; b <= 0b111111111111; b++ {
	//		val := int64(b << ((len(program) - chunks*i - chunks) * 3))
	//
	//		val = (sol) | val
	//		out := machine.calcOut(val)
	//		slStart := len(program) - chunks - chunks*i
	//		//fmt.Printf("AAA: f(%v) = %v \n", intsToBits(int(val)), out)
	//		if len(out) != len(program) {
	//			continue
	//		}
	//
	//		if slices.Equal(out[slStart:], program[slStart:]) {
	//			p := val >> ((len(program) - chunks*i - chunks) * 3)
	//			sols[i] = append(sols[i], p)
	//			//fmt.Printf("FOUND(%d): f(%v) = %v [%d]\n", i, intsToBits(int(val)), out, intsToBits(int(p)))
	//			sol = val
	//			//break
	//		}
	//	}
	//}
	//
	//return sol
}

func intToBin[T int | int64](x T, len int) string {
	return fmt.Sprintf("%0*b", len, x)
}

func bruteForce(regA int, regB int, regC int, program []int, minVal int64, maxVal int64) int64 {
	workers := runtime.NumCPU()
	//space := (maxVal - minVal) / workers
	var wg sync.WaitGroup
	jobs := make(chan int64, workers*2)
	result := make(chan int64, workers)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(jobs <-chan int64, result chan<- int64, ctx context.Context, cancel context.CancelFunc) {
			defer wg.Done()
			mm := &Machine{
				regA:         int64(regA),
				regB:         int64(regB),
				regC:         int64(regC),
				ip:           0,
				program:      program,
				output:       []int{},
				targetOutput: program,
			}
			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-jobs:
					if !ok {
						return
					}
					mm.reset(val)
					mmOut := mm.calcOut(val)
					if slices.Equal(mmOut, mm.program) {
						fmt.Printf("found! %d => out: %v\n", val, mmOut)
						result <- val
						cancel()
						return
					}

				}
			}
		}(jobs, result, ctx, cancel)
	}

	var processed int64 = 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := minVal; i < maxVal+1; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("canceled %d\n", i)
				return
			default:
				jobs <- int64(i)
				atomic.AddInt64(&processed, 1)
			}

		}
		_, ok := <-jobs
		if ok {
			close(jobs)
		}
	}()

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		start := time.Now()
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				now := time.Now()
				elapsed := now.Sub(start)
				processedJobs := atomic.LoadInt64(&processed)
				remainingJobs := maxVal - minVal - processedJobs
				// elapsed/processed = x/remaining
				// x = remaining * elapsed / processed
				eta := elapsed * time.Duration(float64(remainingJobs)/float64(processedJobs))
				fmt.Printf("processed: %d/%d | remaining: %d (%.2f%%), elapsed: %s ; ETA: %s\n",
					processedJobs, maxVal-minVal, remainingJobs, 100*float64(processedJobs)/float64(maxVal-minVal), elapsed, eta)
			}
		}
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	results := []int64{}
	for res := range result {
		results = append(results, res)
	}
	fmt.Printf("results: %d\n", len(results))
	minRes := slices.Min(results)

	wg.Wait()

	return minRes
}

func intsToBits(x int) []int {
	if x == 0 {
		return []int{0}
	}

	var res []int
	for x > 0 {
		res = append(res, x%8)
		x = x / 8
	}
	slices.Reverse(res)
	return res

}

func bitsToInt(bits []int) int {
	result := 0
	for i := 0; i < len(bits); i++ {
		result = result << 3
		result += bits[i]
	}
	return result
}

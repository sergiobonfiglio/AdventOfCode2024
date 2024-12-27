package main

import (
	"strconv"
	"strings"
)

func part1(input string) any {

	secrets := parseInput(input)
	sum := 0
	for _, secret := range secrets {
		curr := secret
		for i := 0; i < 2000; i++ {
			curr = nextSecret(curr)
		}
		sum += curr
	}
	return sum
}

func part2(input string) any {
	secrets := parseInput(input)

	glMap := map[[4]int8]int{}

	maxBananas := 0
	for _, secret := range secrets {
		curr := nextSecret(secret)
		prev := lastDigit(secret)
		currMap := map[[4]int8]int{}

		changes := NewChanges()
		changes.add(lastDigit(curr) - prev)

		prev = lastDigit(curr)
		for i := 1; i < 2000; i++ {
			curr = nextSecret(curr)
			cld := lastDigit(curr)
			change := cld - prev
			changes.add(change)
			prev = cld
			if changes.full {
				key := changes.get()
				if _, ok := currMap[key]; !ok {
					currMap[key] = int(cld)
					glMap[key] += int(cld)
					if glMap[key] > maxBananas {
						maxBananas = glMap[key]
					}
				}
			}
		}
	}
	return maxBananas
}

func parseInput(input string) []int {
	var secrets []int
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		atoi, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		secrets = append(secrets, atoi)
	}
	return secrets
}

func nextSecret(currSecret int) int {

	next := mix(currSecret, currSecret*64)
	next = prune(next)

	next = mix(next, next/32)
	next = prune(next)

	next = mix(next, next*2048)
	next = prune(next)

	return next
}

func prune(next int) int {
	return next % 16777216
}

func mix(secret int, next int) int {
	return secret ^ next
}

func lastDigit(x int) int8 {
	return int8(x % 10)
}

type Changes struct {
	changes [4]int8
	full    bool
	nextI   int
}

func NewChanges() *Changes {
	return &Changes{
		nextI:   0,
		changes: [4]int8{},
	}
}

func (c *Changes) add(change int8) {
	c.changes[c.nextI] = change
	c.nextI = (c.nextI + 1) % len(c.changes)
	if c.nextI == 0 {
		c.full = true
	}
}

func (c *Changes) get() [4]int8 {
	out := [4]int8{}
	for i := 0; i < len(c.changes); i++ {
		out[i] = c.changes[(c.nextI+i)%len(c.changes)]
	}
	return out
}

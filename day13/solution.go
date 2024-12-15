package main

import (
	"regexp"
	"strconv"
	"strings"
)

func part1(input string) any {

	tot := 0
	xa, xb, ya, yb, px, py := 0, 0, 0, 0, 0, 0
	for i, line := range strings.Split(input, "\n") {

		if strings.HasPrefix(line, "Button A:") {
			xa, ya = getButton(line)
		} else if strings.HasPrefix(line, "Button B:") {
			xb, yb = getButton(line)
		} else if strings.HasPrefix(line, "Prize:") {
			px, py = getPrize(line)

			bb := (xa*py - px*ya) / (yb*xa - xb*ya)
			bbF := float64(xa*py-px*ya) / float64(yb*xa-xb*ya)
			if float64(bb) != bbF {
				//fmt.Printf("not solvable bbf %f\n", bbF)
				continue
			}
			ba := (px - bb*xb) / (xa)
			baF := float64(px-bb*xb) / float64(xa)

			if float64(ba) != baF {
				//fmt.Printf("not solvable baf %f\n", baF)
				continue
			}
			tokens := ba*3 + bb
			tot += tokens

			//fmt.Printf("Button A: %d, Button B: %d\n", ba, bb)
		}

		if (i+1)%4 == 0 {
			continue
		}

	}

	return tot
}

func getButton(str string) (int, int) {

	regex := regexp.MustCompile(`.*X\+([0-9]+),\s+Y\+([0-9]+)`)
	submatch := regex.FindAllStringSubmatch(str, -1)
	xstr, ystr := submatch[0][1], submatch[0][2]

	x, err := strconv.Atoi(xstr)
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(ystr)
	if err != nil {
		panic(err)
	}

	return x, y
}

func getPrize(str string) (int, int) {

	regex := regexp.MustCompile(`.*X=([0-9]+),\s+Y=([0-9]+)`)
	submatch := regex.FindAllStringSubmatch(str, -1)
	xstr, ystr := submatch[0][1], submatch[0][2]

	x, err := strconv.Atoi(xstr)
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(ystr)
	if err != nil {
		panic(err)
	}

	return x, y
}

func part2(input string) any {

	offset := 10000000000000
	tot := 0
	xa, xb, ya, yb, px, py := 0, 0, 0, 0, 0, 0
	for i, line := range strings.Split(input, "\n") {

		if strings.HasPrefix(line, "Button A:") {
			xa, ya = getButton(line)
		} else if strings.HasPrefix(line, "Button B:") {
			xb, yb = getButton(line)
		} else if strings.HasPrefix(line, "Prize:") {
			px, py = getPrize(line)
			px += offset
			py += offset

			bb := (xa*py - px*ya) / (yb*xa - xb*ya)
			bbF := float64(xa*py-px*ya) / float64(yb*xa-xb*ya)
			if float64(bb) != bbF {
				//fmt.Printf("not solvable bbf %f\n", bbF)
				continue
			}
			ba := (px - bb*xb) / (xa)
			baF := float64(px-bb*xb) / float64(xa)

			if float64(ba) != baF {
				//fmt.Printf("not solvable baf %f\n", baF)
				continue
			}
			tokens := ba*3 + bb
			tot += tokens

			//fmt.Printf("Button A: %d, Button B: %d\n", ba, bb)
		}

		if (i+1)%4 == 0 {
			continue
		}

	}

	return tot
}

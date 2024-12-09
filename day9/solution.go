package main

import (
	"AdventOfCode2024/utils"
	"strings"
)

func part1(input string) any {
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		diskMap := utils.ToIntArray(line, "")

		if len(diskMap)%2 != 0 {
			//add 0 free space after last file
			diskMap = append(diskMap, 0)
		}

		lastFile := len(diskMap) - 2
		maxId := lastFile / 2

		left := 0
		right := len(diskMap) - 2

		checksum := 0

		leftId := 0
		rightId := maxId
		diskIx := 0
		for left < right {

			//left
			fLen := diskMap[left]
			sLen := diskMap[left+1]
			checksum += chunkChecksum(diskIx, fLen, leftId)
			diskIx += fLen
			leftId++
			left += 2

			// fill space
			for sLen > 0 {
				rFileLen := diskMap[right]
				writtenBlocks := min(sLen, rFileLen)
				checksum += chunkChecksum(diskIx, writtenBlocks, rightId)
				diskIx += writtenBlocks
				sLen -= writtenBlocks
				if writtenBlocks == rFileLen {
					right -= 2
					rightId--
				} else {
					diskMap[right] -= writtenBlocks
				}
			}
		}
		if diskMap[right] > 0 {
			checksum += chunkChecksum(diskIx, diskMap[right], rightId)
		}

		return checksum
	}

	return nil
}

func chunkChecksum(offset int, fLen int, id int) int {
	cs := 0
	for i := 0; i < fLen; i++ {
		cs += (offset + i) * id
	}
	return cs
}

type Block struct {
	len    int
	fileId int
	offset int
}

func (b *Block) checksum() int {
	return chunkChecksum(b.offset, b.len, max(0, b.fileId))
}

func part2(input string) any {
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		diskMap := utils.ToIntArray(line, "")

		blocks := make([]*Block, len(diskMap))

		freeBlocks := []*Block{}

		currFileId := 0
		currOffset := 0
		for i, _ := range diskMap {

			fileId := currFileId
			isFreeSpace := i%2 != 0
			if isFreeSpace {
				fileId = -1
			}

			blocks[i] = &Block{
				len:    diskMap[i],
				fileId: fileId,
				offset: currOffset,
			}

			if isFreeSpace {
				freeBlocks = append(freeBlocks, blocks[i])
			} else {
				currFileId++
			}

			currOffset += diskMap[i]
		}

		// move files
		for i := len(blocks) - 1; i >= 0; i-- {
			if blocks[i].fileId == -1 {
				continue
			}

			fLen := blocks[i].len
			free, freeIx := findLeftmostFree(freeBlocks, fLen, blocks[i].offset)
			if free != nil {

				// move file
				blocks[i].offset = free.offset

				// resize space
				free.len = free.len - fLen
				free.offset += fLen

				if free.len == 0 {
					freeBlocks = append(freeBlocks[0:freeIx], freeBlocks[freeIx+1:]...)
				}
			}
		}

		checksum := 0
		for _, block := range blocks {
			checksum += block.checksum()
		}

		return checksum
	}

	return nil
}

func findLeftmostFree(freeBlocks []*Block, minSize int, maxOffset int) (*Block, int) {
	for i, block := range freeBlocks {
		if block.offset > maxOffset {
			return nil, -1
		}
		if block.len >= minSize {
			return block, i
		}
	}
	return nil, -1
}

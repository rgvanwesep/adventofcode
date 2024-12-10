package day9

import (
	"fmt"
	"log"
	"strconv"
)

type Block struct {
	empty  bool
	fileId int
}

type ExpandedDiskMap struct {
	blocks        []Block
	emptyIndices  []int
	lastFileIndex int
}

func NewExpandedDiskMap() *ExpandedDiskMap {
	return &ExpandedDiskMap{
		blocks:        []Block{},
		emptyIndices:  []int{},
		lastFileIndex: -1,
	}
}

func (m *ExpandedDiskMap) Extend(diskMapIndex int, diskMapValue rune) {
	var (
		empty         bool
		fileId        int
		emptyIndices  []int
		lastFileIndex int
	)
	nBlocks, err := strconv.Atoi(fmt.Sprintf("%c", diskMapValue))
	if err != nil {
		log.Panicf("Could not parse %q", diskMapValue)
	}
	if diskMapIndex%2 == 0 {
		empty = false
		fileId = diskMapIndex / 2
		emptyIndices = []int{}
		lastFileIndex = len(m.blocks) + nBlocks - 1
	} else {
		empty = true
		fileId = -1
		emptyIndices = make([]int, nBlocks)
		for i := 0; i < nBlocks; i++ {
			emptyIndices[i] = m.lastFileIndex + i + 1
		}
		lastFileIndex = m.lastFileIndex
	}
	for range nBlocks {
		m.blocks = append(m.blocks, Block{empty, fileId})
	}
	m.emptyIndices = append(m.emptyIndices, emptyIndices...)
	m.lastFileIndex = lastFileIndex
}

func (m *ExpandedDiskMap) Swap() bool {
	if m.lastFileIndex < m.emptyIndices[0] {
		return false
	}
	m.blocks[m.emptyIndices[0]], m.blocks[m.lastFileIndex] = m.blocks[m.lastFileIndex], m.blocks[m.emptyIndices[0]]
	m.emptyIndices = append(m.emptyIndices, m.lastFileIndex)
	m.emptyIndices = m.emptyIndices[1:]
	for m.blocks[m.lastFileIndex].empty {
		m.lastFileIndex--
	}
	return true
}

func (m *ExpandedDiskMap) CalcChecksum() int {
	sum := 0
	for i, b := range m.blocks {
		if b.empty {
			break
		}
		sum += i * b.fileId
	}
	return sum
}

func ExpandDiskMap(diskMap string) *ExpandedDiskMap {
	expanded := NewExpandedDiskMap()
	for i, r := range diskMap {
		expanded.Extend(i, r)
	}
	return expanded
}

func CalcChecksum(inputs []string) int {
	diskMap := inputs[0]
	expandedDiskMap := ExpandDiskMap(diskMap)
	for expandedDiskMap.Swap() {
	}
	return expandedDiskMap.CalcChecksum()
}

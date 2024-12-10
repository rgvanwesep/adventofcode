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

type File struct {
	start, end int
}

type FreeSpace struct {
	start, end int
}

type ExpandedDiskMap struct {
	blocks        []Block
	files         []File
	freeSpaces    []FreeSpace
	emptyIndices  []int
	lastFileIndex int
}

func NewExpandedDiskMap() *ExpandedDiskMap {
	return &ExpandedDiskMap{
		blocks:        []Block{},
		files:         []File{},
		freeSpaces:    []FreeSpace{},
		emptyIndices:  []int{},
		lastFileIndex: -1,
	}
}

func (m *ExpandedDiskMap) Extend(diskMapIndex int, diskMapValue rune) {
	var (
		empty         bool
		fileId        int
		files         []File
		freeSpaces    []FreeSpace
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
		fileStart := len(m.blocks)
		fileEnd := fileStart + nBlocks - 1
		files = []File{{fileStart, fileEnd}}
		freeSpaces = []FreeSpace{}
		emptyIndices = []int{}
		lastFileIndex = fileEnd
	} else {
		empty = true
		fileId = -1
		files = []File{}
		freeStart := m.lastFileIndex + 1
		freeEnd := freeStart + nBlocks - 1
		freeSpaces = []FreeSpace{{freeStart, freeEnd}}
		emptyIndices = make([]int, nBlocks)
		for i := 0; i < nBlocks; i++ {
			emptyIndices[i] = freeStart + i
		}
		lastFileIndex = m.lastFileIndex
	}
	for range nBlocks {
		m.blocks = append(m.blocks, Block{empty, fileId})
	}
	m.files = append(m.files, files...)
	m.freeSpaces = append(m.freeSpaces, freeSpaces...)
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

func (m *ExpandedDiskMap) UpdateBlocks() {
	blocks := make([]Block, len(m.blocks))
	for i := range blocks {
		blocks[i] = Block{empty: true, fileId: -1}
	}
	for i, f := range m.files {
		for j := f.start; j <= f.end; j++ {
			blocks[j] = Block{empty: false, fileId: i}
		}
	}
	m.blocks = blocks
}

func (m *ExpandedDiskMap) UpdateFreeSpaces() {
	var freeStart, freeEnd int
	inFile := true
	freeSpaces := []FreeSpace{}
	for i, b := range m.blocks {
		switch c := [2]bool{inFile, b.empty}; c {
		case [2]bool{true, true}:
			freeStart = i
			inFile = false
		case [2]bool{true, false}:
			continue
		case [2]bool{false, true}:
			continue
		case [2]bool{false, false}:
			freeEnd = i - 1
			inFile = true
			freeSpaces = append(freeSpaces, FreeSpace{freeStart, freeEnd})
		}
	}
	if !inFile {
		freeEnd = len(m.blocks) - 1
		freeSpaces = append(freeSpaces, FreeSpace{freeStart, freeEnd})
	}
	m.freeSpaces = freeSpaces
}

func (m *ExpandedDiskMap) SwapFile(id int) {
	file := m.files[id]
	fileSize := file.end - file.start + 1
	for _, freeSpace := range m.freeSpaces {
		if freeSpace.start > file.end {
			break
		}
		freeSpaceSize := freeSpace.end - freeSpace.start + 1
		if freeSpaceSize >= fileSize {
			m.files[id] = File{start: freeSpace.start, end: freeSpace.start + fileSize - 1}
			break
		}
	}
	m.UpdateBlocks()
	m.UpdateFreeSpaces()
}

func (m *ExpandedDiskMap) CalcChecksum() int {
	sum := 0
	for i, b := range m.blocks {
		if !b.empty {
			sum += i * b.fileId
		}
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

func CalcChecksumFileSwap(inputs []string) int {
	diskMap := inputs[0]
	expandedDiskMap := ExpandDiskMap(diskMap)
	for i := len(expandedDiskMap.files) - 1; i >= 0; i-- {
		expandedDiskMap.SwapFile(i)
	}
	return expandedDiskMap.CalcChecksum()
}

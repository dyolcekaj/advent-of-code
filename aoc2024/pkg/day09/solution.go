package day09

func PartOne(diskMap []int) int {
	checksum := 0
	minFileIdx := 0
	maxFileIdx := (len(diskMap) - 1) / 2
	idx := 0

	eIdx := len(diskMap) - 1
	if eIdx%2 == 1 {
		eIdx--
	}
	eVal := diskMap[eIdx]

	for i := 0; i < eIdx; i++ {
		size := diskMap[i]

		if i%2 == 0 {
			// file blocks
			for j := 0; j < size; j++ {
				checksum += minFileIdx * idx
				idx++
			}
			minFileIdx++
		} else {
			// empty space
			for j := 0; j < size; j++ {
				if eVal == 0 {
					maxFileIdx--
					eIdx -= 2

					if eIdx < i {
						break
					}

					eVal = diskMap[eIdx]
				}

				checksum += maxFileIdx * idx

				idx++
				eVal--
			}
		}
	}

	if eVal > 0 {
		for i := 0; i < eVal; i++ {
			checksum += maxFileIdx * idx
			idx++
		}
	}

	return checksum
}

func PartTwo(diskMap []int) int {
	checksum := 0
	minFileIdx := 0
	maxFileIdx := (len(diskMap) - 1) / 2
	idx := 0

	eIdx := len(diskMap) - 1
	if eIdx%2 == 1 {
		eIdx--
	}

	moved := make(map[int]struct{})

	for i := 0; i < len(diskMap); i++ {
		size := diskMap[i]
		if _, ok := moved[i]; ok {
			minFileIdx++
			idx += size

			continue
		}

		if i%2 == 0 {
			// file blocks
			for j := 0; j < size; j++ {
				checksum += minFileIdx * idx
				idx++
			}
			minFileIdx++
		} else {
			fi := maxFileIdx

			for e := eIdx; e > i && size > 0; e -= 2 {
				if _, ok := moved[e]; ok || diskMap[e] > size {
					fi--
					continue
				}

				toMove := diskMap[e]
				moved[e] = struct{}{}
				size -= toMove

				for j := 0; j < toMove; j++ {
					checksum += fi * idx
					idx++
				}

				fi--
			}

			if size > 0 {
				idx += size
			}
		}
	}

	return checksum
}

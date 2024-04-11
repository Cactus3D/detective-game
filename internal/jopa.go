package internal

func jopa(devCount int, devRel []int) string {
	relCount := 1
	changed := false

	for i := 0; ; i++ {
		if i == devCount {
			i = 0
		}
		if devRel[i] == 0 {
			continue
		}
		devRel[i] = devRel[i] - 1
		for j, v := range devRel {
			if j != i && v > 0 {
				devRel[j] = devRel[j] - 1
				changed = true
				relCount++
				break
			}
			changed = false
		}
		if relCount == devCount {
			return "Yes"
		}
		if i == devCount-1 && !changed {
			break
		}
	}

	if relCount == devCount {
		return "Yes"
	} else {
		return "No"
	}
}

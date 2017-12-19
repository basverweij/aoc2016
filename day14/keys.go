package main

const keyWindow = 1000

func genKeys(salt string, amount int) []*hash {
	var keys []*hash
	var hashes []*hash

	// generate the first keyWindow hashes
	for i := 0; i < keyWindow; i++ {
		h := genHash(salt, i)
		if h == nil {
			continue
		}

		hashes = append(hashes, h)
	}

	tripletIndex := 0
	quintetIndex := 1
	latestQuintet := 0
	for i := keyWindow; len(keys) < amount; i++ {
		// move tripletIndex to the next hash with a triplet
		for ; tripletIndex < len(hashes) &&
			hashes[tripletIndex].firstTriplet == 0; tripletIndex++ {
		}

		// move quintetIndex to the next hash with a quintet
		for ; quintetIndex < len(hashes) &&
			len(hashes[quintetIndex].quintets) == 0; quintetIndex++ {
		}

		// gen hash for i
		h := genHash(salt, i)
		if h == nil {
			continue
		}

		hashes = append(hashes, h)

		if len(h.quintets) > 0 {
			latestQuintet = h.index
		}

		if hashes[quintetIndex].index+keyWindow >= latestQuintet {
			// generate more hashes to prevent missing additional quintets in the key window
			continue
		}

		if hashes[quintetIndex].index <= hashes[tripletIndex].index+keyWindow {
			// quintet within key key window
			if _, found := hashes[quintetIndex].quintets[hashes[tripletIndex].firstTriplet]; found {
				// found valid key
				keys = append(keys, hashes[tripletIndex])

				// trigger move to next triplet
				tripletIndex++

				// reset quintet index
				quintetIndex = tripletIndex + 1
			} else {
				// trigger move to next quintet
				quintetIndex++
			}
		} else {
			// trigger move to next triplet
			tripletIndex++

			// reset quintet index
			quintetIndex = tripletIndex + 1
		}
	}

	return keys
}

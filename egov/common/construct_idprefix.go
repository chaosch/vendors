package common



func ConstructNextPrefix(pid, cid, did int64) int64 {
	Idprefix := int64(0)
	if cid == 0 {
		Idprefix = int64(((pid+1)<<14 + (0)<<7 + 0) << 32)
	} else {
		if did == 0 {
			Idprefix = int64((pid<<14 + (cid+1)<<7 + (0)) << 32)
		} else {
			Idprefix = int64((pid<<14 + cid<<7 + (did + 1)) << 32)
		}
	}

	return Idprefix
}

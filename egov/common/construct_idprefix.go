package common

import "math"

func Construct(pid, cid, did int64) int64 {
	Idprefix := int64(0)
	//	Idprefix = int64(pid+cid+did) << 32
	Idprefix = int64(pid*(int64(math.Pow10(14))) + cid*int64(math.Pow10(12)) + (did * int64(math.Pow10(10))))

	return Idprefix
}

func ConstructNextPrefix(pid, cid, did int64) int64 {
	Idprefix := int64(0)
	if pid == 0 {
		Idprefix = int64(1*(int64(math.Pow10(14))) + (1)*int64(math.Pow10(12)) + int64(math.Pow10(10)))
	} else {
		if cid == 0 {
			Idprefix = int64(pid*(int64(math.Pow10(14))) + (cid+1)*int64(math.Pow10(12)) + int64(math.Pow10(10)))
		} else {
			Idprefix = int64(pid*(int64(math.Pow10(14))) + cid*int64(math.Pow10(12)) + ((did + 1) * int64(math.Pow10(10))))
		}
	}

	return Idprefix
}

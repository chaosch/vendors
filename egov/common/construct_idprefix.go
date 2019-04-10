package common

func Construct(pid, cid, did int64) int64 {
	Idprefix := int64(0)
	//	Idprefix = int64(pid+cid+did) << 32
	//Idprefix = int64(pid*(int64(math.Pow10(14))) + cid*int64(math.Pow10(12)) + (did * int64(math.Pow10(10))))

	//Id.Provice_id=pid//
	//Id.City_id=cid//
	//Id.District_id=did//
	//Id.Idprefix=0
	//Id.Provice_id=Id.Provice_id<<15
	//Id.City_id=Id.City_id<<8
	if pid >= 80 {
		pid = pid - 55
	}
	Idprefix = int64(pid<<15+cid<<8+did) << 32

	return Idprefix
}

func ConstructNextPrefix(pid, cid, did int64) int64 {
	if pid >= 80 {
		pid = pid - 55
	}
	Idprefix := int64(0)
	if pid == 0 {
		Idprefix = int64((1<<15 + 1<<8 + 1) << 32)
	} else {
		if cid == 0 {
			Idprefix = int64((pid<<15 + (cid+1)<<8 + 1) << 32)
		} else {
			Idprefix = int64((pid<<15 + cid<<8 + (did + 1)) << 32)
		}
	}

	return Idprefix
}

package cmd

import (
	"fmt"
	"syscall"
)

//GetRLimit get the soft limit and hard limit for the number of file
//descriptors available.
func GetRLimit() (softLimit uint64, hardLimit uint64, err error) {
	var rLimit syscall.Rlimit
	var syscallErr error

	softLimit = uint64(0)
	hardLimit = uint64(0)
	syscallErr = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if syscallErr != nil {
		err = fmt.Errorf("ERROR getting rlimit: `%s`\n", syscallErr.Error())
		return
	}
	//err == nil
	softLimit = rLimit.Cur
	hardLimit = rLimit.Max

	return
}

//SetSoftRLimit changes the soft limit number of file descriptors to
//the hard limit value.
func SetSoftRLimit(softLimit uint64) (err error) {
	var (
		hardLimit  uint64
		rLimit     syscall.Rlimit
		syscallErr error
	)
	if _, hardLimit, err = GetRLimit(); err != nil {
		return
	}
	rLimit.Max = hardLimit
	//  Below
	//=========
	//rLimit.Cur is set to hardlimit minus one, because:
	//	RLIMIT_NOFILE
	//      Specifies a value one greater than the maximum file descriptor
	//		number that can  be opened by this process.
	rLimit.Cur = hardLimit - 1
	syscallErr = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if syscallErr != nil {
		err = fmt.Errorf("%s '%d': `%s`\n",
			"ERROR setting soft limit of available file descriptors to ",
			softLimit,
			syscallErr.Error())
		return
	}
	return nil
}

//*
//* Zymkey - sample code to drive zymlib the Zymkey library
//*
//* Written by Maurice Bizzarri, Bizzarri Software
//* January, 2019
//* update March, 2019
//*
//******************************************************************
package main

// #cgo LDFLAGS: -L/usr/lib -lzk_app_utils
// #cgo CFLAGS: -I/usr/include/zymkey
// #include "zk_app_utils.h"
import "C"

import "fmt"
import "os"
import "time"
import "unsafe"
import "log"
import "flag"

var zkptr _Ctype_zkCTX
var flags C.bool

func Check(ret int, mess string) {
	if ret != 0 {
		fmt.Printf("\nError: %s\n", mess)
		os.Exit(0)
	}
}

func main() {

	var retcode int
	boolPtr := flag.Bool("debug", false, "Debug flag")
	flag.Parse()
	debug := *boolPtr
	if !debug {
		fmt.Printf("Zymkey test bed\n")
	}
	retcode = int(C.zkOpen(&zkptr))
	Check(retcode, "Error opening zk")
	defer C.zkClose(zkptr)
	//*
	//* get the public key
	//*
	//*
	var pemkey *_Ctype_uchar
	var pksize _Ctype_int
	retcode = int(C.zkGetECDSAPubKey(zkptr, &pemkey, &pksize, 0))
	pkey := C.GoBytes(unsafe.Pointer(pemkey), pksize)
	Check(retcode, "Retrieving public key")
        if !debug {
		fmt.Printf("pkey: %v\n", pkey)
	}
	//*
	//* turn led on and then off
	//*
	retcode = int(C.zkLEDFlash(zkptr, 500, 1000, 100))
	Check(retcode, "Error turning LED on")
	//*
	//* sleep for 1 second
	//*
	time.Sleep(1 * time.Second)
	//*
	//* set tap sensitivity
	//*
	retcode = int(C.zkSetTapSensitivity(zkptr, 3, 90.0))
	Check(retcode, "error setting tap sensitivity")
	if !debug {
		fmt.Printf("Now waiting for tap!\n")
	}
	//*
	//* wait for tap
	//*
	retcode = int(C.zkWaitForTap(zkptr, 5000))
        if !debug {
	if retcode == 0 {
		fmt.Printf("Tap Detected\n")
	} else {
		fmt.Printf("No tap detected\n")
	}
	}
	//*
	//* lock - lock a data file.  Encrypt with key unique to this
	//* zymkey.  Useful for encrypting a key file.
	//*
	//* first create a random 8 byte string in a file
	//*
	fromf := C.CString("test.in")
	var len C.int
	len = 8
	C.zkCreateRandDataFile(zkptr, fromf, len)
	var rand *_Ctype_uchar
	retcode = int(C.zkGetRandBytes(zkptr, &rand, len))
	if retcode < 0 {
		fmt.Printf("Error creating random bytes\n")
	}
	var randi []byte
	randi = C.GoBytes(unsafe.Pointer(rand), len)
        if !debug {
		fmt.Printf("Random bytes: %v\n", randi[:8])
	}
	tof := C.CString("test.out")
	newto := C.CString("testagain.out")
	flags = false
	retcode = int(C.zkLockDataF2F(zkptr, fromf, tof, flags))
	Check(retcode, "Error encoding file")
	retcode = int(C.zkUnlockDataF2F(zkptr, tof, newto, flags))
	Check(retcode, "Error decoding file")
	//*
	//* read from file
	//
	file, err := os.Open("testagain.out")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data := make([]byte, 48)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
        if !debug {
		fmt.Printf("\nRandom bytes from file: %v\n\n", data[:count])
	}
	//*
	//* get real time from onboard RTC
	//*
	var timer C.uint
	var ctime time.Time
	retcode = int(C.zkGetTime(zkptr, &timer, flags))

	ctime = time.Unix(int64(timer), 0)
        
	fmt.Printf("Date and Time: %s\r", ctime.Format("Mon Jan _2 15:04:05 2006"))
	if !debug {
		fmt.Printf("\n\n")
	}
	//*
	//* close and exit
	//*
	retcode = int(C.zkClose(zkptr))
	Check(retcode, "Error closing zk")
	os.Exit(1)
}

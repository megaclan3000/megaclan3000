package main

import (
	log "github.com/sirupsen/logrus"
	"strconv"
)

func apiHandler(vars map[string]string) []byte {

	var byt []byte

	log.Debug("API request to: ", vars["endpoint"])

	switch vars["endpoint"] {

	case "playerinfo":
		strconv.ParseUint(vars["id"], 10, 64)
		byt = []byte(`Here is info about player: ` + vars["steamid"])
	case "test":
		byt = []byte(`
[
   [
      "2020-08-14T14:58:59.279Z",
      2.6548573028712052
   ],
   [
      "2020-08-14T14:59:00.279Z",
      9.660990312675485
   ],
   [
      "2020-08-14T14:59:01.279Z",
      9.867134407785812
   ],
   [
      "2020-08-14T14:59:02.279Z",
      9.693173367060712
   ],
   [
      "2020-08-14T14:59:03.279Z",
      5.795384070587739
   ],
   [
      "2020-08-14T14:59:04.279Z",
      6.54514491707552
   ],
   [
      "2020-08-14T14:59:05.279Z",
      1.610337192023139
   ],
   [
      "2020-08-14T14:59:06.279Z",
      4.059514075791508
   ],
   [
      "2020-08-14T14:59:07.279Z",
      6.258157166881933
   ],
   [
      "2020-08-14T14:59:08.279Z",
      4.678033662200467
   ]
]`)

	}

	return byt

}

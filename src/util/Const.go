package util

import (
	"os"
)

// 상수 정의

// var CloudConnectionUrl = os.Getenv("SPIDER_URL")
// var TumbleUrl = os.Getenv("TUMBLE_URL")
// var CloudConnectionUrl = os.Getenv("SPIDER_URL")
// var NameSpaceUrl = os.Getenv("TUMBLE_URL")

var SPIDER = os.Getenv("SPIDER_URL")
var TUMBLEBUG = os.Getenv("TUMBLE_URL")

var HTTP_CALL_SUCCESS = 200
var HTTP_POST_SUCCESS = 201

// VM의 상태.  (기타 상태는 UNDEFINED + ETC)
var VM_STATUS_RUNNING = "Running"
var VM_STATUS_RESUMING = "Resuming"
var VM_STATUS_INCLUDE = "include"
var VM_STATUS_SUSPENDED = "Suspended"
var VM_STATUS_TERMINATED = "Terminated"
var VM_STATUS_UNDEFINED = "statusUndefined"
var VM_STATUS_PARTIAL = "partial"
var VM_STATUS_ETC = "etc"

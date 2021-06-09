package service

import (
	// "encoding/base64"
	"fmt"
	"log"

	// "log"
	// "io"
	// "net/http"

	echosession "github.com/go-session/echo-session"
	"github.com/labstack/echo"

	"github.com/cloud-barista/cb-webtool/src/model"
	"github.com/cloud-barista/cb-webtool/src/model/tumblebug"
	util "github.com/cloud-barista/cb-webtool/src/util"
)

// 로그인할 때, NameSpace 저장(Create, Delete, Update) 외에는 이 funtion 사용
// 없으면 tb 조회
func GetStoredNameSpaceList(c echo.Context) ([]tumblebug.NameSpaceInfo, model.WebStatus) {
	fmt.Println("====== GET STORED NAME SPACE ========")
	nameSpaceList := []tumblebug.NameSpaceInfo{}
	nameSpaceErr := model.WebStatus{}
	store := echosession.FromContext(c)

	storedNameSpaceList, isExist := store.Get(util.STORE_NAMESPACELIST)
	if !isExist { // 존재하지 않으면 TB 조회
		nameSpaceList, nameSpaceErr = GetNameSpaceList()
		setError := SetStoreNameSpaceList(c, nameSpaceList)
		if setError != nil {
			log.Println("Set Namespace failed")
		}
	} else {
		log.Println(storedNameSpaceList)
		nameSpaceList = storedNameSpaceList.([]tumblebug.NameSpaceInfo)
		nameSpaceErr.StatusCode = 200
	}
	return nameSpaceList, nameSpaceErr
}

func SetStoreNameSpaceList(c echo.Context, nameSpaceList []tumblebug.NameSpaceInfo) error {
	fmt.Println("====== SET NAME SPACE ========")
	store := echosession.FromContext(c)
	store.Set(util.STORE_NAMESPACELIST, nameSpaceList)
	err := store.Save()
	return err
}

// move to NameSpaceController.go
// func GetNameSpace(c echo.Context) error {
// 	fmt.Println("====== GET NAME SPACE ========")
// 	store := echosession.FromContext(c)

// 	getInfo, ok := store.Get("namespace")
// 	if !ok {
// 		return c.JSON(http.StatusNotAcceptable, map[string]string{
// 			"message": "Not Exist",
// 		})
// 	}
// 	nsId := getInfo.(string)

// 	res := map[string]string{
// 		"message": "success",
// 		"nsID":    nsId,
// 	}

// 	return c.JSON(http.StatusOK, res)
// }

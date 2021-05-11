html template controller



return 시 구현해야 하는 parameter : message, status
model.WebStatus{StatusCode: 500, Message: err.Error()}

if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
    log.Println(" respStatus  ", respStatus)
    return c.JSON(http.StatusBadRequest, map[string]interface{}{
        "message": respStatus.Message,
        "status":  respStatus.StatusCode,
    })
}

-> 최종 변경 : return시 error로 send, error code 도 return받는 respStatus
if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {

    return c.JSON(respStatus.StatusCode, map[string]interface{}{
        "error":  respStatus.Message,
        "status": respStatus.StatusCode,
    })
}
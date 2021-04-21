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
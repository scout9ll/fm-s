package api

import (
	"io/ioutil"
	"net/http"
)

func GetDyRoomInfo(roomId string) string {
	response, err := http.Get("http://open.douyucdn.cn/api/RoomApi/room/" + roomId)
	if err != nil {
		return ""
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

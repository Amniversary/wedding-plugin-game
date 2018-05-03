package server

import (
	"net/http"
	"github.com/Amniversary/wedding-plugin-game/config"
	"encoding/json"
	"log"
	"github.com/Amniversary/wedding-plugin-game/config/mysql"
	"github.com/Amniversary/wedding-plugin-game/components"
)
// todo: 初始化游戏
func (s *Server) NewPlayGame(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code:config.RSP_ERRPR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.NewPlayGame{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("newPlayGame json decode err: [%v]", err)
		Response.Msg = config.RSP_ERROR_MSG
		return
	}
	if req.WeddingId == 0 || req.Seconds == 0{
		log.Printf("params can't be empty: [weddingId: %v, seconds: %v]", req.WeddingId, req.Seconds)
		Response.Msg = "params can't be empty"
		return
	}
	_, Ok := mysql.GetNowPlaySaimaGame(req.WeddingId)
	if Ok {
		Response.Msg = "游戏正在进行中"
		return
	}
	game := &mysql.SaimaGame{WeddingId:req.WeddingId, Seconds:req.Seconds, Status:1}
	if err := mysql.CreateSaimaGame(game); err != nil {
		log.Printf("create saima game err: [%v]", err)
		Response.Msg = config.RSP_ERROR_MSG
		return
	}

	Response.Code = config.RSP_SUCCESS
}

func (s *Server) Test(w http.ResponseWriter, r *http.Request) {
	userId := []int64{84}
	components.SendHLBUserBroadcast(userId, "string json", "HLBUser")
}
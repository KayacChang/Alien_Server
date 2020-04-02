package game

import (
	"encoding/json"

	"server/game/gamerule"

	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/foundation"
	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/foundation/fileload"
	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/iserver"
	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/restfult"
	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/socket"
)

// NewGameServer ...
func NewGameServer(jsStr string) {

	config := foundation.StringToJSON(jsStr)
	baseSetting := iserver.NewSetting()
	baseSetting.SetData(config)

	gamejsStr := fileload.Load("./file/gameconfig.json")
	var gameRule = &gamerule.Rule{}
	if err := json.Unmarshal([]byte(gamejsStr), &gameRule); err != nil {
		panic(err)
	}

	var gameserver = iserver.NewService()
	var game = &Game{
		IGameRule: gameRule,
		Server:    gameserver,
		// ProtocolMap: protocol.NewProtocolMap(),
	}
	gameserver.Restfult = restfult.NewRestfultService()
	gameserver.Socket = socket.NewSocket()
	gameserver.IGame = game

	// start Server
	gameserver.Launch(baseSetting)

	// start DB service
	setting := gameserver.Setting.DBSetting()
	gameserver.LaunchDB("gameDB", setting)
	gameserver.LaunchDB("logDB", setting)
	gameserver.LaunchDB("payDB", setting)

	// start restful service
	go gameserver.LaunchRestfult(game.RESTfulURLs())
	go gameserver.LaunchSocket(game.SocketURLs())

	<-gameserver.ShotDown
}

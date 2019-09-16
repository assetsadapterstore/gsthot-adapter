package openwtester

import (
	"github.com/assetsadapterstore/gsthot-adapter/gsthot"
	"github.com/blocktree/eosio-adapter/eosio"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	cache := eosio.NewCacheManager()

	openw.RegAssets(gsthot.Symbol, gsthot.NewWalletManager(&cache))
}

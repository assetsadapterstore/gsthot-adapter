module github.com/assetsadapterstore/gsthot-adapter

go 1.12

require (
	github.com/astaxie/beego v1.12.0
	github.com/blocktree/eosio-adapter v1.1.0
	github.com/blocktree/go-owcdrivers v1.1.22
	github.com/blocktree/openwallet v1.5.5
	github.com/eoscanada/eos-go v0.8.16
)

replace github.com/eoscanada/eos-go => github.com/blocktree/eos-go v0.8.13-blocktree

module github.com/clucle/golang-nalang/src/console/snake-game/main

go 1.20

replace github.com/clucle/golang-nalang/src/console/snake-game/snakegame => ../snakegame

require (
	github.com/clucle/golang-nalang/src/console/snake-game/snakegame v0.0.0-00010101000000-000000000000
	github.com/eiannone/keyboard v0.0.0-20220611211555-0d226195f203
)

require golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect

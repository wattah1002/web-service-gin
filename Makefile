# wip: まだ動かない
BINARY_NAME=app

db_conn = mysql -ututorial-go -pudPHP_r68.LQ
db_name = tutorial_go
seed_dir = ./db/seed/
album_table = albums

.PHONY: seed
seed:
	$(db_conn)
	use $(db_name)
	source $(seed_dir)$(album_table).sql

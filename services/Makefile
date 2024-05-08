gen_all: gen_saving gen_config gen_database gen_user gen_valorant

gen_saving:
	protoc --go_out=. --twirp_out=. rpc/saving/service.proto

gen_config:
	protoc --go_out=. --twirp_out=. rpc/config/service.proto

gen_database:
	protoc --go_out=. --twirp_out=. rpc/database/service.proto

gen_user:
	protoc --go_out=. --twirp_out=. rpc/user/service.proto

gen_valorant:
	protoc --go_out=. --twirp_out=. rpc/valorant/service.proto

.PHONY: local-db clean-local-env

local-db:
	eval "docker-compose -f docker-compose.yaml down"
	eval "docker-compose -f docker-compose.yaml up -d"

clean-local-db:
	eval "docker-compose -f docker-compose.yaml down"

	docker-compose down
	# shellcheck disable=SC2046
	docker rmi $(docker images -a -q)
	docker-compose up --force-recreate
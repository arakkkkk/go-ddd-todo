run:
	docker compose up -d

testadd:
	curl --request POST localhost:8080/api/v1/todo/create \
		--header 'content-type: application/json' \
		--data '{"title": "test todo", "priority": 5}'

testget:
	curl --request GET localhost:8080/api/v1/todo/list


restart:
	docker compose restart

build:
	docker compose build

down:
	docker compose down

logs:
	docker compose logs -f

downall:
	docker-compose down --rmi all --volumes --remove-orphans

watch:
	watch docker compose ps

uml:
	goplantuml -output doc/uml.md -ignore ent -recursive ./


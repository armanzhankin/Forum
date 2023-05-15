.SILENT:
dbrun: buildbackend
	docker run -dp 5050:5050 --name app back
buildbackend:
	docker build -t back .
dbstop: dbstop
	docker stop app
dbdelete:
	docker rm app
dbclear:
	docker rmi back
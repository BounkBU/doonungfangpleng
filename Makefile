server:
	go run main.go

deploy:
	git push heroku main

.PHONY: server deploy

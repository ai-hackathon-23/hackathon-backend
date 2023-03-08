.PHONY: start
start:
	@echo ローカルサーバーを起動中...
	docker-compose build --no-cache && docker-compose up
	@echo 起動しました.

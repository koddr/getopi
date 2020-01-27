.PHONY: front-start

front-start:
	@cd ./frontend && npm start
	@echo "[OK] Preact app is running!"

front-lint:
	@cd ./frontend && npm run lint
	@echo "[OK] Preact app was linted!"

front-build-prod:
	@cd ./frontend && npm run build
	@echo "[OK] Preact app was builded!"

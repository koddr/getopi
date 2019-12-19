.PHONY: front-start

front-start:
	cd ./frontend \
	&& npm start
	@echo "[✅] Preact app is running!"

front-lint:
	cd ./frontend \
	&& npm run lint
	@echo "[✅] Preact app was linted!"

front-build-prod:
	cd ./frontend \
	&& npm run build
	@echo "[✅] Preact app was builded!"

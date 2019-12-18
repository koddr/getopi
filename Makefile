.PHONY: front-start

front-start:
	cd ./frontend \
	&& npm start
	@echo "[✅] Preact app is running!"

front-build-prod:
	cd ./frontend \
	&& npm run build
	@echo "[✅] Preact app was build!"

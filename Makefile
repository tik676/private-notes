run:
	docker compose up --build

stop:
	docker compose down

logs:
	docker compose logs -f backend

shell:
	docker compose exec backend sh

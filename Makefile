# docker-compose aliases
.PHONY: compose-rs
compose-rs:
	make compose-rm
	make compose-up

.PHONY: compose-up
compose-up:
	docker compose -p tmp up -d

.PHONY: compose-down
compose-down:
	docker compose -p tmp down

.PHONY: compose-rm
compose-rm:
	docker-compose -p tmp rm -fvs

.PHONY: run-full
run-full:
	make compose-up

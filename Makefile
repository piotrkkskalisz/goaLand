# Simple Makefile for a Go project

# Create .env file from .env.example file
env:
	@cp --update=none .env.example .env

# Create DB container
docker-up:
	@docker compose up --build
		
# Shutdown DB container
docker-down:
	@docker compose down

#go to backend and run this command
build:
	@$(MAKE) -C backend build

run:
	@$(MAKE) -C backend run

test:
	@$(MAKE) -C backend test

watch:
	@$(MAKE) -C backend watch
	
.PHONY: env docker-up docker-down build run test watch
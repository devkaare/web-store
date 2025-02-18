MAIN_FILE = cmd/api/main.go

# Build the application
all: build test
templ-install:
	@if ! command -v templ > /dev/null; then \
		read -p "Go's 'templ' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/a-h/templ/cmd/templ@latest; \
			if [ ! -x "$$(command -v templ)" ]; then \
				echo "templ installation failed. Exiting..."; \
				exit 1; \
			fi; \
		else \
			echo "You chose not to install templ. Exiting..."; \
			exit 1; \
		fi; \
	fi

tailwind-install:
	@if [ ! -f tailwindcss ]; then curl -sL https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 -o tailwindcss; fi
	
	@chmod +x tailwindcss

build: tailwind-install templ-install
	@echo "Building..."
	@templ generate
	@./tailwindcss -i view/assets/css/input.css -o cmd/web/assets/css/output.css
	@go build -o main $(MAIN_FILE)

# Run the application
run:
	@go run $(MAIN_FILE)
	#
# Create DB container
docker-build:
	@docker compose up --build

# Run DB container
docker-run:
	@docker compose up

# Shutdown DB container
docker-down:
	@docker-compose down

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the junk
clean:
	@echo "Cleaning..."
	@rm -f main
	@go mod tidy

	@rm -rf views/*_templ.go

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

.PHONY: all build run test clean watch docker-build docker-run docker-down itest templ-install tailwind-install

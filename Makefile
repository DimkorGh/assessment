# -------- APP COMMANDS ------------

docker.app.start:
	make dockerProductionBuild
	docker run \
			--publish 8080:8080 \
			production-build

# -------- TESTS COMMANDS ------------

docker.test.unit:
	make dockerTestBuild
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			test-build go test -tags musl -short -count=1 ./...

docker.test.all:
	make dockerTestBuild
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			test-build go test -tags musl -count=1 ./...

docker.test.all.coverage.withView:
	make dockerTestBuild
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			test-build go test -tags musl -count=1 -v -coverprofile=profile.cov ./... ; go tool cover -func profile.cov && go tool cover -html=profile.cov

# -------- MOCK COMMANDS -------

docker.mock.generate:
	make dockerTestBuild
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			test-build mockgen -source="$(FILE)"

# -------- DOCKER BUILDS -------

dockerProductionBuild:
	@docker build \
			--tag production-build \
			-f ./deployment/production/Dockerfile .

dockerTestBuild:
	@docker build \
			--tag test-build \
			-f ./deployment/test/Dockerfile .

build:
	go build -v -x

install:
	go install -v -x

vendor/clean:  ## Cleanup vendor packages "*_test" files, testdata and nogo files.
	@find ./vendor -type d -name 'testdata' -print | xargs rm -rf
	@find ./vendor -type f -name '*_test.go' -print -exec rm {} ";"
	@find ./vendor \
		\( -name '*.sh' \
		-or -name 'Makefile' \
		-or -name '.gitignore' \
		-or -name '*.yml' \
		-or -name '*.txtr' \
		-or -name '*.vim' \
		-or -name '*.el' \) \
		-type f -print -exec rm {} ";"

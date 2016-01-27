VERSION = $(shell date +'%Y%m%d%H').$(shell git rev-parse --short=10 HEAD)
NAME = go-filelist
LDFLAGS = \
	-X $(NAME)/cli.version=$(VERSION) \
	-B 0x$(shell head -c20 /dev/urandom|od -An -tx1|tr -d ' \n')

tests =
-include "parameters.mk"

.PHONY = all $(NAME) release hooks test

all: hooks
	go get github.com/nsf/gocode
	go get code.google.com/p/rog-go/exp/cmd/godef
	go get github.com/tools/godep
	go get github.com/stretchr/testify/assert
	godep restore
	$(MAKE) $(NAME)

$(NAME): main.go $(shell find src -type f -name '*.go')
	go build -a -ldflags "$(LDFLAGS)" -v

release:
	mkdir -p $(NAME)-"$(VERSION)"/src/$(NAME)

	rsync -avzr --delete \
		--filter='- $(NAME)-*' \
		--filter='- /$(NAME)' \
		--filter='- .*' \
		--filter='- *~' \
		--filter='- *.org' \
		. $(NAME)-"$(VERSION)"/src/$(NAME)

	tar czf $(NAME)-"$(VERSION)".tgz $(NAME)-"$(VERSION)"

.git/hooks/pre-commit:
	cp -f hooks/pre-commit .git/hooks/pre-commit

hooks: .git/hooks/pre-commit
	chmod +x .git/hooks/*

test: $(tests)
	go vet .
	go test $(foreach t,$^,./$(t))

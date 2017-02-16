TOOLS=honnef.co/go/staticcheck/cmd/staticcheck honnef.co/go/simple/cmd/gosimple honnef.co/go/unused/cmd/unused

.DEFAULT_GOAL: test

.PHONY: test
test:
	@# vet or staticcheck errors are unforgivable. gosimple produces warnings
	@go vet
	@staticcheck
	@unused -exported
	@-gosimple
	@go test

.PHONY: install
install:
	go install

.PHONY: bootstrap
bootstrap:
	$(foreach tool,$(TOOLS),$(call goget, $(tool)))

define goget
	go get -u $(1)
	
endef


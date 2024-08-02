.PHONY: generate
generate:
	for i in types client server; do \
		docker run --rm -it -v $$PWD:$$PWD -w $$PWD hidori/oapi-codegen \
			-generate $$i -package contactsapi ./contacts.yml > ./contactsapi/$$i.gen.go; \
	done

.PHONY: clean
clean:
	rm -f ./contactsapi/*.gen.go

.PHONY: run/server
run/server:
	go run ./server/cmd/main.go

.PHONY: run/client
run/client:
	go run ./client/cmd/main.go

get-protos-submodule:
	@echo "initializing protos submodule..." && \
	git submodule add --progress --force https://github.com/exag-community/protos.git protos

update-protos-submodule:
	@echo "initializing protos submodule..." && \
	git submodule update --init --recursive  --remote

gen-protos:
	@echo "generating protos..." && \
	chmod +x scripts/gen-protos.sh && ./scripts/gen-protos.sh

.PHONY: connect-to-admin-gateway get-protos-submodule update-protos-submodule gen-protos

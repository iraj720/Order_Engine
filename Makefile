GOPATH:=$(shell go env GOPATH)

APP_NAME:=$(order_engine)

.PHONY: proto
proto:
	rm -rf api/pb api/swagger
	mkdir -p api/pb api/swagger
	protoc -I submodules/api/aux/ \
	-I submodules/api/proto/ \
	--go_out api/pb \
	--openapiv2_out=logtostderr=true,repeated_path_param_separator=ssv:./api/swagger \
	--openapiv2_opt use_go_templates=true \
	--openapiv2_opt logtostderr=true \
	--openapiv2_opt use_go_templates=true \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	--go_opt paths=source_relative \
	--go-grpc_out api/pb \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out api/pb \
	--grpc-gateway_opt paths=source_relative \
	submodules/api/proto/**/*.proto \
	submodules/api/proto/**/**/*.proto


# .PHONY: extract-translation
# extract-translation:
# 	goi18n extract -outdir=confs/translations
# 	goi18n merge -outdir=confs/translations confs/translations/active.*.toml

# .PHONY: merge-translation
# merge-translation:
# 	goi18n merge -outdir=confs/translations confs/translations/active.*.toml confs/translations/translate.*.toml
# 	rm confs/translations/translate*
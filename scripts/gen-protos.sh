# shellcheck disable=SC2046
PROTO_PATH=protos

# define the output directory
OUT_DIR=gen

# create the gen directory
mkdir -p $OUT_DIR

# generate for auth-server
protoc -I=$PROTO_PATH --go_out=$OUT_DIR --go_opt=paths=source_relative \
  --go-grpc_out=$OUT_DIR --go-grpc_opt=paths=source_relative \
  $(find $PROTO_PATH -name 'email*.proto') common.proto
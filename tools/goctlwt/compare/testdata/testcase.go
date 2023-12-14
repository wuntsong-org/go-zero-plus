package testdata

import _ "embed"

var (
	//go:embed unformat.api
	unformatApi string
	//go:embed kotlin.api
	kotlinApi string
	//go:embed user.sql
	userSql string

	list = Files{
		{
			IsDir: true,
			Path:  "version",
			Cmd:   "goctlwt --version",
		},
		{
			IsDir: true,
			Path:  "api/sample_file/local",
			Cmd:   "goctlwt api --o sample.api",
		},
		{
			IsDir: true,
			Path:  "api/sample_file/local/assign",
			Cmd:   "goctlwt api --o=sample.api",
		},
		{
			IsDir: true,
			Path:  "api/sample_file/local/assign/shorthand",
			Cmd:   "goctlwt api -o=sample.api",
		},
		{
			IsDir: true,
			Path:  "api/sample_file/remote",
			Cmd:   "goctlwt api --o sample.api --remote https://github.com/wuntsong-org/go-zero-plus-template --branch main",
		},
		{
			IsDir: true,
			Path:  "api/sample_file/remote/shorthand",
			Cmd:   "goctlwt api -o sample.api -remote https://github.com/wuntsong-org/go-zero-plus-template -branch main",
		},
		{
			IsDir: true,
			Path:  "api/sample_file/remote/assign",
			Cmd:   "goctlwt api --o=sample.api --remote https://github.com/wuntsong-org/go-zero-plus-template --branch=main",
		},
		{
			IsDir: true,
			Path:  "api/sample_file/remote/assign/shorthand",
			Cmd:   "goctlwt api -o=sample.api -remote https://github.com/wuntsong-org/go-zero-plus-template -branch=main",
		},
		{
			IsDir: true,
			Path:  "api/dart/legacy/true",
			Cmd:   "goctlwt api --o sample.api && goctlwt api dart --api sample.api --dir . --hostname 127.0.0.1 --legacy true",
		},
		{
			IsDir: true,
			Path:  "api/dart/legacy/true/shorthand",
			Cmd:   "goctlwt api -o sample.api && goctlwt api dart -api sample.api -dir . -hostname 127.0.0.1 -legacy true",
		},
		{
			IsDir: true,
			Path:  "api/dart/legacy/true/assign",
			Cmd:   "goctlwt api --o=sample.api && goctlwt api dart --api=sample.api --dir=. --hostname=127.0.0.1 --legacy=true",
		},
		{
			IsDir: true,
			Path:  "api/dart/legacy/true/assign/shorthand",
			Cmd:   "goctlwt api -o=sample.api && goctlwt api dart -api=sample.api -dir=. -hostname=127.0.0.1 -legacy=true",
		},
		{
			IsDir: true,
			Path:  "api/dart/legacy/false",
			Cmd:   "goctlwt api --o sample.api && goctlwt api dart --api sample.api --dir . --hostname 127.0.0.1 --legacy true",
		},
		{
			IsDir: true,
			Path:  "api/dart/legacy/false/shorthand",
			Cmd:   "goctlwt api -o sample.api && goctlwt api dart -api sample.api -dir . -hostname 127.0.0.1 -legacy true",
		},
		{
			IsDir: true,
			Path:  "api/dart/legacy/false/assign",
			Cmd:   "goctlwt api --o=sample.api && goctlwt api dart --api=sample.api --dir=. --hostname=127.0.0.1 --legacy=true",
		},
		{
			IsDir: true,
			Path:  "api/dart/legacy/false/assign/shorthand",
			Cmd:   "goctlwt api -o=sample.api && goctlwt api dart -api=sample.api -dir=. -hostname=127.0.0.1 -legacy=true",
		},
		{
			IsDir: true,
			Path:  "api/doc",
			Cmd:   "goctlwt api --o sample.api && goctlwt api doc --dir . --o .",
		},
		{
			IsDir: true,
			Path:  "api/doc/shorthand",
			Cmd:   "goctlwt api -o sample.api && goctlwt api doc -dir . -o .",
		},
		{
			IsDir: true,
			Path:  "api/doc/assign",
			Cmd:   "goctlwt api --o=sample.api && goctlwt api doc --dir=. --o=.",
		},
		{
			IsDir: true,
			Path:  "api/doc/assign/shorthand",
			Cmd:   "goctlwt api -o=sample.api && goctlwt api doc -dir=. -o=.",
		},
		{
			Path:    "api/format/unformat.api",
			Content: unformatApi,
			Cmd:     "goctlwt api format --dir . --iu",
		},
		{
			Path:    "api/format/shorthand/unformat.api",
			Content: unformatApi,
			Cmd:     "goctlwt api format -dir . -iu",
		},
		{
			Path:    "api/format/assign/unformat.api",
			Content: unformatApi,
			Cmd:     "goctlwt api format --dir=. --iu",
		},
		{
			Path:    "api/format/assign/shorthand/unformat.api",
			Content: unformatApi,
			Cmd:     "goctlwt api format -dir=. -iu",
		},
		{
			IsDir: true,
			Path:  "api/go/style/default",
			Cmd:   "goctlwt api --o sample.api && goctlwt api go --api sample.api --dir .",
		},
		{
			IsDir: true,
			Path:  "api/go/style/default/shorthand",
			Cmd:   "goctlwt api -o sample.api && goctlwt api go -api sample.api -dir .",
		},
		{
			IsDir: true,
			Path:  "api/go/style/assign/default",
			Cmd:   "goctlwt api --o=sample.api && goctlwt api go --api=sample.api --dir=.",
		},
		{
			IsDir: true,
			Path:  "api/go/style/assign/default/shorthand",
			Cmd:   "goctlwt api -o=sample.api && goctlwt api go -api=sample.api -dir=.",
		},
		{
			IsDir: true,
			Path:  "api/go/style/goZero",
			Cmd:   "goctlwt api --o sample.api && goctlwt api go --api sample.api --dir . --style goZero",
		},
		{
			IsDir: true,
			Path:  "api/go/style/goZero/shorthand",
			Cmd:   "goctlwt api -o sample.api && goctlwt api go -api sample.api -dir . -style goZero",
		},
		{
			IsDir: true,
			Path:  "api/go/style/goZero/assign",
			Cmd:   "goctlwt api --o=sample.api && goctlwt api go --api=sample.api --dir=. --style=goZero",
		},
		{
			IsDir: true,
			Path:  "api/go/style/goZero/assign/shorthand",
			Cmd:   "goctlwt api -o=sample.api && goctlwt api go -api=sample.api -dir=. -style=goZero",
		},
		{
			IsDir: true,
			Path:  "api/java",
			Cmd:   "goctlwt api --o sample.api && goctlwt api java --api sample.api --dir .",
		},
		{
			IsDir: true,
			Path:  "api/java/shorthand",
			Cmd:   "goctlwt api -o sample.api && goctlwt api java -api sample.api -dir .",
		},
		{
			IsDir: true,
			Path:  "api/java/assign",
			Cmd:   "goctlwt api --o=sample.api && goctlwt api java --api=sample.api --dir=.",
		},
		{
			IsDir: true,
			Path:  "api/java/shorthand/assign",
			Cmd:   "goctlwt api -o=sample.api && goctlwt api java -api=sample.api -dir=.",
		},
		{
			IsDir: true,
			Path:  "api/new/style/default",
			Cmd:   "goctlwt api new greet",
		},
		{
			IsDir: true,
			Path:  "api/new/style/goZero",
			Cmd:   "goctlwt api new greet --style goZero",
		},
		{
			IsDir: true,
			Path:  "api/new/style/goZero/assign",
			Cmd:   "goctlwt api new greet --style=goZero",
		},
		{
			IsDir: true,
			Path:  "api/new/style/goZero/shorthand",
			Cmd:   "goctlwt api new greet -style goZero",
		},
		{
			IsDir: true,
			Path:  "api/new/style/goZero/shorthand/assign",
			Cmd:   "goctlwt api new greet -style=goZero",
		},
		{
			IsDir: true,
			Path:  "api/ts",
			Cmd:   "goctlwt api --o sample.api && goctlwt api ts --api sample.api --dir . --unwrap --webapi .",
		},
		{
			IsDir: true,
			Path:  "api/ts/shorthand",
			Cmd:   "goctlwt api -o sample.api && goctlwt api ts -api sample.api -dir . -unwrap -webapi .",
		},
		{
			IsDir: true,
			Path:  "api/ts/assign",
			Cmd:   "goctlwt api --o=sample.api && goctlwt api ts --api=sample.api --dir=. --unwrap --webapi=.",
		},
		{
			IsDir: true,
			Path:  "api/ts/shorthand/assign",
			Cmd:   "goctlwt api -o=sample.api && goctlwt api ts -api=sample.api -dir=. -unwrap -webapi=.",
		},
		{
			IsDir: true,
			Path:  "api/validate",
			Cmd:   "goctlwt api --o sample.api && goctlwt api validate --api sample.api",
		},
		{
			IsDir: true,
			Path:  "api/validate/shorthand",
			Cmd:   "goctlwt api -o sample.api && goctlwt api validate -api sample.api",
		},
		{
			IsDir: true,
			Path:  "api/validate/assign",
			Cmd:   "goctlwt api --o=sample.api && goctlwt api validate --api=sample.api",
		},
		{
			IsDir: true,
			Path:  "api/validate/shorthand/assign",
			Cmd:   "goctlwt api -o=sample.api && goctlwt api validate -api=sample.api",
		},
		{
			IsDir: true,
			Path:  "env/show",
			Cmd:   "goctlwt env > env.txt",
		},
		{
			IsDir: true,
			Path:  "env/check",
			Cmd:   "goctlwt env check -f -v",
		},
		{
			IsDir: true,
			Path:  "env/install",
			Cmd:   "goctlwt env install -v",
		},
		{
			IsDir: true,
			Path:  "kube",
			Cmd:   "goctlwt kube deploy --image alpine --name foo --namespace foo --o foo.yaml --port 8888",
		},
		{
			IsDir: true,
			Path:  "kube/shorthand",
			Cmd:   "goctlwt kube deploy -image alpine -name foo -namespace foo -o foo.yaml -port 8888",
		},
		{
			IsDir: true,
			Path:  "kube/assign",
			Cmd:   "goctlwt kube deploy --image=alpine --name=foo --namespace=foo --o=foo.yaml --port=8888",
		},
		{
			IsDir: true,
			Path:  "kube/shorthand/assign",
			Cmd:   "goctlwt kube deploy -image=alpine -name=foo -namespace=foo -o=foo.yaml -port=8888",
		},
		{
			IsDir: true,
			Path:  "model/mongo/cache",
			Cmd:   "goctlwt model mongo --dir . --type user --style goZero -c",
		},
		{
			IsDir: true,
			Path:  "model/mongo/cache/shorthand",
			Cmd:   "goctlwt model mongo -dir . -type user -style goZero -c",
		},
		{
			IsDir: true,
			Path:  "model/mongo/cache/assign",
			Cmd:   "goctlwt model mongo --dir=. --type=user --style=goZero -c",
		},
		{
			IsDir: true,
			Path:  "model/mongo/cache/shorthand/assign",
			Cmd:   "goctlwt model mongo -dir=. -type=user -style=goZero -c",
		},
		{
			IsDir: true,
			Path:  "model/mongo/nocache",
			Cmd:   "goctlwt model mongo --dir . --type user",
		},
		{
			IsDir: true,
			Path:  "model/mongo/nocache/shorthand",
			Cmd:   "goctlwt model mongo -dir . -type user",
		},
		{
			IsDir: true,
			Path:  "model/mongo/nocache/assign",
			Cmd:   "goctlwt model mongo --dir=. --type=user",
		},
		{
			IsDir: true,
			Path:  "model/mongo/nocache/shorthand/assign",
			Cmd:   "goctlwt model mongo -dir=. -type=user",
		},
		{
			Content: userSql,
			Path:    "model/mysql/ddl/user.sql",
			Cmd:     "goctlwt model mysql ddl --database user --dir cache --src user.sql -c",
		},
		{
			Content: userSql,
			Path:    "model/mysql/ddl/shorthand/user.sql",
			Cmd:     "goctlwt model mysql ddl -database user -dir cache -src user.sql -c",
		},
		{
			Content: userSql,
			Path:    "model/mysql/ddl/assign/user.sql",
			Cmd:     "goctlwt model mysql ddl --database=user --dir=cache --src=user.sql -c",
		},
		{
			Content: userSql,
			Path:    "model/mysql/ddl/shorthand/assign/user.sql",
			Cmd:     "goctlwt model mysql ddl -database=user -dir=cache -src=user.sql -c",
		},
		{
			Content: userSql,
			Path:    "model/mysql/ddl/user.sql",
			Cmd:     "goctlwt model mysql ddl --database user --dir nocache --src user.sql",
		},
		{
			Content: userSql,
			Path:    "model/mysql/ddl/shorthand/user.sql",
			Cmd:     "goctlwt model mysql ddl -database user -dir nocache -src user.sql",
		},
		{
			Content: userSql,
			Path:    "model/mysql/ddl/assign/user.sql",
			Cmd:     "goctlwt model mysql ddl --database=user --dir=nocache --src=user.sql",
		},
		{
			Content: userSql,
			Path:    "model/mysql/ddl/shorthand/assign/user.sql",
			Cmd:     "goctlwt model mysql ddl -database=user -dir=nocache -src=user.sql",
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource",
			Cmd:   `goctlwt model mysql datasource --url $DSN --dir cache --table "*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/shorthand",
			Cmd:   `goctlwt model mysql datasource -url $DSN -dir cache -table "*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/shorthand2",
			Cmd:   `goctlwt model mysql datasource -url $DSN -dir cache -t "*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/assign",
			Cmd:   `goctlwt model mysql datasource --url=$DSN --dir=cache --table="*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/shorthand/assign",
			Cmd:   `goctlwt model mysql datasource -url=$DSN -dir=cache -table="*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/shorthand2/assign",
			Cmd:   `goctlwt model mysql datasource -url=$DSN -dir=cache -t="*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource",
			Cmd:   `goctlwt model mysql datasource --url $DSN --dir nocache --table "*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/shorthand",
			Cmd:   `goctlwt model mysql datasource -url $DSN -dir nocache -table "*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/shorthand2",
			Cmd:   `goctlwt model mysql datasource -url $DSN -dir nocache -t "*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/assign",
			Cmd:   `goctlwt model mysql datasource --url=$DSN --dir=nocache --table="*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/shorthand/assign",
			Cmd:   `goctlwt model mysql datasource -url=$DSN -dir=nocache -table="*" -c`,
		},
		{
			IsDir: true,
			Path:  "model/mysql/datasource/shorthand2/assign",
			Cmd:   `goctlwt model mysql datasource -url=$DSN -dir=nocache -t="*" -c`,
		},
		{
			IsDir: true,
			Path:  "rpc/new",
			Cmd:   "goctlwt rpc new greet",
		},
		{
			IsDir: true,
			Path:  "rpc/template",
			Cmd:   "goctlwt rpc template --o greet.proto",
		},
		{
			IsDir: true,
			Path:  "rpc/template/shorthand",
			Cmd:   "goctlwt rpc template -o greet.proto",
		},
		{
			IsDir: true,
			Path:  "rpc/template/assign",
			Cmd:   "goctlwt rpc template --o=greet.proto",
		},
		{
			IsDir: true,
			Path:  "rpc/template/shorthand/assign",
			Cmd:   "goctlwt rpc template -o=greet.proto",
		},
		{
			IsDir: true,
			Path:  "rpc/protoc",
			Cmd:   "goctlwt rpc template --o greet.proto && goctlwt rpc protoc greet.proto --go_out . --go-grpc_out . --zrpc_out .",
		},
		{
			IsDir: true,
			Path:  "rpc/protoc/assign",
			Cmd:   "goctlwt rpc template --o=greet.proto && goctlwt rpc protoc greet.proto --go_out=. --go-grpc_out=. --zrpc_out=.",
		},
	}
)

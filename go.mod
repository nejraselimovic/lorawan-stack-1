module go.thethings.network/lorawan-stack/v3

go 1.14

// Use our fork of grpc-gateway.
replace github.com/grpc-ecosystem/grpc-gateway => github.com/TheThingsIndustries/grpc-gateway v1.14.5-gogo

// Use our fork of otto.
replace github.com/robertkrimen/otto => github.com/TheThingsIndustries/otto v0.0.0-20200507140236-e2ec6b546eb2

// Do not upgrade Echo beyond v4.1.2.
// See https://github.com/TheThingsNetwork/lorawan-stack/issues/977.
replace github.com/labstack/echo/v4 => github.com/labstack/echo/v4 v4.1.2

// Do not upgrade go-sqlmock beyond v1.3.0.
// See https://github.com/heptiolabs/healthcheck/issues/23.
replace gopkg.in/DATA-DOG/go-sqlmock.v1 => gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0

require (
	contrib.go.opencensus.io/exporter/prometheus v0.1.0
	github.com/PuerkitoBio/purell v1.1.1
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/TheThingsIndustries/mystique v0.0.0-20200127144137-4aa959111fe7
	github.com/TheThingsNetwork/go-cayenne-lib v1.0.0
	github.com/aws/aws-sdk-go v1.31.1
	github.com/blang/semver v3.5.1+incompatible
	github.com/chrj/smtpd v0.1.2
	github.com/disintegration/imaging v1.6.2
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/envoyproxy/protoc-gen-validate v0.3.0-java
	github.com/felixge/httpsnoop v1.0.1
	github.com/fsnotify/fsnotify v1.4.9
	github.com/getsentry/sentry-go v0.6.1
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/gobwas/glob v0.2.3
	github.com/gogo/protobuf v1.3.1
	github.com/golang/gddo v0.0.0-20200519224240-a4ebd2f7e574
	github.com/golang/protobuf v1.3.5
	github.com/google/wire v0.4.0 // indirect
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/websocket v1.4.2
	github.com/gotnospirit/makeplural v0.0.0-20180622080156-a5f48d94d976 // indirect
	github.com/gotnospirit/messageformat v0.0.0-20190719172517-c1d0bdacdea2
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.5
	github.com/heptiolabs/healthcheck v0.0.0-20180807145615-6ff867650f40
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c
	github.com/jacobsa/crypto v0.0.0-20190317225127-9f44e2d11115
	github.com/jacobsa/oglematchers v0.0.0-20150720000706-141901ea67cd // indirect
	github.com/jacobsa/oglemock v0.0.0-20150831005832-e94d794d06ff // indirect
	github.com/jacobsa/ogletest v0.0.0-20170503003838-80d50a735a11 // indirect
	github.com/jacobsa/reqtrace v0.0.0-20150505043853-245c9e0234cb // indirect
	github.com/jarcoal/httpmock v1.0.5
	github.com/jaytaylor/html2text v0.0.0-20200412013138-3577fbdbcff7
	github.com/jinzhu/gorm v1.9.12
	github.com/kr/pretty v0.2.0
	github.com/kr/text v0.2.0 // indirect
	github.com/labstack/echo/v4 v4.1.16
	github.com/labstack/gommon v0.3.0
	github.com/lib/pq v1.5.2
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mattn/go-isatty v0.0.12
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mitchellh/mapstructure v1.3.0
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/nats-io/nats-server/v2 v2.1.4
	github.com/nats-io/nats.go v1.9.2
	github.com/oklog/ulid/v2 v2.0.2
	github.com/olekukonko/tablewriter v0.0.4 // indirect
	github.com/openshift/osin v1.0.1
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.5.1
	github.com/robertkrimen/otto v0.0.0-20191219234010-c382bd3c16ff
	github.com/satori/go.uuid v1.2.0
	github.com/sendgrid/rest v2.4.1+incompatible // indirect
	github.com/sendgrid/sendgrid-go v3.5.0+incompatible
	github.com/skip2/go-qrcode v0.0.0-20200519171959-a3b48390827e
	github.com/smartystreets/assertions v1.1.0
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.0.0
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0
	github.com/ssor/bom v0.0.0-20170718123548-6386211fdfcf // indirect
	github.com/valyala/fasttemplate v1.1.0 // indirect
	go.opencensus.io v0.22.3
	go.packetbroker.org/api/v3 v3.0.0
	go.thethings.network/lorawan-stack-legacy/v2 v2.0.2
	gocloud.dev v0.19.0
	gocloud.dev/pubsub/natspubsub v0.19.0
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	golang.org/x/image v0.0.0-20200430140353-33d19683fad8 // indirect
	golang.org/x/mod v0.3.0 // indirect
	golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a
	golang.org/x/sys v0.0.0-20200519105757-fe76b779f299 // indirect
	golang.org/x/tools v0.0.0-20200519205726-57a9e4404bf7 // indirect
	google.golang.org/api v0.24.0
	google.golang.org/genproto v0.0.0-20200507105951-43844f6eee31
	google.golang.org/grpc v1.29.1
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.0.0-00010101000000-000000000000 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/ini.v1 v1.56.0 // indirect
	gopkg.in/mail.v2 v2.3.1
	gopkg.in/square/go-jose.v2 v2.5.1
	gopkg.in/yaml.v2 v2.3.0
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect
)

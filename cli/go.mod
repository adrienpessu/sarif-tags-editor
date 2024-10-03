module sarif-tags-editors

go 1.22

require (
	github.com/google/go-github/v50 v50.0.0
	golang.org/x/oauth2 v0.5.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

require github.com/adrienpessu/sarif-tags-editors/libs/sarif v0.0.0

replace github.com/adrienpessu/sarif-tags-editors/libs/sarif => ./../libs/sarif

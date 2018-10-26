package foobarbaz

import "github.com/google/go-cloud/wire"

var SuperSet = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz)

module github.com/class100/sdk-go

go 1.15

require (
	github.com/class100/core v0.0.2
	github.com/class100/yunke-core v1.0.3
	github.com/storezhang/gox v1.2.8
)

// replace github.com/storezhang/gox => ../../storezhang/gox
// replace github.com/storezhang/replace => ../../storezhang/replace
// replace github.com/class100/yunke-core => ../yunke-core
replace github.com/class100/core => ../core

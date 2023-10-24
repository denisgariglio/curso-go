module github.com/denisgariglio/goexpert/Packaging/4-gomod-replace/sistema

go 1.20

replace github.com/denisgariglio/goexpert/Packaging/4-gomod-replace/math => ../math

require github.com/denisgariglio/goexpert/Packaging/4-gomod-replace/math v0.0.0-00010101000000-000000000000

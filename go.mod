module github.com/ViewlessNote/actions-test

go 1.23

toolchain go1.24.6

replace github.com/ViewlessNote/repo-A => ../repo-A/sdk

require github.com/ViewlessNote/repo-A v0.0.0-20260311140045-c32d085507e3

require google.golang.org/protobuf v1.36.11 // indirect

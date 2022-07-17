# example-go-generate-tests

```
export EXAMPLE_TESTS_ROOT=<YOUR_ROOT_DIR>
export GOTESTS_TEMPLATE_DIR=$EXAMPLE_TESTS_ROOT/testtmpls

go generate ./...
```

### Note
If you add new wire.go, cd the directory which new wire.go exists, then `$ wire`.

Once wire_gen.go is created, you can regenerate it by running go generate.

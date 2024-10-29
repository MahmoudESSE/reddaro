alias r := run
alias b := build
# alias t := tests
# alias te := test_example
# alias d := deps
alias c := clean

bin := "reddaro"
bin_dir := "bin"

srcs := `fd -0augp "**/reddaro/*.go" | xargs --null`

# example_tests := `fd -0augp "**/example/**/*_test.go" | xargs --null`

deps := `fd -E '*main*' .go`

# build then run the program in {bin_dir}/{bin}
run:
    just b
    ./{{bin_dir}}/{{bin}}

# build the program to {bin_dir}/{bin}
build:
    go build -o {{bin_dir}}/{{bin}} {{srcs}}

# run all the tests
# tests: test_example

# test the example
# test_example:
#     go test -v {{example_tests}}

# get all dependencies
# deps:
#     @echo TODO

# clean the project
clean:
    go clean
    @rm -rf bin/*

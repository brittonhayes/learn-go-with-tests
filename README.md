# Learn Go with Tests 🐹🧪

> A personal exercise in improving my testing skills with Go by following
> [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

## Background

<details>
<summary>Expand for a more background context</summary>
<br>
I have been writing Go for a while now. I'm pretty comfortable writing tests, but I could always get better. This repository is
a collection of packages and tests for me to fix gaps in my testing toolkit. It's also a place for me to learn how to write my packages
in a way that is more conducive to testing.

I've found myself too many times writing a package *then* its test suite. This always end up with trying to fit a square peg in a round
hole, and it's just not the best way of doing things. Given that, his repository is also a playground for trying out TDD (test-driven development).

</details>


## Usage ⚡

```shell
# Install the dependencies
go mod tidy

# Run the tests
go test -v -race ./...
```
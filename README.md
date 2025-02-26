# embedvendor

this module exists to reproduce an issue I was having with go embed.

# Problem

When working with the `//go:embed`, I ran into an issue where it performed
differently when working locally with `go.work` vs pulling the module using
`go get`.

For example, I have the [getvendor](https://github.com/jsteenb2/getvendor) module
that depends on `embedvendor`. Everything works as expected when working with `go.work`.
However, when I use `go get` to pull the module into CI for instance, it fails. The failure
I am seeing is that in CI the paths with `**/vendor/**` are excluded. The `go get` cmd does
not pull that directory. However, it pulls in everything else. This led to a very confusing
situation and the presentation of this issue to the go team.
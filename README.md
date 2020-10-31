<h1 align="center">goals</h1>
<p align="center">An automation tool that empowers developers to create more powerful automations in an easy way</p>

---

__Declarative configuration__: Configuration is declared in a YAML file, that contains a list of goals. You can think of a goal as a target in a Makefile. Each goal has at least a description and a list functions to run.

__Programmatic automation logic__: Beside the default functions that goals comes with, additional functions can be created to be used in the same way. A function in goals is nothing but an actual Go function.

__Cross platform support__: Currently goals relies on the [plugins](https://golang.org/pkg/plugin/) package, which is only supported on Linux, FreeBSD, and macOS.

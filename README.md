# Cornucopia

<img src="https://github.com/dvaumoron/cornucopia/raw/main/logo/cornucopialogo.png" width="100">

Cornucopia is a Go code generation tool with [shared scripts](https://github.com/dvaumoron/cornucopiarecipes) (gluing together [starlark-go](https://pkg.go.dev/go.starlark.net/starlark) and [jennifer](https://pkg.go.dev/github.com/dave/jennifer)). Basic text file and marshalling to JSON or YAML are also available.

## Why ?

You don't like writing boilerplate code and find that Golang has no macro preprocessor, admittedly now that it has generics many use cases are gone, for others there is still [`go generate`](https://go.dev/blog/generate), but this one doesn't do anything by itself.

You need a code generation tool, you need Cornucopia.

In this age of generative AI, it might seem useless, but i don't usually use a bazooka to kill flies. More seriously, the probabilistic part of generative AIs makes them unreliable tools, and the fact that it is improving quickly does not make me forget that they use a lot of resources for the simplest tasks.

Moreover, I was not satisfied with jennifer, I like the structuring provided by this library (compared to template engine based solutions), but having to compile at each generator change before launching it lacks meaning, this is not a critical task with a heavy load.

I decided to try something, using a starlark implementation that brings a familiar syntax similar to the Python language (and saves me from reinventing the wheel).

However this leaves a problem, what can I do if I am missing some details about the targeted source code ? Should I fall back to generative AI ? I believe that we can do otherwise, thanks to sharing, that's why I added to Cornucopia the possibility of referring to a common repository of scripts, defaulting to [Cornucopia Recipes](https://github.com/dvaumoron/cornucopiarecipes).

I hope you will enjoy and share.

## Getting started

Get the [last binary](releases) depending on your OS.

Create a file test.crn in your current directory similar to :

```Python
load("hello.crn", "Hello")

Hello()
```

And launch the command `cornucopia test.crn` (or a `go generate` command with a directive `//go:generate cornucopia test.crn` in one of the go file in the current directory). This will generate a hello.go like :

```Go
package main

import "fmt"

func main() {
    fmt.Println("Hello world !")
}
```

You can change the Hello call to add a string argument and retry, see [hello.crn](https://github.com/dvaumoron/cornucopiarecipes/blob/main/hello.crn) to understand the "hidden" recipe.

Other examples could be seen [here](examples) or [there](glu/go/self.crn) (used to generate part of Cornucopia internal boilerplate code)

## Language Reference

On top of the [Starlark specification](https://github.com/google/starlark-go/blob/master/doc/spec.md), Cornucopia add a load implementation and allows to call some part of the [jennifer API](https://pkg.go.dev/github.com/dave/jennifer/jen).

The Cornucopia load implementation try the following resolution to find a an external script :

- read from current directory
- read from local repository path (environment variable `CORNUCOPIA_REPO_PATH`, default is `$HOME/.cornucopia/recipes`)
- download from repository url (environment variable `CORNUCOPIA_REPO_URL`, default link to [Cornucopia Recipes](https://github.com/dvaumoron/cornucopiarecipes)) and write the content in local repository path

## License

The Cornucopia project is released under the Apache 2.0 license. See [LICENSE](LICENSE).

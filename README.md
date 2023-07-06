# Cornucopia

<img src="https://github.com/dvaumoron/cornucopia/raw/main/logo/cornucopialogo.png" width="100">

Cornucopia is a go code generation tool with [shared scripts](https://github.com/dvaumoron/cornucopiarecipes) (gluing together [starlark-go](https://pkg.go.dev/go.starlark.net/starlark) and [jennifer](https://pkg.go.dev/github.com/dave/jennifer))

## Why ?

In this age of generative AI, it might seem useless, but i don't usually use a bazooka to kill flies. More seriously, the probabilistic part of generative AIs makes them unreliable tools, and the fact that it is improving quickly does not make me forget that they use a lot of resources for the simplest tasks.

Moreover, I was not satisfied with jennifer, I like the structuring provided by this library (compared to template engine based solutions), but having to compile at each generator change before launching it lacks meaning, this is not a critical task with a heavy load.

I decided to try something, using a starlark implementation that brings a familiar syntax similar to the Python language (and saves me from reinventing the wheel).

However this leaves a problem, what can I do if I am missing some details about the targeted source code? Should I fall back to generative AI? I believe that we can do otherwise, thanks to sharing, that's why I added to Cornucopia the possibility of referring to a common repository of scripts, defaulting to [Cornucopia Recipes](https://github.com/dvaumoron/cornucopiarecipes).

I hope you will enjoy and share.

## Getting started

In order to install Cornucopia (with the go langage already installed), you can use the command: 

    go install github.com/dvaumoron/cornucopia@latest

With a file test.crn similar to [this example](examples/test.crn) in your current directory, the command :

    cornucopia test.crn

will generate a hello.go like [this one](examples/hello.go).

You can change the Hello call to add a string argument and retry, see [hello.crn](https://github.com/dvaumoron/cornucopiarecipes/blob/main/hello.crn) to understand the "hidden" recipe.

## License

The Cornucopia project is released under the Apache 2.0 license. See [LICENSE](LICENSE).

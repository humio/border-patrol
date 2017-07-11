<img align="right" src="images/logo.png" />

# Border Patrol

__This Project is WIP__

Border Patrol is an Architecture Linter. That means that it helps you to avoid
introducing code that goes against the project's structure. It prevents you from
importing/using packages you shouldn't. `MyProject.Format.Json` should not be
importing anything from `Graphics.*`, so don't allow it - __enforce__ it.

Border Patrol is best using when applying the architectural thinking described in
[The Hexagonal Architecture](http://alistair.cockburn.us/Hexagonal+architecture)
and [Domain-driven design](https://en.wikipedia.org/wiki/Domain-driven_design).

## Configuration File

Border Patrol is configured by placing a _JSON_ file called `boundaries.json` at
the root of your project. It contains the rules defining the boundaries in your
code base.

### Example

```json
{
  "restrictions": {
    "Api": ["Html"],
    "Logic": ["Http", "Api", "SQL"],
    "GUI.Finance": ["Api", "Http", "Logic.Logistics"]
  }
}
```

Here is how to read the rules, e.g.:

```
"GUI": ["Api", "Http", "Logic.Finance"]
```
The code in the packages `GUI.Finance` and `GUI.Finance.*` (any sub-packages of `GUI.Finance`) are restricted from
importing `Api`, `Http` or `Logic.Logistics` or anything from any of their sub-packages.
On the other hand `GUI.Finance` could for instance import `Logic.Finance` without any problems.  

Figuring out restrictions is usually pretty easy. E.g.:

The `Api.*` modules should not be producing any view HTML so has no business
importing the `Html.*` packages.

The `Logic.*` modules should be where the bulk of the business logic lives, it
should not touch networking code or database dependencies.

Anytime you find someone added some dependency to apart of the code where it does not
fit, add a restriction.

## Running

```bash
$ border-patrol
```

If all rules pass it returns with exit code `0`, if any violations are found it
exits with exit code `1`. This makes it easy to make Border Patrol part of your
CI setup.

## Supported Languages

This is a very young project. These are the languages we initial plan to support.

- [ ] Java
- [x] Scala _(No support for single class-restrictions*)_
- [x] Elm
- [ ] NodeJS
- [ ] Go

but suggestions and contributions/PRs are welcome!

* This means you Border Patrol will not detect violations on specific classes when
using the import syntax `import my.{Class, OtherClass}` if you e.g. restricted `my.Class`.

## Rational

_Code Cohesion_ is a measure of how closely code elements inside a
module are related to each other. As a project grows and new contributors join
cohesion tends to decline.

By grouping code in packages by domain/responsibility you can usually quite easily
see what the code should and should _not_ be doing, e.g. code inside the `API` package
should probably be dealing with packages like `network`, `http` or `json` but probably
not be using `html` or `opengl`.

So, `API` should not even be allowed to import those packages. That is common sense.
But conventions like these are hard to maintain in teams without formalizing them as code -
this is what Border Patrol was made for. It maintains the boundaries between code with
different responsibilities in your code base.

## TODOs

- [ ] Allow setting source root directory in config file.

## Project Sponsor - Humio

![Humio](./images/humio.png)

[Humio](https://humio.com/)

Humio is a Distributed Log Aggregation and Monitoring System. Humio has a
powerful query language and makes it feel like using `tail` and `grep` with
aggregation functions and graphs built-in. It crunches through TB of data in no
time at all.

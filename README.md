# Markdown Parser

---
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Package `4hr-markdown` looks to solve converting `Markdown` to `HTML`. Markdown is a simple syntax used to generate
formatted text. It’s used in lots of places, but the one most developers have probably encountered is README files in
GitHub.

## Table Of Contents

---

### Problem

---

Write a program which converts a small subset of markdown to HTML. This implementation can be a command-line program or
a web application. It is up to the engineer as to which is more comfortable

#### Functional Requirements

* Keep the time to around 4 hours or less
* Only looking for a small subset of Markdown to HTML formatting for conversion (listed below)

| Markdown                               | HTML                                              |
|----------------------------------------|---------------------------------------------------|
| `# Heading 1`                          | `<h1>Heading 1</h1>`                              |
| `## Heading 2`                         | `<h2>Heading 2</h2>`                              |
| `...`                                  | `...`                                             |
| `###### Heading 6`                     | `<h6>Heading 6</h6>`                              |
| `Unformatted text`                     | `<p>Unformatted text</p>`                         |
| `[Link text](https://www.example.com)` | `<a href="https://www.example.com">Link text</a>` |
| `Blank line`                           | `Ignored`                                         |

* Handle larger inputs
* Fast execution

#### Non-Functional Requirements

* **Testability**: leveraging unit tests for core logic (**check regex to ensure the transformation is correct**), add
  test markdown files to run through test and compare results
* **Functionality**: ensure it functions as expected, handle edge cases
* **Performance**: the program executes efficiently given the time constraint
* **Readability**: is the code base self expressive - easy to understand
* **Pragmatism**: ensure the program is ***reasonable*** and ***logical*** given the constraints.

## Questions

* What are potential edge cases to handle for markdown --> html ?
* whitespace - how will we handle it? Lexer - deals with the whitespace - discard it?
* Does the source md file contain different syntax's and does it need to be handled (ex. github vs generic)
* Parsing strategies - Tree --> Top-Down OR Bottom-Up

---

## Preparation

---
This package does can either be built by running `go build cmd/cli/main.go` or by running `go run cmd/cli/main.go` **
NOTE**
You must pass in the flag ` -filepath` otherwise **it will hang**.

## Usage

---

1. Clone the repository using `git clone https://jjwoz/4hr-markdown.git`
2. cd into the director `cd 4hr-markdown`
3. either run `go build cmd/cli/main.go -filepath "NAME_OF_MARKDOWN.md"`
   OR `go run cmd/cli/main.go -filepath "NAME_OF_MARKDOWN.md"`
    1. `NAME_OF_MARKDOWN.md` being the markdown file wanting to be converted
4. the html will be rendered in the terminal

### Testing

1. at the base of the project run `go test ./...`

## Approach

---
Given the time constraint as well as the small subset of Markdown, the initial idea is to implement a ***regex-based***
approach that would replace the Markdown syntax with the corresponding HTML syntax directly. However, with this approach
does achieve readability, extensibility, and pragmatism it is not optimized for the full set of markdown grammar and
only supports one output which is HTML.

### Tradeoffs

Attempting to stick to the less than 4hr constraint, There is only a `Parser` with limited testing as well only one way
to output the html.

### If only there was more time!

One thing given more time would be to split the Parser (in its entirety) into more modular components (Lexer, parser,
render). Where the Lexer would be responsible for taking the input from the user and transforming it into tokens, the
Parser would take the stream of tokens outputs from the Lexer and transform it into an Abstract Syntax Tree, Finally the
Render would take the AST and transform it into the configured output.

### Terms

---

* **Parsing**: analysis of an input to organize the data according to the rule of a grammar
* **Regular Expressions**: a sequence of characters that can be defined by a pattern
* **Grammar**: a set of rules that syntactically describes a language. Describes the structure not the meaning.
* **Lexer**: transforms a sequence of characters into a sequence of tokens. scans the input and produces the matching
  tokens - also known as (scanner or tokenizer)
* **Parser**: scans the tokens and produces the parsing result - **proper parser** is responsible for analyzing the
  tokens produced by the lexer. In some contexts "parser" may refer to the software that performs the entire process
  of "parsing"
* **Parse Tree**: also known as **Concrete Syntax Tree** reflects more concretely the syntax of the input
* **Abstract Syntax Tree**: A polished version of the parse tree, which only the information relevant to understanding
  the code is maintained
* **Tree**: a hierarchical data structure that consists of a central node, structural nodes, and sub nodes which are all
  connected via edges. Tree data structure has roots, branches and leaves connected with one another

### How "Parsing" works

Lexer(`input`(md) --> `transform`(token(s)) --> Parser (`input`(token) --> `output`(tree(ast/pt)) --> Renderer (`input`(
ast node) --> `output`(html : configuration))
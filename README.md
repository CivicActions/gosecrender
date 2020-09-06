# Secrender

## Overview

Example tool to render a template using data loaded from a JSON
file.  One intended use case: load an [OSCAL style](https://pages.nist.gov/OSCAL/documentation/schema/ssp/) YAML file and render
a Golang template to produce the markdown for SSP front matter.

## Usage

### Parameters

#### Keys

`Secrender` requires at least one JSON file in keys directory containing the key:value pairs to be used with the templates. For example:

```json
{
  "project": {
    "name": "The Project",
    "name_short": "Project",
    "contact_email": "TheProject@example.net",
  }
}
```

The JSON should have a unique top-level name, `project` in the example above. The variables will be available to your template using `{{.Keys.[TOP-LEVEL].[KEY]}}`. Using the example above to render the project name you would use `{{.Keys.project.name}}`.

You can have as many JSON key files as you like in the directory.

Pass the directory containing your _key files_ using `-k`. For example `-k keys/`.

#### Templates

Templates use the [Golang text/template](https://golang.org/pkg/text/template/) package. The files must have a `.tpl` extension and live in the `templateDir` defined in the `config.json` file. The templates variables use the "dot" notation and are surrounded by `{{ }}`.

```markdown
# Project {{.Keys.project.name}}

{{.Keys.project.name_short}} is a project that does {{.Keys.project.description}}
```

The variables will be replaced with values from the keys JSON files.

Pass the path to your _template file_ using `-t`. For example `-t templates/myTemplate.md.tpl`.

#### Output

Use the `-o` flag to define where you would like the template to be rendered. For example: `-o output/myTemplate.md`.

## Authors

* **Tom Camp** - [Tom-Camp](https://github.com/Tom-Camp)

## TODO

* Oh, so much!

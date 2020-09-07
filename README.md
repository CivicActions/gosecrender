# Go SSP Toolkit

## Overview

A tool to automate a _System Security Plan_.  One intended use case: load an [OSCAL style](https://pages.nist.gov/OSCAL/documentation/schema/ssp/) YAML file and render a Golang template to produce the markdown for SSP front matter.

## Setup

### Keys

`GoSSPTK` requires at least one YAML file in the keys directory containing the ***key:value*** pairs to be used with the templates. For example:

```yaml
project:
    name: The Project
    name_short: Project
    contact_email: TheProject@example.net
```

The YAML should have a unique top-level name, `project` in the example above. The variables will be available to your template using `{{.[TOP-LEVEL].[KEY]}}`. Using the example above to render the project name you would use `{{.project.name}}`.

The directory containing the files can live wherever you like, and can have as many YAMl key files as you like.

Define the directory containing your keys as the value _keyDir_ in the `config.yaml` file.

### Templates

Templates use the [Golang text/template](https://golang.org/pkg/text/template/) package. The files must have a `.tpl` extension and live in the _templateDir_ defined in the `config.yaml` file. The templates variables use the "dot" notation and are surrounded by `{{ }}`.

```markdown
# Project {{.project.name}}

{{.project.name_short}} is a project that does {{.Keys.project.description}}
```

The variables will be replaced with values from the keys YAML files.

## Usage

RenderAll... more info to follow.

## Authors

* **Tom Camp** - [Tom-Camp](https://github.com/Tom-Camp)

## TODO

* Oh, so much!

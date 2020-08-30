# Secrender

## Overview

Example tool to render a template using data loaded from a JSON
file.  One intended use case: load an [OSCAL style](https://pages.nist.gov/OSCAL/documentation/schema/ssp/) YAML file and render
a Golang template to produce the markdown for SSP front matter.

## Usage

### Configuration

Configuration is stored in a file name `config.json` and should be kept in your project root. If the file doesn't exist _Secrender_ will create one with default values for the _keys_, _templates_ and _output_ directory.

Example:

```json
{
  "keyDir": "keys",
  "templateDir": "templates",
  "outputDir": "out"
}
```

#### Keys

`Secrender` requires at least on JSON file in the keyDir directory defined
in the `config.json` file. The keys file contains the keys and values to be
used with the templates. For example:

```json
{
  "project": {
    "name": "The Project",
    "name_short": "Project",
    "contact_email": "TheProject@example.net",
  }
}
```

The file should have a unique top-level name, `project` in the example above. The variables will be available to your template using `{{.Keys.[TOP-LEVEL].[KEY]}}`. Using the example above to render the project name you would use `{{.Keys.project.name}}`.

You can have as many JSON key files as you like in the directory.

#### Templates

Templates use the [Golang text/template](https://golang.org/pkg/text/template/) package. The files must have a `.tpl` extension and live in the `templateDir` defined in the `config.json` file. The templates variables use the "dot" notation and are surrounded by `{{ }}`.

```markdown
# Project {{.Keys.project.name}}

{{.Keys.project.name_short}} is a project that does {{.Keys.project.description}}
```

The variables will be replaced with values from the keys JSON files.

#### Output

Rendered template files will be output to the `outputDir` defined in the `config.json` file without the `.tpl` extension, but will retain their directory structure. In other words a file that lives in `templates/somedir/anotherdir/myTemplate.md.tpl` will be written to `outputDir/somedir/anotherdir/myTemplate.md`

## Authors

* **Tom Camp** - [Tom-Camp](https://github.com/Tom-Camp)

## TODO

* Oh, so much!

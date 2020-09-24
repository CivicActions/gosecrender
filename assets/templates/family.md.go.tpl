## {{ .name }}

{{- range $i, $v := .controls }}

### {{ $v.ctrlkey }}: {{ $v.ctrlname }}

```markdown
{{ $v.description }}
```

{{- range $a, $b := $v.narratives }}

{{ if ne ( index $a )  ( "no key" ) -}}#### {{ $a | ToUpper }}{{- end }}
{{- range $d := $b }}

**{{ $d.Component }}**

{{ $d.Text }}

{{- end }}
{{- end }}
{{- end }}

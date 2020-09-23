## {{ .name }}

{{- range $i, $v := .controls }}

### {{ $v.ctrlkey }}: {{ $v.ctrlname }}

```markdown
{{ $v.description }}
```

{{- range $a, $b := $v.narratives }}

{{ if ne ( index $a )  ( "no key" ) -}}#### {{ $v.ctrlkey }} ({{ $a | ToUpper }}){{ end }}

{{- range $d := $b }}

##### {{ $d.component }}

{{ $d.text }}

{{- end }}
{{- end }}
{{- end }}

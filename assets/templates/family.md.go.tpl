## {{ .name }}

{{- range $i, $v := .controls }}

### {{ $v.ctrlkey }}: {{ $v.ctrlname }}

```markdown
{{ $v.description }}
```

{{- range $a, $b := $v.narratives }}
{{- range $j, $k := $b.key }}

{{ if $j -}}#### {{ $j }}{{ end }}

{{- range $c, $d := $k }}
{{- range $e, $f := $d.text }}

{{ $e }}

{{ $f }}

{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}

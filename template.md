<!-- BEGIN_ACTIONS_HEADER -->
# {{ .Name }}
<!-- END_ACTIONS_HEADER -->

<!-- BEGIN_ACTIONS_DESCRIPTION -->
{{ .Description }}
<!-- END_ACTIONS_DESCRIPTION -->

<!-- BEGIN_ACTIONS_INPUTS -->
## Inputs

| Name | Description | Default | Required |
|------|-------------|---------|----------|
{{- range $name, $value := .Inputs }}
| {{ $name }} | {{ $value.Description}} | `{{ val $value.Default}}` | {{ $value.Required }}
{{- end }}
<!-- END_ACTIONS_INPUTS -->

<!-- BEGIN_ACTIONS_OUTPUTS -->
## Outputs

| Name | Description |
|------|-------------|
{{- range $name, $value := .Outputs }}
| {{ $name }} | {{ $value.Description}} |
{{- end }}
<!-- END_ACTIONS_OUTPUTS -->

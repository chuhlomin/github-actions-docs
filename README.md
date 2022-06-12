# github-actions-docs

A Github action for generating docs for GitHub actions using Go template.

This action has the ability to auto commit docs to an open PR or after a push to a specific branch.

## Usage

To use `github-actions-docs` GitHub action, configure a YAML workflow file, e.g. `.github/workflows/documentation.yml`, with the following:

```yaml
name: Generate GitHub Action docs
on:
  - pull_request
jobs:
  docs:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
      with:
        ref: ${{ github.event.pull_request.head.ref }}

    - name: Render GitHub Actions docs inside the README.md and push changes back to PR branch
      uses: chuhlomin/github-actions-docs@v1
      with:
        working-dir: .
        output-file: README.md
```

> **NOTE:** If `output-file` (default README.md) already exists it will need to be updated, with the block delimeters `<!-- BEGIN_ACTIONS_DOCS -->` and `<!-- END_ACTIONS_DOCS -->`, where the generated Markdown will be injected. Otherwise the generated content will be appended at the end of the file.

> **NOTE:** If `output-file` (default README.md) doesn't exist it will created and the content will be saved into it.

---

Inspired by [terraform-docs/gh-actions](https://github.com/terraform-docs/gh-actions).

Alternative: [npalm/action-docs](https://github.com/npalm/action-docs).

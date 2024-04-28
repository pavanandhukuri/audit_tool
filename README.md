# Security Audit Tool

A simple rule evaluator tool that can be used to evaluate security settings. At the moment, the tool supports evaluating
github repository configuraiton. But it can be extended to support more.

## Usage

Create a copy of `.env.example` and rename it to `.env`. Fill in the required values. Alternatively, environment
variables can be set.

```bash
GITHUB_PERSONAL_ACCESS_TOKEN=generate_personal_access_token
GITHUB_ORG_NAME=your_org_name
```

Run the tool using the following command:

```bash
go run main.go
```

## Extending the tool

`VersionControlData` entity present in `domain/entities/versionControl.go` shows what properties are supported by the
tool as of now. https://github.com/go-playground/validator is used to run rules against the data represented by that
entity. Additional properties can be added by adding the relevant fetching code to the adapters
implementing `VersionControlPort`. Rules can then be added to `resources/version_control_system_rues.yml` to be
evaluated.

## Testing

Mocks for the ports are generated using https://github.com/vektra/mockery. If more ports are added, mocks can be
generated using the following command

```bash
mockery --all
```

To run tests, use the following command:

```bash
go test ./...
```

To view coverage, use

```bash
go test -cover ./...
```
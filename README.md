# Security Audit Tool

A simple rule evaluator tool that can be used to audit the configuration of a GitHub organisation. The tool is designed
to be extensible and can be easily extended to add more rules.

## Configuration

The tool is configured using two files:

* `resources/schema_config.yml`: This file contains the schema definition for the data model to be created based by
  calling different endpoints of GitHub. The schema is defined in YAML format.
* `resources/version_control_system_rules.yml`: This file contains the rules that need to be evaluated. The rule
  operations are inherited from https://github.com/go-playground/validator. Custom validators are defined
  in `adapters/ruleEvaluator/customValidators.go` and linked in `adapters/ruleEvaluator/validatorService.go`.

New rules can be added by adding them to the `resources/version_control_system_rules.yml` file. If they require a
modification to the schema, it needs to be first updated in the `resources/schema_config.yml` file.

## Usage

**Step 1:** Create a copy of `.env.example` and rename it to `.env`. Fill in the required values. Alternatively,
environment
variables can be set.

```bash
GITHUB_PERSONAL_ACCESS_TOKEN=generate_personal_access_token
GITHUB_ORG_NAME=your_org_name
```

Run the tool using the following command:

```bash
go run main.go
```

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
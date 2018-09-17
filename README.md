# terraform-provider-bugsnag

Simple provider for terraform to read Bugsnag project information. This is useful for injecting
project API keys into deployments/configuration.

## Configuration

Set environment variable BUGSNAG_ACCESS_TOKEN with a Bugsnag personal access token.

Alternatively you can set the "credentials" parameter in the bugsnag provider block.

## Use

```
data "bugsnag_project" "my_project" {
    project_id = "515fb9337c1074f6fd000001"
}
```

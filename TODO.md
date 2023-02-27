# Things that need doing

## Security

* Switch to federated identity management: 
    * https://cloud.google.com/iam/docs/workload-identity-federation
    * https://cloud.google.com/blog/products/identity-security/how-to-authenticate-service-accounts-to-help-keep-applications-secure
* Secure Service Accounts

    >Important: When a default service account is created, it is automatically granted the Editor role (roles/editor) on your project. This role includes a very large number of permissions. To follow the principle of least privilege, we strongly recommend that you either disable the automatic role grant by adding a constraint to your organization policy, or revoke the Editor role manually. If you disable or revoke the role grant, you must decide which roles to grant to the default service accounts. [1](https://cloud.google.com/iam/docs/service-accounts)
* Enable Dependabot version updates 
    * dependabot.yml

        >To get started with Dependabot version updates, you'll need to specify which package ecosystems to update and where the package manifests are located. Please see the documentation for all configuration options: https://docs.github.com/github/administering-a-repository/configuration-options-for-dependency-updates

## Architecture

* Switch to Cloud Run

## Testing

* Add Integration Test

## Infrastructure

* Deploy project using Pulumi
* Implement InSpec 

## Observabilty

* Setup metrics and alerts for infrastructure and application

## Path to Prod

* Unit test gate
* Integration test gate
* A/B deployments
* Rollbacks

## Ops

* Health checks

## Process

* [todo-issue](https://github.com/marketplace/actions/todo-issue) GitHub Action

## Code Quality

* TODO Add Linter(s)
* Investigate and adopt a linter for Go, CSS, HTML
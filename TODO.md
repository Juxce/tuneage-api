# Things that need doing

## Security

* TODO Implement API Authentication
* Require authentification to access the API.

* TODO Switch to federated identity management: 
* Review the docs and implement federated identity management:
* [Identity Security](https://cloud.google.com/iam/docs/workload-identity-federation)
* [Service Account Authentication](https://cloud.google.com/blog/products/identity-security/how-to-authenticate-service-accounts-to-help-keep-applications-secure)

* Secure Service Accounts
* Determine least priviledge for all service accounts.
* "Important: When a default service account is created, it is automatically granted the Editor role (roles/editor) on your project. This role includes a very large number of permissions. To follow the principle of least privilege, we strongly recommend that you either disable the automatic role grant by adding a constraint to your organization policy, or revoke the Editor role manually. If you disable or revoke the role grant, you must decide which roles to grant to the default service accounts. [1](https://cloud.google.com/iam/docs/service-accounts)"

* TODO Enable Dependabot version updates 
* Make the the modifications to dependabot.yml needed to enable the full automation of version updates.
* "To get started with Dependabot version updates, you'll need to specify which package ecosystems to update and where the package manifests are located. Please see the documentation for all configuration options: https://docs.github.com/github/administering-a-repository/configuration-options-for-dependency-updates"

* TODO Rotate GCP service account keys
* Use the Github "Rotate GCP service account keys" action to rotate gcp service account keys and update github secrets. Determine the appropriate rotation schedule (deployment, day, week, month).

## Architecture

* TODO Spike deploying into Cloud Run
* Determine if Cloud Run is a better hosting platform.

## Testing

* TODO Load Testing
* Have the ability to apply a load to the API to determine how much load can be serviced and for testing infrastructure alerting.

## Infrastructure

* TODO Deploy project using Pulumi
* Use Pulumi to provision and deploy the Tunage API

* TODO Evaluate InSpec 
* Spike InSpec for IaC testing. 

## Observabilty

* TODO Setup metrics and alerts for infrastructure and application
* Have metrics for infrastructure that alert the team when the provisioned infrastructure approaches configured limits for memory, cpu, storage, bandwidth, etc.
* Have metrics for the API that support identifying usage patterns,

## Path to Prod

* TODO Deploy prod stack to tunage-api gcp project
* Create a new poc stack and deploy the prod stack to a new gcp project, tuneage-api. 
* Have automation for all actions required to standup the tuneage-api project. 
* Verify that the automation works by provisioning a new project and deploy the stack with no errors and no manual intervention. Use temporary projects until the automation is fully vetted. 
* Destroy all temporary projects when nolonger needed.

* TODO Unit test gate
* Have unit tests that all functionallity implemented is correct, and that error conditions are handled properly.

* TODO Integration test gate
* Have tests that run against a deployed stack that verify that the deployed infrastructure and application is working end-to-end.

* TODO A/B deployments
* Have two production stacks. When A is live, deploy to B, test B and then make B live

* TODO Rollback to previous version 
* If automated deployment verification tests fail switch back to previous version

## Ops

* TODO Health checks
* Contiuously verify that the API is up and returning data

## Process

* [todo-issue](https://github.com/marketplace/actions/todo-issue) GitHub Action

## Code Quality

* TODO Add Linter(s)
* Investigate and adopt a linter for Go, CSS, HTML
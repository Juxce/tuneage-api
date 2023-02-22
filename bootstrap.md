# Getting setup to contribute to the Tuneage API

* [Go](#go)
* [IDE](#ide)
* [gcloud cli](#gcloud)
* [GitHub Workflow](#github)

## <a name="go"></a>Go

The API is written in [Go](https://github.com/topics/golang) version 1.20. The latest release is available [here](https://go.dev/dl/). As Go is a cross platform development language you can work on Linux, Mac OS, Windows, or WSL. The Tuneage API is currently deployed on Ubuntu in GCP.

### Additional Tools

* TBD

## <a name="ide"></a>IDE

The API is IDE agnostic, but you may be more effective if you configure your IDE to use the reccomended extensions.

### Visual Studio Code Extensions

* [GitHub Pull Requests and Issues](https://marketplace.visualstudio.com/items?itemName=GitHub.vscode-pull-request-github)
* [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)
* [Google Cloud Code](https://marketplace.visualstudio.com/items?itemName=GoogleCloudTools.cloudcode)

## <a name="gcloud"></a>gcloud cli

The API is deployed in the Google Cloud. To work effectively with GCP you should have [gcloud cli](https://cloud.google.com/sdk/docs/install) installed.

## <a name="github"></a>GitHub Workflow

The API uses a continuous deployment workflow to deploy to GCP. See `.github\workflows\main.yml`


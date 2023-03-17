# :warning: caution

This deploys resources to Azure App Service that will churn (start, fail, stop, repeat) and cost money while doing so. Only run this if you are actively troubleshooting an issue. 

Be sure to delete the resources when you are done troubleshooting.

# To deploy

build the and publish the docker image:

```shell
docker build -t YOUR_IMAGE_TAG .
docker push YOUR_IMAGE_TAG
```

create a resource group

```shell
az group create -l eastus -n YOUR_RG_NAME
```

Update `arm-template/appservice-parameters.json` with the required values

```shell
az deployment group create --resource-group YOUR_RG_NAME  --template-file appservice-template.json --parameters @appservice-parameters.json
```

Destroy app services resources:

```shell
az deployment group deeete --resource-group YOUR_RG_NAME  --name YOUR_APPSERVICE_NAME
```

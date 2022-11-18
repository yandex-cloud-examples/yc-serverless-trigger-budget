# Шаг 1. Создание Service Account
## Создание сервисного аккаунта

Создайте сервисный аккаунт с именем `service-account-for-budget`:

    export SERVICE_ACCOUNT=$(yc iam service-account create \
    --name service-account-for-budget \
    --description "service account for buget" \
    --format json | jq -r .)

Проверьте текущий список сервисных аккаунтов:

    yc iam service-account list
    echo $SERVICE_ACCOUNT

После проверки запишите ID, созданного сервисного аккаунта, в переменную `SERVICE_ACCOUNT_ID`:

    echo "export SERVICE_ACCOUNT_ID=<ID>" >> ~/.bashrc && . ~/.bashrc 
    echo $SERVICE_ACCOUNT_ID

## Назначение роли сервисному аккаунту

Добавим вновь созданному сервисному аккаунту необходимые роли `compute.admin`,  `iam.serviceAccounts.user`, `serverless.functions.invoker`: 

    echo "export FOLDER_ID=$(yc config get folder-id)" >> ~/.bashrc && . ~/.bashrc
    echo $FOLDER_ID

    yc resource-manager folder add-access-binding $FOLDER_ID \
    --subject serviceAccount:$SERVICE_ACCOUNT_ID \
    --role compute.admin

    yc resource-manager folder add-access-binding $FOLDER_ID \
    --subject serviceAccount:$SERVICE_ACCOUNT_ID \
    --role iam.serviceAccounts.user

    yc resource-manager folder add-access-binding $FOLDER_ID \
    --subject serviceAccount:$SERVICE_ACCOUNT_ID \
    --role serverless.functions.invoker

Все перечисленные роли можно заменить ролью `editor`:

    yc resource-manager folder add-access-binding $FOLDER_ID \
    --subject serviceAccount:$SERVICE_ACCOUNT_ID \
    --role editor

## Видео

https://youtu.be/q8LJ0MfUhV4
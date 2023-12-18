# Шаг 2. Создание экземпляра Serverless Function

Мы будем использовать версию рантайма `golang119`. При создании функции нам понадобится задать две переменные `FOLDER_ID` и `TAG`, первая задаст каталог в котором надо будет останавливать виртуальные машины, а вторая тег, который должен быть у виртуальной машины которую мы хотим останавливать автоматически при достижении некоторого порога в бюджете. 

Находясь в директории с исходными файлами, упакуем все нужные нам файлы в zip-архив:

    zip src.zip index.go go.mod 

Создадим нашу функцию `function-for-budget`, при этом сразу зададим сервисный аккаунт:

    yc serverless function create \
    --name function-for-budget \
    --description "function function-for-budget for trigger example"

    yc serverless function version create \
    --function-name function-for-budget \
    --memory=512m \
    --execution-timeout=5s \
    --runtime=golang119 \
    --entrypoint=index.StopComputeInstances \
    --service-account-id=$SERVICE_ACCOUNT_ID \
    --environment FOLDER_ID=$FOLDER_ID \
    --environment TAG=target-for-stop \
    --source-path=src.zip

    yc serverless function allow-unauthenticated-invoke function-for-budget

Если вызвать нашу функцию, то в указанном в переменной фолдере будут проверены все виртуальные машины. Если у виртуальной машины будет `label` `target-for-stop` равный `true` она будет остановлена.

## Видео

https://youtu.be/NIbph15KLqs
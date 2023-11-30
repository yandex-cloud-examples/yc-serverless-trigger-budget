# Шаг 4. Создание экземпляра Serverless Function для обработки триггера для бюджета

Находясь в директории с исходными файлами, упакуем все нужные нам файлы в zip-архив:

    zip -FS src.zip budget_trigger_handler.go go.mod

Создадим функцию `budget-trigger-handler`, при этом зададим сервисный аккаунт и нужные переменные окружения:

    yc serverless function create \
    --name budget-trigger-handler \
    --description "budget trigger handler"

    yc serverless function version create \
    --function-name budget-trigger-handler \
    --memory=512m \
    --execution-timeout=5s \
    --runtime=golang119 \
    --entrypoint=budget_trigger_handler.Handler \
    --service-account-id=$SERVICE_ACCOUNT_ID \
    --environment FOLDER_ID=$FOLDER_ID \
    --environment TAG=target-for-stop \
    --environment AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
    --environment AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
    --environment BUDGET_QUEUE_URL=$BUDGET_QUEUE_URL \
    --source-path=src.zip
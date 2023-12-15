# Шаг 8. Создание триггера для Message Queue

Вы можете создать триггер согласно этой [инструкции](https://cloud.yandex.ru/docs/functions/operations/trigger/ymq-trigger-create). Введите имя тригера `budget-queue-trigger`, тип триггера `Message Queue`, очередь сообщений `budget-queue`, функцию `budget-queue-handler`, для Message Queue и фукции укажите сервисный аккаунт созданный ранее.

Также создание триггера можно произвести из CLI:

    yc serverless trigger create message-queue \
        --name budget-queue-trigger \
        --queue $BUDGET_QUEUE_ID \
        --queue-service-account-id $SERVICE_ACCOUNT_ID \
        --invoke-function-id $BUDGET_QUEUE_HANDLER_FUNCTION_ID \
        --invoke-function-service-account-id $SERVICE_ACCOUNT_ID \
        --batch-size 1 \
        --batch-cutoff 10s

Проверьте наличие созданного тригера:

    yc serverless trigger list
# Шаг 7. Создание экземпляра функции для триггера Message Queue

По аналогии с предыдущей функцией создайте функцию для обработки сообщений, которые приходят в Message Queue.

Находясь в директории с исходными файлами, упакуйте все нужные нам файлы в zip-архив:

    zip -FS src.zip budget_queue_handler.go utils.go go.mod

    yc serverless function create \
        --name budget-queue-handler \
        --description "handles messages from budge-message-queue"

Запишите идентификатор созданной функции в переменную окружения:
    echo "export BUDGET_QUEUE_HANDLER_FUNCTION_ID=<function_id>" >> ~/.bashrc && . ~/.bashrc 

Создайте экземпляр функции с переменной окружения TELEGRAM_BOT_API_TOKEN и TELEGRAM_BOT_CHAT_ID.

    yc serverless function version create \
        --function-name budget-queue-handler \
        --memory 512m \
        --execution-timeout 5s \
        --runtime golang119 \
        --entrypoint budget_queue_handler.HandleBudgetQueueMessage \
        --service-account-id $SERVICE_ACCOUNT_ID \
        --environment TELEGRAM_BOT_API_TOKEN=$TELEGRAM_BOT_API_TOKEN \
        --environment TELEGRAM_BOT_CHAT_ID=$TELEGRAM_BOT_CHAT_ID \
        --source-path src.zip
# Шаг 3. Создание Message Queue

Создайте очередь используя CLI:

    echo "export BUDGET_QUEUE_URL=$(aws sqs create-queue \
    --queue-name budget-queue \
    --endpoint https://message-queue.api.cloud.yandex.net/ | jq -r ".QueueUrl")" >> ~/.bashrc && . ~/.bashrc

Также можно создать очередь в консоли управления используя [эту инструкцию](https://cloud.yandex.ru/docs/message-queue/operations/message-queue-new-queue) указав имя очереди `budget-queue`, тип `Стандартная`. Открыв созданную очередь и перейдя во вкладку `Обзор` в блоке `Общая информация` скопируйте URL очереди и сохраните в переменную окружения:

    echo "export BUDGET_QUEUE_URL=<queue_url>" >> ~/.bashrc && . ~/.bashrc

Проверьте текущий список очередей:

    aws sqs list-queues --endpoint https://message-queue.api.cloud.yandex.net/
    
Сохраните идентификатор очереди:

    echo "export BUDGET_QUEUE_ID=$(aws sqs get-queue-attributes \
    --queue-url $BUDGET_QUEUE_URL \
    --attribute-names QueueArn \
    --endpoint https://message-queue.api.cloud.yandex.net/ | jq -r ".Attributes.QueueArn")" >> ~/.bashrc && . ~/.bashrc

# Шаг 5. Создание Триггера

Пока вы находитесь в web-интерфейсе Yandex Cloud вы можете создать триггер согласно этой [инструкции](https://cloud.yandex.ru/docs/functions/operations/trigger/budget-trigger-create). Выбирайте тип триггера `Бюджет`, указывайте свой платежный аккаунт и бюджет созданный на прерыдущем шаге. В качестве функции выберете функцию `function-for-budget` и в качестве сервисного аккаунта `service-account-for-budget`.   

После создания триггера в веб-интерфейсе можно проверить его видимость в CLI:

    yc serverless trigger list

Если в следующую команду подставить `ID` созданного триггера то вы получите описание триггера. Среди параметров вы увидите  `billing_account_id` и `budget_id`, которые были использованы вами

    yc serverless trigger get <ID>

Также создание триггера можно произвести из CLI. Для этого подставьте в следующую команду `billing_account_id` и `budget_id`:

    yc serverless trigger create billing-budget \
    --name trigger-for-budget-from-yc \
    --invoke-function-name budget-trigger-handler \
    --invoke-function-service-account-id $SERVICE_ACCOUNT_ID  \
    --billing-account-id <id_вашего_биллинг_аккаута>  \
    --budget-id <id_вашего_бюджета>

    yc serverless trigger list

Но для работы нашего сценария два триггера не нужны, оставьте один. Удаление вы можете выполнить как в CLI, так и в web-интерфейсе. 

## Видео

https://youtu.be/XZ6tRHhG37w
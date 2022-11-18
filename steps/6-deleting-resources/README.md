# Шаг 6. Удаление ресурсов

По завершению работы удалите все используемые вами ресурсы.

Остановите и удалите экземпляры виртуальных машин:

    yc compute instance stop target-instance-3
    yc compute instance delete --name=target-instance-1
    yc compute instance delete --name=target-instance-2
    yc compute instance delete --name=target-instance-3

Удалите бюджет в веб-интерфейсе, триггер и экземпляр функции:

    yc serverless trigger list
    yc serverless trigger delete --name=trigger-for-budget
    yc serverless trigger delete --name=trigger-for-budget-from-yc

    yc serverless function list
    yc serverless function delete --name=function-for-budget

Удалите сервисный аккаунт:

    yc iam service-account delete --id $SERVICE_ACCOUNT_ID

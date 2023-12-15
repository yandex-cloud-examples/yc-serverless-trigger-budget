# Шаг 9. Создание Serverless Container из Docker-образа

Создадим несколько виртуальных машин, чтобы имитировать потребление в облаке. Для этого нам потребуется уточнить значение сетей и подсетей.

    yc vpc network list
    yc vpc subnet list

Создадим три виртуальные машины `target-instance-1`, `target-instance-2` и `target-instance-3`. Сделаем это в зоне `ru-central1-a` и в подсети `default-ru-central1-a` (возможно вам прийдется подставить ваши значения). Укажем `labels` для первых двух виртуальных машин `target-for-stop=true` , а для третьей `target-for-stop=false`. Значение  `labels` подсветят какие виртуальные машины надо остановить.

    yc compute instance create \
    --name target-instance-1 \
    --labels target-for-stop=true \
    --zone ru-central1-a \
    --network-interface subnet-name=default-ru-central1-a,nat-ip-version=ipv4 \
    --create-boot-disk image-folder-id=standard-images,image-family=ubuntu-2004-lts \
    --ssh-key ~/.ssh/id_ed25519.pub

    yc compute instance create \
    --name target-instance-2 \
    --labels target-for-stop=true \
    --zone ru-central1-a \
    --network-interface subnet-name=default-ru-central1-a,nat-ip-version=ipv4 \
    --create-boot-disk image-folder-id=standard-images,image-family=ubuntu-2004-lts \
    --ssh-key ~/.ssh/id_ed25519.pub

    yc compute instance create \
    --name target-instance-3 \
    --labels target-for-stop=false \
    --zone ru-central1-a \
    --network-interface subnet-name=default-ru-central1-a,nat-ip-version=ipv4 \
    --create-boot-disk image-folder-id=standard-images,image-family=ubuntu-2004-lts \
    --ssh-key ~/.ssh/id_ed25519.pub

    yc compute instance list

После этого следует подождать пока потребление в каталоге перевалит пороговое значение, сработает триггер, запуститься функция и часть виртуальных машин будет остановлена. Возможно вам захочется ускорить процесс, можете создать больше виртуальных машин или другой бюджет.

Помните, что `target-instance-3` продолжает работать, перейдите к следующему шагу и удалите все созданные в рамках сценария ресурсы.

## Видео

https://youtu.be/lsYml46Eydg
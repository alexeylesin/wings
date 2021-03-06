# LesinPanel Wings

Wings - демон для LesinPanel. Он написан на Go, основан на [Pterodactyl](https://github.com/pterodactyl/panel), а также устанавливается на все ноды, которые подключаются к панели и служит для их корректной работы. Для подключения демона к панели, вам **потребуется** уникальный конфиг, созданный под вашу конфигурацию сервера, который я вам предоставлю, если скажу установить этот скрипт.

# Версии

Демон имеет собственную систему версий, на данный момент это **0.0.1**, основанная на **1.3.1**.

# Установка

Для скачивания и установки демона, используйте данную команду:
```
$ bash <(curl -s https://cdn.alexeylesin.me/panel/deploy.sh)
```

# Сборка

Если вы хотите осуществить сборку демона вручную, вам понадобится установить [Go](https://golang.org/doc/install). После этого, вы можете перейти в специальную папку и скачать туда этот репозиторий с GitHub:

```
$ cd /srv && git clone https://github.com/alexeylesin/wings && cd wings
```

Теперь вы можете приступать к сборке:

```
$ cd /srv/wings/ && go build -o /usr/local/bin/wings && chmod +x /usr/local/bin/wings
```

После этого, вы можете запустить демон при помощи **systemd** или через данную команду:

```
$ sudo wings --debug
```

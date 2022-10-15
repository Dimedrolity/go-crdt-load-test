# Нагрузка счетчиков gcounter-crdt

Нагрузка счетчика [gcounter-crdt](https://github.com/Dimedrolity/gcounter-crdt) с помощью отправки HTTP-запросов.



# TODO

Параметры:
- Политика обращения к серверам. Round-robin/random.
- Кол-во вызовов Count.
- IncPerCount.

Сначала IncPerCount вызовов inc, потом вызов Count.

Вывод статистики по времени запросов, к какому хосту и сколько времени затрачено. Также среднее и медиана. Можно вывод в файл.
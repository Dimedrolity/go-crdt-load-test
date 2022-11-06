# Нагрузка счетчиков gcounter-crdt

Нагрузка счетчика [gcounter-crdt](https://github.com/Dimedrolity/gcounter-crdt) с помощью отправки HTTP-запросов.

Параметры config.yml:
- Политика обращения к серверам. 
  - [x] Round-robin (по умолчанию)
  - [ ] Random
- [x] Кол-во вызовов Count
- [x] IncPerCount

Сначала IncPerCount вызовов inc, потом вызов Count, и так CountsCount раз.

Вывод статистики по времени запросов, к какому хосту и сколько времени затрачено в файл.

Статистика:
- [x] Среднее
- [x] 1, 2, 3 квартили


## TODO

нормальный логгер
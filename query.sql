1. Выборки всех уникальных eventType у которых более 1000 событий.

SELECT COUNT(*) as c, eventType FROM events group by eventType HAVING c > 1000

2. Выборки событий которые произошли в первый день каждого месяца.

SELECT *, toDayOfMonth(eventTime) as day FROM events where day = 1

3. Выборки пользователей которые совершили более 3 различных eventType.

SELECT COUNT(DISTINCT eventType) as et, userID FROM events group by userID HAVING et > 3



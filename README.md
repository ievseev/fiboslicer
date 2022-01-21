# Fibonacci slicer

В сервисе реализован POST метод, позволяющий получить срез из последовательности Фибоначчи.
В теле запроса необходимо указать индексы границ среза: переменные x и y. Предполагается, что индексация начинается с 0 элемента. 

Для запуска приложения можно использовать Docker. Из корня репозитория необходимо выполнить:
```
docker build . -t fibonacci_slicer
docker run fibonacci_slicer
```

Пример запроса среза последовательности Фибоначчи:

```
curl --location --request POST 'http://localhost:80/api/fibonacci-slice' \
--header 'Content-Type: application/json' \
--data-raw '{"x":5,
"y":20}'
```

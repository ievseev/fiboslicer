# Fiboslicer

Сервис поддерживает grpc и http запросы, слушает 80й и 81й порт соответственно. Реализован grpc метод GetFibonacciSlice и соответствующий ему
POST метод get-fibonacci-slice, позволяющий получить срез из последовательности Фибоначчи.
В теле запроса необходимо указать индексы границ среза: переменные x и y. Предполагается, что индексация начинается с 0 элемента. 

Для запуска приложения можно использовать Docker. Из корня репозитория необходимо выполнить:
```
docker build . -t fiboslicer
docker run fiboslicer
```

Пример запроса среза последовательности Фибоначчи:

```
curl --location --request POST 'http://localhost:81/get-fibonacci-slice' \
--header 'Content-Type: application/json' \
--data-raw '{"x":5,
"y":20}'
```

Реализовать контейнер - аналог OrderedDict в python.
Операции C++/Java/Go:
get(key)
add(key, value)
delete(key)
keys() - список ключей, отсортированный по времени добавления.

При реализации можно использовать существующие библиотечные ассоциативные
контейнеры без учета порядка (dict, unordered_map, HashMap и тд).
Сложность удаления/добавления - константа

Обновление значения существующего ключа на порядок не влияет никак
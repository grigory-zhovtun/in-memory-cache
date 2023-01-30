# Добро пожаловать в пакет In-Memory Cache!

Этот пакет предназначен для использования стандартных функций кэширования в памяти.

## Установка

Для установки используйте go get:

```
$ go get -u github.com/username/in-memory-cache
```

## Использование

Чтобы начать использовать пакет, импортируйте его в ваш проект:

```
import "github.com/username/in-memory-cache"
```

Далее создайте новый экземпляр кеша:

```
cache := cache.New()
```

Теперь вы можете добавлять значения в кеш, извлекать их и удалять их:

```
cache.Set("userId", 42)
userId := cache.Get("userId")
cache.Delete("userId")
```
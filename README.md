# 🎯 Guess the Number

Консольная игра на Go: угадай секретное число за ограниченное число попыток.

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-green)

## Возможности

- Три уровня сложности: Easy (1–50, 15 попыток), Medium (1–100, 10), Hard (1–200, 5)
- Подсказки «больше / меньше» и «горячо / тепло / холодно»
- Проверка ввода: нечисловые значения не тратят попытку
- История попыток в конце партии
- Повтор игры без перезапуска
- Сохранение результата (дата, исход, число попыток) в `results.json`
- Цветной вывод в терминале

## Запуск

```bash
git clone https://github.com/hakurai-cmd/guess-the-number.git
cd guess-the-number
go run .
```

## Как устроено

| Функция | Что делает |
|---|---|
| `choise` | меню и выбор сложности |
| `generateNumber` | генерирует секретное число |
| `readNumber` | читает и проверяет ввод |
| `compare` | сравнивает с секретным числом |
| `hint` | подсказка по расстоянию |
| `history` | выводит историю попыток |
| `saveResult` | пишет результат в JSON |

## Стек

Go, стандартная библиотека (`bufio`, `strconv`, `math/rand/v2`, `encoding/json`, `time`), [fatih/color](https://github.com/fatih/color)

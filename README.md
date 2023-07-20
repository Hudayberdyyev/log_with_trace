# Логирование и трассировка ошибок
В данном репозитории показывается пример того как можно логировать с возможностью трассировки ошибок.</br>
В качестве примера был взят, сервис у которого есть два эндпоинта:
1. **/phrase-of-the-day** - Возвращает случайную знаменитую фразу и его автора. 
2. **/greet** - Возвращает сообщение приветствия (тестовый эндпоинт, наподобие пинг-а).

### Сущности:
- **domain** - описываются доменные сущности(интерфейсы и модели).
- **handler** - обработчик запросов.
- **repository** - репозиторий который работает с хранилищем. 

### Заметки:
- В папке **data** храниться файл **phrases.data**. Внутри этого файла находиться фразы знаменитых людей. В нашем примере этот файл будут считаться хранилищем и репозиторий будет работать именно с ним.

### Пакеты:
- **github.com/rs/zerolog** - пакет логирования, с возможностью установки глобального уровня логирования.
- **github.com/pkg/errors** - пакет обработки ошибок, с возможностью трассировать ошибки.   


# Typo Настройки

---

[ ru ]/ [eng](..%2Feng%2Fsettings.md)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[/](..%2F..%2FREADME.ru.md) настройки

---

## Функция, указывающая каким словом выводить null в виде строки

`SetStringToWriteNull(str string)`

Вызывая данную функцию, у нас есть возможность указать, каким образом мы будем печатать null. По умолчанию, стоит вывод 
как `null`

---

## Функция, указывающая имя тега поля структуры, из которого мы пытаемся получить имя параметра для сопоставления

`SetStructureTagContainingNameTheElementJson(name string)`

По умолчанию, данный параметр установлен как `json`

---

## Функция, задающая список служебных слов, которые могут быть указаны в теге, и не являются названием

`SetSpecialStringsInStructureTag(words []string)`

По умолчанию данный список состоит из 2 слов.  `[]string{"-", "omitempty"}`

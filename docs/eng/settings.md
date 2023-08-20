
# Typo Настройки

---

[ eng ] / [ ru ](..%2Fru%2Fsettings.md)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[/](..%2F..%2FREADME.md) настройки

---

<style>
table {
    padding: 0;
    margin: 0;
}
td {
    padding: 0;
    vertical-align: top;
}
th {
    vertical-align: top;
    text-align: left;
    padding: 0;
}
tr {
    padding: 0;
    margin: 0;
}
td div {
    text-align: right;
}
</style>

## Function indicating which word to display null as a string

`SetStringToWriteNull(str string)`

By calling this function, we have the opportunity to specify how we will print null. By default, the output is set to `null`

---

## A function that specifies the tag name of the structure field from which we are trying to get the parameter name to match

`SetStructureTagContainingNameTheElementJson(name string)`

By default, this parameter is set to `json`

---

## A function that specifies a list of service words that can be specified in a tag, and are not a name.

`SetSpecialStringsInStructureTag(words []string)`

By default, this list consists of 2 words. `[]string{"-", "omitempty"}`
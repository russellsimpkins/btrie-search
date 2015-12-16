### Btrie Search

The name might not be the best. This program lets you search a file of strings for a file of strings.

``` go run trie.go -i <input file> -s <search file> ```

```<input file>``` should be a list of strings to search through that are new line separated e.g.

```
/foo/bar/baz
/some/other/one
/and/yet/another
this.is.another
```

```<search file>``` is also a list of strings to search the input file e.g.

```
this.is.another
search/for/me
/some/other/one
```

The program will print out each entry in <search file> with the match true false as a csv e.g.

```
this.is.another, true
search/for/me, false
/some/other/one, true
```


The code is a basic btrie data structure.

TODO - add more tests.

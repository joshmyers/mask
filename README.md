mask
====

### Usage

```
$ mask -h

Usage: mask [OPTIONS] COMMAND [ARG...]

Run command with masked stdout

Options:
  -h	show help
  -s string
    	set secret word
  -v	show version

```

### Example

```
$ mask echo hoge fuga piyo hogera hogehoge
Enter secret word: hoge
**** fuga piyo ****ra ********

$ mask -s hoge echo hoge fuga piyo hogera hogehoge
**** fuga piyo ****ra ********

$ echo hoge fuga piyo hogera hogehoge | mask -s hoge
**** fuga piyo ****ra ********

$ (read -s -p "Enter secret word: " secret && echo && echo hoge fuga piyo hogera hogehoge | ./mask -s $secret)
Enter secret word: hoge
**** fuga piyo ****ra ********

$ mask bash
Enter secret word: hoge
$ echo hoge fuga piyo hogera hogehoge
**** fuga piyo ****ra ********
```

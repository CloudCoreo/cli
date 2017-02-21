# Using Coreo CLI tool with autocomplete

You can generate a autocomplete script with `coreo`. Here is how:
```
$ coreo completion > coreo-autocomplete.sh
```
This command will generate a shell script named `coreo-autocomplete.sh`. You'll need to install `bash-completion` in your system to enable it.

## Installing bash-completion
### Linux ( apt-get )
```
$ sudo apt-get install bash-completion
```
### Linux ( yum )
```
$ sudo yum install bash-completion
```
### Mac
```
$ brew install bash-completion
```

## Place `coreo-autocomplete.sh` in `bash_completion.d`
### Linux
```
$ mv coreo-autocomplete.sh /etc/bash_completion.d/
```
### Mac
```
$ mv coreo-autocomplete.sh $(brew --prefix)/etc/bash_completion.d/
```

Good to go. On your next login, you should be able autocomplete `coreo` commands:
```
$ coreo [TAB] [TAB]
cloud      composite  configure  git-key    plan       team       token      version
$ coreo version
```
Happy tabbing...
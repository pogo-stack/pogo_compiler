# pogo_compiler

```
go get github.com/pogo-stack/pogo_compiler
```


## `.bash_aliases` sample to manage go environment and DB connection strings

We use environmental variables to tell compiler what database to use or what to do.

```
cd () { 
    if builtin cd "$@"
    then
        export GOPATH=$HOME/go
        export POGO_FLAG_ENV='-e d'

        if [[ "$PWD" =~ /pogo_project(/|$) ]]
        then
            export POGO_ENV=pogo3dev
            export GOPATH=$HOME/go:~/pogo_project
            export PSQL_HOST=localhost
            export PSQL_DB=db_pogo3
            export PSQL_USER=pogo3_user
            export PSQL_PASS=pogo3_password
            export PSQL_SCHEMA=pogo3
            export POGO_FLAG_ENV='-e d'
            export POGO_FLAG_DEBUG='-debug'
        fi

        return 0
    else
        return $?
    fi
}

```
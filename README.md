# SillyNames
A silly names generator written in Go

Build and execute

```
> .\sillynames.exe -?

flag provided but not defined: -?
Usage of X:\Dev Work\golang\Projects\SillyNames\sillynames.exe:
  -h    Prefix honorific
  -n int
        <int> - number of entries (default 5)
  -t string
        <string> - { acme | band | business | character | drug | eatery | fantasy | morpheme | team } (default "character")
```

Default character
```          
.\sillynames.exe
Lyn Xeno
Xannon Lotto
Queenie Sketch
Fanarina Varqua
Heaven Clarity
```

Different data set
```
> .\sillynames.exe -t drug
Carprexar
Curdrize
Akamalole
Soluyxana
iNcilgen
```

Default with honorifcs
```
> .\sillynames.exe -h
          
Cllr Clovie Snow
Viscount Zeam Lainey
Major Cloud Tawny
Major Torie Fostre
Ms Groover Vicki
```

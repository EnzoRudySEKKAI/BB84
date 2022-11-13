## BB84 ##
### building the executable ###
```
$ go build
```

### Launch ###
```
$ ./BB84 
```

### Usage ###

List of arguments
```
$ ./BB84
```

#### Server-client mode ####
```
$ ./BB84 server
```

```
$ ./BB84 client
```

#### Tests mode ####
Test of the 75% (Alice and Bob)
```
$ ./BB84 75
```

Test of the 62% (Alice and Bob with Eve in the middle)
```
$ ./BB84 62
```

Test with 4 users (Alice and Bob with 2 people in the middle)
```
$ ./BB84 4users
```

Test with N users (at least 2)
```
$ ./BB84 Nusers
```

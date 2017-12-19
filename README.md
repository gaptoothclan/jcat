# Jcat

A simple golang command line application to investigate json files

# Compile
```go build jcat```

Then move to where ever you need jcat

# Usage
```jcat package.json```

Will list the whole file

```jcat package.json scripts```
Will list the sub json of package.json

```jcat -k package.json``` 
Will list all of the keys in the json

```cat ./package.json | jcat scripts```

Will attempt to read the file from stdin also works with curl
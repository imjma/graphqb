# graphqb
GraphQL Query Builder

## Example

### Query

```
q := &Query{
  Type: "query",
  Name: "HeroNameAndFriends",
}

q.SetFields(
  NewField("hero").SetFields(
    NewField("name"),
    NewField("friends").SetFields(
      NewField("name"),
    ),
  ),
)

fmt.Println(q.Stringify())
```

```
query HeroNameAndFriends{hero{name,friends{name}}}
```

### Query with argument

```
q := &Query{
  Type: "query",
}

q.SetFields(
  NewField("human").SetArguments(
    &Argument{Name: "id", Value: 100},
  ).SetFields(
    NewField("name"),
    NewField("height"),
  ),
)

fmt.Println(q.Stringify())
```

```
query{human(id:100){name,height}}
```

### Query with variable

```
q := &Query{
  Type: "query",
  Name: "HeroNameAndFriends($episode: Episode)",
}

var arg ArgumentVariable = "episode"
q.SetFields(
  NewField("hero").SetArguments(
    &Argument{Name: "episode", Value: arg},
  ).SetFields(
    NewField("name"),
    NewField("friends").SetFields(
      NewField("name"),
    ),
  ),
)

fmt.Println(q.Stringify())
```

```
query HeroNameAndFriends($episode: Episode){hero(episode:$episode){name,friends{name}}}
```

## TODO

- [] Query by functional options
- [] Fragment

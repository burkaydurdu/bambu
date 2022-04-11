# Bambu

Bambu is a hash map library for golang. It uses single linked methodology.

### API

#### Create Hash Map

```go
hash := New[string, string]()
```

#### Add a new item

```go
hash.Add("name", "burkay")
hash.Add("surname", "durdu")
```

#### IndexOf

```go
k, v := hash.indexOf(0) // name, burkay
```

#### Get item with key

```go
val := hash.get("surname") // durdu
```

#### Get first/last item
```go
k, v := hash.first() // name, burkay
k, v = hash.last() // surname, durdu
```

#### Clear
```go
hash.clear()
```

#### Length
```go
length := hash.length() // 2
```

#### Remove item
```go
hash.Remove("name")
```
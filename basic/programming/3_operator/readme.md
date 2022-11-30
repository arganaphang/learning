# Operator

## Arithmetic Operators

### Add (+)

```go
3+4 // 7
```

### Subtract (-)

```go
5-3 // 2
```

### Multiply (\*)

```go
7*4 // 28
```

### Divide (/)

```go
54/6 // 9
```

### Modulus (%)

Modulus/modulo operator is a way to determine the remainder of a division operation.

```go
6%4 // 2
8%5 // 3
13%4 // 1
```

### Increment (++)

Increment the value by 1
but you need to learn about `prefix` and `postfix` operator

```go
a := 1
a++
fmt.Println(a) // 2
```

How Postfix increment actually works? `a := 1` and the we print while increment it using `fmt.Println(a++) // 1`

1. The program will use the value first
2. After the value being used, variable `a` will be increment it

And then, how Prefix increment actually works? `a := 1` and the we print while increment it using `fmt.Println(++a) // 2`

1. The value will be increment by 1 first
2. Then value will be used by print function

### Decrement (--)

Decrement the value by 1

```go
a := 0
a--
fmt.Println(a) // -1
```

## Relational Operators

This Operators commonly use for comparing two or more values and will return boolean as the result when the condition is pass

### Equal (==)

Will return `true` if left value is equal with right value

```go
a := 5
b := 3
fmt.Println(a == b) // false
```

### Not Equal (!=)

Will return `true` if left value is not equal with right value

```go
a := 5
b := 3
fmt.Println(a != b) // true
```

### Less Than (<)

Will return `true` if left value is less than with right value

```go
a := 5
b := 3
fmt.Println(a < b) // false
```

### Less Than Equal (<=)

Will return `true` if left value is less than equal with right value

```go
a := 3
b := 3
fmt.Println(a <= b) // true
```

### Greater Than (>)

Will return `true` if left value is greater than with right value

```go
a := 5
b := 3
fmt.Println(a > b) // true
```

### Greater Than Equal (>=)

Will return `true` if left value is greater than with right value

```go
a := 5
b := 3
fmt.Println(a >= b) // true
```

## Logical Operators

### And (&&)

| p       | q       | p && q  |
| ------- | ------- | ------- |
| `true`  | `true`  | `true`  |
| `true`  | `false` | `false` |
| `false` | `true`  | `false` |
| `false` | `false` | `false` |

```go
a := true
b := false
fmt.Println(a && b) // false
```

### Or (||)

| p       | q       | p \|\| q |
| ------- | ------- | -------- |
| `true`  | `true`  | `true`   |
| `true`  | `false` | `true`   |
| `false` | `true`  | `true`   |
| `false` | `false` | `false`  |

```go
a := true
b := false
fmt.Println(a || b) // true
```

### Not (!)

| p       | !p      |
| ------- | ------- |
| `true`  | `false` |
| `false` | `true`  |

## Bitwise Operators

| p   | q   | p & q | p \| q | p ^ q |
| --- | --- | ----- | ------ | ----- |
| 0   | 0   | 0     | 0      | 0     |
| 0   | 1   | 0     | 1      | 1     |
| 1   | 1   | 1     | 1      | 0     |
| 1   | 0   | 0     | 1      | 1     |

### Binary And (&)

### Binary Or (|)

### Binary Xor (^)

### Binary One's (~)

### Binary Left Shift Operator (<<)

### Binary Right Shift Operator (>>)

## Assignment Operators

### Simple Assign (=)

```go
a := 5
fmt.Println(a) // 5
a = 101        // <- assign new value
fmt.Println(a) // 101
```

### Add and Assign (+=)

```go
a := 5
fmt.Println(a) // 5
a += 3         // <- this is equal with `a = a + 5`, which is `a = 5 + 3`
fmt.Println(a) // 8
```

### Subtract and Assign (-=)

```go
a := 5
fmt.Println(a) // 5
a -= 3         // <- this is equal with `a = a - 5`, which is `a = 5 - 3`
fmt.Println(a) // 2
```

### Multiply and Assign (\*=)

```go
a := 5
fmt.Println(a) // 5
a *= 3         // <- this is equal with `a = a * 5`, which is `a = 5 * 3`
fmt.Println(a) // 15
```

### Divide and Assign (/=)

```go
a := 5
fmt.Println(a) // 5
a /= 3         // <- this is equal with `a = a / 5`, which is `a = 5 / 3`
fmt.Println(a) // 1 (this result might be different in other programming language (e.g in js will be print 1.6666666666666667))
```

### Modulus and Assign (%=)

```go
a := 5
fmt.Println(a) // 5
a %= 3         // <- this is equal with `a = a % 5`, which is `a = 5 % 3`
fmt.Println(a) // 2
```

### Left Shift and Assign (<<=)

### Right Shift and Assign (>>=)

### Bitwise and Assign (&=)

### Bitwise Exclusive and Assign (^=)

### Bitwise Inclusive and Assign (|=)

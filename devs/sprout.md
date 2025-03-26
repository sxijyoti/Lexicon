# Sprout : simple programming language with a playful design
    - variables ( data types ) 
        - like go :=
    - comments
    - print statement
    - conditional
    - switch case
    - expressions (arith, logical)


## Print
```sh
echo "hello"
echo 1
```

## Comments
```sh
#this is a comment
```

## Variable Declaration
```sh
sprout x = 10
sprout x int = 10
x = 10
```

## Conditional statement
```sh
if <exp>{
    <statement>
} 
else {
    <statement>
}
```

## expressions
```sh
# logical: and or not
sprout y = true || false;
z = x and y;
```
```sh
# arithmetic: + - * / 
sprout z = x + y;
a = 3**2;
```
```sh
# combination of operators
z = x * (y + 2);
result = (a + b) and (c * d);
```


## Switch case
```sh
choose x { 
    when 1  
        echo "One selected"  
    when 2  
        echo "Two chosen"  
    when 3  
        echo "Three picked"  
    otherwise  
        echo "Something else!"

}
```

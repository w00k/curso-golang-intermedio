# Curso Go/Golang Intermedio

Curso intermediode go/golang de platzi

## Módulos 

1. Usar módulos en Go/Golang.

```bash
go mod init github.com/user_name/module_name
```

Se genera un archivo **go.mod**. 

2. Para instalar un módulo nuevo usamos el siguiente comando.

```bash 
go get github.com/user_name/module_name
```

En el **go.mod** se agregará una nueva línea y se creará un archivo **sum.go** con las versiones del módulo.

3. Para eliminar dependencias que no se usan.

```bash
go mod tidy
```

4. Para usar una versión especifica del módulo, se tiene que especificar usando mod con el siguiente comando, el **/2** es lo que le da la versión.

```bash 
go get https://github.com/donvito/hellomod/v2
```

## Testing 

1. Para ejecutar testing, es necesario crear los testcase y ejecutar el comando:

```bash
go test
```

2. Para ejecutar el coverage de código es necesario ir a los testcase y ejecutar el siguiente coamndo:

```bash
go test -cover
```

3. Para ver el reporte de código que ha sido cubierto, es necesario generar un archivo y pasarlo por un tool para ser interpretado a lenguaje más legible.

Comandos

```bash 
go test -coverprofile=coverage.out
go tool cover -func=coverage.out
```

Salida

```bash 
github.com/w00k/curso-golang-intermedio/src/main.go:5:	Sum		100.0%
github.com/w00k/curso-golang-intermedio/src/main.go:9:	GetMax		0.0%
github.com/w00k/curso-golang-intermedio/src/main.go:16:	main		0.0%
total:							(statements)	16.7%
```

Otra forma de generar el reporte, pero en formato HTML es de la siguiente manera:

```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

Se abrira nuestro navegador web con el reporte en HTML.

## Profiling

```bash 
go test -cpuprofile=cou.out
```

```bash
go tool pprof cou.out
go tool pprof cou.out 
Type: cpu
Time: Jan 31, 2022 at 10:36pm (-03)
Duration: 67.61s, Total samples = 60.41s (89.35%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top // comando escrito
Showing nodes accounting for 60.22s, 99.69% of 60.41s total
Dropped 20 nodes (cum <= 0.30s)
      flat  flat%   sum%        cum   cum%
    60.22s 99.69% 99.69%     60.34s 99.88%  github.com/w00k/curso-golang-intermedio/src.Fibonacci
         0     0% 99.69%     60.34s 99.88%  github.com/w00k/curso-golang-intermedio/src.TestFibonacci
         0     0% 99.69%     60.34s 99.88%  testing.tRunner
(pprof) list Fibonacci // comando escrito
Total: 60.41s
ROUTINE ======================== github.com/w00k/curso-golang-intermedio/src.Fibonacci in /xxxx/xxxxxx/go/src/github.com/w00k/curso-golang-intermedio/src/main.go
    60.22s     96.55s (flat, cum) 159.82% of Total
         .          .     11:		return x
         .          .     12:	}
         .          .     13:	return y
         .          .     14:}
         .          .     15:
    25.74s     25.79s     16:func Fibonacci(n int) int {
     2.07s      2.07s     17:	if n <= 1 {
     8.32s      8.35s     18:		return n
         .          .     19:	}
    24.09s     60.34s     20:	return Fibonacci(n-1) + Fibonacci(n-2) // mayor tiempo de ejecución en el código
         .          .     21:}
         .          .     22:
         .          .     23:func main() {
         .          .     24:	x := Sum(5, 5)
         .          .     25:	fmt.Println(x)
ROUTINE ======================== github.com/w00k/curso-golang-intermedio/src.TestFibonacci in /xxxx/xxxxxx/go/src/github.com/w00k/curso-golang-intermedio/src/main_test.go
         0     60.34s (flat, cum) 99.88% of Total
         .          .     54:		{8, 21},
         .          .     55:		{50, 12586269025},
         .          .     56:	}
         .          .     57:
         .          .     58:	for _, item := range tables {
         .     60.34s     59:		fibonacci := Fibonacci(item.a) // mayor tiempo de ejecución en el test
         .          .     60:		if fibonacci != item.n {
         .          .     61:			t.Errorf("Fibonacci war incorrect, got %d and except %d", fibonacci, item.n)
         .          .     62:		}
         .          .     63:	}
         .          .     64:}
(pprof) web // también puede ser pdf 

```

# Mocks

Ejemplo en el testing como hacer un mock en los tables

```bash
func TestGetFullTimeEmployeeById(t testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEMployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{Id: 1, Position: "CEO"}, nil
				}

				GetPersonByDNI = func(id string) (Person, error) {
					return Person{Name: "Andy", Age: 35, DNI: "1"}, nil
				}
			},
			expectedEMployee: FullTimeEmployee{},
		},
	}
}
```

# curso-golang-intermedio
curso intermediode go/golang de platzi

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

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

4. Para usar una versión especifica del módulo, se tiene que especificar usando mod con el siguiente comando.

```bash 
go get mod https://github.com/donvito/hellomod/v2
```

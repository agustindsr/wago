# Nombre del archivo wasm resultante
WASM_OUTPUT = public/play.wasm

# Comando para compilar Go a WebAssembly
build:
	GOOS=js GOARCH=wasm go build -o public/main.wasm main.go

# Comando para optimizar el archivo WebAssembly
optimize: build
	wasm-opt public/main.wasm --enable-bulk-memory -Oz -o $(WASM_OUTPUT)

# Regla por defecto que ejecuta ambas tareas
all: optimize

# Limpiar los archivos generados
clean:
	rm -f main.wasm $(WASM_OUTPUT)

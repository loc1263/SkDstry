# Descripcion 

Busca archivos en un arbol de directorios, dando la opcion de eliminarlos.  

# Uso

\<nombre programa> \<directorio_raiz> \<pattern>  
go run skdstry.go d:\msys64 cnf  


# Ejemplo

```sh
D:\Code\GoLang\SkDstry>go run skdstry.go d:\msys64 cnf
Buscando..... 
Archivos encontrados:
d:\msys64\mingw64\etc\ssl\ct_log_list.cnf     
d:\msys64\mingw64\etc\ssl\ct_log_list.cnf.dist
d:\msys64\mingw64\etc\ssl\openssl.cnf
d:\msys64\mingw64\etc\ssl\openssl.cnf.dist    
d:\msys64\mingw64\include\iiscnfg.h
d:\msys64\ucrt64\include\iiscnfg.h
d:\msys64\usr\share\texinfo\htmlxref.cnf
d:\msys64\usr\ssl\ct_log_list.cnf
d:\msys64\usr\ssl\ct_log_list.cnf.dist
d:\msys64\usr\ssl\openssl.cnf
d:\msys64\usr\ssl\openssl.cnf.dist
Peso total en KB: 135.33 KB
Peso total en MB: 0.13 MB
¿Desea borrar los archivos encontrados? (Y/N): n
No se han borrado los archivos.
```

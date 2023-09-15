let filtrarCadenas (subcadena:string) (cadenas:string list) =
    List.filter (fun (cadena : string)-> cadena.Contains(subcadena)) cadenas

// Ejemplo de uso:
let listaCadenas = ["la casa"; "el perro"; "pintando la cerca"]
let resultado = filtrarCadenas "la" listaCadenas

printfn "%A" resultado // ["la casa"; "pintando la cerca"]

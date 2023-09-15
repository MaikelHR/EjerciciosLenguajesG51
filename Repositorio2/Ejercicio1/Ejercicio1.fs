let desplazar (direccion:string) (posiciones:int) (lista:int list) =
    let longitud = List.length lista
    match direccion with
    | "izq" ->
        let desplazamiento = posiciones % longitud
        List.init longitud (fun i -> if i < desplazamiento then List.nth lista (longitud - desplazamiento + i) else 0)
    | "der" ->
        let desplazamiento = posiciones % longitud
        List.init longitud (fun i -> if i < desplazamiento then 0 else List.nth lista (i - desplazamiento))
    | _ ->
        List.replicate longitud 0



// Ejemplos de uso:
let resultado1 = desplazar "izq" 2 [1; 2; 3; 4; 5]
let resultado2 = desplazar "der" 2 [1; 2; 3; 4; 5]
let resultado3 = desplazar "izq" 6 [1; 2; 3; 4; 5]

printfn "%A" resultado1 // [4; 5; 0; 0; 0]
printfn "%A" resultado2 // [0; 0; 1; 2; 3]
printfn "%A" resultado3 // [0; 0; 0; 0; 0]


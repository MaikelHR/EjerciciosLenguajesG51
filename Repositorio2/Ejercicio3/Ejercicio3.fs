let n_esimo n lista =
    let indices = [0..List.length lista - 1]
    let elementos = List.map (fun i -> List.nth lista i) indices
    match List.tryFindIndex (fun x -> x = n) elementos with
    | Some indice -> true, indice + 1 // El índice se ajusta a partir de 1
    | None -> false, 0

// Ejemplos de uso:
let resultado1, indice1 = n_esimo 2 [1; 2; 3; 4; 5]
let resultado2, indice2 = n_esimo 3 [1; 2; 3; 4; 5]
let resultado3, _ = n_esimo 6 [1; 2; 3; 4; 5]

printfn "Posición n-esimo 2: %d" indice1 // Posición n-esimo 2: 3
printfn "Posición n-esimo 3: %d" indice2 // Posición n-esimo 3: 4
printfn "Posición n-esimo 6: No encontrado" // Posición n-esimo 6: No encontrado
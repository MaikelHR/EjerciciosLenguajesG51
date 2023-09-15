type GrafoPonderado = (char * (char * int) list) list

let dfsPonderado grafo inicio objetivo =
    let rec dfs_aux ruta objetivo =
        if List.isEmpty ruta then
            None // No se encontró un camino
        elif List.head ruta = objetivo then
            Some (List.rev ruta) // Se encontró el objetivo
        else
            let verticeActual = List.head ruta
            match List.tryFind (fun (vecino, peso) -> vecino = verticeActual) grafo with
            | Some (_, vecinos) ->
                let extensiones = List.choose (fun (vecino, peso) ->
                    if not (List.contains vecino ruta) then
                        Some (vecino :: ruta, peso)
                    else
                        None
                ) vecinos
                let resultados = List.choose (fun (ruta', peso) ->
                    match dfs_aux ruta' objetivo with
                    | Some camino -> Some (camino, peso)
                    | None -> None
                ) extensiones
                match List.minBy snd resultados with
                | (mejorCamino, mejorPeso) -> Some mejorCamino
            | None -> None

    dfs_aux [inicio] objetivo

// Ejemplo de uso:
let grafoPonderado : GrafoPonderado = [
    ('A', [('B', 3); ('C', 1)]);
    ('B', [('D', 2)]);
    ('C', [('D', 6)]);
    ('D', []);
]

let resultado = dfsPonderado grafoPonderado 'A' 'D'
match resultado with
| Some camino ->
    printfn "Camino más corto: %A" camino
| None ->
    printfn "No se encontró un camino."

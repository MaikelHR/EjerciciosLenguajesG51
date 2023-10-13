open System
open Microsoft.FSharp.Core

let matriz = 
    [
        ['A'; 'B'; 'I'; 'N'; 'G'; 'P'];
        ['G'; 'B'; 'I'; 'J'; 'K'; 'L'];
        ['M'; 'N'; 'I'; 'P'; 'Q'; 'R'];
        ['S'; 'T'; 'U'; 'N'; 'W'; 'X'];
        ['Y'; 'Z'; 'A'; 'B'; 'G'; 'D'];
        ['E'; 'F'; 'G'; 'H'; 'I'; 'O'];
    ]

let palabraABuscar = "BINGO"

let actualizarVisitadas x y visitadas valor =
    let fila = List.nth visitadas x
    let filaActualizada = List.mapi (fun j _ -> if j = y then valor else List.nth fila j) fila
    List.mapi (fun i _ -> if i = x then filaActualizada else List.nth visitadas i) visitadas

let inicializarVisitadas filas columnas =
    List.init filas (fun _ -> List.init columnas (fun _ -> false))

let extraerSubcadena (cadena: string) (indiceInicial: int) (longitud: int) =
    let rec loop index remaining acc =
        if index >= String.length cadena || remaining <= 0 then
            acc
        else
            let charToAdd = cadena.[index]
            loop (index + 1) (remaining - 1) (acc + charToAdd.ToString())
    loop indiceInicial longitud ""

let rec buscarDesdeCoordenada (x: int) (y: int) (letrasRestantes: string) (visitadas: bool list list) =
    if String.IsNullOrEmpty letrasRestantes then
        [(x, y)]
    else
        let letraActual = letrasRestantes.[0]
        let posiblesMovimientos =
            [
                (x + 1, y); (x - 1, y); (x, y + 1); (x, y - 1);
                (x + 1, y + 1); (x + 1, y - 1); (x - 1, y + 1); (x - 1, y - 1)
            ]
        let coordenadasValidas =
            posiblesMovimientos
            |> List.filter (fun (nx, ny) ->
                nx >= 0 && nx < List.length matriz &&
                ny >= 0 && ny < List.length (List.head matriz) &&
                not (List.nth (List.nth visitadas nx) ny)
            )
        let resultado =
            coordenadasValidas
            |> List.tryPick (fun (nx, ny) ->
                if List.nth (List.nth matriz nx) ny = letraActual then
                    let subcadena = extraerSubcadena letrasRestantes 1 (String.length letrasRestantes - 1)
                    let coordenadasEncontradas = buscarDesdeCoordenada nx ny subcadena (actualizarVisitadas x y visitadas true)
                    if List.length coordenadasEncontradas > 0 then
                        Some((x, y)::coordenadasEncontradas)
                    else
                        None
                else
                    None
            )
        match resultado with
        | Some(coordenadas) -> coordenadas
        | None -> actualizarVisitadas x y visitadas false; []

let rec iniciarBusquedaDesdeCoordenadas (x: int) (y: int) =
    if List.nth (List.nth matriz x) y = palabraABuscar.[0] then
        let subcadena = extraerSubcadena palabraABuscar 1 (String.length palabraABuscar - 1)
        let resultado = buscarDesdeCoordenada x y subcadena (inicializarVisitadas (List.length matriz) (List.length (List.head matriz)))
        if List.length resultado > 0 then
            Some(resultado)
        else
            None
    else
        None

let rec buscarEnMatrizDesdeTodosPuntos () =
    let filas = List.length matriz
    let columnas = List.length (List.head matriz)
    let rec buscarDesdePuntoInicial x y =
        match iniciarBusquedaDesdeCoordenadas x y with
        | Some(coordenadas) -> Some(coordenadas)
        | None ->
            if x < filas - 1 then
                if y < columnas - 1 then
                    buscarDesdePuntoInicial x (y + 1)
                else
                    buscarDesdePuntoInicial (x + 1) 0
            else
                None
    let rec buscarDesdeTodosLosPuntos x y =
        match buscarDesdePuntoInicial x y with
        | Some(coordenadas) -> Some(coordenadas)
        | None ->
            if x < filas - 1 then
                if y < columnas - 1 then
                    buscarDesdeTodosLosPuntos x (y + 1)
                else
                    buscarDesdeTodosLosPuntos (x + 1) 0
            else
                None
    buscarDesdeTodosLosPuntos 0 0

match buscarEnMatrizDesdeTodosPuntos() with
| Some(coordenadas) -> printfn "Palabra encontrada en coordenadas: %A" coordenadas
| None -> printfn "Palabra no encontrada en la matriz"

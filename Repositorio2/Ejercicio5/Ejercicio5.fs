type Casilla = { x: int; y: int }
type Laberinto = char[,]

let esCasillaValida (laberinto: Laberinto) (casilla: Casilla) =
    let rows = Array2D.length1 laberinto
    let cols = Array2D.length2 laberinto
    casilla.x >= 0 && casilla.x < rows && casilla.y >= 0 && casilla.y < cols

let obtenerAdyacentes (laberinto: Laberinto) (casilla: Casilla) =
    let direcciones = [(1, 0); (0, 1); (-1, 0); (0, -1)] // Abajo, Derecha, Arriba, Izquierda
    direcciones
    |> List.map (fun (dx, dy) -> { x = casilla.x + dx; y = casilla.y + dy })
    |> List.filter (esCasillaValida laberinto)

let laberintoComoGrafo (laberinto: Laberinto) =
    let rows = Array2D.length1 laberinto
    let cols = Array2D.length2 laberinto
    let grafo = Array2D.create rows cols []

    for x = 0 to rows - 1 do
        for y = 0 to cols - 1 do
            let casilla = { x = x; y = y }
            let adyacentes = obtenerAdyacentes laberinto casilla
            Array2D.set grafo x y adyacentes

    grafo

let rec dfs laberinto grafo inicio objetivo ruta =
    if inicio = objetivo then
        ruta @ [inicio]
    else
        let adyacentes = Array2D.get grafo inicio.x inicio.y
        let rutaExtendida = ruta @ [inicio]
        let resultados =
            adyacentes
            |> List.filter (fun casilla -> not (List.contains casilla ruta))
            |> List.map (fun vecino -> dfs laberinto grafo vecino objetivo rutaExtendida)
        List.tryFind (fun resultado -> List.exists (fun c -> c = objetivo) resultado) resultados |> Option.defaultValue []

let encontrarRutas laberinto inicio objetivo =
    let grafo = laberintoComoGrafo laberinto
    dfs laberinto grafo inicio objetivo []

// Ejemplo de uso:
let laberinto : Laberinto = 
    [|
        ['S'; ' '; ' '; ' '; ' '; ' '; ' '];
        ['X'; 'X'; 'X'; 'X'; 'X'; 'X'; ' '];
        [' '; ' '; ' '; ' '; ' '; 'X'; ' '];
        ['X'; 'X'; 'X'; 'X'; ' '; 'X'; ' '];
        ['X'; ' '; ' '; ' '; ' '; 'X'; ' '];
        ['X'; 'X'; 'X'; 'X'; 'X'; 'X'; ' '];
        [' '; ' '; ' '; ' '; ' '; ' '; 'E']
    |]

let inicio = { x = 0; y = 0 }
let objetivo = { x = Array2D.length1 laberinto - 1; y = Array2D.length2 laberinto - 1 }

let rutas = encontrarRutas laberinto inicio objetivo

printfn "Rutas encontradas:"
rutas |> List.iteri (fun i ruta -> printfn "Ruta %d: %A" (i + 1) ruta)

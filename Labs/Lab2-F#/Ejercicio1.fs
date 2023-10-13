let cromosomaSujeto = [2; 10; 5; 9; 7; 6; 4; 2; 1; 1]
let muestra= [5; 10; 5; 2; 3; 9; 8; 7; 1; 1]
let personas = [("Persona1", [7; 2; 4; 4; 10; 6; 7; 15; 9; 10]);
        ("Persona2", [0; 7; 3; 4; 5; 8; 4; 8; 6; 3]);
        ("Persona3", [5; 3; 3; 4; 5; 10; 7; 8; 9; 0]);
    ]

let cromosomaSujetoCandidato (sujeto:int list) (candidato:int list) : float =
    let rec contarCoincidencias suj candid acc =
        match suj, candid with
        | [], [] -> acc
        | h1::t1, h2::t2 when h1 = h2 -> contarCoincidencias t1 t2 (acc + 1)
        | _::t1, _::t2 -> contarCoincidencias t1 t2 acc
        | _, _ -> acc
    let coincidencias = contarCoincidencias sujeto candidato 0
    let porcentaje = float coincidencias / float (List.length sujeto) * 100.0
    porcentaje


let sujetoMasParecidoAMuestra (muestra:int list) (personas:(string * int list) list) : string option =
    match personas with
    | [] -> None
    | _ ->
        let mejoresCoincidencias =
            personas
            |> List.map (fun (nombre, cromosoma) -> (nombre, cromosomaSujetoCandidato muestra cromosoma))
            |> List.filter (fun (_, porcentaje) -> porcentaje >= 0.0) // Filtrar por coincidencias positivas
        match mejoresCoincidencias with
        | [] -> None
        | _ ->
            let sujetoMasParecido =
                mejoresCoincidencias
                |> List.maxBy (fun (_, porcentaje) -> porcentaje)
            Some (fst sujetoMasParecido)

let porcentajeExacto = cromosomaSujetoCandidato cromosomaSujeto muestra
printf "Porcentaje de exactitud: %f%%" porcentajeExacto

match sujetoMasParecidoAMuestra muestra personas with
| Some(nombre) -> printf "\nEl sujeto mas parecido a la muestra es: %s" nombre
| None -> printf "\nNo se encontró un sujeto parecido a la muestra"
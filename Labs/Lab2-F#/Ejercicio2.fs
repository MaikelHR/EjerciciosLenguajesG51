let rec esIsograma (palabra:string) : bool =
    let palabraNormalizada = palabra.ToLower() // Convierte la palabra a minúsculas
    match palabraNormalizada with
    | "" | " " -> true // Caso base: palabra vacía o solo espacios
    | _ ->
        let primeraLetra = palabraNormalizada.[0]
        let restoPalabra = palabraNormalizada.Substring(1)
        if restoPalabra.Contains(primeraLetra) then
            false
        else
            esIsograma restoPalabra

let palabra1 = "murcielago"
let palabra2 = "repetido"
let palabra3 = "FSharp"
let palabra4 = "Isograma"

printfn "%s es un isograma: %b" palabra1 (esIsograma palabra1)
printfn "%s es un isograma: %b" palabra2 (esIsograma palabra2)
printfn "%s es un isograma: %b" palabra3 (esIsograma palabra3)
printfn "%s es un isograma: %b" palabra4 (esIsograma palabra4)
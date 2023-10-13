%Repositorio3 Prolog Maikel Hernandez G51
%1)Defina sumlist(L, S) que es verdadero si S es la suma de los elementos de L.

sumlist([], 0).
sumlist([H|T], S) :- sumlist(T, Resto), S is H + Resto.

%sumlist([1, 2, 3, 4, 5], S). S=15 true

%2)Defina la relación subconj(S, S1), donde S y S1 son listas representando conjuntos, que es verdadera si S1 es subconjunto de S.

subconj([], _).
subconj([H|T], S1) :- member(H, S1), subconj(T, S1).

%?- subconj([2, 3],[1, 2, 3, 4, 5]). true.

%?- subconj([6, 7],[1, 2, 3, 4, 5]). false.

%3)Defina la función aplanar(L,L2). Esta función recibe una lista con múltiples listas anidadas como elementos
% otra lista que contendría los mismos elementos de manera lineal (sin listas).

aplanar([], []).
aplanar([H|T], L2) :- is_list(H), aplanar(H, FlatH), aplanar(T, FlatT), append(FlatH, FlatT, L2).
aplanar([H|T], [H|FlatT]) :- \+ is_list(H), aplanar(T, FlatT).

%?- aplanar([1, [2, 3, [4, 5], 6], [7, [8], 9]], L2).
%L2 = [1, 2, 3, 4, 5, 6, 7, 8, 9].

%4)Defina un predicado llamado partir(Lista, Umbral, Menores, Mayores) para dividir una lista respecto un
% umbral dado, dejando los valores menores a la izquierda y los mayores
% a la derecha. Por ejemplo, el resultado
% de partir la lista [2,7,4,8,9,1] respecto al umbral 6 serían las listas[2,4,1] y [7,8,9].

partir([], _, [], []).  % Caso base: lista vacía, ambas listas resultantes son vacías.
partir([H|T], Umbral, [H|Menores], Mayores) :- H < Umbral, partir(T, Umbral, Menores, Mayores).
partir([H|T], Umbral, Menores, [H|Mayores]) :- H >= Umbral, partir(T, Umbral, Menores, Mayores).

%?- partir([2, 7, 4, 8, 9, 1], 6, Menores, Mayores).
%Menores = [2, 4, 1],
%Mayores = [7, 8, 9].

%5)Implemente un predicado que, a partir de una lista de cadenas string, filtre aquellas que
%contengan una subcadena que el usuario indique en otro argumento. Ej

% sub_cadenas(“la”, [“la casa, “el perro”, “pintando lacerca”],Filtradas).
%True
%Filtradas = [“la casa, “pintando la cerca”]

% Caso base: la lista vacía resulta en una lista vacía.
sub_cadenas(_, [], []).

% Si la cadena H contiene la subcadena, agregarla a la lista filtrada y continuar con el resto de la lista.
sub_cadenas(Subcadena, [H|T], [H|Filtradas]) :-
    sub_atom(H, _, _, _, Subcadena), % Verifica si Subcadena está contenida en H
    sub_cadenas(Subcadena, T, Filtradas).

% Si la cadena H no contiene la subcadena, simplemente continuar con el resto de la lista.
sub_cadenas(Subcadena, [_|T], Filtradas) :-
    sub_cadenas(Subcadena, T, Filtradas).

% ?- sub_cadenas("la", ["la casa", "el perro", "pintando la cerca"], Filtradas).
%Filtradas = ["la casa", "pintando la cerca"].

%EJERCICIOS DEL VIDEO DE CLASE

conectado(i,a,50).
conectado(i,b,20).
conectado(a,b,15).
conectado(a,c,25).
conectado(b,f,35).
conectado(b,c,40).
conectado(c,f,5).

conectados(X,Y,Peso):- conectado(X,Y,Peso).
conectados(X,Y,Peso):- conectado(Y,X,Peso).


ruta(Ini,Fin,Ruta,Peso):- ruta2(Ini,Fin,[],Ruta,Peso).

ruta2(Fin,Fin,_,[Fin],0).
ruta2(Ini,Fin,Visitados,[Ini|Resto],PesoTotal):-
    conectados(Ini,Alguien,Peso),
    not(member(Alguien,Visitados)),
    ruta2(Alguien,Fin,[Ini|Visitados],Resto,PesoRestante),
    PesoTotal is Peso + PesoRestante.

%ruta(i,f,Ruta,Peso
%Ruta = [i,a,b,f]
%Peso = 100

% Definición de personas con sus cromosomas
persona(maikel, [1,0,0,0,0,1,0,0,1,1]).
persona(ernesto, [1,0,0,0,0,1,0,0,1,0]).
persona(lucia, [0,1,0,1,0,1,1,0,1,1]).
persona(ana, [1,0,0,0,0,1,0,0,0,1]).
persona(juan, [0,0,1,1,1,0,1,1,0,0]).
persona(maria, [1,1,0,0,0,1,1,0,0,1]).
persona(roberto, [1,0,1,0,0,1,0,0,1,1]).
persona(laura, [1,0,1,1,0,1,0,0,1,0]).
persona(carlos, [0,0,1,0,1,0,1,1,0,0]).
persona(susana, [1,1,1,0,0,0,0,1,1,1]).

% Predicado para calcular la similitud entre dos cromosomas
similitud_cromosomas([], [], 100).  % Si ambos cromosomas están vacíos, la similitud es 100%.
similitud_cromosomas([H1|T1], [H2|T2], Similitud) :-
    similitud_cromosomas(T1, T2, SimilitudRestante),
    (H1 =:= H2 -> Similitud is SimilitudRestante; Similitud is SimilitudRestante - 10).

% Predicado para calcular la similitud entre dos personas
similitud_persona(Nombre1, Nombre2, Similitud) :-
    persona(Nombre1, Cromosomas1),
    persona(Nombre2, Cromosomas2),
    similitud_cromosomas(Cromosomas1, Cromosomas2, Similitud).

% Ejemplo de uso
%similitud_persona(maikel, ernesto, Similitud).
%Similitud = 90.

%similitud_persona(lucia, ana, Similitud).
%Similitud = 50.

%similitud_persona(juan, maria, Similitud).
%Similitud = 20.

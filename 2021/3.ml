open Core

let rec pow x y = if y = 0 then 1 else x * pow x (y-1)

let explode s =
  let rec exp i l =
    if i < 0 then l else exp (i-1) ( match s.[i] with '1' -> 1::l | _-> 0::l) in
  exp (String.length s - 1) []

let rec lst_add a b =
  match a, b with
  | [], _ -> b
  | _, [] -> a
  | h1 :: t1, h2 :: t2 -> (h1+h2)::(lst_add t1 t2)

let sumIns lst =
  let rec sum acc lst =
    match lst with
    | [] -> acc
    | h :: tl -> sum (lst_add h acc)  tl in
  match lst with
  | []->[]
  | h::tl -> sum h tl


let part1 ins =
  (sumIns ins, List.length ins) |> fun (acc, total) ->
  let g = ref 0 in let e = ref 0 in
  let rec compute ac t =
    match ac with
    | [] -> ()
    | h :: tl -> (if 2*h > total
                  then g:= !g + pow 2 (t-1)
                  else e:= !e + pow 2 (t-1)); compute tl (t-1) in
  compute acc (List.length acc);
  (!g * !e)


let rec computeO lst oxy  =
  if List.length lst = 1 then match lst with
    | x::_ -> oxy @ x
    | _ -> oxy else
    let zero = ref [] in let one = ref [] in
    let rec res lst  =
      match lst with
      | [] -> ()
      | h::tl -> match h with
        | 1::r -> one := r::!one; res tl
        | 0::r -> zero:= r::!zero; res tl
        | _ -> () in
    res lst;
    if ((List.length !one) >= (List.length !zero))
    then computeO !one (oxy @ [1])
    else computeO !zero (oxy @ [0])


let rec computeS lst oxy  =
  if List.length lst = 1 then match lst with
    | x::_ -> oxy @ x
    | _ -> oxy else
    let zero = ref [] in let one = ref [] in
    let rec res lst  =
      match lst with
      | [] -> ()
      | h::tl -> match h with
        | 1::r -> one := r::!one; res tl
        | 0::r -> zero:= r::!zero; res tl
        | _ -> () in
    res lst;
    if ((List.length !one) < (List.length !zero))
    then computeS !one (oxy @ [1])
    else computeS !zero (oxy @ [0])

let rec to_decimal lst acc =
  match lst with
  | [] -> acc
  | x::tl -> to_decimal tl (acc + x * pow 2 (List.length tl))

let ()=
  let instructions =
    In_channel.read_lines "/mnt/d/input.txt"
    (* ["00100"; *)
    (*  "11110"; *)
    (*  "10110"; *)
    (*  "10111"; *)
    (*  "10101"; *)
    (*  "01111"; *)
    (*  "00111"; *)
    (*  "11100"; *)
    (*  "10000"; *)
    (*  "11001"; *)
    (*  "00010"; *)
    (*  "01010"] *)
    |> List.map  ~f:(fun l-> Scanf.sscanf l "%s" (fun l -> explode l)) in
  (to_decimal (computeO instructions []) 0) * (to_decimal (computeS instructions [] ) 0) |>
  Printf.printf "%d \n%!"

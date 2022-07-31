open Core

let first = function [] -> 0 | x::_ -> x
let rest = function [] -> [] | _::tl -> tl

let m_int ch str = String.split str ~on:ch |> List.map ~f:int_of_string;;

let rec remove x lst =
  match lst with
  | [] -> []
  | e::tl -> if x = e then tl else e::(remove x tl)

let rec transpose = function
  | [] -> []
  | []::tl -> transpose tl
  | (x::tl)::tl' ->
    (x:: List.map tl' ~f:first ) ::  transpose (tl:: List.map ~f:rest tl')

let play n bingo =
  let rec play_aux n acc = function
    | [] -> acc
    | x::tl -> if x<>n then play_aux n (x::acc) tl else acc @ tl in
  List.map bingo ~f:(play_aux n [])

let check bingo =
  let rec check_empty = function
    | [] -> false
    | x::tl -> if List.length x = 0 then true else check_empty tl
  in check_empty bingo

let sum bingo =
  let rec sum_aux ?acc:(acc=0) = function
    | [] -> acc
    | x::tl -> sum_aux ~acc:(acc+x) tl in
  bingo |> List.map ~f:sum_aux |> sum_aux

let rec printArr = function
  | [] -> ()
  | x::tl ->Printf.printf "%d " x ; printArr tl


let _ =
  (In_channel.read_all "/mnt/d/input.txt") |>
  Str.split (Str.regexp "\n\n")
  |> function
  | [] -> ()
  | x::h ->
    let numbers  = m_int ','  x in
    let bingos = h |> List.map ~f:(fun s ->
        (String.split s ~on:'\n')
        |> List.map ~f:(fun l -> Str.split (Str.regexp " +") l
                                 |> List.map ~f:int_of_string ))
                 |> List.map ~f:(fun b-> (b, transpose b)) in
    let rec win_first aB = function
      | [] -> 0
      | n::tl ->
        let curr = List.map aB ~f:(fun (r, c) -> (play n r, play n c )) in
        let rec onBingos = function
          | [] -> 0
          | (r, c)::os  -> if (check r || check c) then n * sum r else onBingos os in
        let result = onBingos curr in
        if result <> 0 then result else win_first curr tl in

    let rec win_last aB = function
      | [] -> 0
      | n::tl ->
        let curr = List.map aB ~f:(fun (r, c) -> (play n r, play n c )) in
        let rec onBingos acc = function
          | [] -> (0, acc)
          | (r, c)::os  -> let win = check r || check c in
            if ( win && (List.length curr = 1)) then (n * sum r, acc) else
            if win then onBingos acc os  else onBingos ((r,c)::acc) os in
        let (result, remainder) = onBingos [] curr in
        if result <> 0 then result else win_last remainder tl in

    Printf.printf "Silver %d \n %!" (win_first bingos numbers);
    Printf.printf "Gold %d \n %!" (win_last bingos numbers);

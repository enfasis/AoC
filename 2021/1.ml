open Core;;

let rec sum1 numbers acc =
  match numbers with
  | x::y::r -> if x < y then sum1 (y::r) acc+1 else sum1 (y::r) acc
  | _ -> acc

let rec sum2 numbers acc prev=
  match numbers with
  | x::y::z::r -> let curr = x+y+z in
    if curr > prev && (prev <> 0) then sum2 (y::z::r) (acc+1) curr  else sum2 (y::z::r) acc curr
  | _ -> acc

let () =
  let numbers = In_channel.read_lines "/mnt/d/input.txt" |> List.map ~f:int_of_string in
  sum1 numbers 0 |> Printf.printf "Silver: %i\n%!";
  sum2 numbers 0  0 |> Printf.printf "Gold: %i\n%!"

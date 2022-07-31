open Core

let get_fuel1 l =
  let median = ListLabels.nth l (List.length l / 2) in
  List.fold l ~init:0 ~f:(fun acc x -> acc + (x-median |> abs))

let get_fuel2 l =
  let mean = float_of_int (List.fold l ~init:0 ~f:(+)) /. float_of_int (List.length l)
             |> round ~dir:`Nearest |> int_of_float in
  let sum n =  let s = ref 0 in for i = 1 to n do s:= !s+i done; !s in
  List.fold l ~init:0 ~f:(fun acc x -> acc + (x-mean+1 |> abs |> sum ))

let _ =
  let crabs =
    In_channel.read_all "input.txt" |> StringLabels.trim |> String.split_on_chars ~on:[',']
    |> List.map ~f:int_of_string |> List.sort ~compare:(fun a b -> if a>b then 1 else 0) in
  crabs |> get_fuel1 |> Printf.printf "Silver %d\n%!";
  crabs |> get_fuel2 |> Printf.printf "Gold %d\n%!"

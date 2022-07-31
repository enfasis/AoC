open Core

let to_fishes lst = let fs = StdLabels.Array.make 9 0 in
  let rec to_aux = function
    | [] -> fs
    | x::tl -> fs.(x) <- (fs.(x) +1); to_aux tl in
  to_aux lst |> StdLabels.Array.to_list

let rec day j fs =
  let day_aux = function
    | x0::x1::x2::x3::x4::x5::x6::x7::x8::_ -> [x1;x2;x3;x4;x5;x6;x7+x0;x8;x0]
    | _->[] in
  if j<>0 then day (j-1) (day_aux fs) else fs

let _ = let fishes =
          In_channel.read_all "/mnt/d/input.txt" |> StringLabels.trim |> String.split_on_chars ~on:[',']
          |> List.map ~f:int_of_string |> to_fishes in
  fishes |> day 80 |> List.fold ~init:0 ~f:(+)  |> Printf.printf "Silver: %d \n%!";
  fishes |> day 256 |> List.fold ~init:0 ~f:(+) |> Printf.printf "Gold: %d \n%!"

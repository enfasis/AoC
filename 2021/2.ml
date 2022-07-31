open Core;;

let rec sumPos1 instructions accx accy =
  match instructions with
  | ("forward", v) :: r -> sumPos1 r (accx+v) accy
  | ("up", v) :: r -> sumPos1 r accx (accy-v)
  | ("down", v) :: r -> sumPos1 r accx (accy+v)
  | _ -> (accx, accy)

let rec sumPos2 instructions accx accy aim =
  match instructions with
  | ("forward", v) :: r -> sumPos2 r (accx+v) (accy+aim*v) aim
  | ("up", v) :: r -> sumPos2 r accx (accy) (aim-v)
  | ("down", v) :: r -> sumPos2 r accx (accy) (aim+v)
  | _ -> (accx, accy)


let () =
  let instructions =
    In_channel.read_lines "/mnt/d/input.txt"
    |> List.map  ~f:(fun l-> Scanf.sscanf l "%s %i" (fun x y -> x, y)) in
  sumPos1 instructions 0 0  |> fun (x, y) -> Printf.printf "Silver: %i\n%!"  (x*y);
  sumPos2 instructions 0 0 0 |> fun (x, y) -> Printf.printf "Gold: %i\n%!"  (x*y)

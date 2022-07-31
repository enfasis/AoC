open Core

let parse_data = function
  | x::y::_ -> (x |> StringLabels.trim |> String.split_on_chars ~on:[' '] |> List.to_array,
                y |> StringLabels.trim |> String.split_on_chars ~on:[' '] |> List.to_array)
  | _ -> ([||],[||])

let count = Array.fold ~init: 0  ~f:(
    fun acc v ->
      let l = String.length v in
      if l = 2 || l = 4 || l= 3 || l = 7
      then acc+1
      else acc
  )

let contains str1 str2 = String.fold ~init: 0  ~f:(
    fun acc c ->
      if String.exists str2 ~f:(
          fun ch -> Char.equal ch c)
      then acc+1
      else acc
  ) str1 |> (=) (String.length str1)

let decode str =
  let mask= StdLabels.Array.make 10 "" in
  let fives = ref [] in
  let sixs = ref [] in
  let () = StdLabels.Array.iteri ~f:(fun i s ->
      match String.length s with
      | 2 -> mask.(1) <- s
      | 4 -> mask.(4) <- s
      | 3 -> mask.(7) <- s
      | 7 -> mask.(8) <- s
      | 6 -> sixs:= s::!sixs
      | 5 -> fives:= s::!fives
      | _ -> ()
    ) str in
  let rec update n1 n2 acc ?inv:(inv=false) = function
    | x::tl ->  if (if inv then contains x mask.(n1) else contains mask.(n1) x)
      then (mask.(n2) <- x; acc@tl)
      else update n1 n2 (x::acc) tl ~inv:inv
    | [] -> acc in
  let rest = update 4 9 [] !sixs in
  let _ = match update 7 0 [] rest with | x::_ -> mask.(6) <- x | _ -> () in
  let rest = update 1 3 [] !fives in
  let _ = match update 6 5 [] rest ~inv:true with | x::_ -> mask.(2) <- x | _ -> () in
  mask

let get_key str lst = lst
                      |> Array.findi_exn ~f:(fun i v -> contains str v && contains v str)
                      |> function | (i,_) -> i

let rec process key acc= function
  | [] -> acc
  | x::tl -> process key (10*acc + get_key x key) tl


let ()=
  let ins =
    In_channel.read_lines "/mnt/d/input.txt" |> List.map ~f:(fun l -> l |> String.split_on_chars ~on:['|'] |> parse_data) in
  ins |> List.fold ~init: 0 ~f:(fun acc (_,o) -> acc + count o) |> Printf.printf "Silver %d\n%!";
  ins |> List.fold ~init:0 ~f:(fun acc (i,o)-> acc + process (decode i) 0 (Array.to_list o) ) |> Printf.printf "Gold %d\n%!"

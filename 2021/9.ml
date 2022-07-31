open Core

module IntPair = struct
  module T = struct
    type t = int * int
    let compare x y = Tuple2.compare ~cmp1:Int.compare ~cmp2:Int.compare x y
    let sexp_of_t = Tuple2.sexp_of_t Int.sexp_of_t Int.sexp_of_t
    let t_of_sexp = Tuple2.t_of_sexp Int.t_of_sexp Int.t_of_sexp
    let hash = Hashtbl.hash
  end
  include T
  include Comparable.Make(T)
end

let get i j matrix =
  match matrix.(i).(j)  with
  | x -> Some x
  | exception  _ -> None

let is_low i j v matrix =
  [get (i-1) (j) matrix;
   get (i) (j-1) matrix;
   get (i+1) (j) matrix;
   get (i) (j+1) matrix]
  |> List.exists ~f:(fun x ->
      match x with
      | None -> false
      | Some x' -> v >= x') |> not

let rec check acc ins = function
  | [] -> acc
  | (x,y,_)::tl ->
    let visited = Hashtbl.create (module IntPair) in
    let find p = Hashtbl.find visited p in
    let points = ref [(x,y)] in
    let insert pts (x,y) = match get x y ins with
      | Some 9 | None ->   ()
      | _ -> match find (x,y) with
        | None -> pts:=(x,y)::!pts;
          Hashtbl.set visited (x,y) true
        | _ -> () in
    while not (List.is_empty !points) do
      let pts = ref []; in
      List.iter !points ~f:(fun (x,y) ->
          insert pts (x+1, y);
          insert pts (x-1, y);
          insert pts (x, y+1);
          insert pts(x, y-1)
        );
      points := !pts
    done;
    check ((Hashtbl.length visited)::acc) ins tl

let _ =
  let ins =In_channel.read_lines "/mnt/d/input.txt" |> List.to_array |>  Array.map ~f:(fun l -> l |> String.to_array|> Array.map ~f:(fun ch -> int_of_char ch - 48 ))  in
  let low_points =  let lp = ref [] in Array.iteri ~f:(
      fun i arr -> Array.iteri arr ~f:(
          fun j v ->
            if is_low i j v ins then lp:= (i,j,v)::!lp )
    ) ins; !lp in
  low_points |> List.fold ~init:0 ~f:(fun acc (_,_,v) -> acc+v+1) |> Printf.printf "Silver %d\n%!";
  check [] ins low_points |> List.sort ~compare:(fun a b -> if a<b then 1 else 0) |> function
  | a::b::c::_ -> a*b*c |> Printf.printf "Gold %d \n%!"
  |_ -> ()

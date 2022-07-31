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

let h = Hashtbl.create (module IntPair)

let insert k = Hashtbl.update h k ~f:(function | None -> 1 | Some x-> x+1); ()
let find k = Hashtbl.find h k

let count_overlap _ = Hashtbl.count h ~f:(fun x -> x>1)

let rec from_h ((x1,y1) as p1) ((_,y2) as p2) =
  insert p1;
  if y1>y2 then from_h (x1, y1-1) p2
  else if y1<y2 then from_h (x1, y1+1) p2
  else ()

let rec from_v ((x1,y1) as p1) ((x2,_) as p2) =
  insert p1;
  if x1>x2 then from_v (x1-1, y1) p2
  else if x1<x2 then from_v (x1+1, y1) p2
  else ()

let rec from_d ((x1,y1) as p1) ((x2,y2) as p2) =
  insert p1;
  if x1>x2 && y1>y2 then from_d (x1-1, y1-1) p2
  else if x1>x2 && y1<y2 then from_d (x1-1, y1+1) p2
  else if x1<x2 && y1<y2 then from_d (x1+1, y1+1) p2
  else if x1<x2 && y1>y2 then from_d (x1+1, y1-1) p2
  else ()

let from_h_v ((x1,y1) as p1) ((x2,y2) as p2) =
  if x1 = x2 then from_h p1 p2
  else if y1 = y2 then from_v p1 p2
  else ()

let from_all ((x1,y1) as p1) ((x2,y2) as p2) =
  if  abs(x1-x2) = abs(y1-y2) then from_d p1 p2
  else ()

let rec getPoints f = function
  | [] -> ()
  | (p1,p2)::tl -> f p1 p2;  getPoints f tl

let () =
  let instructions =
    In_channel.read_lines "/mnt/d/input.txt"
    |> List.map  ~f:(fun l-> Scanf.sscanf l "%d,%d -> %d,%d" (fun x1 y1 x2 y2 -> ((x1, y1), (x2, y2)))) in
  getPoints from_h_v instructions |> count_overlap |> Printf.printf "Silver %d \n%!";
  getPoints from_all instructions |> count_overlap |> Printf.printf "Silver %d \n%!";

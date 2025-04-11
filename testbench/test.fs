
// let rec mergeSort L = 
//     match L with
//     | [] -> []
//     | head::[] -> [head]
//     | _ -> let N = List.length L
//            let (left, right) = List.splitAt(N/2) L
//            let merge (mergeSort left) (mergeSort right)


// Merge: Given two sorted lists, produce one sorted list.
// let rec merge left right =
//     match left, right with
//     | [], _ -> right
//     | _, [] -> left
//     | x::xs, y::ys ->
//         if x <= y then x :: merge xs right
//         else y :: merge left ys
//
// // Merge sort: recursively sort the list.
// let rec mergeSort L = 
//     match L with
//     | [] -> []
//     | [_] -> L
//     | _ -> 
//         let (left, right) = List.splitAt (List.length L / 2) L
//         merge (mergeSort left) (mergeSort right)

let lastElemHigher L =
    List.reduce (fun _ x -> x) L

let rec lastElemRec L =
    match L with
    | [] -> 0
    | [x] -> x
    | _ :: tl -> lastElemRec tl

let lastElemTailRec L =
    match L with
    | [] -> 0
    | hd :: tl -> 
        let rec helper acc xs =
            match xs with
            | [] -> acc
            | h :: t -> helper h t
        helper hd tl




// let makeSetHigher L =
//     List.filter (fun x y -> if x = y then  )



// let rec isPalindromeRec L =
//     match L with
//     | [] -> true
//     | [x] -> true
//     | hd :: tl -> if hd = 

let rec reverseList L = 
    match L with
    | [] -> []
    | hd :: tl -> ( reverseList tl ) @ [ hd ]



let rec myMap F L =
    match L with
    | [] -> []
    | hd :: tl -> (F hd) :: myMap F tl

let median L =
    match L with
    | [] -> 0.0
    | _ -> 
        let len = List.length L
        let midIdx = len / 2
        if ( len % 2 ) = 0 then
            float (List.item (midIdx - 1) L + List.item (midIdx) L) / 2.0
        else
            float ( List.item midIdx L )











let rec sumOdd L =
    match L with
    | [] -> 0
    | [x] -> x
    | hd :: sc :: tl -> 
        sumOdd tl + hd




[<EntryPoint>]
let main argv =
    let L = [1;2;3;4]
    let L2 = [1;2;3;4;5]
    // printfn "Higer order %A" (lastElemHigher L)
    // printfn "Tailrec %A" (lastElemTailRec L)
    // printfn "rec %A" (lastElemRec L)
    printfn "myMap output: %A" (median L)
    printfn "expected %A" (median L2) 
    // printfn "sumOdd %A" (sumOdd L2)
    0

























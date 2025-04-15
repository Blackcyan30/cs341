//
// Write a header comment here!
// 

namespace compiler

module parser =
  //
  // NOTE: all functions in the module must be indented.
  //

  let beginsWith (pattern: string) (literal: string) =
    literal.StartsWith (pattern)


  //
  // matchToken
  //
  let private matchToken expected_token tokens =
    //
    // if the next token matches the expected token,  
    // keep parsing by returning the rest of the tokens.
    // Otherwise throw an exception because there's a 
    // syntax error, effectively stopping compilation
    // at the first error.
    //
    let next_token = List.head tokens

    if expected_token = next_token then  
      List.tail tokens
    else
      failwith ("expecting " + expected_token + ", but found " + next_token)

  let empty tokens = 
    matchToken ";" tokens

  let stmt tokens = 
    printfn "stmt %A" tokens
    match tokens with
    | ";" :: _ -> empty tokens 
    | "cout" :: _ -> [] // output tokens
    | _ -> tokens

  let rec moreStmts tokens = 
    match tokens with
    | "}" :: _ -> tokens 
    | _ -> stmt tokens |> moreStmts



  let rec stmts tokens = 
    // match token with
    // | [] -> []
    // | x -> 
    // | statement :: moreStatements ->
    //   matchToken ""
    //   moreStmts moreStatements
    printfn "stmts: %A" tokens
    tokens
    |> stmt
    |> moreStmts

  let exprValue tokens =
    match tokens with
    | "true" -> matchToken "true" tokens
    | "false" -> matchToken "false" tokens
    | _ -> if beginsWith "str_literal" tokens then
             List.tail tokens
           else if beginsWith "int_literal" tokens then
             List.tail tokens
           else
            failwith("expecting literal, but found" + token)




    // let t2 = stmt tokens
    // let t3 = moreStmts t2
    // t3
  
  //
  // verySimpleC
  //
  let private verySimpleC tokens = 
    // 
    // TODO: Start here by filling it in and
    //       creating your other functions!
    //
    let t2 = matchToken "void" tokens
    let t3 = matchToken "main" t2
    let t4 = matchToken "(" t3
    let t5 = matchToken ")" t4
    let t6 = matchToken "{" t5
    let t7 = stmts t6
    let t8 = matchToken "}" t7
    let t9 = matchToken "$" t8
    t9

    // matchToken "$" tokens

  //
  // parse tokens
  //
  // Given a list of tokens, parses the list and determines
  // if the list represents a valid VerySimpleC program.  Returns
  // the string "success" if valid, otherwise returns a 
  // string of the form "syntax_error:...".
  //
  let parse tokens = 
    try
      let result = verySimpleC tokens
      "Success!"
    with 
      | ex -> "syntax_error: " + ex.Message

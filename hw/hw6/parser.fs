//
// Write a header comment here!
// 

namespace compiler

module parser =
  //
  // NOTE: all functions in the module must be indented.
  //

  let beginsWith (pattern: string) (literal: string) =
    literal.StartsWith(pattern)

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

  let exprValue tokens =
    match tokens with
    | token :: _ when token = "true" -> matchToken "true" tokens
    | token :: _ when token = "false" -> matchToken "false" tokens
    | token :: _ ->
        if beginsWith "str_literal" token then
          List.tail tokens
        else if beginsWith "int_literal" token then
          List.tail tokens
        else
          failwith ("expecting literal, but found " + token)
    | [] -> failwith "no tokens available when expecting literal"

  let outputValue tokens =
    match tokens with
    | token :: _ when token = "endl" -> matchToken "endl" tokens
    | _ -> exprValue tokens

  let output tokens = 
    if (List.head tokens) = "cout" then
      let nxtList = List.tail tokens
      if (List.head nxtList) = "<<" then
        List.tail nxtList
        |> outputValue 
      else
        failwith ("expecting <<, but found " + List.head nxtList)
    else
      failwith ("expecting cout, but found " + List.head tokens)

  let strip (s: string) = s.Trim()

  let empty tokens =
    matchToken ";" tokens

  let stmt tokens = 
    match tokens with
    | token :: _ when token = ";" -> empty tokens 
    | token :: _ when token = "cout" -> output tokens
    | token :: _ when token = "}" -> 
         failwith ("expecting statement, but found " + token)
    | [] -> failwith "unexpected end of input when expecting statement"

  let rec moreStmts tokens = 
    match tokens with
    | "}" :: _ -> tokens 
    | _ -> stmt tokens |> moreStmts

  let rec stmts tokens = 
    tokens
    |> stmt
    |> moreStmts

  let private verySimpleC tokens = 
    let t2 = matchToken "void" tokens
    let t3 = matchToken "main" t2
    let t4 = matchToken "(" t3
    let t5 = matchToken ")" t4
    let t6 = matchToken "{" t5
    let t7 = stmts t6
    let t8 = matchToken "}" t7
    let t9 = matchToken "$" t8
    t9

  //
  // parse tokens
  // Given a list of tokens, parses the list and determines
  // if the list represents a valid VerySimpleC program. Returns
  // the string "Success!" if valid, otherwise returns a string
  // of the form "syntax_error: ...".
  //
  let parse tokens = 
    try
      let _ = verySimpleC tokens
      "Success!"
    with 
      | ex -> "syntax_error: " + ex.Message

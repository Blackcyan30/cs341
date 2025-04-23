//
// Parser for SimpleC programs.  This component checks 
// the input program to see if it meets the syntax rules
// of SimpleC.  The parser returns a string denoting
// success or failure. 
//
// Returns: the string "success" if the input program is
// legal, otherwise the string "syntax_error: ..." is
// returned denoting an invalid SimpleC program.
//
// Talha Khan
// CS 341, Fall 2025
// NetID: mkhan387
//
// Original author:
//   Prof. Joe Hummel
//   U. of Illinois, Chicago
//   CS 341, Spring 2022
//

namespace compiler

module parser =
    //
    // NOTE: all functions in the module must be indented.
    //

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

    let private beginswith (pattern: string) (literal: string) =
        literal.StartsWith (pattern)

    let private empty tokens = 
        matchToken ";" tokens
    let private expr_op tokens =
        let token = List.head tokens
        match token with
        | "+" -> matchToken "+" tokens
        | "-" -> matchToken "-" tokens
        | "*" -> matchToken "*" tokens
        | "/" -> matchToken "/" tokens
        | "^" -> matchToken "^" tokens
        | "<" -> matchToken "<" tokens
        | "<=" -> matchToken "<=" tokens
        | ">" -> matchToken ">" tokens
        | ">=" -> matchToken ">=" tokens
        | "==" -> matchToken "==" tokens
        | "!=" -> matchToken "!=" tokens
        | _ -> failwith ("expecting expression operator, but found " + token)


    let private expr_value tokens = 
        match tokens with
        | hd :: tl when hd = "true" -> matchToken "true" tokens
        | hd :: tl when hd = "false" -> matchToken "false" tokens
        | hd :: tl when  beginswith "identifier" hd -> tl
        | hd :: tl when  beginswith "str_literal" hd -> tl
        | hd :: tl when  beginswith "int_literal" hd -> tl
        | _ -> failwith ("expecting identifier or literal, but found " + List.head tokens)


    let private expr tokens = 
        let tail = expr_value tokens
        let token = List.head tail
        match token with
        | "+" | "-" | "*" | "/" | "^" 
        | "<" | "<=" | ">" | ">=" 
        | "==" | "!=" -> expr_op tail |> expr_value
        | _ -> tail

    let private output_value tokens = 
        let token = List.head tokens
        if token = "endl" then 
            matchToken "endl" tokens
        else
            expr_value tokens


    let private output tokens = 
        matchToken "cout" tokens
        |> matchToken "<<"
        |> output_value 
        |> matchToken ";"

    let private input tokens = 
        let t1 = matchToken "cin" tokens
        let t2 = matchToken ">>" t1
        if beginswith "identifier" (List.head t2) then
            let t3 = List.tail t2
            let t4 = matchToken ";" t3
            t4
        else
            failwith ("expecting identifier, but found " + List.head t2)
            // failwith ("in input expecting identifier, but found " + List.head t2)

    let private vardecl tokens = 
        let token = List.head tokens
        if token = "int" then
            let t1 = matchToken "int" tokens
            if beginswith "identifier" (List.head t1) then
                let t2 = List.tail t1
                let t3 = matchToken ";" t2
                t3
            else
                failwith ("expecting identifier, but found " + List.head t1)
                // failwith("in vardecl expecting identifier, but found " + List.head t1)
        else
            failwith ("expecting int, but found " + token)
            // failwith ("in vardecl expecting int, but found " + token)


    let private assignment tokens =
        if beginswith "identifier" (List.head tokens) then
            let t1 = List.tail tokens
            let t2 = matchToken "=" t1
            let t3 = expr t2
            let t4 = matchToken ";" t3
            t4
        else
            failwith ("expecting identifier, but found " + List.head tokens)
            // failwith ("in assignment expecting identifier, but found " + List.head tokens)

    let rec private stmt tokens =
        let hd = List.head tokens
        match hd with
        | ";" -> empty tokens
        | "int" -> vardecl tokens
        | "cin" -> input tokens
        | "cout" -> output tokens
        | "if" -> ifstmt tokens
        | _ when beginswith "identifier" hd -> assignment tokens
        | _ -> failwith ("expecting statement, but found " + hd)


    and private ifstmt tokens = 
        matchToken "if" tokens
        |> matchToken "("
        |> condition
        |> matchToken ")"
        |> then_part
        |> else_part

    and private condition tokens = 
        expr tokens

    and private then_part tokens =
        stmt tokens

    and private else_part tokens =
        match tokens with
        | hd :: tail when hd = "else" ->
            stmt tail
        | _ -> tokens


    let rec private moreStmts tokens =
        match tokens with
        | "}" :: _ -> tokens 
        | _ -> stmt tokens |> moreStmts

    let rec stmts tokens = 
        tokens
        |> stmt
        |> moreStmts


  //
  // simpleC
  //
    let private simpleC tokens = 
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
    // if the list represents a valid SimpleC program.  Returns
    // the string "success" if valid, otherwise returns a 
    // string of the form "syntax_error:...".
    //
    let parse tokens = 
        try
        let result = simpleC tokens
        "Success!"
        with 
        | ex -> "syntax_error: " + ex.Message

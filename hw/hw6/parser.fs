//
// Write a header comment here!
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

  //
  // verySimpleC
  //
  let private verySimpleC tokens = 
    // 
    // TODO: Start here by filling it in and
    //       creating your other functions!
    //
    matchToken "$" tokens

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

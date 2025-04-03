// @author Muhammad Talha Adnan Khan
// @file Program.fs
// @description This program takes in a csv file and outputs
// statistics regarding that file for divvy bike data of chicago.
// --------------------------------------------------------------
// 
// ParseLine and ParseInput
//
// Given a sequence of strings representing Divvy data, 
// parses the strings and returns a list of lists. Each
// sub-list denotes one bike ride.
// 
// Example:
//   [ [21864678; 6035; 3000; 23; 317; 338; 1; 1; 1974]; ... ]
//
// The values are:
//    trip_id
//    bike_id
//    trip_duration   (seconds)
//    starting_hour   (between 0 and 23)
//    from_station_id
//    to_station_id
//    is_subscriber   (1 if yes, 0 if no)
//    gender          (0 if not specified, 1 if identifies as male, 2 if identifies as female)
//    birth_year      (0 if not specified)
// 


let ParseLine (line:string) = 
  let tokens = line.Split(',')
  let ints = Array.map System.Int32.Parse tokens
  Array.toList ints

let rec ParseInput lines = 
  let rides = Seq.map ParseLine lines
  Seq.toList rides

let printHead list : unit = 
  list 
  |> List.truncate 5
  |> List.iter (fun innerList -> printfn "%A" innerList)

let riderInfo ridedata (totalRides: int) (gender: int) = 
  let numPersons = 
    ridedata
    |> List.filter (fun ride -> List.item 7 ride = gender)
    |> List.length
  
  let percentage =
    (float numPersons / float totalRides) * 100.0

  (numPersons, percentage)

let filterData data func = 
  List.filter func data

let rec ages data ( acc: int ) =
  match data with
  | [] -> acc
  | hd::tl -> ages tl (acc + ( System.DateTime.Now.Year - (List.item 8 hd) ))

let avgAge ridedata: float =
  let cleaneddata = filterData ridedata ( fun ride -> (List.item 8 ride > 0) )
  let totalAge = ages cleaneddata 0
  let numRides = List.length cleaneddata
  float totalAge / float numRides

let ridesByTripDuration ridedata ( startTime: int ) ( endTime: int ): int =
  ridedata
  |> List.filter (fun ride -> List.item 2 ride > startTime && List.item 2 ride <= endTime)
  |> List.length

let ridesOverTwoHours ridedata: int =
  ridedata
  |> List.filter (fun ride -> List.item 2 ride > 7200)
  |> List.length

let ridesByHour ridedata ( hour: int ): int = 
  ridedata
  |> List.filter( fun ride -> List.item 3 ride = hour )
  |> List.length

let histDisplayUnit ridedata ( hour: int ): unit =
  let rec displayStars numStars: unit =
    if numStars > 0 then
      printf "*"
      displayStars (numStars - 1)

  let baseDisplay ( hour: int ): unit =
    printf "   %A: " hour

  let numRides: int = ridesByHour ridedata hour
  baseDisplay hour
  displayStars (numRides / 100)
  printfn "%A" numRides

let rec histDisplay ridedata ( acc: int ) = 
  if acc >= 24 then 
    ()
  else
    histDisplayUnit ridedata acc
    histDisplay ridedata (acc + 1)


// --------------------------------------------------------------
// main function
// 

[<EntryPoint>]
let main argv =

  printfn "Project 3: Divvy Rides Data Analysis with F#"
  printfn "CS 341, Spring 2025"
  printfn "0.0"
  printfn "This application allows you to analyze and visualize"
  printfn "information about Divvy bike rides in Chicago, such as"
  printfn "the number of male/female riders, the average age, etc."
  printfn ""

  printf "Enter the name of the file with the Divvy ride data: "
  let filename = System.Console.ReadLine()
  let contents = System.IO.File.ReadLines(filename)
  let ridedata = ParseInput contents

  let numTrips = ( ridedata |> List.length )
  printfn ""
  printfn "Number of Trips: %A" numTrips
  printfn ""

  let (numMales, percentageMale) = riderInfo ridedata numTrips 1
  let (numFemales, percentageFemale) = riderInfo ridedata numTrips 2
  printfn "Number of Riders Identifying as Male: %A (%A%%)" numMales percentageMale
  printfn "Number of Riders Identifying as Female: %A (%A%%)" numFemales percentageFemale
  printfn ""
 
  let age = avgAge ridedata
  printfn "Average Age: %A" age
  printfn ""

  let firstInterval = ridesByTripDuration ridedata 0 ( 30 * 60 )
  let secondInterval = ridesByTripDuration ridedata ( 30 * 60 ) ( 60 * 60 )
  let thirdInterval = ridesByTripDuration ridedata ( 60 * 60 ) ( 120 * 60 )
  let fourthInterval = ridesOverTwoHours ridedata
  printfn "Ride Durations: "
  printfn "   0-30 mins: %A (%A%%)" firstInterval ( (float firstInterval / float numTrips) * 100.0)
  printfn "   30-60 mins: %A (%A%%)" secondInterval ( (float secondInterval / float numTrips) * 100.0)
  printfn "   60-120 mins: %A (%A%%)" thirdInterval ( (float thirdInterval / float numTrips) * 100.0)
  printfn "   > 2 hours: %A (%A%%)" fourthInterval ( (float fourthInterval / float numTrips) * 100.0)
  printfn ""
  
  printfn "Histogram of Start Times:"
  histDisplay ridedata 0
  printfn ""
  printfn "Exiting program."
  0

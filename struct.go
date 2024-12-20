// Define a Car struct with the following int type fields:

// battery
// batteryDrain
// speed
// distance
// Allow creating a remote controlled car by defining a function NewCar that takes the speed of the car in meters, and the battery drain percentage as its two parameters (both of type int) and returns a Car instance:

// speed := 5
// batteryDrain := 2
// car := NewCar(speed, batteryDrain)
// // => Car{speed: 5, batteryDrain: 2, battery:100, distance: 0}

// Stuck? Reveal Hints
// Opens in a modal
// Define another struct type called Track with the field distance of type integer. Allow creating a race track by defining a function NewTrack that takes the track's distance in meters as its sole parameter (which is of type int):

// distance := 800
// track := NewTrack(distance)
// // => Track{distance: 800}

// Stuck? Reveal Hints
// Opens in a modal
// Implement the Drive function that updates the number of meters driven based on the car's speed, and reduces the battery according to the battery drainage. If there is not enough battery to drive one more time the car will not move:

// speed := 5
// batteryDrain := 2
// car := NewCar(speed, batteryDrain)
// car = Drive(car)
// // => Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}

// Stuck? Reveal Hints
// Opens in a modal
// To finish a race, a car has to be able to drive the race's distance. This means not draining its battery before having crossed the finish line. Implement the CanFinish function that takes a Car and a Track instance as its parameter and returns true if the car can finish the race; otherwise, return false.

// Assume that you are currently at the starting line of the race and start the engine of the car for the race. Take into account that the car's battery might not necessarily be fully charged when starting the race:

// speed := 5
// batteryDrain := 2
// car := NewCar(speed, batteryDrain)

// distance := 100
// track := NewTrack(distance)

// CanFinish(car, track)
// // => true

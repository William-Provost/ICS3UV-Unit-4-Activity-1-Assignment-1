/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-25
 * @fileoverview This program calculates the average of a set of marks
 * and displays a performance message.
 */

// variables
let numberOfMarks: number = 0;
let markInput: number = 0;
let sum: number = 0;
let average: number = 0;

// ask for number of marks
numberOfMarks = parseInt(prompt("How many marks will you enter today? ") || "0");

// loop to get all marks
for (let counter = 1; counter <= numberOfMarks; counter++) {
  markInput = parseFloat(prompt("Enter mark " + counter + ": ") || "0");
  sum = sum + markInput;
}

// calculate average
average = sum / numberOfMarks;

// display results
console.log(
  "You have entered " +
  numberOfMarks +
  " marks. The student's average is " +
  Math.round(average) +
  "%."
);

// performance message
if (average <= 49) {
  console.log("The student is failing.");
} else if (average >= 50 && average <= 69) {
  console.log("The student's performance is just under average.");
} else if (average >= 70 && average <= 79) {
  console.log("The student's performance is average.");
} else {
  console.log("The student is on the honour roll.");
}

console.log("\nDone.");

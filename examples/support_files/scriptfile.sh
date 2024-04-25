#!/bin/bash

# A simple bash script to demonstrate basic scripting concepts

echo "Starting the script execution..."

# Define a variable
GREETING="Hello, world!"

# Print the variable
echo $GREETING

# Define a function
print_date() {
    echo "Today's date is: $(date)"
}

# Call the function
print_date

# Demonstrate a for loop
for i in {1..5}; do
    echo "Iteration number $i"
done

echo "Script execution completed."

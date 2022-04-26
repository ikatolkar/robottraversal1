package main

import "fmt"

type Coordinates struct{
  X int
  Y int
  Dir string
}

type Robot struct{
  Cur Coordinates
}

type Plane struct{
  Min Coordinates
  Max Coordinates
}

// Update modifies a coordinate based on command.
// Command is a single letter  M, L, R.
func (c* Coordinates) Update(cmd string) {
  if cmd == "M" {
    // Maintain current direction and move one step
    switch c.Dir {
      case "N":
        c.Y++
      case "S":
        c.Y--
      case "E":
        c.X++
      case "W":
        c.X--
    }
  } else if cmd == "L" {
    // Turn left without moving any steps
    switch c.Dir {
      case "N":
        c.Dir = "W"
      case "S":
        c.Dir = "E"
      case "E":
        c.Dir = "N"
      case "W":
        c.Dir = "S"
    }
  } else if cmd == "R" {
    // Turn right without moving any steps
    switch c.Dir {
      case "N":
        c.Dir = "E"
      case "S":
        c.Dir = "W"
      case "E":
        c.Dir = "S"
      case "W":
        c.Dir = "N"
    }
  }
}

// CheckInsideBounds returns true only if input coordinates are inside bounds of plane.
func (p Plane) CheckInsideBounds(cur Coordinates) bool {
  if cur.X > p.Max.X || cur.X < p.Min.X || cur.Y > p.Max.Y || cur.Y < p.Min.Y {
    // Out of bounds of plane
    return false
  }
  return true
}

// Move updates robots coordinates to move it inside a plane.
// It validates coordinates to ensure robot never goes out of bounds of plane.
func (r *Robot) Move(p Plane, cmd string) {
  tmp := Coordinates{}
  for i := 0; i < len(cmd); i++ {
    tmp = r.Cur
    tmp.Update(string(cmd[i]))
    if p.CheckInsideBounds(tmp) {
      r.Cur = tmp
    }else{
      break
    }
  }
}

func main() {
  // Take inputs
  var max, cur Coordinates
  var command string

  // Size of plane
  fmt.Scan(&max.X)
  fmt.Scan(&max.Y)

  // Robot's start coordinates and direction
  fmt.Scan(&cur.X)
  fmt.Scan(&cur.Y)
  fmt.Scan(&cur.Dir)

  // Command for moving robot
  fmt.Scan(&command)

  // Create plane and robot
  plane := Plane{Max: max}
  robot := Robot{Cur: cur}

  // Move Robot in plane based on input command
  robot.Move(plane, command)

  // Output final coordinates
  fmt.Println(robot.Cur.X, robot.Cur.Y, robot.Cur.Dir)
}

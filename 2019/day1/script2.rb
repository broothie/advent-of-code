
lines = File.read('input.txt').split("\n")

def find_fuel(mass)
  fuel = mass / 3 - 2
  return 0 if fuel < 0

  fuel + find_fuel(fuel)
end

masses = lines.map(&:to_i)
total_mass = masses.map(&method(:find_fuel)).sum
puts total_mass
